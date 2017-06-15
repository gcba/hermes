package parser

import (
	"fmt"

	"github.com/labstack/echo"
)

type (
	user struct {
		Name   string `json:"name"`
		MiBAID uint   `json:"mibaId"`
	}

	platform struct {
		Key     string `json:"key"`
		Version string `json:"version"`
	}

	device struct {
		Name   string `json:"name"`
		Brand  string `json:"brand"`
		Screen screen `json:"screen"`
	}

	screen struct {
		Width  int `json:"width"`
		Height int `json:"height"`
		PPI    int `json:"ppi"`
	}

	browser struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}

	Request struct {
		Rating   int      `json:"rating"`
		Comment  string   `json:"comment"`
		App      string   `json:"app"`   // The app key
		Range    string   `json:"range"` // The range key
		User     user     `json:"user"`
		Platform platform `json:"platform"`
		Device   device   `json:"device"`
		Browser  browser  `json:"browser"`
	}
)

func Parse(context echo.Context) (*Request, error) {
	request := new(Request)

	if err := context.Bind(request); err != nil {
		fmt.Println("Error parsing request:", err)

		return request, err
	}

	return request, nil
}
