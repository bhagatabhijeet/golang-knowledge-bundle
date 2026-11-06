// Command httpserver is a small, production-shaped JSON API built only with
// the standard library. It shows:
//
//   - Go 1.22+ ServeMux patterns:  "GET /items/{id}"  and  r.PathValue("id")
//   - JSON decoding/encoding with proper status codes
//   - middleware (logging + panic recovery) as func(http.Handler) http.Handler
//   - server timeouts and graceful shutdown
//
// So that `go run` terminates on its own, main starts the server, calls it a
// few times as a client, then shuts it down gracefully.
//
// Run from assets/code:
//
//	go run ./ch05/httpserver
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

// Item is the resource our API serves.
type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// store is a tiny concurrency-safe in-memory database.
type store struct {
	mu    sync.RWMutex
	items map[string]Item
}

func newStore() *store { return &store{items: map[string]Item{}} }

func (s *store) get(id string) (Item, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	it, ok := s.items[id]
	return it, ok
}

func (s *store) put(it Item) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[it.ID] = it
}

// routes wires handlers onto a ServeMux using method + wildcard patterns.
func routes(s *store) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "ok\n")
	})

	mux.HandleFunc("GET /items/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id") // the {id} wildcard
		it, ok := s.get(id)
		if !ok {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "item not found"})
			return
		}
		writeJSON(w, http.StatusOK, it)
	})

	mux.HandleFunc("POST /items", func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // cap request bodies at 1 MiB
		var it Item
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		if err := dec.Decode(&it); err != nil || it.ID == "" {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid item"})
			return
		}
		s.put(it)
		writeJSON(w, http.StatusCreated, it)
	})

	// Middleware wraps the whole mux: outermost runs first.
	return recoverer(logger(mux))
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// statusRecorder captures the status code a handler writes.
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

// logger is middleware: it logs method, path, status and duration.
func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rec, r)
		slog.Info("request",
			"method", r.Method, "path", r.URL.Path,
			"status", rec.status, "dur", time.Since(start).Round(time.Microsecond))
	})
}

// recoverer turns a panicking handler into a 500 instead of a dead connection.
func recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if v := recover(); v != nil {
				slog.Error("panic", "value", v)
				http.Error(w, "internal error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey { // drop timestamps so output is stable
				return slog.Attr{}
			}
			return a
		},
	})))

	// Listen on a random free port so the demo never collides with anything.
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		slog.Error("listen", "err", err)
		os.Exit(1)
	}

	srv := &http.Server{
		Handler:           routes(newStore()),
		ReadHeaderTimeout: 5 * time.Second, // defends against slowloris
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	go func() {
		if err := srv.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("serve", "err", err)
		}
	}()

	base := "http://" + ln.Addr().String()
	client := &http.Client{Timeout: 3 * time.Second} // never use a client without a timeout

	call(client, base, "GET", "/healthz", "")
	call(client, base, "POST", "/items", `{"id":"42","name":"gopher plush"}`)
	call(client, base, "GET", "/items/42", "")
	call(client, base, "GET", "/items/7", "")
	call(client, base, "DELETE", "/items/42", "") // no DELETE route -> 405

	// Graceful shutdown: stop accepting, let in-flight requests finish.
	// In a real service you would trigger this from signal.NotifyContext.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("shutdown", "err", err)
	}
	fmt.Println("server stopped cleanly")
}

// call performs one request and prints a one-line summary of the response.
func call(c *http.Client, base, method, path, body string) {
	req, _ := http.NewRequest(method, base+path, strings.NewReader(body))
	if body != "" {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := c.Do(req)
	if err != nil {
		fmt.Println("client error:", err)
		return
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	fmt.Printf("%-6s %-9s -> %d %s\n", method, path, resp.StatusCode, strings.TrimSpace(string(b)))
}
