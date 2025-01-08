package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Movies      MovieStore
	Users       UserStore
	Tokens      TokenStore
	Permissions PermissionStore
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

		Permissions: PermissionModel{
			db,
		},
	}
}

func NewMockModels() Models {
	return Models{
		Movies: MockMovieModel{},
	}
}
