package movies

import (
	"net/http"
	"strconv"

	"github.com/talgat-ruby/lessons-go/projects/movie-reservation/internal/db/movie"
	"github.com/talgat-ruby/lessons-go/projects/movie-reservation/pkg/httputils/response"
)

type FindMoviesResponse struct {
	Data []movie.ModelMovie `json:"data"`
}

func (h *Movies) FindMovies(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := h.logger.With("method", "FindMovies")

	query := r.URL.Query()
	offset, err := strconv.Atoi(query.Get("offset"))
	if err != nil {
		log.ErrorContext(
			ctx,
			"fail parse query offset",
			"error", err,
		)
		http.Error(w, "invalid query offset", http.StatusBadRequest)
		return
	}
	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil {
		log.ErrorContext(
			ctx,
			"fail parse query limit",
			"error", err,
		)
		http.Error(w, "invalid query limit", http.StatusBadRequest)
		return
	}

	dbResp, err := h.db.FindMovies(ctx, offset, limit)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp := FindMoviesResponse{
		Data: dbResp,
	}

	if err := response.JSON(
		w,
		http.StatusOK,
		resp,
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
		"number_of_movies", len(resp.Data),
	)
	return
}
