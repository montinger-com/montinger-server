package hashing

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var p = &params{
	memory:      64 * 1024,
	iterations:  3,
	parallelism: 2,
	saltLength:  16,
	keyLength:   32,
}

var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

func VerifyPassword(password string, hashEncoded string) (bool, error) {
	pa, salt, hash, err := decodeHash(hashEncoded)
	if err != nil {
		return false, err
	}

	newHash := argon2.IDKey([]byte(password), salt, pa.iterations, pa.memory, pa.parallelism, pa.keyLength)

	if subtle.ConstantTimeCompare(hash, newHash) != 1 {
		return false, nil
	}

	return true, nil
}

func decodeHash(hashEncoded string) (*params, []byte, []byte, error) {
	values := strings.Split(hashEncoded, "$")
	if len(values) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}

	var version int
	_, err := fmt.Sscanf(values[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}

	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	pa := &params{}
	_, err = fmt.Sscanf(values[3], "m=%d,t=%d,p=%d", &pa.memory, &pa.iterations, &pa.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err := base64.RawStdEncoding.Strict().DecodeString(values[4])
	if err != nil {
		return nil, nil, nil, err
	}
	pa.saltLength = uint32(len(salt))

	hash, err := base64.RawStdEncoding.Strict().DecodeString(values[5])
	if err != nil {
		return nil, nil, nil, err
	}
	pa.keyLength = uint32(len(hash))

	return pa, salt, hash, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func HashPassword(password string) (string, error) {
	salt, err := generateRandomBytes(p.saltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.memory, p.iterations, p.parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}
