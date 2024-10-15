package main

import (
	"fmt"
	"net/http"

	"github.com/Prakhar256/RSS_aggregator/internal/auth"
	"github.com/Prakhar256/RSS_aggregator/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Authentication Error %v", err))
			return
		}
		user, err := apiCfg.DB.GetUserNyAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf(" Couldn't get the user  %v", err))
			return
		}
		handler(w,r,user)
		// If authentication is successful, the middleware passes the control to the actual handler (handler(w, r, user)) along with the User object obtained from the database.
	}
}
