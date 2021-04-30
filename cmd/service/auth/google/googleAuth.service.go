package googleauth

import (
	"fmt"
	"net/http"

	"github.com/markbates/goth/gothic"
)

func HandleGoogleAuth(res http.ResponseWriter, req *http.Request) {
	gothic.BeginAuthHandler(res, req)
}

func HandleGoogleAuthCallback(res http.ResponseWriter, req *http.Request) {
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}

	fmt.Println("user", user)

}
