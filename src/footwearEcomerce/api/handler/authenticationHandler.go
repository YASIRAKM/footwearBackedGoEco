package handler

import (
	"GoProjects/src/footweearEcomerce/src/footwearEcomerce/db"
	"GoProjects/src/footweearEcomerce/src/footwearEcomerce/model"
	"GoProjects/src/footweearEcomerce/src/footwearEcomerce/utils"
	"database/sql"
	"fmt"
	"net/http"

	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var secretKey = []byte("your-secret-key")

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user := model.User{}

	stmt, err := db.DB.Prepare("SELECT id, username, password, role, email, name FROM User WHERE username = ? AND password = ?")
	if err != nil {
		return utils.Response(c, http.StatusOK, false, "Invalid username or pasword", err.Error())
	}
	defer stmt.Close()

	err = stmt.QueryRow(username, password).Scan(&user.ID, &user.Username, &user.Pasword, &user.Role, &user.Email, &user.Name)
	if err == sql.ErrNoRows {
		return utils.Response(c, http.StatusOK, false, "Invalid username or pasword", err.Error())
	} else if err != nil {
		return utils.Response(c, http.StatusBadGateway, false, "Invalid username or pasword", err.Error())
	}

	token, err := GenerateJwtToken(user)
	if err != nil {
		return utils.Response(c, http.StatusNonAuthoritativeInfo, true, "Unexpected Error", "")

	}

	return c.JSON(http.StatusOK, map[string]interface{}{"status": true, "message": "Login Successful", "data": user, "token": token})

}

func GenerateJwtToken(user model.User) (string, error) {
	claims := model.JwtCustomClaims{
		UserID: user.ID,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := rawToken.SignedString(secretKey)
	if err != nil {
		return "", err

	}

	return token, nil

}

func GetUserByUsername(c echo.Context) error {
	username := c.FormValue("username")

	user := model.User{}
	fmt.Println("Looking for user:", username)

	// Prepare the query to fetch the user by username
	stmt, err := db.DB.Prepare("SELECT id, username, password, role, email, name FROM User WHERE username = ?")
	if err != nil {
		return utils.Response(c, http.StatusOK, false, "Unexpected Error", err.Error())
	}
	defer stmt.Close()

	// Execute the query and scan the result into the user struct
	err = stmt.QueryRow(username).Scan(&user.ID, &user.Username, &user.Pasword, &user.Role, &user.Email, &user.Name)
	if err == sql.ErrNoRows {
		// Return response if no user is found
		return utils.Response(c, http.StatusOK, false, "User not found", "")
	} else if err != nil {
		// Handle other errors
		return utils.Response(c, http.StatusBadGateway, false, "Database Error", err.Error())
	}

	// Return a success response with the user data
	return c.JSON(http.StatusOK, map[string]interface{}{"status": true, "message": "User found", "data": user})
}
