package dto

import "chatbox-app/models"

type LoginUserOutput struct {
	UserId       uint   `json:"userId"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Avatar       string `json:"avatar"`
	Intro        string `json:"intro"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func NewLoginUserOutput(user *models.User, aToken, rToken string) *LoginUserOutput {
	return &LoginUserOutput{
		UserId:       user.ID,
		Username:     user.Username,
		Email:        user.Email,
		Avatar:       user.Avatar,
		Intro:        user.Intro,
		AccessToken:  aToken,
		RefreshToken: rToken,
	}
}
