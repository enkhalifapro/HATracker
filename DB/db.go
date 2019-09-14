package DB

type IPersistence interface {
  Insert(i interface{}) error
  Update(i interface{}) error
}
