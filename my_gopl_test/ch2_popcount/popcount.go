package popcount

// pc[i] is the population count of i.
var pc [256]byte = func() (pc [256]byte) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return
}()

// func init() {
// 	for i := range pc {
// 		pc[i] = pc[i/2] + byte(i&1)
// 	}
// }

// PopCount returns the population count (number of set bits) of x.
// look-up table
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))],
	)
}

// loop
func PopCount2(x uint64) int {
	var count int
	for i := uint(0); i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}

// offset
func PopCount3(x uint64) int {
	var count int
	for i := uint(0); i < 64; i++ {
		if (x & 1) == 1 {
			count += 1
		}
        x >>= 1
	}
	return count
}
