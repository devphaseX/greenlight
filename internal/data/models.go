package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type MovieStore interface {
	Insert(movie *Movie) error
	Get(id int64) (*Movie, error)
	GetAll(title string, genres []string, filters Filters) ([]*Movie, Metadata, error)
	Update(movie *Movie) error
	Delete(id int64) error
}

type Models struct {
	Movies MovieStore
	Users  UserStore
	Tokens TokenStore
}

func NewModels(db *sql.DB) *Models {
	return &Models{
		Movies: MovieModel{
			db,
		},

		Users: UserModel{
			db,
		},

		Tokens: TokenModel{
			db,
		},
	}
}

func NewMockModels() Models {
	return Models{
		Movies: MockMovieModel{},
	}
}
