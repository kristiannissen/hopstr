package domain

import (
	"encoding/json"
	"log"
	"os"
)

type Characteristics struct {
	Name        string `json:"name"`
	Value       string `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
}

type Hop struct {
	Uuid            string            `json:"uuid"`
	Name            string            `json:"name"`
	Description     string            `json:"description,omitempty"`
	Slug            string            `json:"slug"`
	Characteristics []Characteristics `json:"characteristics"`
}

type Hoplist struct {
	Hoplist []Hop `json:"hoplist"`
}

// repository pattern but for simplicity it is kept in the same
// package as the domain
type HopRepository interface {
	List() (*Hoplist, error)
	Find(slug string) (*Hop, error)
	Download(url string) ([]byte, error)
}

// specific implementation of the repository interface
// for testing purposes we will start out with a mock
type MockHopRepository struct {
	url string
}

// implement the methods
func (m *MockHopRepository) Download() ([]byte, error) {
	var bytes []byte
	var err error

	if bytes, err = os.ReadFile(m.url); err != nil {
		log.Fatal(err)
	}

	return bytes, err
}

func (m *MockHopRepository) Find(slug string) (*Hop, error) {
	var err error

	hop := &Hop{}
	hoplist := &Hoplist{}

	if err = json.Unmarshal(m.Download(), &hoplist); err != nil {
		log.Fatal(err)
	}

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
	hoplist := &Hoplist{}

	if err = json.Unmarshal(m.Download(), &hoplist); err != nil {
		log.Fatal(err)
	}

	return hoplist, err
}
