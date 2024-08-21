package config

type PG struct {
	Master      string `envconfig:"MASTER" default:""`
	Replica     string `envconfig:"REPLICA" default:""`
	UseCloudSQL bool   `envconfig:"USE_CLOUDSQL" default:"false"`
}
