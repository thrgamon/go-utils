package authentication

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
  userRepo "github.com/thrgamon/go-utils/repo/user"
)

var Store *sessions.CookieStore
var Logger *log.Logger
var UserRepo *userRepo.UserRepo


type Profile struct {
	Nickname string
	Sub      string
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	sessionState, _ := Store.Get(r, "auth")
	authenticator, _ := New()

	if queryValues.Get("state") != sessionState.Values["state"] {
		http.Error(w, "Oh No", http.StatusBadRequest)
		return
	}

	token, err := authenticator.Exchange(r.Context(), queryValues["code"][0])

	if err != nil {
		http.Error(w, "Oh No", http.StatusUnauthorized)
		return
	}

	idToken, err := authenticator.VerifyIDToken(r.Context(), token)

	if err != nil {
		http.Error(w, "There was an unexpected error", http.StatusInternalServerError)
		Logger.Println(err.Error())
		return
	}

	var profile Profile
	if err := idToken.Claims(&profile); err != nil {
		http.Error(w, "There was an unexpected error", http.StatusInternalServerError)
		Logger.Println(err.Error())
		return
	}

	exists, err := UserRepo.Exists(r.Context(), Auth0ID(profile.Sub))
	if err != nil {
		http.Error(w, "There was an unexpected error", http.StatusInternalServerError)
		Logger.Println(err.Error())
		return
	}

	if !exists {
		err := UserRepo.Add(r.Context(), Username(profile.Nickname), Auth0ID(profile.Sub))
		if err != nil {
			http.Error(w, "There was an unexpected error", http.StatusInternalServerError)
			Logger.Println(err.Error())
			return
		}
	}

	sessionState.Values["access_token"] = token.AccessToken
	sessionState.Values["user_id"] = profile.Sub
	if err := sessionState.Save(r, w); err != nil {
		http.Error(w, "There was an unexpected error", http.StatusInternalServerError)
		Logger.Println(err.Error())
		return
	}

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
