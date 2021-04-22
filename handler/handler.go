package handler

import (
	"encoding/json"
	"go_bookshelf/database"
	"go_bookshelf/models"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	// Explicitly end reading request body before exiting function
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	// Declare a new book models and assigned it using value from the request body
	var book models.Books
	_ = json.NewDecoder(r.Body).Decode(&book)

	// Check wether the requested body is valid, function available at helper.go
	if validateBody(book, w) {
		// Generate a unique random id
		book.Id = rand.Intn(1000000)
		book.Finished = book.ReadPage == book.PageCount

		// Appending the newly created book, to Books database
		database.Books = append(database.Books, book)

		// Helper function to help me write a response in JSON, available at helper.go
		respondWithJSON(w, 201, map[string]string{
			"status":  "success",
			"message": "Buku berhasil ditambahkan",
			"bookId":  strconv.Itoa(book.Id),
		})
	}
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// Declaring struct, available at struct.go
	var shortResult []shortDetail

	// Getting the URL Query and assigned them to their respected variables
	query := r.URL.Query()
	name := query.Get("name")
	reading := query.Get("reading")
	finished := query.Get("finished")

	for _, book := range database.Books {
		// Logic that will make sure we will not append book that did not contain query name
		if name != "" && !strings.Contains(strings.ToLower(book.Name), strings.ToLower(name)) {
			continue
		}

		// Logic that will make sure we will not append book that have unmatched code
		if (reading == "1" && !book.Reading) || (reading == "0" && book.Reading) {
			continue
		}

		if (finished == "1" && !book.Finished) || (finished == "0" && book.Finished) {
			continue
		}

		shortResult = append(shortResult, shortDetail{
			Name:   book.Name,
			Author: book.Author,
			BookId: book.Id,
		})
	}
	respondWithJSON(w, 200, &shortResult)
}

func GetBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// Get the bookId value from the URL parameter
	bookId := mux.Vars(r)["bookId"]

	// Check wether bookId exists in database, this will return -1 if the bookId is not exists
	index := checkIdExists(bookId)
	if index != -1 {
		respondWithJSON(w, 200, &database.Books[index])
		return
	}

	// Template helper function that write 404 response
	respondWith404(w)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	// Get the bookId value from the URL parameter
	bookId := mux.Vars(r)["bookId"]

	// Check wether bookId exists in database, this will return -1 if the bookId is not exists
	index := checkIdExists(bookId)
	if index != -1 {
		book := database.Books[index]
		_ = json.NewDecoder(r.Body).Decode(&book)

		if validateBody(book, w) {
			database.Books[index] = book

			respondWithJSON(w, 200, response{
				Status:  "success",
				Message: "Buku berhasil diubah",
			})
			return
		}
	}
	respondWith404(w)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// Get the bookId value from the URL parameter
	bookId := mux.Vars(r)["bookId"]

	// Check wether bookId exists in database, this will return -1 if the bookId is not exists
	index := checkIdExists(bookId)
	if index != -1 {
		// Swapping the deleted element with the last element of this slice
		database.Books[index] = database.Books[len(database.Books)-1]

		// Overwriting the database with the same value but without the last element
		database.Books = database.Books[:len(database.Books)-1]
		respondWithJSON(w, 200, response{
			Status:  "success",
			Message: "Buku berhasil dihapus",
		})
		return
	}
	respondWith404(w)
}
