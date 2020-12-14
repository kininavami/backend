package product

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vmware/vending/external/middleware"
	"net/http"
)

func (p Product) CreateProduct(w http.ResponseWriter, r *http.Request)  {
	decoder := json.NewDecoder(r.Body)
	var product Product
	if err := decoder.Decode(&product); err != nil {
		middleware.RespondError(w, http.StatusBadRequest, err)
		return
	}
	if err := product.Save(); err != nil {
		middleware.RespondError(w, http.StatusBadRequest, err)
		return
	}
	middleware.RespondJSON(w, http.StatusCreated, product)
}

func (p *Product) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	var products Products
	if err := products.GetAllProducts(); err != nil {
		middleware.RespondError(w, http.StatusBadRequest, err)
		return
	}
	middleware.RespondJSON(w, http.StatusOK, products)
}

func (p *Product) GetProductByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username, _ := vars["name"]
	p.Name = username
	if err := p.FetchByName(); err != nil {
		middleware.RespondError(w, http.StatusBadRequest, err)
		return
	}
	middleware.RespondJSON(w, http.StatusOK, p)
}

func (u *Product) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username, _ := vars["name"]
	u.Name = username
	if err := u.DeleteByName(); err != nil {
		middleware.RespondError(w, http.StatusBadRequest, err)
		return
	}
	middleware.RespondJSON(w, http.StatusNoContent, nil)
}