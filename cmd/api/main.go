package main

import (
	"CatalogItems/internal/products"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"createdAt"`
}

func getProductsHandle(w http.ResponseWriter, r *http.Request, repo products.Repository) {

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	p, err := repo.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{"products": p})
}
func getProductHandle(w http.ResponseWriter, r *http.Request, repo products.Repository) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	idStr := strings.TrimPrefix(r.URL.Path, "/products/")
	if idStr == "" {
		http.NotFound(w, r)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	prod, err := repo.GetProduct(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{"product": prod})
}
func searchProductsHandle(w http.ResponseWriter, r *http.Request, repo products.Repository) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	q := r.URL.Query().Get("q")
	if q == "" {
		http.Error(w, "Query parameter is required", http.StatusBadRequest)
		return
	}
	products, err := repo.Search(r.Context(), q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{"products": products})
}
func main() {
	var err error

	_ = godotenv.Load(".env")
	dsn := os.Getenv("DATABASE_URL")

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	productsRepo := products.NewPostgresRepository(db)
	log.Println("Connected to postgres")

	mux := http.NewServeMux()
	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		getProductsHandle(w, r, productsRepo)
	})
	mux.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		getProductHandle(w, r, productsRepo)
	})
	mux.HandleFunc("/products/search", func(w http.ResponseWriter, r *http.Request) {
		searchProductsHandle(w, r, productsRepo)
	})
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
