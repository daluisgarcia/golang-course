package main

import "fmt"

type HashAlgorithm interface { // Strategy interface
	Hash(p *PasswordProtector)
}

type PasswordProtector struct {
	user          string
	passwordName  string
	hashAlgorithm HashAlgorithm // Strategy field
}

func NewPasswordProtector(user string, passwordName string, hashAlgorithm HashAlgorithm) *PasswordProtector { // Constructor
	return &PasswordProtector{
		user:          user,
		passwordName:  passwordName,
		hashAlgorithm: hashAlgorithm,
	}
}

func (p *PasswordProtector) SetHashAlgorithm(hash HashAlgorithm) { // Strategy setter
	p.hashAlgorithm = hash
}

func (p *PasswordProtector) Hash() {
	p.hashAlgorithm.Hash(p)
}

type SHA struct{}

func (SHA) Hash(p *PasswordProtector) { // Concrete strategy
	fmt.Printf("Hashing password '%s' for user '%s' with SHA algorithm\n", p.passwordName, p.user)
}

type MD5 struct{}

func (MD5) Hash(p *PasswordProtector) { // Concrete strategy
	fmt.Printf("Hashing password '%s' for user '%s' with MD5 algorithm\n", p.passwordName, p.user)
}

func main() {
	sha := SHA{}
	md5 := MD5{}

	passwordProtector := NewPasswordProtector("user", "gmail password", sha)
	passwordProtector.Hash()

	passwordProtector.SetHashAlgorithm(md5)
	passwordProtector.Hash()
}
