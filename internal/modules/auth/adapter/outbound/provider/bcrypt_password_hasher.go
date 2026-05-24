package provider

import "golang.org/x/crypto/bcrypt"

type BcryptPasswordHasher struct{}

func NewBcryptPasswordHasher() *BcryptPasswordHasher {
	return &BcryptPasswordHasher{}
}

func (b *BcryptPasswordHasher) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	return string(bytes), err
}

func (b *BcryptPasswordHasher) Compare(hashed, plain string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashed),
		[]byte(plain),
	)
}
