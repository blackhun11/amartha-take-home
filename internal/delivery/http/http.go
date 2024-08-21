package http

import (
	"amartha-loan-system/internal/model/request"
	"amartha-loan-system/internal/usecase/loan"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Common struct {
	Data any `json:"data,omitempty"`
}

type HTTPHandler struct {
	loan.LoanUsecase
}

func NewHTTPHandler(loanUsecase loan.LoanUsecase) HTTPHandler {
	return HTTPHandler{
		LoanUsecase: loanUsecase,
	}
}

func (h *HTTPHandler) CreateLoan(c echo.Context) error {
	var req request.CreateLoanRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	loan, err := h.LoanUsecase.CreateLoan(c.Request().Context(), req.ToModel())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, Common{
		Data: loan,
	})

}

func (h *HTTPHandler) ApproveLoan(c echo.Context) error {
	var req request.LoanApprovalRequest

	idInt, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	loan, err := h.LoanUsecase.ApproveLoan(c.Request().Context(), req.ToModel(), idInt)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, Common{
		Data: loan,
	})
}

func (h *HTTPHandler) RecordLoanInvestment(c echo.Context) error {
	var req request.LoanInvestmentRequest

	idInt, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	loan, err := h.LoanUsecase.RecordLoanInvestment(c.Request().Context(), req.ToModel(), idInt)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, Common{
		Data: loan,
	})

}

func (h *HTTPHandler) RecordLoanDisbursement(c echo.Context) error {
	var req request.LoanDisbursementRequest

	idInt, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	loan, err := h.LoanUsecase.RecordDisburseLoan(c.Request().Context(), req.ToModel(), idInt)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, Common{
		Data: loan,
	})

}
