package storeConsent

import (
	"net/http"

	"github.com/oskaremilsson/spotify-controller/database"
	"github.com/oskaremilsson/spotify-controller/utils/decodeSpotifyToken"
	"github.com/oskaremilsson/spotify-controller/utils/infoJson"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	token := r.Form.Get("token")
	allow_user := r.Form.Get("allow_user")

	username := decodeSpotifyToken.GetUsername(token)

	if username == "bad_token" || allow_user == "" {
		info := infoJson.Parse("Missing username or allow_user", false)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(info)
		return
	}

	if database.StoreConsent(username, allow_user) {
		info := infoJson.Parse(username+" now allow "+allow_user, true)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(info)
		return
	}

	info := infoJson.Parse("Could not save to database", false)
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write(info)
}
