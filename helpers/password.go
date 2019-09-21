package helpers

import "golang.org/x/crypto/bcrypt"

// Password  to HashPassword and  CheckPasswordHash
type Password struct {
}

//_________________________________________________________________________________________________________________________Password

// HashPassword to encrypt the password
func (p *Password) Hash(password string) string {
  bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
  return string(bytes)
}

// CheckPasswordHash to decrypt the password
func (p *Password) CheckHash(password, hash string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

  return err == nil //if True return 1
}
