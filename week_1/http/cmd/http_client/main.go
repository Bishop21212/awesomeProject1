/*package main

import (
	"encoding/json"
	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	"log"
	"time"
)

const (
	baseUrl       = "localhost:8081"
	createPostfix = "/notes"
	getPostfix    = "/notes/%d"
)

type NoteInfo struct {
	Title    string `json:"title"`
	Context  string `json:"context"`
	Author   string `json:"author"`
	IsPublic bool   `json:"is_public"`
}

type Note struct {
	ID        int64     `json:"id"`
	Info      NoteInfo  `json:"info"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func createNoteClient() (Note, error) {
	note := NoteInfo{
		Title:    gofakeit.BeerName(),
		Context:  gofakeit.IPv4Address(),
		Author:   gofakeit.Name(),
		IsPublic: gofakeit.Bool(),
	}
	data, err := json.Marshal(note)
	if err != nil {
		return Note{}, err
	}
}

func main() {
	note, err := createNoteClient()
	if err != nil {
		log.Fatal("failed to create note:", err)
	}

	log.Printf(color.RedString("Note created:\n"), color.GreenString("%+v", note))

	note, err = getNoteClient(note.ID)
	if err != nil {
		log.Fatal("failed to get note:", err)

		log.Printf(color.RedString("Note info got:\n"), color.GreenString("%+v", note))
	}
}*/
