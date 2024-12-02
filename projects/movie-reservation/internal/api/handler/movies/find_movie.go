package movies

import (
	"net/http"
	"strconv"

	"github.com/talgat-ruby/lessons-go/projects/movie-reservation/internal/db/movie"
	"github.com/talgat-ruby/lessons-go/projects/movie-reservation/pkg/httputils/response"
)

type FindMovieResponse struct {
	Data *movie.ModelMovie `json:"data"`
}

func (h *Movies) FindMovie(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "FindMovie")

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to convert id to int",
			"error", err,
		)
		http.Error(w, "failed to convert id to int", http.StatusBadRequest)
		return
	}

	dbResp, err := h.db.FindMovie(ctx, int64(id))

	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to query from db",
			"error", err,
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := FindMovieResponse{
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
		"success find movie",
		"movie id", resp.Data.ID,
	)
	return
}
