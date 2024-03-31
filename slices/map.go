package slices

func Map[T, U any](slice []T, transform func(T) U) []U {
	out := []U{}
	for _, val := range slice {
		out = append(out, transform(val))
	}
	return out
}
