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

func rowToUser(row pgx.Row) (*model.User, error) {
	var user model.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PGUserRepository) GetUser(ctx context.Context, username string) (*model.User, error) {
	query := "SELECT * FROM Users WHERE Username = $1;"

	rows := r.driver.QueryRow(ctx, query, username)

	return rowToUser(rows)
}

func (r *PGUserRepository) RegisterUser(ctx context.Context, username string, email string, password string) (*model.User, error) {
	query := `
	INSERT INTO Users (Username, Email, Password) VALUES ($1, $2, $3) 
		RETURNING ID, Username, Email, Password;
	`
	rows := r.driver.QueryRow(ctx, query, username, email, password)

	return rowToUser(rows)
}

func (r *PGUserRepository) ChangeUserPassword(ctx context.Context, username string, newPassword string) (*model.User, error) {
	query := `
		UPDATE Users 
		SET Password = $2
		WHERE Username = $1 RETURNING ID, Username, Email, Password;
	`
	row := r.driver.QueryRow(ctx, query, username, newPassword)

	return rowToUser(row)
}
