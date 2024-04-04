package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/go-basic-struct-note/note"
)

func main() {
	noteTitle, noteContent := getNoteData()

	userNote, err := note.New(noteTitle, noteContent)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()

	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving the note is failed!")
	}
	fmt.Println("Saving the note is successful!")
}

func getUserInput(promptText string) string {
	fmt.Print(promptText)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}
