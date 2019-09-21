package DB

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //???
)

type PostgresHelper struct {
	db *sqlx.DB
}

//Database Connection
/*
func init() {
	//fmt.Println("hi")

	postgres := PostgresHelper{}

	db, err := sqlx.Connect("postgres", "user=HA  dbname=HATracker password=1  host=localhost sslmode=disable")

	if err != nil {
		log.Fatalln("couden't connect to Database : %v", err)
	}

	postgres.db = db
}
*/
//_________________________________________________________________________________________________________________________Connect

// Connect to  Database
func (postgres *PostgresHelper) Connect() error {

	//--if db  is Already defined--
	if postgres.db != nil {

		return nil
	}

	//------if not-------
	db, err := sqlx.Connect("postgres", "user=HA  dbname=HATracker password=1  host=localhost sslmode=disable")
	if err != nil {

		return (err)
	}

	postgres.db = db
	return nil

}

//_________________________________________________________________________________________________________________________CreateQuery

// CreateQuery for Database
func (postgres *PostgresHelper) CreateQuery(q interface{}, operation string) (string, error) {

	//---Table Name ---
	operation = strings.ToUpper(operation)
	switch operation {
	case "INSERT":

		if reflect.ValueOf(q).Kind() == reflect.Struct {
			tableName := strings.ToLower(reflect.TypeOf(q).Name())

			query := fmt.Sprintf("INSERT INTO %s  VALUES(", tableName)
			elements := reflect.ValueOf(q)
			/*
				query := fmt.Sprintf("INSERT INTO %s", tableName)

				elements := reflect.ValueOf(q)
				//---if has been send unorganized Struct such as Database Table
				//---Get columns
				for i := 0; i < elements.NumField(); i++ {
					if i != 0 {
						query = fmt.Sprintf("%s, %s", query, elements.Type().Field(i).Name)
					} else {
						query = fmt.Sprintf("%s(%s", query, elements.Type().Field(i).Name)
					}
				}
				query = fmt.Sprintf("%s) VALUES(", query)
			*/
			// Creat Scurity Object for sha Passowrd
			scurity := Scurity{}

			//---Value of the elements
			for i := 0; i < elements.NumField(); i++ {
				switch elements.Field(i).Kind() {
				case reflect.Int:
					if i != 0 {
						query = fmt.Sprintf("%s, %d", query, elements.Field(i).Int())
					} else {
						query = fmt.Sprintf("%s%d", query, elements.Field(i).Int())
					}
				case reflect.String:
					if i != 0 {
						if elements.Type().Field(i).Name != "Password" {
							query = fmt.Sprintf("%s, '%s'", query, elements.Field(i).String())
						} else {
							query = fmt.Sprintf("%s, '%s'", query, scurity.HashPassword(elements.Field(i).String()))
						}

					} else {
						query = fmt.Sprintf("%s'%s'", query, elements.Field(i).String())
					}
				}
			}
			query = fmt.Sprintf("%s)", query)
			return query, nil
		}
	case "Update":

	}
	return "", fmt.Errorf("unsupported Type in: PostgresHelper.CreateQuery")
}

//___________________________________________________________________________________________________________________________Insert

// Insert a new item into db
func (postgres *PostgresHelper) Insert(i interface{}) error {

	//---Creat SQL Query ---
	query, err := postgres.CreateQuery(i, "insert")
	fmt.Println(query)

	if err == nil {
		//---Insert----

		tx := postgres.db.MustBegin()

		resulte := tx.MustExec(query)

		if resulte == nil {

			return fmt.Errorf(" couden't INSERT INTO PostgresHelper SQL Table: %s", reflect.TypeOf(i).Name())
		}
		tx.Commit()
		return nil

	}

	return err

}

//___________________________________________________________________________________________________________________________Select

// Select an existing item from Database
func (postgres *PostgresHelper) Select(query string) ([]map[string]interface{}, error) {

	row, err := postgres.db.Queryx(query)
	if err != nil {
		return nil, fmt.Errorf("PostgresHelper Select Method %s", err)
	}

	m := make(map[string]interface{})
	var e []map[string]interface{}

	//	i := 0
	for row.Next() {
		row.MapScan(m)
		e = append(e, m) //Change it to Pointer
		//i++
	}

	return e, nil
}

//___________________________________________________________________________________________________________________________Update

// Update an existing item into db
func (postgres *PostgresHelper) Update(i interface{}) error {

	// Add your implementation here

	return nil
}

//___________________________________________________________________________________________________________________________Close

// Close db
func (postgres *PostgresHelper) Close() {
	postgres.db.Close()
}
