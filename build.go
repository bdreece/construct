package construct

// A Factory defines the contract for a function that constructs a values
// from a single parameter input.
type Factory[P, V any] func(param P) *V

// A Builder provides a mechanism for procedurally configuring the construction
// of a value.
//
// Deferred construction is achieved by delegating to a [Factory]. [Option]
// implementors supplied to the Builder must modify the factory parameter, not
// the resulting value.
type Builder[P, V any] struct {
	param  P
	create Factory[P, V]
}

// Applies the provided options to the contained parameter.
func (b *Builder[P, V]) Apply(opts ...Option[P]) *Builder[P, V] {
	_ = Apply(&b.param, opts...)
	return b
}

// Constructs the target value using the contained parameter.
func (b *Builder[P, V]) Build() *V {
	return b.create(b.param)
}

// Initializes a new Builder for use in procedurally constructing values.
//
// See [Builder] for more details.
func NewBuilder[P, V any](factory Factory[P, V]) *Builder[P, V] {
	return &Builder[P, V]{
		create: factory,
	}
}

// Build constructs the target value by initializing a new [Builder] and applying
// the provided options.
func Build[P, V any](factory Factory[P, V], opts ...Option[P]) *V {
	return NewBuilder(factory).Apply(opts...).Build()
}
