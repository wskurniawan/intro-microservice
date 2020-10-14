package handler

import (
	"github.com/wskurniawan/intro-microservice/utils"
	"net/http"
)

// AddMenuHandler handle add menu
func AddMenuHandler(w http.ResponseWriter, r *http.Request) {
	utils.WrapAPISuccess(w, r, "success", 200)
}
