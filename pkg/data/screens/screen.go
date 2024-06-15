package screens

import (
	"encoding/json"
	"log"
	"slices"
)

type ScreenType string

const (
	InitialScreen       string = "initial_screen"
	InputScreen                = "input_screen"
	MenuScreen                 = "menu_screen"
	QuitScreen                 = "quit_screen"
	HttpScreen                 = "http_screen"
	RouterScreen               = "router_screen"
	UpdateSessionScreen        = "update_session_screen"
	FunctionScreen             = "function_screen"
)

type UssdScreenEntry struct {
	InitialScreen string             `yaml:"initial_screen"`
	Screens       map[string]*Screen `yaml:"screens"`
}

type Screen struct {
	Type       string `yaml:"type" validate:"required"`
	NextScreen string `yaml:"next_screen,omitempty" validate:"required"`
	Text       string `yaml:"text,omitempty" validate:"required"`
}

func (a *Screen) String() string {
	jb, _ := json.Marshal(a)
	return string(jb)
}

func ValidateScreenEntry(entry map[string]*Screen) (bool, error) {
	for _, screen := range entry {
		validTypes := []string{InitialScreen}
		if !slices.Contains(validTypes, screen.Type) {
			log.Println("Checking screen type ::::: |", screen.Type)
			//return false, errors.New(string("your entry must contain: " + screen.Type))
		}
	}

	return true, nil
}
