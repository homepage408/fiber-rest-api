package repository

import (
	"context"
	"database/sql"
	"errors"

	"fiber-rest-api/helper"
	"fiber-rest-api/model/domain"
)

type UserRespository interface {
	Login(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, user domain.User)
	FindByEmail(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
}

type ClientUserRepository struct {
}

func NewUserRepository() UserRespository {
	return &ClientUserRepository{}
}

func (repository *ClientUserRepository) Login(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	query := "select email, password, deleted_at from users where email = ? and password = ?"
	rows, err := tx.QueryContext(ctx, query, user.Email, user.Password)
	helper.PanicIfError(err)
	defer rows.Close()

	var deletedAt *sql.NullString
	userDb := domain.User{}

	if rows.Next() {
		err := rows.Scan(&userDb.Email, &userDb.Password, deletedAt)
		helper.PanicIfError(err)

		return userDb, nil
	} else {
		return userDb, errors.New("user not found")
	}
}

func (repository *ClientUserRepository) Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	query := "insert into users(username, email, password) values (?, ?, ?)"
	resutl, err := tx.ExecContext(ctx, query, user.Username, user.Email, user.Password)
	helper.PanicIfError(err)

	id, err := resutl.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)
	return user
}

func (repository *ClientUserRepository) FindByEmail(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	query := "select email, password, deleted_at from users where email = ?"
	rows, err := tx.QueryContext(ctx, query, user.Email)
	helper.PanicIfError(err)
	defer rows.Close()

	var deletedAt *sql.NullString
	userDB := domain.User{}

	if rows.Next() {
		err := rows.Scan(&userDB.Email, &userDB.Password, deletedAt)
		helper.PanicIfError(err)

		return userDB, nil
	} else {
		return userDB, errors.New("user not found")
	}
}

func (repository *ClientUserRepository) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	query := "delete from users where email = ?"
	_, err := tx.ExecContext(ctx, query, user.Email)
	helper.PanicIfError(err)

}
