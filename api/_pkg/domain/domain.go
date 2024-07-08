package domain

import (
	"encoding/json"
	"log"
	"os"

	"github.com/rpdg/vercel_blob"
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
	Characteristics []Characteristics `json:"characteristics,omitempty"`
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

// Vercel specific repository
type VercelHopRepository struct {
	url string
}

func NewVercelHopRepository() *VercelHopRepository {
	return &VercelHopRepository{
		url: "https://ajorqiotomntgglo.public.blob.vercel-storage.com/hops/hops.json",
	}
}

func NewMockHopRepository() *MockHopRepository {
	return &MockHopRepository{
		url: "./../../../hops.json",
	}
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

func (v *VercelHopRepository) Download() ([]byte, error) {
	var bytes []byte
	var err error
	client := vercel_blob.NewVercelBlobClient()

	if bytes, err = client.Download(v.url, vercel_blob.DownloadCommandOptions{}); err != nil {
		log.Fatal(err)
	}

	return bytes, err
}

func (m *MockHopRepository) Find(slug string) (*Hop, error) {
	var err error

	hop := &Hop{}
	hoplist := &Hoplist{}

	var bytes []byte
	if bytes, err = m.Download(); err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(bytes, &hoplist); err != nil {
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

func (v *VercelHopRepository) Find(slug string) (*Hop, error) {
	var err error

	hop := &Hop{}
	hoplist := &Hoplist{}

	var bytes []byte
	if bytes, err = v.Download(); err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(bytes, &hoplist); err != nil {
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
	var bytes []byte
	hoplist := &Hoplist{}

	if bytes, err = m.Download(); err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(bytes, &hoplist); err != nil {
		log.Fatal(err)
	}

	return hoplist, err
}

func (v *VercelHopRepository) List() (*Hoplist, error) {
	var err error
	var bytes []byte
	hoplist := &Hoplist{}

	if bytes, err = v.Download(); err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(bytes, &hoplist); err != nil {
		log.Fatal(err)
	}

	return hoplist, err
}
