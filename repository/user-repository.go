package repository

import (
	"cinemy-auth-api/db"
	"cinemy-auth-api/models"
	"cinemy-auth-api/utils"
)


func UserRegister(email string, pass string) (bool, error)  {
	conn := db.Connect()
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	hashedPassword, _ := utils.GenerateFromPassword(pass)

	if _, err:=utils.IsValidEmail(email); err != nil {
		return false, err
	}

	if _ , err:= conn.DB.Exec(query, email, hashedPassword); err != nil {
		return false, err
	}

	defer conn.DB.Close()
	return true, nil
}

func UserLogin(email string, pass string)(models.User, error){
		var hashedPassword string
		var user models.User
		
		conn := db.Connect()
		query := "SELECT id, email, password FROM users WHERE email = ?"
		
		if err:= conn.DB.QueryRow(query, email).Scan(&user.Id, &user.Email, &hashedPassword); err != nil {
			return models.User {}, err
		}

		if err:=utils.CompareHashAndPassword(pass, hashedPassword); err != nil {
			return models.User {}, err
		}

		return user, nil
}