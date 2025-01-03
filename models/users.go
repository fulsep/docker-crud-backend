package models

import (
	"context"

	"github.com/fulsep/docker-crud-backend/tree/main/lib"
	"github.com/jackc/pgx/v5"
)

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Users []User

func FindAllUsers() (Users, error) {
	db, err := lib.DB()

	if err != nil {
		return Users{}, err
	}

	defer db.Close(context.Background())

	query := `
		SELECT 
		id, email, password
		FROM users
	`

	rows, err := db.Query(context.Background(), query)

	if err != nil {
		return Users{}, err
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[User])

	if err != nil {
		return Users{}, err
	}

	return users, nil

}

func FindOneUser(id int) (User, error) {
	db, err := lib.DB()

	if err != nil {
		return User{}, err
	}

	defer db.Close(context.Background())

	query := `
		SELECT 
		id, email, password
		FROM users
		WHERE id = $1
	`

	row := db.QueryRow(context.Background(), query, id)

	var user User

	err = row.Scan(
		&user.Id,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		return User{}, err
	}

	return user, nil

}

func InsertUser(data User) (User, error) {
	db, err := lib.DB()

	if err != nil {
		return User{}, err
	}

	defer db.Close(context.Background())

	query := `
		INSERT INTO users 
		(email, password)
		VALUES
		($1,$2)
		RETURNING
		id, email, password
		
	`

	row := db.QueryRow(context.Background(), query, data.Email, data.Password)

	var user User

	err = row.Scan(
		&user.Id,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func UpdateUser(data User) (User, error) {
	db, err := lib.DB()

	if err != nil {
		return User{}, err
	}

	defer db.Close(context.Background())

	query := `
		UPDATE users
		SET
		email = $1,
		password = $2
		WHERE
		id = $3
		RETURNING
		id, email, password
		
	`

	row := db.QueryRow(context.Background(), query, data.Email, data.Password, data.Id)

	var user User

	err = row.Scan(
		&user.Id,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func DeleteUser(id int) (User, error) {
	db, err := lib.DB()

	if err != nil {
		return User{}, err
	}

	defer db.Close(context.Background())

	query := `
		DELETE FROM users WHERE id = $1
		RETURNING
		id, email, password
	`

	row := db.QueryRow(context.Background(), query, id)

	var user User

	err = row.Scan(
		&user.Id,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		return User{}, err
	}

	return user, nil

}
