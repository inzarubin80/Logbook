package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/inzarubin80/Logbook/auth/db"
	"github.com/inzarubin80/Logbook/auth/env"
	"github.com/inzarubin80/Logbook/auth/errors"
	"github.com/inzarubin80/Logbook/auth/server/jwt"
	"github.com/inzarubin80/Logbook/auth/server/write"
)

type loginRequest struct {
	Email string
	Pass  string
}

func Login(env env.Env, user *db.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	decoder := json.NewDecoder(r.Body)
	req := &loginRequest{}
	err := decoder.Decode(&req)
	if err != nil || req == nil {
		return write.Error(errors.NoJSONBody)
	}

	u, err := env.DB().FindUserByEmail(r.Context(), req.Email)
	if err != nil {
		if isNotFound(err) {
			return write.Error(errors.FailedLogin)
		}
		return write.Error(err)
	}

	if !checkPasswordHash(req.Pass, u.Salt, u.Pass) {
		return write.Error(errors.FailedLogin)
	}

	jwt.WriteUserCookie(w, &u)
	return write.JSON(&u)
}

func Logout(env env.Env, user *db.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	u := &db.User{}
	jwt.WriteUserCookie(w, u)
	return write.Success()
}
