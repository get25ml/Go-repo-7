package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const notesFile = "notes.txt"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: notes [command] [args]")
		fmt.Println("Commands: add, view, edit, delete")
		fmt.Println("add [note] - adds a new note to the notes file")
		fmt.Println("view - displays all notes in the notes file")
		fmt.Println("edit [old note] [new note] - replaces an existing note with a new note")
		fmt.Println("delete [note] - removes a note from the notes file")
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "add":
		if len(args) < 1 {
			fmt.Println("Usage: notes add [note]")
			os.Exit(1)
		}
		addNote(strings.Join(args, " "))
	case "view":
		viewNotes()
	case "edit":
		if len(args) < 2 {
			fmt.Println("Usage: notes edit [old note] [new note]")
			os.Exit(1)
		}
		oldNote := strings.Join(args[0:1], " ")
		newNote := strings.Join(args[1:], " ")
		if !noteExists(oldNote) {
			fmt.Println("Note not found: ", oldNote)
			os.Exit(1)
		}
		if newNote == "" {
			fmt.Println("New note text not provided.")
			os.Exit(1)
		}
		editNote(oldNote, newNote)
	case "delete":
		if len(args) < 1 {
			fmt.Println("Usage: notes delete [note]")
			os.Exit(1)
		}
		deleteNote(strings.Join(args, " "))
	default:
		fmt.Println("Unknown command: ", command)
		os.Exit(1)
	}
}

func addNote(note string) {
	// Check if the note already exists
	if noteExists(note) {
		fmt.Println("Note already exists: ", note)
		return
	}

	// Open the file in append mode to add the new note
	file, err := os.OpenFile(notesFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening notes file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Write the new note to the file
	_, err = file.WriteString(note + "\n")
	if err != nil {
		fmt.Println("Error writing to notes file:", err)
		os.Exit(1)
	}

	fmt.Println("Note added successfully.")
}

func viewNotes() {
	file, err := os.Open(notesFile)
	if err != nil {
		fmt.Println("Error opening notes file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading notes file:", err)
		os.Exit(1)
	}
}

func editNote(oldNote, newNote string) {
	// Check if the old note exists before trying to edit it
	if !noteExists(oldNote) {
		fmt.Println("Note not found: ", oldNote)
		return
	}

	// Read the entire file into memory
	file, err := os.Open(notesFile)
	if err != nil {
		fmt.Println("Error opening notes file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == oldNote {
			line = newNote // Replace the old note with the new note
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading notes file:", err)
		os.Exit(1)
	}

	// Close the file before deleting it
	file.Close()

	// Delete the file
	err = os.Remove(notesFile)
	if err != nil {
		fmt.Println("Error deleting notes file:", err)
		os.Exit(1)
	}

	// Recreate the file and write the updated lines
	file, err = os.Create(notesFile)
	if err != nil {
		fmt.Println("Error creating notes file:", err)
		os.Exit(1)
	}
	defer file.Close()

	for _, line := range lines {
		_, err = file.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to notes file:", err)
			os.Exit(1)
		}
	}

	fmt.Println("Note edited successfully.")
}

func deleteNote(note string) {
	// Check if the note exists before trying to delete it
	if !noteExists(note) {
		fmt.Println("Note not found: ", note)
		return
	}

	// Read the entire file into memory
	file, err := os.Open(notesFile)
	if err != nil {
		fmt.Println("Error opening notes file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line != note {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading notes file:", err)
		os.Exit(1)
	}

	// Close the file before deleting it
	file.Close()

	// Delete the file
	err = os.Remove(notesFile)
	if err != nil {
		fmt.Println("Error deleting notes file:", err)
		os.Exit(1)
	}

	// Recreate the file and write the updated lines
	file, err = os.Create(notesFile)
	if err != nil {
		fmt.Println("Error creating notes file:", err)
		os.Exit(1)
	}
	defer file.Close()

	for _, line := range lines {
		_, err = file.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to notes file:", err)
			os.Exit(1)
		}
	}

	fmt.Println("Note deleted successfully.")
}

func noteExists(note string) bool {
	file, err := os.Open(notesFile)
	if err != nil {
		fmt.Println("Error opening notes file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == note {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading notes file:", err)
		os.Exit(1)
	}

	return false
}
