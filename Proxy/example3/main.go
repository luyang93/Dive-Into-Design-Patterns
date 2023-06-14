package main

import "fmt"

type DatabaseService interface {
	Query(query string) string
}

type RealDatabase struct{}

func (r *RealDatabase) Query(query string) string {
	return fmt.Sprintf("Results for query: %s", query)
}

type DatabaseProxy struct {
	realDatabase *RealDatabase
}

func (p *DatabaseProxy) Query(query string) string {
	if p.realDatabase == nil {
		fmt.Println("Initializing real database...")
		p.realDatabase = &RealDatabase{}
	}
	fmt.Println("Querying real database...")
	return p.realDatabase.Query(query)
}

func main() {
	database := &DatabaseProxy{}

	// The real database won't be initialized until we need to query.
	fmt.Println(database.Query("SELECT * FROM users"))
}
