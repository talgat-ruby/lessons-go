package movies

import (
	"net/http"

	"github.com/talgat-ruby/lessons-go/projects/movie-reservation/internal/db/movie"
	"github.com/talgat-ruby/lessons-go/projects/movie-reservation/pkg/httputils/response"
)

type FindMoviesResponse struct {
	Data []movie.ModelMovie `json:"data"`
}

func (h *Movies) FindMovies(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := h.logger.With("method", "FindMovies")

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

	dbResp, err := h.db.FindMovies(ctx)

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
