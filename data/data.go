package data

import (
	"encoding/json"
	"os"
)

type Description struct {
	Short string `json:"short"`
	Long  string `json:"long"`
}

type Button struct {
	Label string `json:"label"`
	Url   string `json:"url"`
}

type Config struct {
	Name        string      `json:"name"`
	Title       string      `json:"title"`
	Description Description `json:"description"`
	Buttons     []Button    `json:"buttons"`
	Link        string      `json:"link"`
}

var Data Config

func Init(Path string) {
	bytes, err := os.ReadFile(Path)

	if err != nil {
		panic(err)
	}

	json.Unmarshal(bytes, &Data)
}
