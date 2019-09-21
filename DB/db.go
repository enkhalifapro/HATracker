package DB

type IPersistence interface {
	Connect() error

	Insert(i interface{}) error

	Update(i interface{}) error

	Select(query string) ([]map[string]interface{}, error)

	Close()
}