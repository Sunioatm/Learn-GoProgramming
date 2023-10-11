package handlers

import (
	"encoding/json"
	"fmt"
	"goarch/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type customerHandler struct {
	customerService services.CustomerService
}

func NewCustomerHandler(customerService services.CustomerService) customerHandler {
	return customerHandler{customerService: customerService}
}

func (h customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.customerService.GetCustomers()
	if err != nil {
		handleError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)

}

func (h customerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	// strid := mux.Vars(r)["id"]
	vars := mux.Vars(r)
	strid := vars["id"]
	id, err := strconv.Atoi(strid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	customer, err := h.customerService.GetCustomer(id)
	if err != nil {
		handleError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
