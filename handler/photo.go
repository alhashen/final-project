package handler

import (
	"errors"
	"finalproj/helper"
	"finalproj/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/v4"
	"gorm.io/gorm"
)

// CreatePhoto godoc
//	@Summary		Upload photo
//	@Description	post a json data to upload photo
//	@Description  Validates: (Photo Title, Photo URL)	
//	@Tags			Photo
//	@Accept			json
//	@Produce		json
//	@Param			photo			body	model.SwagPhoto true	"photo"
//	@Param			Authorization	header	string		true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Security		Bearer
//	@Success		201	{object}	model.Photo
//	@Failure		400	{object}	helper.Response
//	@Failure		401	{object}	helper.Response
//	@Failure		500	{object}	helper.Response
//	@Router			/photos [post]
func (h HttpServer) CreatePhoto(c *gin.Context) {

	userId, err := helper.GetTokenValue(c, "id")
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	photo := model.Photo{}
	err = c.BindJSON(&photo)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	photo.UserID = uint(userId.(float64))

	res, err := h.app.CreatePhoto(photo)
	if err != nil {

		if errors.As(err, &validation.Errors{}) {
			helper.BadRequest(c, err.Error())
			return
		}

		helper.InternalServerError(c, err.Error())
		return
	}

	helper.StatusCreated(c, res)
	return

}

// GetPhoto godoc
//	@Summary		Get photo by photo id
//	@Description	Input valid photo id through parameter
//	@Tags			Photo
//	@Accept			json
//	@Produce		json
//	@Param			id				path	int		true	"photo id"
//	@Param			Authorization	header	string	true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Security		Bearer
//	@Success		200	{object}	model.Photo
//	@Failure		401	{object}	helper.Response
//	@Failure		404	{object}	helper.Response
//	@Failure		500	{object}	helper.Response
//	@Router			/photos/{id} [get]
func (h HttpServer) GetPhoto(c *gin.Context) {
	photoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	photo, err := h.app.GetPhoto(uint(photoId))

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.NotFound(c, err.Error())
			break
		default:
			helper.InternalServerError(c, err.Error())
			break
		}
		return
	}

	helper.Ok(c, photo)
}

// GetAllPhoto godoc
//	@Summary	Return all photos from this user
//	@Description Get all photos from specific user ID that was taken from the token
//	@Tags		Photo
//	@Produce	json
//	@Param		Authorization	header	string	true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Security	Bearer
//	@Success	200	{object}	model.Photo
//	@Failure	401	{object}	helper.Response
//	@Failure	500	{object}	helper.Response
//	@Router		/photos [get]
func (h HttpServer) GetAllPhoto(c *gin.Context) {
	userId, err := helper.GetTokenValue(c, "id")
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	res, err := h.app.GetAllPhoto(uint(userId.(float64)))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helper.NoContent(c)
			return
		}
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

// UpdatePhoto godoc
//	@Summary		Update certain photo that is specified by the id
//	@Description	Return forbidden if the user is not the owner of the photo
//	@Description  Validates: (Photo Title, Photo URL)	
//	@Tags			Photo
//	@Accept			json
//	@Produce		json
//	@Param			id				path	int			true	"photo id"
//	@Param			photo			body	model.SwagPhoto true	"body"						
//	@Param			Authorization	header	string		true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Security		Bearer
//	@Success		200	{object}	model.Photo
//	@Failure		400	{object}	helper.Response
//	@Failure		401	{object}	helper.Response
//	@Failure		403	{object}	helper.Response
//	@Failure		404	{object}	helper.Response
//	@Failure		500	{object}	helper.Response
//	@Router			/photos/{id} [put]
func (h HttpServer) UpdatePhoto(c *gin.Context) {

	photoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	userId, err := helper.GetTokenValue(c, "id")
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	photo := model.Photo{}
	err = c.BindJSON(&photo)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.UpdatePhoto(photo, uint(userId.(float64)), uint(photoId))
	if err != nil {
		if errors.As(err, &validation.Errors{}) {
			helper.BadRequest(c, err.Error())
			return
		}

		switch err {
		case helper.ErrForbiddenAccess:
			helper.Forbidden(c)
			break
		case gorm.ErrRecordNotFound:
			helper.NotFound(c, err.Error())
			break
		default:
			helper.InternalServerError(c, err.Error())
			break
		}
		return
	}

	helper.Ok(c, res)

}

// DeletePhoto godoc
//	@Summary		Delete certain photo that is specified by the id
//	@Description	Return forbidden if the user is not the owner of the photo
//	@Description  Validates: (Photo Title, Photo URL)	
//	@Tags			Photo
//	@Produce		json
//	@Param			id				path	int		true	"photo id"
//	@Param			Authorization	header	string	true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Security		Bearer
//	@Success		200	{object}	helper.Response
//	@Failure		401	{object}	helper.Response
//	@Failure		403	{object}	helper.Response
//	@Failure		404	{object}	helper.Response
//	@Failure		500	{object}	helper.Response
//	@Router			/photos/{id} [delete]
func (h HttpServer) DeletePhoto(c *gin.Context) {
	photoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	userId, err := helper.GetTokenValue(c, "id")
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	res, err := h.app.DeletePhoto(uint(userId.(float64)), uint(photoId))
	if err != nil {
		switch err {
		case helper.ErrForbiddenAccess:
			helper.Forbidden(c)
			break
		case gorm.ErrRecordNotFound:
			helper.NotFound(c, err.Error())
			break
		default:
			helper.InternalServerError(c, err.Error())
			break
		}
		return
	}

	helper.OkWithMessage(c, res)
}
