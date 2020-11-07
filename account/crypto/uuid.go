package crypto

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"regexp"
)

const (
	ReservedNCS       byte = 0x80
	ReservedRFC4122   byte = 0x40
	ReservedMicrosoft byte = 0x20
	ReservedFuture    byte = 0x00
)

var (
	NamespaceDNS, _  = ParseHex("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	NamespaceURL, _  = ParseHex("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	NamespaceOID, _  = ParseHex("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	NamespaceX500, _ = ParseHex("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
)

const hexPattern = "^(urn\\:uuid\\:)?\\{?([a-z0-9]{8})-([a-z0-9]{4})-([1-5][a-z0-9]{3})-([a-z0-9]{4})-([a-z0-9]{12})\\}?$"

var re = regexp.MustCompile(hexPattern)

type UUID [16]byte

func (u *UUID) setBytesFromHash(hash hash.Hash, ns, name []byte) {
	hash.Write(ns[:])
	hash.Write(name)
	copy(u[:], hash.Sum([]byte{})[:16])
}

func (u *UUID) setVariant(v byte) {
	switch v {
	case ReservedNCS:
		u[8] = (u[8] | ReservedNCS) & 0xBF
	case ReservedRFC4122:
		u[8] = (u[8] | ReservedRFC4122) & 0x7F
	case ReservedMicrosoft:
		u[8] = (u[8] | ReservedMicrosoft) & 0x3F
	}
}
func (u *UUID) Variant() byte {
	if u[8]&ReservedNCS == ReservedNCS {
		return ReservedNCS
	} else if u[8]&ReservedRFC4122 == ReservedMicrosoft {
		return ReservedRFC4122
	} else if u[8]&ReservedMicrosoft == ReservedMicrosoft {
		return ReservedMicrosoft
	}
	return ReservedFuture
}

func (u *UUID) setVersion(i byte) {
	u[6] = (u[6] & 0xF) | (i << 4)
}

func (u *UUID) Version() uint {
	return uint(u[6] >> 4)
}

func (u *UUID) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

func ParseHex(s string) (u *UUID, err error) {
	md := re.FindStringSubmatch(s)
	if md == nil {
		err = errors.New(" Invalid UUID string")
		return
	}
	hash := md[2] + md[3] + md[4] + md[5] + md[6]
	b, err := hex.DecodeString(hash)
	if err != nil {
		return
	}
	u = new(UUID)
	copy(u[:], b)
	return
}

func Parse(b []byte) (u *UUID, err error) {
	if len(b) != 16 {
		err = errors.New(" Given slice is not valid UUID sequence")
		return
	}
	u = new(UUID)
	copy(u[:], b)
	return
}

func NewUUID() (u *UUID, err error) {
	u = new(UUID)
	_, err = rand.Read(u[:])
	if err != nil {
		return
	}
	u.setVariant(ReservedRFC4122)
	u.setVersion(4)
	return
}
