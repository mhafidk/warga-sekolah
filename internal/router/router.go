package router

import (
	"net/http"

	"github.com/mhafidk/warga-sekolah/internal/auth"
	"github.com/mhafidk/warga-sekolah/internal/user"
)

type Handlers struct {
	User *user.Handler
}

func New(jwtSecret []byte, h Handlers) http.Handler {
	mainMux := http.NewServeMux()

	mainMux.HandleFunc("POST /api/v1/users", h.User.Create)
	mainMux.HandleFunc("POST /api/v1/users/login", h.User.Login)

	protectedMux := http.NewServeMux()

	protectedMux.HandleFunc("GET /api/v1/users/me", h.User.Me)

	authMiddleware := auth.RequireAuth(jwtSecret)
	mainMux.Handle("/api/v1/", authMiddleware(protectedMux))

	return mainMux
}
