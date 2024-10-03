package hashing

import (
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
