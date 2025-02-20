package seeds

import (
	"database/sql"
	"time"
)

type movie struct {
	id          int
	title       string
	description string
	posterUrl   string
	createdAt   *time.Time
	updatedAt   *time.Time
}

var movies []*movie

func (s *seeder) movie(tx *sql.Tx) error {
	movies = []*movie{
		{
			title:       "Lord of the Rings",
			description: "Lord of the Rings",
			posterUrl:   "https://www.amazon.com/Lord-Rings-Movie-Poster-24x36/dp/B07D96K2QK",
		},
		{
			title:       "Back to the future",
			description: "Back to the future",
			posterUrl:   "https://www.amazon.com/Back-Future-Movie-Poster-Regular/dp/B001CDQF8A",
		},
		{
			title:       "I, Robot",
			description: "I, Robot",
			posterUrl:   "https://www.cinematerial.com/movies/i-robot-i343818",
		},
	}

	sqlQuery := `
		INSERT INTO movie (title, description, posterurl)
		VALUES ($1, $2, $3)
		RETURNING id;
	`
	sqlStmt, err := tx.Prepare(sqlQuery)
	if err != nil {
		return err
	}

	s.cleanups = append(
		s.cleanups, func() error {
			return sqlStmt.Close()
		},
	)

	for _, g := range movies {
		err := sqlStmt.
			QueryRow(g.title, g.description, g.posterUrl).
			Scan(&g.id)
		if err != nil {
			return err
		}
	}

	return nil
}
