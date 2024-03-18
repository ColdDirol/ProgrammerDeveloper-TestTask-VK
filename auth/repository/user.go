package repository

import (
	"ProgrammerDeveloper-TestTask-VK/models"
	"ProgrammerDeveloper-TestTask-VK/utils"
	"errors"
)

const (
	userRole  = "user"
	adminRole = "admin"
)

func AddUser(user models.User) error {
	if user.Role != userRole && user.Role != adminRole {
		return errors.New("role is incorrect")
	}
	_, err := utils.DB.Exec(`
		INSERT INTO users (first_name, last_name, email, password, balance, role) 
		VALUES ($1, $2, $3, $4, $5, $6)
	`, user.FirstName, user.LastName, user.Email, user.Password, user.Balance, user.Role)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByEmail(email string) (*models.User, error) {
	query := "SELECT id, first_name, last_name, email, password, balance, role FROM users WHERE email = $1"
	var user models.User
	err := utils.DB.QueryRow(query, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Balance, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteUserByID(id int) error {
	_, err := utils.DB.Exec(`
		DELETE FROM users 
		WHERE id = $1
	`, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAllUsers() error {
	_, err := utils.DB.Exec(`
		DELETE FROM users
	`)
	if err != nil {
		return errors.New("failed to delete actors: " + err.Error())
	}
	return nil
}
