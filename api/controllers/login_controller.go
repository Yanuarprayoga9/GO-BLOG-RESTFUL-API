package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Yanuarprayoga9/GO-BLOG-RESTFUL-API/api/auth"
	"github.com/Yanuarprayoga9/GO-BLOG-RESTFUL-API/api/models"
	"github.com/Yanuarprayoga9/GO-BLOG-RESTFUL-API/api/responses"
	"github.com/Yanuarprayoga9/GO-BLOG-RESTFUL-API/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	// Create a response user object without the password
	responseUser := struct {
		ID        uint32 `json:"id"`
		Nickname  string `json:"nickname"`
		Email     string `json:"email"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}{
		ID:        user.ID,
		Nickname:  user.Nickname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	responses.JSON(w, http.StatusOK, map[string]interface{}{"token": token, "user": responseUser})
}

func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}