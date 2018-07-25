package uuid

/*
 * Implementations (that are not mocked) are expected to generate the RFC 4122 36 byte String form with 4 dashes
 */
type UUIDsource func() (string, error)

type UUIDsourceFactory interface {
	GetSource() UUIDsource
}
