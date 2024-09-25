package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/ErichBerger/erichberger.com/ui/templates"
)

func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

	server := http.Server{
		Handler: Routes(),
		Addr:    ":8080",
	}

	log.Info("Starting server...", slog.String("Addr", ":8080"))

	err := server.ListenAndServe()

	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

}

func Routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))

	mux.Handle("GET /favicon.ico", http.FileServer(http.Dir("./ui/static/assets")))

	mux.Handle("/{$}", IndexHandler())

	return mux
}

func IndexHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templates.HXRender(w, r, templates.Home(), templates.Base(templates.Home(), "App"))
	})
}
