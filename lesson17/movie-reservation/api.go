package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strconv"
)

type Api struct {
	logger *slog.Logger
	router *http.ServeMux
}

func newApi(logger *slog.Logger) *Api {
	mux := http.NewServeMux()

	return &Api{
		logger: logger,
		router: mux,
	}
}

func (a *Api) Start(ctx context.Context) error {
	a.MoviesRouter(ctx)

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return err
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: a.router,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	fmt.Printf("Starting server on :%d\n", port)
	if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

type ModelMovie struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PosterURL   string `json:"poster_url"`
}

func (a *Api) MoviesRouter(ctx context.Context) {
	a.router.HandleFunc("GET /movies", a.FindMovies)
}

func (a *Api) FindMovies(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := a.logger.With("method", "FindMovies")

	resp := struct {
		Results []*ModelMovie `json:"results"`
	}{
		Results: []*ModelMovie{
			{
				Title:       "Lord of the Rings",
				Description: "Lord of the Rings",
				PosterURL:   "https://www.amazon.com/Lord-Rings-Movie-Poster-24x36/dp/B07D96K2QK",
			},
			{
				Title:       "Back to the future",
				Description: "Back to the future",
				PosterURL:   "https://www.amazon.com/Back-Future-Movie-Poster-Regular/dp/B001CDQF8A",
			},
			{
				Title:       "I, Robot",
				Description: "I, Robot",
				PosterURL:   "https://www.cinematerial.com/movies/i-robot-i343818",
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.ErrorContext(ctx, "error happened in JSON write", "error", err)
		w.Write([]byte(err.Error()))
	}
	if _, err := w.Write(jsonResp); err != nil {
		log.ErrorContext(ctx, "error happened in JSON write", "error", err)
		w.Write([]byte(err.Error()))
		return
	}

	log.InfoContext(
		ctx,
		"success find movies",
		"number_of_movies", len(resp.Results),
	)
	return
}
