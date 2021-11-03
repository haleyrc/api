package service

const MaxResults = 1000

func clampUint(min, val, max uint) uint {
	switch {
	case val < min:
		return min
	case val > max:
		return max
	default:
		return val
	}
}
