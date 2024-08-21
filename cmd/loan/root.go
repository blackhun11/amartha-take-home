package loan

import (
	"amartha-loan-system/internal/pkg/config"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpHandler "amartha-loan-system/internal/delivery/http"
	"amartha-loan-system/internal/pg"
	loanRepo "amartha-loan-system/internal/repository/loan"
	loanUsecase "amartha-loan-system/internal/usecase/loan"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type application struct {
	httpHandler.HTTPHandler
}

func newApplication() application {
	return application{}
}

func (a application) config() application {
	config.Load()
	return a
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func (a application) serveHTTP() {
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{},
		AllowOriginFunc: func(origin string) (bool, error) {
			return true, nil
		},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))

	e.GET("/healthcheck", func(c echo.Context) error {
		response := map[string]string{
			"status": "healthy",
		}
		return c.JSON(http.StatusOK, response)
	})

	loanGroup := e.Group("loan")
	loanGroup.POST("/loans", a.CreateLoan)
	loanGroup.PUT("/loans/:id/approve", a.ApproveLoan)
	loanGroup.POST("/loans/:id/invest", a.RecordLoanInvestment)
	loanGroup.PUT("/loans/:id/disburse", a.RecordLoanDisbursement)

	// TODO: Create API To get loan detail
	loanGroup.GET("/loans/:id", nil)
	loanGroup.GET("/loans", nil)

	h2s := &http2.Server{}
	h1s := &http.Server{
		Addr:    ":" + config.Instance().App.ServerPort,
		Handler: h2c.NewHandler(e, h2s),
	}

	// Start server
	go func() {
		fmt.Println("Server Started at:", config.Instance().App.ServerPort)
		if err := h1s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	// Accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGTERM (Ctrl+/) is emitted by Docker stop command
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("Shutting down server...")
	// Attempt to gracefully shut down the server
	if err := h1s.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	fmt.Println("Server gracefully stopped")
}

func (a application) init() application {
	loanPG := pg.Initialize(config.Instance().PG.Master, config.Instance().PG.Replica)

	err := pg.MigrateUp(loanPG.MasterConn, "")
	if err != nil {
		panic(err)
	}
	loanRepo := loanRepo.NewPG(loanPG)
	a.HTTPHandler = httpHandler.NewHTTPHandler(loanUsecase.NewloanUsecase(loanRepo))
	return a
}

func Execute() {
	newApplication().config().init().serveHTTP()
}
