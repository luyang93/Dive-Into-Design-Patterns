package main

import "fmt"

type User struct {
	name string
	age  int
}

type Iterator interface {
	hasNext() bool
	getNext() *User
}

type UserIterator struct {
	Index int
	Users []*User
}

func (u *UserIterator) hasNext() bool {
	if u.Index < len(u.Users) {
		return true
	}
	return false
}

func (u *UserIterator) getNext() *User {
	if u.hasNext() {
		user := u.Users[u.Index]
		u.Index++
		return user
	}
	return nil
}

type Collection interface {
	CreateIterator() Iterator
}

type UserCollection struct {
	Users []*User
}

func (u *UserCollection) CreateIterator() Iterator {
	return &UserIterator{
		Users: u.Users,
	}
}

func main() {
	user1 := &User{
		name: "a",
		age:  30,
	}
	user2 := &User{
		name: "b",
		age:  20,
	}
	userCollection := &UserCollection{
		Users: []*User{user1, user2},
	}
	iterator := userCollection.CreateIterator()
	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
