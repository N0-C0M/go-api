package store

import "github.com/N0-C0M/go-api/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	// Implement the logic to create a user in the database
	// For example, you can use r.store.db to execute SQL queries
	if err := r.store.db.QueryRow(
		"INSERT INTO users (email, encripted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncriptedPassword,
	).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	// Implement the logic to find a user by email in the database
	// For example, you can use r.store.db to execute SQL queries
	return nil, nil
}
