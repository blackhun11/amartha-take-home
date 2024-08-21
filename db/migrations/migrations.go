package migrations

import "embed"

//go:embed pg/*
var PGMigrationFS embed.FS
