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
	FirstName  string `json:"first_name" form:"first_name" binding:"required,min=1,max=32,excludesall=\\\"'<>!@#$%^&*()_+=:;?/0x2C"` // 0x2C = comma (,)
	SecondName string `json:"second_name" form:"second_name" binding:"required,min=1,max=32,excludesall=\\\"'<>!@#$%^&*()_+=:;?/0x2C"`
	AKA        string `json:"aka" form:"aka" binding:"required,min=1,max=64"`
	ImageURL   string `json:"img_url" form:"img_url" binding:"required,url"`
}

type CreateQuoteArgs struct {
	Quote       string `json:"quote" form:"quote" binding:"required,min=8,max=256,excludesall=\\\"'<>!@#$%^&*()_+=?/"`
	SmokingRoom bool   `json:"smoking_room" form:"smoking_room" validate:"required"`
	AuthorID    string `json:"author" form:"author_id" binding:"omitempty,uuid"`
	// AuthorID    uuid.UUID `json:"author_id" form:"author_id" binding:"omitempty,uuid"`
}

func CreateAuthorHandler(c *gin.Context) {
	var args CreateAuthorArgs

	if err := c.ShouldBind(&args); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	da := domain.CreateAuthor(args.FirstName, args.SecondName, args.AKA, args.ImageURL)
	sa := storage.AuthorCreate(da)
	dto := createAuthorDTO(sa)
	c.JSON(http.StatusCreated, dto)
}

func ShowAuthorHandler(c *gin.Context) {
	aid, err := uuid.Parse(c.Param("author_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	if lia, found := storage.AuthorRead(aid); found {
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

	author, err := GetAuthorDTO(args.AuthorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	}

	dq := domain.CreateQuote(args.Quote, args.SmokingRoom, args.AuthorID)
	sq := storage.QuoteCreate(dq)
	dto := createQuoteDTO(sq, author)
	c.JSON(http.StatusCreated, dto)
}

func ShowQuoteHandler(c *gin.Context) {
	qid, err := uuid.Parse(c.Param("quote_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	if lia, found := storage.QuoteRead(qid); found {
		author, err := GetAuthorDTO(lia.AuthorID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		}

		dto := createQuoteDTO(lia, author)
		c.JSON(http.StatusOK, dto)
		return
	}

	c.JSON(http.StatusNotFound, make(map[string]string))
}
