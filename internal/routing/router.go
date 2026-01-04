package routing

import (
	"CatalogItems/internal/products"
	"fmt"
	"net/http"
)

func NewRouter(productsSvc products.Service) http.Handler {
	mux := http.NewServeMux()
	fmt.Println("Start routing")
	mux.Handle("/products/search", products.SearchProducts(productsSvc))
	mux.Handle("/products/", products.GetProduct(productsSvc))
	mux.Handle("/products", products.GetProducts(productsSvc))

	return mux
}
