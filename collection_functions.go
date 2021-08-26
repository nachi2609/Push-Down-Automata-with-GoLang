// https://gobyexample.com/collection-functions
package main

func Index(vs []string, t string) int {
	// index of t in vs
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	// true if t in vs
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
	// tru if any s in vs makes f(s) true
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	// true if all s in vs make f(s) true
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	// slice of s in vs making f(s) true
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	// slice of applying f to each element of vs
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
