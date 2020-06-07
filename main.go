package main

/**
 *  A sample class for RestApi with Golang
 *  @author anjani joshi
 */

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Record struct (Model)
type Record struct {
	ID          string `json:"id"`
	TITLE       string `json:"title"`
	DESCRIPTION string `json:"description"`
	IMGLINK     string `json:"imglink"`
	TECHTYPE    string `json:"techtype"`
	UPVOTE      string `json:"upvote"`
	DOWNVOTE    string `json:"downvote"`
	UPLOADEDBY  string `json:"uploadedby"`
}

// Init records var as a slice record struct
var records []Record

// Get all records
func getRecords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(records)
}

// Get single record
func getRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through records and find one with the id from the params
	for _, item := range records {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Record{})
}

// Add new record
func createRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var record Record
	_ = json.NewDecoder(r.Body).Decode(&record)
	record.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	records = append(records, record)
	json.NewEncoder(w).Encode(record)
}

// Delete record
func deleteRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range records {
		if item.ID == params["id"] {
			records = append(records[:index], records[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(records)
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	records = append(records, Record{ID: "1", TITLE: "Kotlin funcation", DESCRIPTION: "Functions are declared with the fun keyword. For the parameters, you must declare not only their names, but also their types, and you must declare the type of the value the function is intending to return. The body of the function is usually a block, which is enclosed in curly braces",
		IMGLINK: "https://en.wikipedia.org/wiki/Kotlin_(programming_language)#/media/File:Kotlin-logo.svg", TECHTYPE: "kotlin", UPVOTE: "60", DOWNVOTE: "10", UPLOADEDBY: "anjani joshi"})

	records = append(records, Record{ID: "2", TITLE: "Kotlin funcation", DESCRIPTION: "Functions are declared with the fun keyword. For the parameters, you must declare not only their names, but also their types, and you must declare the type of the value the function is intending to return. The body of the function is usually a block, which is enclosed in curly braces",
		IMGLINK: "https://en.wikipedia.org/wiki/Kotlin_(programming_language)#/media/File:Kotlin-logo.svg", TECHTYPE: "kotlin", UPVOTE: "60", DOWNVOTE: "10", UPLOADEDBY: "anjani joshi"})

	records = append(records, Record{ID: "3", TITLE: "Kotlin funcation", DESCRIPTION: "Functions are declared with the fun keyword. For the parameters, you must declare not only their names, but also their types, and you must declare the type of the value the function is intending to return. The body of the function is usually a block, which is enclosed in curly braces",
		IMGLINK: "https://en.wikipedia.org/wiki/Kotlin_(programming_language)#/media/File:Kotlin-logo.svg", TECHTYPE: "kotlin", UPVOTE: "60", DOWNVOTE: "10", UPLOADEDBY: "anjani joshi"})

	records = append(records, Record{ID: "4", TITLE: "Kotlin funcation", DESCRIPTION: "Functions are declared with the fun keyword. For the parameters, you must declare not only their names, but also their types, and you must declare the type of the value the function is intending to return. The body of the function is usually a block, which is enclosed in curly braces",
		IMGLINK: "https://en.wikipedia.org/wiki/Kotlin_(programming_language)#/media/File:Kotlin-logo.svg", TECHTYPE: "kotlin", UPVOTE: "60", DOWNVOTE: "10", UPLOADEDBY: "anjani joshi"})

	// Route handles & endpoints
	r.HandleFunc("api/records", getRecords).Methods("GET")
	r.HandleFunc("api/records/{id}", getRecord).Methods("GET")
	r.HandleFunc("api/records", createRecord).Methods("POST")
	r.HandleFunc("api/records/{id}", deleteRecord).Methods("DELETE")

	//check records into log
	//fmt.Println("records", records)

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
