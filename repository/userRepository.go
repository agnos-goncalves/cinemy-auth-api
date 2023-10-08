package userRepository

import (
	"cinemy-auth-api/db"
	"cinemy-auth-api/model"
)


func InsertUser(email string, pass string)(bool, error){
	conn := db.Connect()
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	_, err := conn.DB.Exec(query, email, pass)
	defer conn.DB.Close()

	if err != nil {
		return false, err;
	}
	return true, err;
}

func SelectByEmail(email string)(model.UserFull, error){
	conn := db.Connect()
	query := "SELECT id, email, active, password FROM users WHERE email = ?"
	var user model.UserFull;
	err := conn.DB.QueryRow(query, email).Scan(&user.Id, &user.Email, &user.Active, &user.Password)
	defer conn.DB.Close()

	if err != nil {
		return model.UserFull{}, err;
	}
	return user, err;
}


func UpdateActive(email string) (bool, error) {
	conn := db.Connect()
	query := "UPDATE users SET active = 1 WHERE email = ?"
	_, err := conn.DB.Exec(query, email)
	defer conn.DB.Close()

	if err != nil {
		return false, err;
	}
	return true, err;
}

func UpdatePassword(email string, pass string) (bool, error) {
	conn := db.Connect()
	query := "UPDATE users SET password = ? WHERE email = ?"
	_, err := conn.DB.Exec(query, pass, email)
	defer conn.DB.Close()

	if err != nil {
		return false, err;
	}
	return true, err;
}

