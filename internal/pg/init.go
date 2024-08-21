package pg

import (
	"amartha-loan-system/internal/pkg/config"
	"database/sql"

	"cloud.google.com/go/cloudsqlconn/postgres/pgxv4"
	_ "github.com/lib/pq"
)

type DB struct {
	MasterConn  *sql.DB
	ReplicaConn *sql.DB
}

func Initialize(master, replica string) DB {
	driverName := "postgres"

	var err error

	if config.Instance().PG.UseCloudSQL {
		driverName = "cloudsql-postgres"
		// Register a driver using whatever name you like.
		cleanup, err := pgxv4.RegisterDriver(
			driverName,
			// any desired options go here, for example:
		)
		if err != nil {
			panic(err)
		}
		// call cleanup to close the underylying driver when you're done with the
		// db.
		defer cleanup()
	}

	masterConn, err := sql.Open(
		driverName,
		master,
	)
	if err != nil {
		panic(err)
	}
	replicaConn, err := sql.Open(
		driverName,
		replica,
	)
	if err != nil {
		panic(err)
	}

	err = masterConn.Ping()
	if err != nil {
		panic(err)
	}

	return DB{
		MasterConn:  masterConn,
		ReplicaConn: replicaConn,
	}
}
