package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"radiant-protal-api/entity"
	"radiant-protal-api/models"
	u "radiant-protal-api/utils"

	"github.com/dgrijalva/jwt-go"
)

var Login = func(w http.ResponseWriter, r *http.Request) {

	user := &models.UserListModel{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request", err.Error()))
		return
	}

	if user.UserCode == "" {
		u.Respond(w, u.Message(false, "Please enter SPA ID", ""))
		return
	}

	if user.Password == "" {
		u.Respond(w, u.Message(false, "Please enter password", ""))
		return
	}

	var userResult []entity.LoginResult
	err_db := entity.Login(user.UserCode, user.Password, &userResult)
	if err_db != nil {
		u.Respond(w, u.Message(false, "Database Server error!", err_db.Error()))
		return
	}

	if len(userResult) == 0 {
		u.Respond(w, u.Message(false, "Please check your SAP id and password", ""))
		return
	}

	//Create JWT token
	tk := &models.TokenModel{UserId: uint(userResult[0].ID)}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	userResult[0].Token = tokenString //Store the token in the response

	resp := u.Message(true, "Data found!", "")
	resp["result"] = userResult[0]
	u.Respond(w, resp)

}

var CheckUser = func(w http.ResponseWriter, r *http.Request) {

	user := &models.UserListModel{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request", err.Error()))
		return
	}

	if user.UserCode == "" {
		u.Respond(w, u.Message(false, "Please enter SPA ID", ""))
		return
	}

	var userResult []entity.LoginResult
	err_db := entity.CheckUser(&userResult, user.UserCode)
	if err_db != nil {
		u.Respond(w, u.Message(false, "Database Server error!", err_db.Error()))
		return
	}

	if len(userResult) == 0 {
		u.Respond(w, u.Message(false, "User not found", ""))
		return
	}

	resp := u.Message(true, "Data found!", "")
	resp["result"] = userResult[0]
	u.Respond(w, resp)

}
