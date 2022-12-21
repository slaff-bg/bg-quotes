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

	var tests []map[string]string
	tests = append(tests, map[string]string{
		"first_name":  "Йордан",
		"second_name": "Радичков",
		"aka":         "Йордан Радичков",
		"img_url":     "https://upload.wikimedia.org/wikipedia/en/thumb/1/1a/Yordan_Radichkov.jpg/200px-Yordan_Radichkov.jpg",
	})
	tests = append(tests, map[string]string{
		"first_name":  "Yordan",
		"second_name": "Radichkov",
		"aka":         "Yordan Radichkov",
		"img_url":     "https://upload.wikimedia.org/wikipedia/en/thumb/1/1a/Yordan_Radichkov.jpg/200px-Yordan_Radichkov.jpg",
	})

	for _, tc := range tests {
		t.Run(fmt.Sprintf("[%v]", tc), func(t *testing.T) {
			jpl, _ := json.Marshal(tc)

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

			assert.Equal(t, umAuthorResp["first_name"], tc["first_name"])
		})
	}
}

func TestCreateQuoteHandler_StatusOK_BodyContent(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupRouter()

	author := domain.CreateAuthor("Йордан", "Радичков", "Йордан Радичков", "https://upload.wikimedia.org/wikipedia/en/thumb/1/1a/Yordan_Radichkov.jpg/200px-Yordan_Radichkov.jpg")

	tests := []struct {
		quote        string
		smoking_room bool
		author       string
		respCode     int
		respAuthor   interface{}
	}{
		{"If I feel like working, I sit down and wait for it to pass.", true, "", 201, map[string]interface{}{}},
		{"If I feel like working, I sit down and wait for it to pass.", false, "", 201, map[string]interface{}{}},
		{"Ако ми се работи, сядам и чекам да ми мине.", true, "", 201, map[string]interface{}{}},
		{"Ако ми се работи, сядам и чекам да ми мине.", true, "123", 400, nil},
		{"Човек е дълго изречение, написано с много любов и вдъхновение, ала пълно с правописни грешки.", false, author.AuthorID.String(), 201, map[string]interface{}{}},
		{"A person is a long sentence written with much love and inspiration but full of spelling mistakes.", false, author.AuthorID.String(), 201, map[string]interface{}{}},
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

func TestShowAuthorHandler_StatusOK_BodyContent(t *testing.T) {
	// Creates Author in storage for the current test purposes only.
	da := domain.CreateAuthor("Йордан", "Радичков", "Йордан Радичков", "https://upload.wikimedia.org/wikipedia/en/thumb/1/1a/Yordan_Radichkov.jpg/200px-Yordan_Radichkov.jpg")
	sa := storage.AuthorCreate(da)

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

	assert.Equal(t, umAuthorResp["aka"], da.AKA)
}

/*
func TestShowQuoteHandler_StatusOK_BodyContent(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupRouter()

	author := domain.CreateAuthor("Йордан", "Радичков", "Йордан Радичков", "https://upload.wikimedia.org/wikipedia/en/thumb/1/1a/Yordan_Radichkov.jpg/200px-Yordan_Radichkov.jpg")

	tests := []struct {
		quote        string
		smoking_room bool
		author       string
		respCode     int
		respQuote    interface{}
	}{
		{"If I feel like working, I sit down and wait for it to pass.", true, "", 201, map[string]interface{}{}},
		{"If I feel like working, I sit down and wait for it to pass.", false, "", 201, map[string]interface{}{}},
		{"Ако ми се работи, сядам и чекам да ми мине.", true, "", 201, map[string]interface{}{}},
		{"Ако ми се работи, сядам и чекам да ми мине.", true, "123", 400, nil},
		{"Човек е дълго изречение, написано с много любов и вдъхновение, ала пълно с правописни грешки.", false, author.AuthorID.String(), 201, map[string]interface{}{}},
		{"A person is a long sentence written with much love and inspiration but full of spelling mistakes.", false, author.AuthorID.String(), 201, map[string]interface{}{}},
	}

	// Seeds predefined data.
	for _, tc := range tests {
		tc.respQuote = domain.CreateQuote(tc.quote, tc.smoking_room, tc.author)
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("[%v]", tc), func(t *testing.T) {

			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/quotes/", nil)
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
*/
