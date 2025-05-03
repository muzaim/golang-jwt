package controllers

import (
	"golang-jwt/helpers"
	"golang-jwt/models"
	"net/http"
)

func Me(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("userinfo").(*helpers.MyCustomcClaims)

	userResponse := &models.MyProfile{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	helpers.Response(w, 200, "My Profile", userResponse)
}
