package userService

import (
	"cinemy-auth-api/model"
	userRepository "cinemy-auth-api/repository"
	"cinemy-auth-api/utils"

	"github.com/dgrijalva/jwt-go"
)

func Register(email string, pass string) (bool, error)  {
	hashedPassword, _ := utils.GenerateFromPassword(pass)
	_ , err:= userRepository.InsertUser(email, hashedPassword)

	if err != nil {
		return false, err
	}

	return true, nil
}

func Login(email string, pass string)(string, error){
	var user model.UserFull
	user, err := userRepository.SelectByEmail(email)
	
	if err != nil {
		return "", err
	}

	if err:=utils.CompareHashAndPassword(pass, user.Password); err != nil {
		return "", err
	}

	payload := jwt.MapClaims{
		"sub": user.Id,
		"email": user.Email,
		"active": user.Active,
	}

	token, err:=utils.GenerateJWT(payload)
	if err != nil {
		return "", err
	}

	return token, nil
}

func Active(token string)(bool, error) {

	claims, err := utils.DecodeJWT(token)

	if err != nil {
		return false, err
	}
	
	username := claims["email"].(string)
	if _, err:= userRepository.UpdateActive(username); err != nil {
		return false, err
	}
	
	return true, nil
}

func PasswordChange(email string, currentPass string, newPass string)(bool, error) {
	user, err:= userRepository.SelectByEmail(email)
	if err != nil {
		return false, err
	}
	if err := utils.CompareHashAndPassword(currentPass, user.Password); err != nil {
		return false, err
	}
	if _, err := userRepository.UpdatePassword(user.Email, newPass); err != nil {
		return false, err
	}
	return true, nil
}