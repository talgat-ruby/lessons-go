package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/talgat-ruby/lessons-go/projects/movie-reservation/pkg/httputils/response"
)

type Api struct {
	logger *slog.Logger
	router *http.ServeMux
	db     *DB
}

func newApi(logger *slog.Logger, db *DB) *Api {
	mux := http.NewServeMux()

	return &Api{
		logger: logger,
		router: mux,
		db:     db,
	}
}

func (a *Api) Start(ctx context.Context) error {
	a.MoviesRouter(ctx)

	port, err := strconv.Atoi(os.Getenv("API_PORT"))
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

func (a *Api) MoviesRouter(ctx context.Context) {
	a.router.HandleFunc("GET /movies", a.FindMovies)
}

func (a *Api) FindMovies(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := a.logger.With("method", "FindMovies")

	//resp := struct {
	//	Results []*ModelMovie `json:"results"`
	//}{
	//	Results: []*ModelMovie{
	//		{
	//			Title:       "Lord of the Rings",
	//			Description: "Lord of the Rings",
	//			PosterURL:   "https://www.amazon.com/Lord-Rings-Movie-Poster-24x36/dp/B07D96K2QK",
	//		},
	//		{
	//			Title:       "Back to the future",
	//			Description: "Back to the future",
	//			PosterURL:   "https://www.amazon.com/Back-Future-Movie-Poster-Regular/dp/B001CDQF8A",
	//		},
	//		{
	//			Title:       "I, Robot",
	//			Description: "I, Robot",
	//			PosterURL:   "https://www.cinematerial.com/movies/i-robot-i343818",
	//		},
	//	},
	//}

	dbResp, err := a.db.FindMovies(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := response.JSON(
		w,
		http.StatusOK,
		dbResp,
	); err != nil {
		log.ErrorContext(
			ctx,
			"fail json",
			"error", err,
		)
		return
	}

	log.InfoContext(
		ctx,
		"success find movies",
		"number_of_movies", len(dbResp),
	)
	return
}
