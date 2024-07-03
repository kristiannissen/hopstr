package domain

import (
	"encoding/json"
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

// utility function to load json data
func loaddata() *Hoplist {
	hoplist := &Hoplist{}
	abs, _ := filepath.Abs("./../../../hops.json")
	if file, err := os.Open(abs); err != nil {
		log.Fatal(err)
	} else {
		json.NewDecoder(file).Decode(&hoplist)
	}

	return hoplist
}

// implement the methods
func (m *MockHopRepository) Find(slug string) (*Hop, error) {
	var err error
	hoplist := loaddata()
	hop := &Hop{}

	for _, h := range hoplist.Hoplist {
		if h.Slug == slug {
			hop.Uuid = h.Uuid
			hop.Name = h.Name
			hop.Slug = h.Slug
			hop.Characteristics = h.Characteristics
			break
		}
	}

	return hop, err
}

func (m *MockHopRepository) List() (*Hoplist, error) {
	var err error
	hoplist := loaddata()

	return hoplist, err
}
