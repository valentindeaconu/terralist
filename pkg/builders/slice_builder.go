package builders

// SliceBuilder is an interface to describe how a builder for slices
// should look like
type SliceBuilder[T any] interface {
	// Add adds a new element to the slice
	Add(T) SliceBuilder[T]

	// Build returns the slice
	Build() []T
}

// sliceBuilder is a concrete implementation of SliceBuilder
type sliceBuilder[T any] struct {
	args []T
}

// NewSliceBuilder returns a new SliceBuilder object using
// the sliceBuilder implementation
func NewSliceBuilder[T any]() SliceBuilder[T] {
	return &sliceBuilder[T]{
		args: []T{},
	}
}

func (b *sliceBuilder[T]) Add(arg T) SliceBuilder[T] {
	b.args = append(b.args, arg)

	return b
}

func (b *sliceBuilder[T]) Build() []T {
	return b.args
}
