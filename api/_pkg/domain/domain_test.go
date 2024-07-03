package domain

import "testing"

func TestFind(t *testing.T) {
	repo := NewMockHopRepository()
	want := "Columbia"

	if got, err := repo.Find("columbia"); err != nil {
		t.Fatal(err)
	} else {
		if got.Name != want {
			t.Errorf("want %s got %s", want, got.Name)
		}
	}
}

func TestList(t *testing.T) {
	hoplist := NewMockHopRepository()
	want := 0

	if got, err := hoplist.List(); err != nil {
		t.Fatal(err)
	} else {
		if len(got.Hoplist) == want {
			t.Errorf("want %d got %d", want, len(got.Hoplist))
		}
	}
}
