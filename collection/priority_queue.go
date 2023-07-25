package collection

type Comparator[T any] interface {
	Compare(base, compare T) bool
}
type PriorityQueue[T any] struct {
	data       []T
	comparator Comparator[T]
}

func NewPriorityQueue[T any](comparator Comparator[T]) *PriorityQueue[T] {
	if comparator == nil {
	}
	return &PriorityQueue[T]{comparator: comparator}
}

func (p *PriorityQueue[T]) Push(t T) {

}
