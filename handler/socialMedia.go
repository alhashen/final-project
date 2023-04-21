package handler

import (
	"errors"
	"finalproj/helper"
	"finalproj/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

// CreateSocialMedia godoc
//	@Summary		Add social media to a user	
//	@Description	Specified user was taken from the token. Post a json data to add social media. 
//	@Description  Validates: (Name, Social Media URL)	
//	@Description  Error: No duplicates	
//	@Tags			Social Media	
//	@Accept			json
//	@Produce		json
//	@Param			socialMedia		body	model.SwagSocialMedia true	"-"							
//	@Param			Authorization	header	string				true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Security		Bearer
//	@Success		201	{object}	model.SocialMedia
//	@Failure		400	{object}	helper.Response
//	@Failure		401	{object}	helper.Response
//	@Failure		404	{object}	helper.Response
//	@Failure		500	{object}	helper.Response
//	@Router			/social_medias [post]
func (h HttpServer) CreateSocialMedia(c *gin.Context) {

	userId, err := helper.GetTokenValue(c, "id")
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	social := model.SocialMedia{}
	err = c.BindJSON(&social)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.CreateSocialMedia(social, uint(userId.(float64)))
	if err != nil {

		pgerr := &pgconn.PgError{}
		if errors.As(err, &validation.Errors{}) {
			helper.BadRequest(c, err.Error())
			return
		} else if errors.As(err, &pgerr) {
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

// GetSocialMedia godoc
//	@Summary		Get social media by social media id
//	@Description	Input valid social media id into parameter
//	@Tags			Social Media
//	@Accept			json
//	@Produce		json
//	@Param			id				path	int		true	"Social Media ID"
//	@Param			Authorization	header	string	true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Security		Bearer
//	@Success		200	{object}	model.SocialMedia
//	@Failure		401	{object}	helper.Response
//	@Failure		404	{object}	helper.Response
//	@Failure		500	{object}	helper.Response
//	@Router			/social_medias/{id} [get]
func (h HttpServer) GetSocialMedia(c *gin.Context) {

	socialMediaId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.GetSocialMedia(uint(socialMediaId))
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

// GetAllSocialMedia godoc
//	@Summary	Return all social medias
//	@Description
//	@Tags		Social Media
//	@Produce	json
//	@Param		Authorization	header	string	true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Security	Bearer
//	@Success	200	{object}	model.SocialMedia
//	@Failure	401	{object}	helper.Response
//	@Failure	500	{object}	helper.Response
//	@Router		/social_medias [get]
func (h HttpServer) GetAllSocialMedia(c *gin.Context) {

	res, err := h.app.GetAllSocialMedia()
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

// UpdateSocialMedia godoc
//	@Summary		Update certain social media that is specified by the id
//	@Description	Return forbidden if the user is not the owner of the social media
//	@Description  Validates: (Name, Social Media URL)	
//	@Tags			Social Media
//	@Accept			json
//	@Produce		json
//	@Param			id				path	int					true	"Social Media ID"
//	@Param			socialMedia		body	model.SwagSocialMedia	true	"body"						
//	@Param			Authorization	header	string				true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Security		Bearer
//	@Success		200	{object}	model.SocialMedia
//	@Failure		400	{object}	helper.Response
//	@Failure		401	{object}	helper.Response
//	@Failure		403	{object}	helper.Response
//	@Failure		404	{object}	helper.Response
//	@Failure		500	{object}	helper.Response
//	@Router			/social_medias/{id} [put]
func (h HttpServer) UpdateSocialMedia(c *gin.Context) {

	userId, err := helper.GetTokenValue(c, "id")
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	socialId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	social := model.SocialMedia{}
	err = c.BindJSON(&social)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.UpdateSocialMedia(social, uint(userId.(float64)), uint(socialId))
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

// DeleteSocialMedia godoc
//	@Summary		Delete certain social media that is specified by the id
//	@Description	Return forbidden if the user is not the owner of the social media
//	@Description  Validates: (Name, Social Media URL)	
//	@Tags			Social Media
//	@Produce		json
//	@Param			id				path	int		true	"-"							
//	@Param			Authorization	header	string	true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Security		Bearer
//	@Success		200	{object}	helper.Response
//	@Failure		401	{object}	helper.Response
//	@Failure		403	{object}	helper.Response
//	@Failure		404	{object}	helper.Response
//	@Failure		500	{object}	helper.Response
//	@Router			/social_medias/{id} [delete]
func (h HttpServer) DeleteSocialMedia(c *gin.Context) {

	userId, err := helper.GetTokenValue(c, "id")
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	socialId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	res, err := h.app.DeleteSocialMedia(uint(userId.(float64)), uint(socialId))
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
