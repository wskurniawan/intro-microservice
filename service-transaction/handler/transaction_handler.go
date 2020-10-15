package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/context"
	"github.com/wskurniawan/intro-microservice/service-transaction/database"
	"github.com/wskurniawan/intro-microservice/utils"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type Transaction struct {
	Db *gorm.DB
}

func (transaction *Transaction) AddTransaction(w http.ResponseWriter, r *http.Request) {
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

	var transactionReq database.Transaction
	err = json.Unmarshal(body, &transactionReq)
	if err != nil {
		utils.WrapAPIError(w, r, "error unmarshal : "+err.Error(), http.StatusInternalServerError)
		return
	}
	username := context.Get(r, "user")
	transactionReq.Username = fmt.Sprintf("%v",username)
	transactionReq.MenuTotalAmount = transactionReq.MenuQuantity * transactionReq.MenuPrice
	err = transactionReq.AddTransaction(transaction.Db)
	if err != nil {
		utils.WrapAPIError(w, r, "insert menu error : "+err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WrapAPISuccess(w, r, "success", 200)
}

func (transaction *Transaction) GetTransaction (w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	username := context.Get(r, "user")
	res,err := database.GetTransactions(fmt.Sprintf("%v",username), transaction.Db); if err != nil{
		utils.WrapAPIError(w, r, "insert menu error : "+err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WrapAPIData(w,r,res,http.StatusOK,"success")
}