package postgres

import (
	"github.com/jmoiron/sqlx"
	"start/internal/model"
	"start/internal/repository"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) repository.AuthorizationRepository {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUserStorageAuth(user model.UserAuth) error {
	sqlStatement := `INSERT INTO t_users (email, username, password_hash) VALUES ($1,$2,$3) `
	_, err := r.db.Exec(sqlStatement, user.Email, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthPostgres) GetUserStorageAuth(username, password, email string) (model.UserAuth, error) {
	var user model.UserAuth

	sqlStatement := `SELECT id FROM t_users WHERE username=$1 AND password_hash=$2`
	err := r.db.Get(&user, sqlStatement, username, password)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *AuthPostgres) CheckUserByUsernameAndPassword(username, password string) (bool, error) {
	sqlStatement := `SELECT COUNT(*) FROM t_users WHERE username=$1 AND password_hash=$2`

	var count int
	err := r.db.QueryRow(sqlStatement, username, password).Scan(&count)

	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
func (r *AuthPostgres) CheckUserByUsername(username string) (bool, error) {
	sqlStatement := `SELECT COUNT(*) FROM t_users WHERE username=$1`

	var count int
	err := r.db.QueryRow(sqlStatement, username).Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil

}

func (r *AuthPostgres) CheckUserByEmail(email string) (bool, error) {
	sqlStatement := `SELECT COUNT(*) FROM t_users WHERE email=$1`

	var count int
	err := r.db.QueryRow(sqlStatement, email).Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
