package router

import (
	"io/fs"
	"net/http"
	"strings"

	"github.com/mhafidk/warga-sekolah/frontend"
	"github.com/mhafidk/warga-sekolah/internal/auth"
	"github.com/mhafidk/warga-sekolah/internal/user"
)

type Handlers struct {
	User *user.Handler
}

func New(jwtSecret []byte, h Handlers) http.Handler {
	mainMux := http.NewServeMux()

	// Register API Routes
	mainMux.HandleFunc("POST /api/v1/users", h.User.Create)
	mainMux.HandleFunc("POST /api/v1/users/login", h.User.Login)

	protectedMux := http.NewServeMux()
	protectedMux.HandleFunc("GET /api/v1/users/me", h.User.Me)

	authMiddleware := auth.RequireAuth(jwtSecret)
	mainMux.Handle("/api/v1/", authMiddleware(protectedMux))

	// Get the embedded frontend filesystem
	frontendFS := frontend.GetFS()
	fileServer := http.FileServer(http.FS(frontendFS))

	// Catch-all handler for static frontend assets & SPA routing fallback
	mainMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// API path safety catch (if something fell through)
		if strings.HasPrefix(r.URL.Path, "/api/") {
			http.NotFound(w, r)
			return
		}

		cleanedPath := strings.TrimPrefix(r.URL.Path, "/")
		if cleanedPath == "" {
			cleanedPath = "index.html"
		}

		// Try opening the requested file in the embedded filesystem
		f, err := frontendFS.Open(cleanedPath)
		if err != nil {
			// If file does not exist, serve index.html (SPA client-side router fallback)
			indexHTML, err := fs.ReadFile(frontendFS, "index.html")
			if err != nil {
				http.Error(w, "Internal Server Error: Frontend assets missing", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write(indexHTML)
			return
		}
		f.Close()

		// Serve the actual static file
		fileServer.ServeHTTP(w, r)
	})

	return mainMux
}
