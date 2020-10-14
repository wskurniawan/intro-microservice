package handler

import (
	"encoding/json"
	"github.com/wskurniawan/intro-microservice/service-product/config"
	"github.com/wskurniawan/intro-microservice/utils"
	"io/ioutil"
	"net/http"
)

type AuthMiddleware struct {
	AuthService config.AuthService
}

func (auth *AuthMiddleware) ValidateAuth(nextHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := http.NewRequest("POST", auth.AuthService.Host+"/admin-auth", nil)
		if err != nil {
			utils.WrapAPIError(w, r, "failed to create request : "+err.Error(), http.StatusInternalServerError)
			return
		}

		request.Header = r.Header
		authResponse, err := http.DefaultClient.Do(request)
		if err != nil {
			utils.WrapAPIError(w, r, "validate auth failed : "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer authResponse.Body.Close()

		body, err := ioutil.ReadAll(authResponse.Body)
		if err != nil {
			utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		var authResult map[string]interface{}
		err = json.Unmarshal(body, &authResult)

		if authResponse.StatusCode != 200 {
			utils.WrapAPIError(w, r, authResult["error_details"].(string), authResponse.StatusCode)
			return
		}

		nextHandler(w, r)
	}
}
