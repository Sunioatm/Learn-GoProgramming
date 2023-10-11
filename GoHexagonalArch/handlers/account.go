package handlers

import (
	"encoding/json"
	"goarch/errs"
	"goarch/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type accountHandler struct {
	accountService services.AccountService
}

func NewAccountHandler(accountService services.AccountService) accountHandler {
	return accountHandler{accountService: accountService}
}

func (h accountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if r.Header.Get("Content-Type") != "application/json" {
		handleError(w, errs.NewValidationError("request body incorrect format"))
		return
	}

	req := services.NewAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handleError(w, errs.NewValidationError("request body incorrect format"))
		return
	}

	res, err := h.accountService.NewAccount(id, req)
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h accountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	response, err := h.accountService.GetAccounts(id)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
