package parser

import (
	"fmt"

	"github.com/labstack/echo"
)

type Request struct {
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
	App     string `json:"app"`   // The app key
	Range   string `json:"range"` // The range key
	User    struct {
		Name   string `json:"name"`
		MiBAID uint   `json:"mibaId"`
	}
	Platform struct {
		Key     string `json:"key"`
		Version string `json:"version"`
	}
	Device struct {
		Name   string `json:"name"`
		Brand  string `json:"brand"`
		Screen struct {
			Width  int `json:"width"`
			Height int `json:"height"`
			PPI    int `json:"ppi"`
		}
	}
	Browser struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}
}

func Parse(context echo.Context) (*Request, error) {
	request := new(Request)

	if err := context.Bind(request); err != nil {
		fmt.Println("Error parsing request:", err)

		return request, err
	}

	return request, nil
}
