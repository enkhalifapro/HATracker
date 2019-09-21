package DB

import "golang.org/x/crypto/bcrypt"

type IPersistence interface {
	Connect() error

	Insert(i interface{}) error

	Update(i interface{}) error

	Select(query string) ([]map[string]interface{}, error)

	Close()
}

// Scurity  to HashPassword and  CheckPasswordHash
type Scurity struct {
}

//_________________________________________________________________________________________________________________________Password

// HashPassword to encrypt the password
func (p Scurity) HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

// CheckPasswordHash to decrypt the password
func (p Scurity) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil //if True return 1
}
