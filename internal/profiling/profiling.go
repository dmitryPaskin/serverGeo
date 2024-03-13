package profiling

import (
	"github.com/go-chi/chi"
	"net/http/pprof"
)

func NewProfRouter() *chi.Mux {
	router := chi.NewRouter()

	router.HandleFunc("/allocs", pprof.Handler("allocs").ServeHTTP)
	router.HandleFunc("/block", pprof.Handler("block").ServeHTTP)
	router.HandleFunc("/cmdline", pprof.Cmdline)
	router.HandleFunc("/goroutine", pprof.Handler("goroutine").ServeHTTP)
	router.HandleFunc("/heap", pprof.Handler("heap").ServeHTTP)
	router.HandleFunc("/mutex", pprof.Handler("mutex").ServeHTTP)
	router.HandleFunc("/profile", pprof.Profile)
	router.HandleFunc("/threadcreate", pprof.Handler("threadcreate").ServeHTTP)
	router.HandleFunc("/trace", pprof.Trace)
	router.HandleFunc("/goroutine_full", pprof.Handler("goroutine").ServeHTTP)

	return router
}
