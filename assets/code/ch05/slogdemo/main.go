// Command slogdemo shows structured logging with log/slog (Go 1.21+):
// levels, key/value attributes, groups, child loggers and JSON output.
//
// Run from assets/code:
//
//	go run ./ch05/slogdemo
package main

import (
	"log/slog"
	"os"
)

func main() {
	// Drop the time attribute so the output is identical on every run.
	noTime := func(_ []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey {
			return slog.Attr{}
		}
		return a
	}

	// Human-friendly text for local development...
	text := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{ReplaceAttr: noTime}))
	text.Info("server starting", "port", 8080, "env", "dev")

	// ...machine-friendly JSON for production log pipelines.
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:       slog.LevelDebug, // show Debug and above
		ReplaceAttr: noTime,
	}))

	// A child logger carries context into every line it writes.
	reqLog := logger.With("request_id", "req-7f3a", "user", "ada")
	reqLog.Debug("cache lookup", "key", "items:42", "hit", false)
	reqLog.Info("order placed",
		slog.Int("items", 3),
		slog.Group("payment", slog.String("method", "card"), slog.Float64("amount", 59.97)),
	)
	reqLog.Warn("slow query", "table", "orders", "ms", 812)
	reqLog.Error("charge failed", "err", "card declined")
}
