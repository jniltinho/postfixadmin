package handler

import (
	"fmt"
	"math/rand"
	"time"

	"postfixadmin/util/crypt"
	_ "postfixadmin/util/crypt/md5_crypt"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePass(password string) string {
	c := crypt.MD5.New()
	hash, err := c.Generate([]byte(password), getSalt())
	if err != nil {
		panic(err)
	}
	return hash
}

func ComparePass(hashedPassword, password string) bool {
	c := crypt.MD5.New()
	err := c.Verify(hashedPassword, []byte(password))
	//fmt.Println(err)
	return err == nil
}

func getSalt() []byte {
	// Generate a random string for use in the salt
	const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := make([]byte, 8)
	for i := range s {
		s[i] = charset[seededRand.Intn(len(charset))]
	}
	salt := []byte(fmt.Sprintf("$1$%s", s))
	return salt
}

func getSalt2(password string) []byte {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 4)
	salt := string(bytes)[len(string(bytes))-11:]
	salt = fmt.Sprintf("$1$%s", salt)
	return []byte(salt)
}
