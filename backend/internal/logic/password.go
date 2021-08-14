package logic

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strings"
)

const defaultPasswordLength = 16
const defaultPasswordCharSet = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"0123456789" +
	"=!$%&+#-_.,;:"

func (l *Logic) defaultPassword() string {
	builder := strings.Builder{}
	for i := 0; i < defaultPasswordLength; i++ {
		builder.WriteRune(rune(defaultPasswordCharSet[rand.Intn(len(defaultPasswordCharSet))]))
	}
	return builder.String()
}

func (l *Logic) hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func (l *Logic) checkPassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
