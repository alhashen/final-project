package handler

import (
	"errors"
	"finalproj/helper"
	"finalproj/model"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/v4"
	"gorm.io/gorm"
)

// CreateComment godoc
//	@Summary		Add comment to a photo
//	@Description	post a json data to comment
//	@Description  Validates: (Message)	
//	@Tags			Comment
//	@Accept			json
//	@Produce		json
//	@Param			id				path	int		true	"Photo ID here"
//	@Param			comment			body	model.SwagComment true	"-"
//	@Param			Authorization	header	string			true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Security		Bearer
//	@Success		201	{object}	model.Comment
//	@Failure		400	{object}	helper.Response
//	@Failure		401	{object}	helper.Response
//	@Failure		403	{object}	helper.Response
//	@Failure		404	{object}	helper.Response
//	@Failure		500	{object}	helper.Response
//	@Router			/comments/photo/{id} [post]
func (h HttpServer) CreateComment(c *gin.Context) {

	userId, err := helper.GetTokenValue(c, "id")
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	photoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

  fmt.Println(photoId)

	comment := model.Comment{}
	err = c.BindJSON(&comment)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.CreateComment(comment, uint(userId.(float64)), uint(photoId))
	if err != nil {

		if errors.As(err, &validation.Errors{}) {
			helper.BadRequest(c, err.Error())
			return
		}

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

	helper.StatusCreated(c, res)
	return

}

// GetComment godoc
//	@Summary		Get comment from specific comment id
//	@Description	Input valid comment id through parameter
//	@Tags			Comment
//	@Accept			json
//	@Produce		json
//	@Param			id				path	int		true	"Comment ID"
//	@Param			Authorization	header	string	true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Security		Bearer
//	@Success		200	{object}	model.Comment
//	@Failure		401	{object}	helper.Response
//	@Failure		404	{object}	helper.Response
//	@Failure		500	{object}	helper.Response
//	@Router			/comments/{id} [get]
func (h HttpServer) GetComment(c *gin.Context) {

	commentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	res, err := h.app.GetComment(uint(commentId))
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

	helper.Ok(c, res)

}

// GetAllComment godoc
//	@Summary	Return all comments from specific photo ID
//	@Description Input valid photo ID in the parameter
//	@Tags		Comment
//	@Produce	json
//	@Param			id				path	int		true	"Photo ID here"
//	@Param		Authorization	header	string	true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Security	Bearer
//	@Success	200	{object}	model.Comment
//	@Failure	401	{object}	helper.Response
//	@Failure	500	{object}	helper.Response
//	@Router		/comments/photo/{id} [get]
func (h HttpServer) GetAllComment(c *gin.Context) {

	photoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	res, err := h.app.GetAllComment(uint(photoId))
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

	helper.Ok(c, res)

}

// UpdateComment godoc
//	@Summary		Update certain comment that is specified by the id
//	@Description	Return forbidden if the user is not the owner of the comment
//	@Description  Validates: (Message)	
//	@Tags			Comment
//	@Accept			json
//	@Produce		json
//	@Param			id				path	int				true	"Comment ID"							
//	@Param			comment			body	model.SwagComment true	"body"						
//	@Param			Authorization	header	string			true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Security		Bearer
//	@Success		200	{object}	model.Comment
//	@Failure		400	{object}	helper.Response
//	@Failure		401	{object}	helper.Response
//	@Failure		403	{object}	helper.Response
//	@Failure		404	{object}	helper.Response
//	@Failure		500	{object}	helper.Response
//	@Router			/comments/{id} [put]
func (h HttpServer) UpdateComment(c *gin.Context) {
	commentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	userId, err := helper.GetTokenValue(c, "id")
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	comment := model.Comment{}
	err = c.BindJSON(&comment)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.UpdateComment(comment, uint(userId.(float64)), uint(commentId))
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

// DeleteComment godoc
//	@Summary		Delete certain comment that is specified by the id
//	@Description	Return forbidden if the user is not the owner of the comment
//	@Description  Validates: (Message)	
//	@Tags			Comment
//	@Produce		json
//	@Param			id				path	int		true	"Comment ID"							
//	@Param			Authorization	header	string	true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Security		Bearer
//	@Success		200	{object}	helper.Response
//	@Failure		401	{object}	helper.Response
//	@Failure		403	{object}	helper.Response
//	@Failure		404	{object}	helper.Response
//	@Failure		500	{object}	helper.Response
//	@Router			/comments/{id} [delete]
func (h HttpServer) DeleteComment(c *gin.Context) {

	commentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	userId, err := helper.GetTokenValue(c, "id")
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	res, err := h.app.DeleteComment(uint(userId.(float64)), uint(commentId))
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
