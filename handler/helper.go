package handler

import (
	"encoding/json"
	"go_bookshelf/database"
	"go_bookshelf/models"
	"net/http"
	"strconv"
)

func validateBody(book models.Books, w http.ResponseWriter) bool {
	if book.Name == "" {
		respondWithJSON(w, 400, response{
			Status:  "fail",
			Message: "Gagal menambahkan buku. Mohon isi nama buku",
		})
		return false
	}

	if book.ReadPage > book.PageCount {
		respondWithJSON(w, 400, response{
			Status:  "fail",
			Message: "Gagal menambahkan buku. readPage tidak boleh lebih besar dari pageCount",
		})
		return false
	}

	return true
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func checkIdExists(bookId string) int {
	Id, _ := strconv.Atoi(bookId)
	for index, book := range database.Books {
		if book.Id == Id {
			return index
		}
	}
	return -1
}

func respondWith404(w http.ResponseWriter) {
	respondWithJSON(w, 404, response{
		Status:  "fail",
		Message: "Buku dengan Id itu tidak ditemukan.",
	})
}
