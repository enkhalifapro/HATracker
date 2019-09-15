package DB

import "fmt"

type PostgresHelper struct {

}

// Insert a new item into db
func (p *PostgresHelper) Insert(i interface{}) error {
  // Add your implementation here
  fmt.Println("I'm inside posgresHelper")

  return nil
}

// Update an existing item into db
func (p *PostgresHelper) Update(i interface{}) error {

  // Add your implementation here

  return nil
}
