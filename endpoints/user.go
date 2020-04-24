package endpoints

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	. "github.com/Prabandham/cost-tracker/objects"
)

type FindUserParams struct {
	Email string `form:"email"`
}

type LoginUserParams struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ValidateSessionHeader struct {
	Authorization string `header:"Authorization"`
}

func (e Endpoints) RegisterUser(c *gin.Context) {
	user := User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": FormatErrors(err)})
		return
	}

	e.Connection.FirstOrCreate(&user, user)
	user.Password = ""
	user.PasswordConformation = ""
	c.JSON(http.StatusCreated, user)
}

func (e Endpoints) FindUserByEmail(c *gin.Context) {
	findUser := FindUserParams{}
	user := User{}
	if err := c.ShouldBindQuery(&findUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": FormatErrors(err)})
		return
	}

	e.Connection.Where(&User{Email: findUser.Email}).First(&user)
	if user.Email != "" {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusNotFound, gin.H{})
	}

}

func (e Endpoints) Login(c *gin.Context) {
	loginParams := LoginUserParams{}
	user := User{}
	if err := c.ShouldBindJSON(&loginParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": FormatErrors(err)})
		return
	}

	e.Connection.Where(&User{Email: loginParams.Email}).First(&user)
	if user.Email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
		return
	}

	if loginErr := user.CheckPassword(loginParams.Password); loginErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
		return
	}
	token, _ := user.GenerateJwtToken()
	c.JSON(http.StatusOK, gin.H{
        "authToken": token,
        "userInfo": user,
        "userId": user.ID,
	})
}

func (e Endpoints) ValidateSession(c *gin.Context) {
	headers := ValidateSessionHeader{}
	if err := c.ShouldBindHeader(&headers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Malformed Auth Headers"})
		return
	}
	splitToken := strings.Split(headers.Authorization, "Bearer")
	reqToken := strings.TrimSpace(splitToken[1])
	user := User{}
	if parseErr := user.IsValidSession(reqToken); parseErr != nil {
		c.String(http.StatusUnauthorized, parseErr.Error())
		return
	}
	c.String(http.StatusOK, "valid")
}
