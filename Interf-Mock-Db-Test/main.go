package main

import (
	"fmt"
	"log"
)

type User struct {
	ID    int
	First string
}

type MockDatastore struct {
	Users map[int]User
}

func (md MockDatastore) GetUser(id int) (User, error) {
	user, ok := md.Users[id]
	if !ok {
		return User{}, fmt.Errorf("user %v not found", id)
	}
	return user, nil
}

func (md MockDatastore) SaveUser(u User) error {
	_, ok := md.Users[u.ID]
	if ok {
		return fmt.Errorf("user %v already exists", u.ID)
	}
	md.Users[u.ID] = u
	return nil
}

type DataStore interface {
	GetUser(int) (User, error)
	SaveUser(User) error
}

type Service struct {
	db DataStore
}

func (s Service) GetUser(id int) (User, error) {
	return s.db.GetUser(id)
}

func (s Service) SaveUser(u User) error {
	return s.db.SaveUser(u)
}

func main() {
	db := MockDatastore{
		Users: make(map[int]User),
	}

	serv := Service{
		db: db,
	}
	ui := User{
		First: "Ange",
		ID:    1,
	}
	err := serv.SaveUser(ui)
	if err != nil {
		log.Fatalf("error %s", err)
	}
	userRet, err := serv.GetUser(ui.ID)
	if err != nil {
		log.Fatalf("error %s", err)
	}
	fmt.Println(userRet)

}
