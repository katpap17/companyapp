package restserver

import (
	"net/http"

	"github.com/katpap17/companyapp/auth"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := ExtractToken(r)
		if token == "" {
			handleErrorResponse(w, "No token provided", http.StatusBadRequest)
			return
		}
		claims, valid := auth.ValidateToken(token)
		if !valid || claims == nil {
			handleErrorResponse(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func ExtractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	if len(bearerToken) > 7 && bearerToken[:7] == "Bearer " {
		return bearerToken[7:]
	}
	return ""
}
