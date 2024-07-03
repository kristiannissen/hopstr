package domain

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Characteristics struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

type Hop struct {
	Uuid            string            `json:"uuid"`
	Name            string            `json:"name"`
	Description     string            `json:"description"`
	Slug            string            `json:"slug"`
	Characteristics []Characteristics `json:"characteristics"`
}

type Hoplist struct {
	Hoplist []Hop `json:"hoplist"`
}

// repository pattern but for simplicity it is kept in the same
// package as the domain
type HopRepository interface {
	// TODO: add support for pagination
	List() (*Hoplist, error)
	Find(slug string) (*Hop, error)
}

// specific implementation of the repository interface
// for testing purposes we will start out with a mock
type MockHopRepository struct{}

// implement the methods
func (m *MockHopRepository) Find(slug string) (*Hop, error) {
	var err error

	return &Hop{Name: "Kitty"}, err
}

func (m *MockHopRepository) List() (*Hoplist, error) {
	var err error
	var file io.Reader
	hoplist := &Hoplist{}

	// Open file and decode
	abs, _ := filepath.Abs("./../../../hops.json")
	if file, err = os.Open(abs); err != nil {
		log.Fatalf("List %s", err)
	}

	json.NewDecoder(file).Decode(&hoplist)

	return hoplist, err
}
