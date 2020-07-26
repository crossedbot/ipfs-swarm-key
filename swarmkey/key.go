package swarmkey

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

const (
	Protocol     = "/key/swarm/psk/1.0.0/"
	BaseEncoding = "base16"
	KeyLength    = 32
)

type Key interface {
	Bytes() []byte
	String() string
}

type key struct {
	Proto string
	Enc   string
	Key   []byte
}

func New() (Key, error) {
	b := make([]byte, KeyLength)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return &key{
		Proto: Protocol,
		Enc:   BaseEncoding,
		Key:   b,
	}, nil
}

func (k *key) Bytes() []byte {
	return []byte(k.String())
}

func (k *key) String() string {
	return fmt.Sprintf(
		"%s\n%s\n%s",
		k.Proto,
		k.Enc,
		hex.EncodeToString(k.Key),
	)
}
