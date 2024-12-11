package user

import (
	"database/sql"
	"fmt"

	"github.com/Siddhant6674/ECOM/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// get user by there email
func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT*FROM user WHERE email=?", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = sacnRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}

// function for scanning user to find user is exist or not
func sacnRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Store) GetUserByID(ID int) (*types.User, error) {
	return nil, nil
}

func (s *Store) CreateUser(user types.User) error {
	return nil
}
