package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/go-basic-struct-note/note"
	"example.com/go-basic-struct-note/todo"
)

type saver interface {
	Save() error
}

type outputInterface interface {
	saver
	Display()
}

func main() {
	noteTitle, noteContent := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(noteTitle, noteContent)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}
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

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		return err
	}
	fmt.Println("Saving data is successful!")
	return nil
}

func outputData(data outputInterface) error {
	data.Display()
	return saveData(data)
}
