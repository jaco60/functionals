package functionals

type Number interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 |
		float64 | float32
}

func Map[T, S any](xs []T, fct func(x T) S) []S {
	res := make([]S, len(xs))
	for i, e := range xs {
		res[i] = fct(e)
	}
	return res
}

func Filter[T any](xs []T, fct func(x T) bool) []T {
	res := []T{}
	for _, e := range xs {
		if fct(e) {
			res = append(res, e)
		}
	}
	return res
}

func FoldLeft[T, S any](xs []T, accu S, fct func(acc S, e T) S) S {
	for _, e := range xs {
		accu = fct(accu, e)
	}
	return accu
}

func Sum[T Number](xs []T) T {
	var accu T
	return FoldLeft(xs, accu, func(x, y T) T { return x + y })
}

func Product[T Number](xs []T) T {
	var accu = T(1)
	return FoldLeft(xs, accu, func(x, y T) T { return x * y })
}

func All[T any](xs []T, fct func(x T) bool) bool {
	for i := 0; i < len(xs); i++ {
		if !fct(xs[i]) {
			return false
		}
	}
	return true
}

func Any[T any](xs []T, fct func(x T) bool) bool {
	for i := 0; i < len(xs); i++ {
		if fct(xs[i]) {
			return true
		}
	}
	return false
}
