package validate

import "fmt"

type SourcedData interface {
	Found() bool
	Get() string
	Source() fmt.Stringer
}
