package router

import (
	_ "GeoAPI/docs"
	"GeoAPI/internal/controller"
	"GeoAPI/internal/controller/Auth"
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Router struct {
	patternGeo    string
	patternSearch string
	rout          *chi.Mux
}

func NewRouter(patternGeo, patternSearch string) Router {
	return Router{
		patternGeo,
		patternSearch,
		chi.NewRouter(),
	}
}

func (r *Router) StartRouter() {
	r.rout.Use(middleware.Recoverer)
	r.rout.Use(middleware.Logger)

	r.rout.Post("/api/register", controlerAuth.SingUpHandler)
	r.rout.Post("/api/login", controlerAuth.SingInHandler)

	r.rout.With(controlerAuth.TokenMiddleware).Post(r.patternSearch, controller.SearchAddressHandler)
	r.rout.With(controlerAuth.TokenMiddleware).Post(r.patternGeo, controller.GeocodeHandler)

	r.rout.Get("/swagger/*", httpSwagger.WrapHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r.rout,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	sigChan := make(chan os.Signal, 1)
	defer close(sigChan)
	signal.Notify(sigChan, syscall.SIGINT)

	go func() {
		log.Println("Starting server...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	<-sigChan
	stopCTX, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := server.Shutdown(stopCTX); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
}
