package container

type empty struct{}

type SetItem interface{}

type Set map[SetItem]empty

func (s Set) Has(item SetItem) bool {
	_, exists := s[item]
	return exists
}

func (s Set) Insert(item SetItem) {
	s[item] = empty{}
}

func (s Set) Delete(item SetItem) {
	delete(s, item)
}

func (s Set) Len() int {
	return len(s)
}
