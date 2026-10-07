// Package migrations embeds the forward-only SQL migration files so the api
// binary can apply them at startup without shipping the .sql files separately.
package migrations

import "embed"

// FS holds every numbered migration (NNN_*.sql), applied in lexical order.
//
//go:embed *.sql
var FS embed.FS
