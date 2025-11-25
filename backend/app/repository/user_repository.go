package repository

import (
	"context"
	"database/sql"

	"github.com/Kutukobra/FinproKemjar_5/backend/app/model"
)

type UserRepository interface {
	GetUser(ctx context.Context, username string) (*model.User, error)
	RegisterUser(ctx context.Context, username string, email string, password string) (*model.User, error)
	ChangeUserPassword(ctx context.Context, username string, newPassword string) (*model.User, error)
}

type PGUserRepository struct {
	driver *sql.DB
}

func NewPGUserRepository(driver *sql.DB) *PGUserRepository {
	return &PGUserRepository{driver: driver}
}

func rowsToUser(rows *sql.Rows) (*model.User, error) {
	var user model.User
	err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PGUserRepository) GetUser(ctx context.Context, username string) (*model.User, error) {
	query := "SELECT * FROM Users WHERE Username = ?;"
	rows, err := r.driver.QueryContext(ctx, query, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return rowsToUser(rows)
}

func (r *PGUserRepository) RegisterUser(ctx context.Context, username string, email string, password string) (*model.User, error) {
	query := `
	INSERT INTO Users (Username, Email, Password) VALUES (?, ?, ?) 
		RETURNING (ID, Username, Email, Password);
	`
	rows, err := r.driver.QueryContext(ctx, query, username, email, password)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return rowsToUser(rows)
}

func (r *PGUserRepository) ChangeUserPassword(ctx context.Context, username string, newPassword string) (*model.User, error) {
	query := `
		UPDATE Users 
		SET Password = ?
		WHERE Username = ? RETURNING (ID, Username, Email, Password);
	`
	rows, err := r.driver.QueryContext(ctx, query, newPassword, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return rowsToUser(rows)
}
