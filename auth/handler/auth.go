package handler

import (
	"encoding/json"
	"github.com/wskurniawan/intro-microservice/auth/database"
	"github.com/wskurniawan/intro-microservice/auth/utils"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type Auth struct {
	Db *gorm.DB
}

func (db *Auth) ValidateAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	authToken := r.Header.Get("Authorization")

	res, err := database.Validate(authToken,db.Db); if err != nil{
		utils.WrapAPIError(w, r, err.Error(), http.StatusForbidden)
		return
	}

	utils.WrapAPIData(w,r,database.Auth{
		Username: res.Username,
		Token:    res.Token,
	},http.StatusOK,"success")
	return
}

func (db *Auth) SignUp (w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		utils.WrapAPIError(w, r, "can't read body", http.StatusBadRequest)
		return
	}

	var signup database.Auth

	err = json.Unmarshal(body, &signup)
	if err != nil {
		utils.WrapAPIError(w, r, "error unmarshal : "+err.Error(), http.StatusInternalServerError)
		return
	}
	signup.Token = utils.IdGenerator()

	err = signup.SignUp(db.Db); if err != nil{
		utils.WrapAPIError(w,r,err.Error(),http.StatusBadRequest)
		return
	}

	utils.WrapAPISuccess(w,r,"success",http.StatusOK)
}

func (db *Auth) Login (w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}


	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		utils.WrapAPIError(w, r, "can't read body", http.StatusBadRequest)
		return
	}

	var login database.Auth

	err = json.Unmarshal(body, &login)
	if err != nil {
		utils.WrapAPIError(w, r, "error unmarshal : "+err.Error(), http.StatusInternalServerError)
		return
	}

	res,err := login.Login(db.Db);if err != nil {
		utils.WrapAPIError(w, r, "error unmarshal : "+err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WrapAPIData(w,r,database.Auth{
		Username: res.Username,
		Token:    res.Token,
	},http.StatusOK,"success")
	return
}