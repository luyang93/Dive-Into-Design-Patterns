package main

import (
	"fmt"
	"sync"
)

var (
	once     sync.Once
	instance *Database
)

// Database defines the singleton database class
type Database struct {
	ID int
	// Other fields can be added here
}

// getInstance returns the singleton instance of the Database
func (d *Database) getInstance() *Database {
	once.Do(func() {
		instance = &Database{}
		// Initialization code here
	})
	return instance
}

// query executes a database query
func (d *Database) query(sql string) {
	fmt.Println(sql)
	// Query logic here
}

// Application represents the main application
type Application struct {
	// Other fields can be added here
}

// main is the entry point of the application
func (a *Application) main() {
	foo := instance.getInstance()
	foo.query("SELECT foo ...")
	// ...
	bar := instance.getInstance()
	bar.query("SELECT bar ...")
	// The variable `bar` will contain the same object as
	// the variable `foo`.
}

func main() {
	app := &Application{}
	app.main()
}
