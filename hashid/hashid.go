package hashid

import (
	"errors"

	"github.com/speps/go-hashids/v2"
)

type Key struct {
	h *hashids.HashID
}

func NewKey(salt string, minLen int) *Key {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLen

	h, err := hashids.NewWithData(hd)
	if err != nil {
		panic(err)
	}

	return &Key{
		h: h,
	}
}

func (k *Key) Encode(prefix, key int64) (string, error) {
	return k.h.EncodeInt64([]int64{prefix, key})
}

func (k *Key) Decode(prefix int64, key string) (int64, error) {
	d, err := k.h.DecodeInt64WithError(key)
	if err != nil {
		panic(err)
	}
	if prefix != d[0] {
		return 0, errors.New("invalid prefix")
	}

	return d[1], nil
}
