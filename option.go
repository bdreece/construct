package construct

// An Option is an opaque interface used to modify the construction or
// initialization of a value of type T.
type Option[T any] interface {
	apply(*T)
}

type option[T any] func(*T)

func (fn option[T]) apply(val *T) { fn(val) }

// With creates a new [Option], applying the provided modifiers to the target
// value when used.
func With[T any](modifiers ...func(val *T)) Option[T] {
	return option[T](func(val *T) {
		for _, modify := range modifiers {
			modify(val)
		}
	})
}
