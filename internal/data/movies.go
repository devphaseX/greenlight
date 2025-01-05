package data

import (
	"encoding/json"
	"fmt"
	"time"
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
