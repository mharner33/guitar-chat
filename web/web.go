// Package web holds the embedded static assets for the chat UI, served by the
// api binary via go:embed (no framework, no build step).
package web

import "embed"

//go:embed index.html app.js style.css
var Assets embed.FS
