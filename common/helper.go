package common

func Filter[T any](slice []T, fn func(it T) bool) []T {
	var ret []T
	for _, it := range slice {
		if fn(it) {
			ret = append(ret, it)
		}
	}
	return ret
}

func FilterNotNil[T any](slice []T) []T {
	return Filter(slice, func(it T) bool {
		var anyIt any = it
		return anyIt != nil
	})
}

// Iter applies the given function fn to each element of the slice
// and returns a new slice of type N with the results.
func Iter[T any, N any](slice []T, fn func(it T) N) []N {
	retSlice := make([]N, 0, len(slice))
	for index := range slice {
		retSlice = append(retSlice, fn(slice[index]))
	}

	return retSlice
}

// UniqBy removes duplicates from the input slice based on a computed value from each element.
// It uses the provided 'fn' function to generate a comparable value for each element
// and ensures that only unique computed values are kept in the result.
func UniqBy[T any, C comparable](arr []T, fn func(T) C) []T {
	seen := make(map[C]struct{}, len(arr))
	var result []T

	for _, item := range arr {
		key := fn(item)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		result = append(result, item)
	}

	return result
}
