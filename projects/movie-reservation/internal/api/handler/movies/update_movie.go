package movies

import (
	"net/http"
	"strconv"

	"github.com/talgat-ruby/lessons-go/projects/movie-reservation/internal/db/movie"
	"github.com/talgat-ruby/lessons-go/projects/movie-reservation/pkg/httputils/request"
	"github.com/talgat-ruby/lessons-go/projects/movie-reservation/pkg/httputils/response"
)

type UpdateMovieRequest struct {
	Data *movie.ModelMovie `json:"data"`
}

func (h *Movies) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "UpdateMovie")

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

	// request parse
	requestBody := &UpdateMovieRequest{}

	if err := request.JSON(w, r, requestBody); err != nil {
		log.ErrorContext(
			ctx,
			"failed to parse request body",
			"error", err,
		)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	// db request
	if err := h.db.UpdateMovie(ctx, int64(id), requestBody.Data); err != nil {
		log.ErrorContext(
			ctx,
			"failed to query from db",
			"error", err,
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := response.JSON(
		w,
		http.StatusNoContent,
		nil,
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
		"success update movie",
		"id", id,
	)
	return
}
