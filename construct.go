// Package construct is a generic library for implementing the functional
// options pattern in Go.
package construct

// Constructs a new value of type V, applying the provided options.
func New[V any](opts ...Option[V]) *V {
	v := new(V)
	for _, opt := range opts {
		opt.apply(v)
	}

	return v
}

// Applies the provided options to the existing value v, returning the
// result.
func Apply[V any](v *V, opts ...Option[V]) *V {
	for _, opt := range opts {
		opt.apply(v)
	}

	return v
}
