package service

import "golang.org/x/crypto/bcrypt"

type Hasher interface {
	Hash(string) (string, error)
	Check(string, string) bool
}

type BcryptHasher struct{}

func (b *BcryptHasher) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (b *BcryptHasher) Check(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func PasswordHasher() Hasher {
	return &BcryptHasher{}
}
