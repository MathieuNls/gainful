package indexable

type HasStringIndex interface {
	StringIndex() string
}

type indexableImpl struct {
	value string
}

func (indexable indexableImpl) StringIndex() string {
	return indexable.value
}

func New(str string) *indexableImpl {
	return &indexableImpl{
		value: str,
	}
}
