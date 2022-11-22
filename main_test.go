package main

import (
	"bg-quotes/api"
	"bg-quotes/domain"
	"bg-quotes/storage"
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMainHandler_StatusOK_BodyContent(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"BG Quotes":"main page"}`, w.Body.String())
}

func TestCreateAuthorHandler_StatusOK_BodyContent(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupRouter()

	jpl, _ := json.Marshal(map[string]string{
		"first_name":  "Henry",
		"second_name": "Bukowski",
		"aka":         "Charles Bukowski",
		"img_url":     "https://upload.wikimedia.org/wikipedia/en/e/e2/Charles_Bukowski_smoking.jpg",
	})

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/authors", bytes.NewReader(jpl))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	var umAuthorResp map[string]string
	err = json.Unmarshal([]byte(w.Body.Bytes()), &umAuthorResp)
	if err != nil {
		fmt.Printf("Could not unmarshal json: %s\n", err)
		return
	}

	_, ok := uuid.FromString(umAuthorResp["author_id"])
	assert.Nil(t, ok)
}

func TestCreateQuoteHandler_StatusOK_BodyContent(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupRouter()

	author := domain.CreateAuthor("Henry", "Bukowski", "Charles Bukowski", "https://upload.wikimedia.org/wikipedia/en/e/e2/Charles_Bukowski_smoking.jpg")

	tests := []struct {
		quote        string
		smoking_room bool
		author       string
		respCode     int
		respAuthor   interface{}
	}{
		{"Ако ми се работи, сядам и чекам да ми мине.", true, "", 201, map[string]interface{}{}},
		{"Ако ми се работи, сядам и чекам да ми мине.", false, "", 201, map[string]interface{}{}},
		{"Ако ми се работи, сядам и чекам да ми мине.", true, "123", 400, nil},
		{"Ако ми се работи, сядам и чекам да ми мине.", true, author.AuthorID.String(), 201, map[string]interface{}{}},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("[%v]", tc), func(t *testing.T) {
			jpl, _ := json.Marshal(map[string]interface{}{
				"quote":        tc.quote,
				"smoking_room": tc.smoking_room,
				"author":       tc.author,
			})

			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPost, "/quotes", bytes.NewReader(jpl))
			req.Header.Add("Content-Type", "application/json")
			if err != nil {
				t.Fatal(err)
			}
			r.ServeHTTP(w, req)
			assert.Equal(t, tc.respCode, w.Code)

			var umQuoteResp api.QuoteDTO
			err = json.Unmarshal([]byte(w.Body.Bytes()), &umQuoteResp)
			if err != nil {
				t.Fatalf("Could not unmarshal json: %s\n", err)
				return
			}
			assert.Equal(t, umQuoteResp.Author, tc.respAuthor)
		})
	}
}

func TestShowAuthorHandlerr_StatusOK_BodyContent(t *testing.T) {
	// Creates Author in storage for the current test purposes only.
	da := domain.CreateAuthor("Henry", "Bukowski", "Charles Bukowski", "https://upload.wikimedia.org/wikipedia/en/e/e2/Charles_Bukowski_smoking.jpg")
	sa := storage.AuthorAdd(da)

	gin.SetMode(gin.TestMode)
	r := setupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/authors/"+sa.AuthorID.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var umAuthorResp map[string]string
	err = json.Unmarshal([]byte(w.Body.Bytes()), &umAuthorResp)
	if err != nil {
		fmt.Printf("Could not unmarshal json: %s\n", err)
		return
	}

	_, ok := uuid.FromString(umAuthorResp["author_id"])
	assert.Nil(t, ok)

	assert.Equal(t, "Charles Bukowski", umAuthorResp["aka"])
}
