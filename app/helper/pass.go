package helper

import "golang.org/x/crypto/bcrypt"

func CreatePass(pass string) ([]byte, error) {
	byte := []byte(pass)

	hash, err := bcrypt.GenerateFromPassword(byte, bcrypt.DefaultCost)
	if err != nil {
		return hash, err
	}

	return hash, nil
}

func ComparePass(hash []byte, input string) bool {
	byteInput := []byte(input)

	// check
	if err := bcrypt.CompareHashAndPassword(hash, byteInput); err != nil {
		return false
	}

	return true
}
