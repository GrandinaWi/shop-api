package app

import (
	"CatalogItems/internal/db"
	"CatalogItems/internal/products"
	"CatalogItems/internal/routing"
	"context"
	"database/sql"
	"net/http"
)

type App struct {
	server *http.Server
	db     *sql.DB
}

func NewApp() (*App, error) {
	postgres, err := db.NewPostgresDB()
	if err != nil {
		panic(err)
	}

	productsRepo := products.NewPostgresRepository(postgres)

	productsService := products.NewService(productsRepo)

	// Router
	router := routing.NewRouter(productsService)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	return &App{server: server, db: postgres}, nil
}
func (a *App) Run() error {
	return a.server.ListenAndServe()
}
func (a *App) Shutdown(ctx context.Context) error {
	if err := a.server.Shutdown(ctx); err != nil {
		return err
	}
	if err := a.db.Close(); err != nil {
		return err
	}
	return nil

}
