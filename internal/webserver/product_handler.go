package webserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ericgpinto/imersao-fullcycle-go-api/internal/entity"
	"github.com/ericgpinto/imersao-fullcycle-go-api/internal/service"
	"github.com/go-chi/chi/v5"
)

type WebProductHandler struct {
	ProductService *service.ProductService
}

func NewWebProductHandler(productService *service.ProductService) *WebProductHandler{
	return &WebProductHandler{ProductService: productService}
}

func (wph *WebProductHandler) GetProducts(w http.ResponseWriter, r *http.Request){
	products, err := wph.ProductService.GetProducts()
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}


func (wph *WebProductHandler) GetProduct(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "id")
	if id == ""{
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	product, err := wph.ProductService.GetProcuct(id)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)

}

func (wph *WebProductHandler) GetProductsCategoryID(w http.ResponseWriter, r *http.Request){
	categoryID := chi.URLParam(r, "categoryID")
	if categoryID == ""{
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	product, err := wph.ProductService.GetProcuctByCategoryID(categoryID)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)

}

func (wph *WebProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request){
	var product entity.Product
	fmt.Println(product)
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := wph.ProductService.CreateProduct(
		product.Name, product.Description, product.CategoryID, 
		product.ImageURL, product.Price)
	fmt.Println(result)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)


}