package helper

import (
	"github.com/ryhnfhrza/Golang-To-Do-List-API/model/domain"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/model/web"
)

func ToAuthResponse(user domain.Users)web.AuthResponse{
	return web.AuthResponse{
		Username: user.Username,
	}

}

func ToAuthResponses(users [] domain.Users)[]web.AuthResponse{
	var authResponses []web.AuthResponse
	for _,lf := range users{
		authResponses = append(authResponses, ToAuthResponse(lf))
	}
	return authResponses
}

func ToLoginResponse(user domain.Users)web.AuthResponse{
	return web.AuthResponse{
		Username: user.Username,	
	}

}

