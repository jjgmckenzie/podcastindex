package internal

import "unsafe"

// BoolToInt is a helper function to convert boolean to int (1 for true, 0 for false)
func BoolToInt(b bool) int {
	return int(*(*byte)(unsafe.Pointer(&b)))
}
