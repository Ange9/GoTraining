package main

import "testing"

func TestGetUser(t *testing.T) {
	md := &MockDatastore{
		Users: map[int]User{
			2: {First: "Ange", ID: 2},
		},
	}
	s := &Service{
		db: md,
	}
	want := "Ange"
	got, err := s.GetUser(2)
	if err != nil {
		t.Errorf("got error %v", err)
	}
	if got.First != want {
		t.Errorf("got: %v, want %v", got.First, want)
	}

}
