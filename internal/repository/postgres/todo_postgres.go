package postgres

import (
	"github.com/jmoiron/sqlx"
	"start/internal/model"
)

type TodoPostgres struct {
	db *sqlx.DB
}

func NewTodoPostgres(db *sqlx.DB) *TodoPostgres {
	return &TodoPostgres{db: db}
}

func (r *TodoPostgres) SetUserStorage(user *model.UserTodo) error {
	sqlStatement := `INSERT INTO t_users_todo(age, name) VALUES ($1,$2)`
	_, err := r.db.Exec(sqlStatement, user.Age, user.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoPostgres) GetUserByIDStorage(id int) (*model.UserTodo, error) {
	sqlStatement := `SELECT id,age,name FROM t_users_todo WHERE id=$1`

	var user model.UserTodo

	err := r.db.QueryRow(sqlStatement, id).Scan(&user.Id, &user.Age, &user.Name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *TodoPostgres) CheckUserByIDStorage(id int) (bool, error) {
	sqlStatement := `SELECT COUNT(*) FROM t_users_todo WHERE id=$1`

	var count int
	err := r.db.QueryRow(sqlStatement, id).Scan(&count)

	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (r *TodoPostgres) DeleteUserByIdStorage(id int) error {
	sqlStatement := ` DELETE FROM t_users_todo WHERE id=$1`

	_, err := r.db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoPostgres) GetAllUserIDStorage() ([]int, error) {
	sqlStatement := `SELECT id FROM t_users_todo`

	rows, err := r.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	var ids []int
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ids, nil
}
