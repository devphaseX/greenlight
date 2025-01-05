package data

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/devphaseX/greenlight.git/internal/validator"
)

type Movie struct {
	ID        int64     `json:"id"` // Unique integer ID for the movie
	CreatedAt time.Time `json:"-"`  // Timestamp for when the movie is added to our database
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"` // Slice of genres for the movie (romance, comedy, etc.)
	Version   int32     `json:"version"`          // The version number starts at 1 and will be incremented each
	// time the movie information is updated
}

func (m Movie) MarshalJSON() ([]byte, error) {
	var runtime string

	if m.Runtime > 0 {
		runtime = fmt.Sprintf("%d mins", m.Runtime)
	}

	aux := struct {
		ID      int64    `json:"id"` // Unique integer ID for the movie
		Title   string   `json:"title"`
		Year    int32    `json:"year,omitempty"`
		Runtime string   `json:"runtime,omitempty"`
		Genres  []string `json:"genres,omitempty"` // Slice of genres for the movie (romance, comedy, etc.)
		Version int32    `json:"version"`          // The version number starts at 1 and will be incremented each
		// time the movie information is updated
	}{
		ID:      m.ID,
		Title:   m.Title,
		Year:    m.Year,
		Runtime: runtime,
		Genres:  m.Genres,
		Version: m.Version,
	}

	return json.Marshal(aux)
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
v.Check(movie.Title != "", "title", "must be provided")
v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")
v.Check(movie.Year != 0, "year", "must be provided")
v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")
v.Check(movie.Runtime != 0, "runtime", "must be provided")
v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")
v.Check(movie.Genres != nil, "genres", "must be provided")
v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
}
