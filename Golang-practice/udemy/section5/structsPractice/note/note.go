package note

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
	// convert content to Json file
	"encoding/json"
)

type Note struct {
	Title     string ` json:"title"` // This is how you can assign metadata or make it lowercase or what you want
	Content   string
	CreatedAt time.Time
}

// This is a method function
func (n Note) Display() {
	fmt.Printf("Your note titled %v has the following content: %v \n\n", n.Title, n.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	// Only able to access struct that is global
	json, err := json.Marshal(note)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// This will return nil is everything going right
	// This will return an error if something goes wrong
	return os.WriteFile(fileName, json, 0644)
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid Input")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
