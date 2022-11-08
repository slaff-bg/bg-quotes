package api

import (
	"bg-quotes/domain"
	"bg-quotes/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateAuthorArgs struct {
	FirstName  string `json:"first_name" form:"first_name" binding:"required"`
	SecondName string `json:"second_name" form:"second_name" binding:"required"`
	AKA        string `json:"aka" form:"aka" binding:"required"`
	ImageURL   string `json:"img_url" form:"img_url" binding:"required"`
}

func CreateAuthorHandler(c *gin.Context) {
	var args CreateAuthorArgs

	if err := c.ShouldBind(&args); err != nil {
		// @TODO - I have to add a validator. The best way probably is to use middleware for that.
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	author := domain.CreateAuthor(args.FirstName, args.SecondName, args.AKA, args.ImageURL)
	lia := storage.AuthorAdd(author)
	dto := createAuthorDTO(lia)
	c.JSON(http.StatusOK, dto)
}

func ShowAuthorHandler(c *gin.Context) {
	aid, err := uuid.Parse(c.Param("author_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
	lia, _ := storage.AuthorGet(aid)
	dto := createAuthorDTO(lia)
	c.JSON(http.StatusOK, dto)
}
