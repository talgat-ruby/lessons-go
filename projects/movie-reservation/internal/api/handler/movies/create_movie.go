package movies

import (
	"net/http"

	"github.com/talgat-ruby/lessons-go/projects/movie-reservation/internal/db/movie"
	"github.com/talgat-ruby/lessons-go/projects/movie-reservation/pkg/httputils/request"
	"github.com/talgat-ruby/lessons-go/projects/movie-reservation/pkg/httputils/response"
)

type CreateMovieRequest struct {
	Data *movie.ModelMovie `json:"data"`
}

type CreateMovieResponse struct {
	Data *movie.ModelMovie `json:"data"`
}

func (h *Movies) CreateMovie(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "CreateMovie")

	// request parse
	requestBody := &CreateMovieRequest{}

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
	dbResp, err := h.db.CreateMovie(ctx, requestBody.Data)

	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to query from db",
			"error", err,
		)
		http.Error(w, "failed to query from db", http.StatusInternalServerError)
		return
	}

	if dbResp == nil {
		log.ErrorContext(
			ctx,
			"row is empty",
		)
		http.Error(w, "row is empty", http.StatusInternalServerError)
		return
	}

	// response
	resp := CreateMovieResponse{
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
		"success insert movie",
		"movie id", resp.Data.ID,
	)
	return
}
