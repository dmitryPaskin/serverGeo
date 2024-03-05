package router

import (
	_ "GeoAPI/docs"
	"GeoAPI/internal/controller"
	"GeoAPI/internal/controller/responder"
	"GeoAPI/internal/service"
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
	chi           *chi.Mux
	handler       controller.Handler
}

func New(patternGeo, patternSearch string) Router {
	var router Router

	router.patternSearch = patternSearch
	router.patternGeo = patternGeo
	router.chi = chi.NewRouter()

	s := service.New(&http.Client{})
	r := responder.New()
	router.handler = controller.New(s, r)

	return router
}

func (r *Router) StartRouter() {
	r.chi.Use(middleware.Recoverer)
	r.chi.Use(middleware.Logger)

	r.chi.Post(r.patternSearch, r.handler.SearchAddressHandler)
	r.chi.Post(r.patternGeo, r.handler.GeocodeHandler)

	r.chi.Get("/swagger/*", httpSwagger.WrapHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r.chi,
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
