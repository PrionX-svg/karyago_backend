package pkg

import "strconv"

// ParseUint mengembalikan 0 kalau string kosong / tidak valid.
func ParseUint(s string) uint {
	if s == "" {
		return 0
	}
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return uint(v)
}

// ParseUintE versi yang mengembalikan error (kalau perlu validasi ketat).
func ParseUintE(s string) (uint, error) {
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(v), nil
}
