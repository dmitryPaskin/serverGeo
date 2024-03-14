package router

import (
	_ "GeoAPI/docs"
	"GeoAPI/internal/controller"
	"GeoAPI/internal/controller/responder"
	"GeoAPI/internal/metrics"
	"GeoAPI/internal/profiling"
	"GeoAPI/internal/service"
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)

	metrics.New()

	r.chi.Use(middleware.Recoverer)
	r.chi.Use(middleware.Logger)

	r.chi.Get("/swagger/*", httpSwagger.WrapHandler)
	r.chi.Post("/", controller.Login)

	r.chi.Handle("/metrics", promhttp.Handler())

	r.chi.Group(func(rout chi.Router) {
		rout.Use(jwtauth.Verifier(tokenAuth))
		rout.Use(jwtauth.Authenticator)

		rout.Mount("/mycustompath/pprof", profiling.NewProfRouter())

		r.chi.Post(r.patternSearch, r.handler.SearchAddressHandler)
		r.chi.Post(r.patternGeo, r.handler.GeocodeHandler)
	})

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
