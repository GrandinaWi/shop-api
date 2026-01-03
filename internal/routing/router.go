package http

import (
	"CatalogItems/internal/products"
	"net/http"
)

func NewRouter(productsSvc products.Service) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/products/search", func(w http.ResponseWriter, r *http.Request) {
		searchProductsHandle(w, r, productsSvc)
	})
	mux.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		getProductHandle(w, r, productsSvc)
	})
	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		getProductsHandle(w, r, productsSvc)
	})

	return mux
}
