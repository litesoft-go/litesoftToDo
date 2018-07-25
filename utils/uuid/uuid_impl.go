package uuid

import (
	"fmt"
	"io"
	"crypto/rand"
)

/*
 * Implementations (that are not mocked) are expected to generate the RFC 4122 36 byte String form with 4 dashes
 */
type UUIDsourceImpl struct {
}

func NewImpl() UUIDsourceFactory {
	return &UUIDsourceImpl{}
}

func (this *UUIDsourceImpl) GetSource() UUIDsource {
	return asSourceImpl
}

// Taken from https://play.golang.org/p/4FkNSiUDMg
func asSourceImpl() (rUUID string, err error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if err == nil {
		if n != len(uuid) {
			err = fmt.Errorf("expected %d random bytes, but got: %d", len(uuid), n)
		} else {
			// variant bits; see section 4.1.1
			uuid[8] = uuid[8]&^0xc0 | 0x80
			// version 4 (pseudo-random); see section 4.1.3
			uuid[6] = uuid[6]&^0xf0 | 0x40

			rUUID = fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
		}
	}
	return
}
