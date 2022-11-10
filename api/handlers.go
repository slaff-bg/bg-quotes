package api

import (
	"bg-quotes/domain"
	"bg-quotes/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// This struct will be extracted and used by the validator later.
type CreateAuthorArgs struct {
	FirstName  string `json:"first_name" form:"first_name" binding:"required"`
	SecondName string `json:"second_name" form:"second_name" binding:"required"`
	AKA        string `json:"aka" form:"aka" binding:"required"`
	ImageURL   string `json:"img_url" form:"img_url" binding:"required"`
}

type CreateQuoteArgs struct {
	Quote       string `json:"quote" form:"quote" binding:"required,min=8,max=256"`
	SmokingRoom bool   `json:"smoking_room" form:"smoking_room"`
	AuthorID    string `json:"author_id" form:"author_id" binding:"omitempty,uuid"`
	// AuthorID    uuid.UUID `json:"author_id" form:"author_id" binding:"omitempty,uuid"`
}

func CreateAuthorHandler(c *gin.Context) {
	// @TODO - I have to add a validator. The best way probably is to use middleware for that.

	var args CreateAuthorArgs

	if err := c.ShouldBind(&args); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	da := domain.CreateAuthor(args.FirstName, args.SecondName, args.AKA, args.ImageURL)
	sa := storage.AuthorAdd(da)
	dto := createAuthorDTO(sa)
	c.JSON(http.StatusOK, dto)
}

func ShowAuthorHandler(c *gin.Context) {
	aid, err := uuid.Parse(c.Param("author_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	if lia, found := storage.AuthorGet(aid); found {
		dto := createAuthorDTO(lia)
		c.JSON(http.StatusOK, dto)
		return
	}

	c.JSON(http.StatusNotFound, make(map[string]string))
}

func CreateQuoteHandler(c *gin.Context) {
	var args CreateQuoteArgs
	if err := c.ShouldBind(&args); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	var author interface{}
	if len(args.AuthorID) > 0 {
		aid, err := uuid.Parse(args.AuthorID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
			return
		}

		if a, found := storage.AuthorGet(aid); found {
			dto := createAuthorDTO(a)
			author = dto
		} else {
			author = make(map[string]string)
		}
	}

	dq := domain.CreateQuote(args.Quote, args.SmokingRoom, args.AuthorID)
	sq := storage.QuoteAdd(dq)
	dto := createQuoteDTO(sq, author)
	c.JSON(http.StatusOK, dto)
}
