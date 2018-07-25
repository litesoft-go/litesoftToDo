package sets

// taken from https://www.davidkaya.sk/2017/12/10/sets-in-golang/

var exists = struct{}{}

type StringSet struct {
	mMap map[string]struct{}
}

func NewString() StringSet {
	return StringSet{mMap: make(map[string]struct{})}
}

func (this *StringSet) Add(pValue string) {
	this.mMap[pValue] = exists
}

func (this *StringSet) Remove(pValue string) {
	delete(this.mMap, pValue)
}

func (this *StringSet) Contains(pValue string) bool {
	_, zFound := this.mMap[pValue]
	return zFound
}
