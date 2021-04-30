package googlerouter

import (
	"github.com/MohitVachhani/go-learn/pkg/utils/env"
	"github.com/gorilla/mux"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"

	googleAuth "github.com/MohitVachhani/go-learn/cmd/service/auth/google"
)

func InitializeGoogleAuthRouter(router *mux.Router) *mux.Router {

	var googleClientId = env.Get("GOOGLE_CLIENT_ID")
	var googleClientSecret = env.Get("GOOGLE_CLIENT_SECRET")
	var googleCallbackUrl = env.Get("GOOGLE_CALLBACK_URL")

	goth.UseProviders(
		google.New(googleClientId, googleClientSecret, googleCallbackUrl, "email", "profile"),
	)

	router.HandleFunc("/{provider}", googleAuth.HandleGoogleAuth).Methods("GET")
	router.HandleFunc("/{provider}/callback", googleAuth.HandleGoogleAuthCallback).Methods("GET")

	return router
}
