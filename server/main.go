package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Card struct {
	SideA string `json: "sideA"`
	SideB string `json: "sideB"`
}

func getPhrases(w http.ResponseWriter, r *http.Request) {

	cards := []Card{}

	database, err := sql.Open("sqlite3", "../data/flash_card_data")
	if err != nil {
		log.Fatal(err)
	}
	// TODO: get random rows in a more efficient way
	rows, err := database.Query("select jpn_text, jpn_transcription || '-' || eng_text from jpn_eng_phrases order by random() limit 10")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		c := Card{}
		err = rows.Scan(&c.SideA, &c.SideB)
		if err != nil {
			log.Fatal(err)
		}
		cards = append(cards, c)
	}

	fmt.Println(cards)

	json.NewEncoder(w).Encode(cards)
	database.Close()
}

func main() {

	router := mux.NewRouter()
	// Return 10 random phrase pairs
	router.HandleFunc("/api/1.0/phrases", getPhrases).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
