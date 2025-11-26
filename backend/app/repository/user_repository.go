package repository

import (
	"context"

	"github.com/Kutukobra/FinproKemjar_5/backend/app/model"
	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	GetUser(ctx context.Context, username string) (*model.User, error)
	RegisterUser(ctx context.Context, username string, email string, password string) (*model.User, error)
	ChangeUserPassword(ctx context.Context, username string, newPassword string) (*model.User, error)
}

type PGUserRepository struct {
	driver *pgx.Conn
}

func NewPGUserRepository(driver *pgx.Conn) *PGUserRepository {
	return &PGUserRepository{driver: driver}
}

func rowsToUser(rows pgx.Rows) (*model.User, error) {
	var user model.User
	rows.Next()
	err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	rows.Close()
	return &user, nil
}

func (r *PGUserRepository) GetUser(ctx context.Context, username string) (*model.User, error) {
	query := "SELECT * FROM Users WHERE Username = $1;"

	rows, err := r.driver.Query(ctx, query, username)
	if err != nil {
		return nil, err
	}

	return rowsToUser(rows)
}

func (r *PGUserRepository) RegisterUser(ctx context.Context, username string, email string, password string) (*model.User, error) {
	query := `
	INSERT INTO Users (Username, Email, Password) VALUES ($1, $2, $3) 
		RETURNING ID, Username, Email, Password;
	`
	rows, err := r.driver.Query(ctx, query, username, email, password)
	if err != nil {
		return nil, err
	}

	return rowsToUser(rows)
}

func (r *PGUserRepository) ChangeUserPassword(ctx context.Context, username string, newPassword string) (*model.User, error) {
	query := `
		UPDATE Users 
		SET Password = $2
		WHERE Username = $1 RETURNING ID, Username, Email, Password;
	`
	rows, err := r.driver.Query(ctx, query, username, newPassword)
	if err != nil {
		return nil, err
	}

	return rowsToUser(rows)
}
