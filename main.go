package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Note struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

var notes []Note
const filename = "notes.json"

func main() {
	loadNotes()


	if len(os.Args) < 2 {
		fmt.Println("Usage: notes-cli [add|list|delete] [options]")
		return
	}

	switch os.Args[1] {
		case "add":
			if len(os.Args) < 4 {
				fmt.Println("Usage: notes-cli add <title> <content>")
				return
			}
			addNote(os.Args[2], os.Args[3])

		case "list":
			listNotes()

		case "delete":
			if len(os.Args) < 3 {
				fmt.Println("Usage: notes-cli delete <title>")
				return
			}
			deleteNote(os.Args[2])

		default:
			fmt.Println("Command not recognized.")
	}
}

func loadNotes() {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			notes = []Note{}
			return
		}
		panic(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&notes)
	if err != nil {
		fmt.Println("Error loading notes:", err)
	}
}

func saveNotes() {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(notes)
	if err != nil {
		fmt.Println("Error saving notes:", err)
	}
}

func addNote(title, content string) {
	notes = append(notes, Note{Title: title, Content: content})
	saveNotes()
	fmt.Println("Note added successfully!")
}

func listNotes() {
	if len(notes) == 0 {
		fmt.Println("No notes found.")
		return
	}
	for i, note := range notes {
		fmt.Printf("[%d] %s: %s\n", i+1, note.Title, note.Content)
	}
}

func deleteNote(title string) {
	for i, note := range notes {
		if note.Title == title {
			notes = append(notes[:i], notes[i+1:]...)
			saveNotes()
			fmt.Println("Note removed:", title)
			return
		}
	}
	fmt.Println("Not not found.")
}
