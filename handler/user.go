package handler

import (
	"errors"
	"finalproj/helper"
	"finalproj/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/jackc/pgx/v5/pgconn"
)

// RegisterUser godoc
//	@Summary		Register a user
//	@Description	post a json data to register
//	@Description  Validates: (Username, Email, Password, Age[13 >])	
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			User	body		model.SwagUserRegister	true	"user"
//	@Success		201		{object}	model.User
//	@Failure		400		{object}	helper.Response
//	@Failure		404		{object}	helper.Response
//	@Failure		500		{object}	helper.Response
//	@Router			/users/register [post]
func (h HttpServer) RegisterUser(c *gin.Context) {
	user := model.User{}

	err := c.BindJSON(&user)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.RegisterUser(user)
	if err != nil {

		pgerr := &pgconn.PgError{}
		if errors.As(err, &validation.Errors{}) {
			helper.BadRequest(c, err.Error())
			return
		} else if errors.As(err, &pgerr) {
			helper.BadRequest(c, err.Error())
			return
		} else {
			helper.InternalServerError(c, err.Error())
			return
		}
	}

	helper.StatusCreated(c, res)

}

// LogInUser godoc
//	@Summary		Log certain user in
//	@Description	post a json data to login
//	@Description  Validates: (Email, Password)	
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user	body		model.SwagUserLogin true	"Email must be valid, Password longer than 5 characters"
//	@Success		200		{object}	string{status=int{200},token=string}
//	@Failure		400		{object}	helper.Response
//	@Failure		401		{object}	helper.Response
//	@Router			/users/login [post]
func (h HttpServer) LogInUser(c *gin.Context) {
	user := make(map[string]any)

	err := c.BindJSON(&user)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

  err = validation.Validate(user, validation.Map(
    validation.Key("user_email", validation.Required, validation.Length(5, 30), is.EmailFormat),
    validation.Key("user_password", validation.Required, validation.Length(5, 30)),
  ))
  if err != nil {
    helper.BadRequest(c, err.Error())
    return
  }

	err = h.app.LogInUser(user)
	if err != nil {
		helper.Unauthorized(c)
		return
	}

	token := helper.GenerateToken(user["user_id"].(uint), user["user_email"].(string))
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "token": token})
}
