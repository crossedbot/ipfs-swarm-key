package swarmkey

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

const (
	Protocol  = "/key/swarm/psk/1.0.0/"
	KeyLength = 32

	// Encodings
	Base16Encoding = "/base16/"
)

// Key represents the interface to a IPFS swarm key
type Key interface {
	Bytes() []byte
	String() string
}

// key represents an IPFS swarm key
type key struct {
	Proto string
	Enc   string
	Key   []byte
}

// New returns an IPFS swarm key
func New() (Key, error) {
	b := make([]byte, KeyLength)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return &key{
		Proto: Protocol,
		Enc:   Base16Encoding,
		Key:   b,
	}, nil
}

// Bytes returns the bytes representation of key
func (k *key) Bytes() []byte {
	return []byte(k.String())
}

// String returns the string representation of key
func (k *key) String() string {
	return fmt.Sprintf(
		"%s\n%s\n%s",
		k.Proto,
		k.Enc,
		hex.EncodeToString(k.Key),
	)
}
