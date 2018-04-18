package main

import (
	"fmt"
)

const (
	Byte = 1 << (3 + 10*iota)
	KB   // 1024 * 8
	MB   // 1024 * 1024 * 8
	GB   // 1024 * 1024 * 1024 * 8
	TB   // pow(1024, 4) * 8
	PB   // pow(1024, 5) * 8
	EB   // pow(1024, 6) * 8
	ZB   // pow(1024, 7) * 8
	YB   // pow(1024, 8) * 8
)

func main() {
	fmt.Printf(" 1B = %d bits\n", Byte)
	fmt.Printf("1KB = %d bits\n", KB)
	fmt.Printf("1MB = %d bits\n", MB)
	fmt.Printf("1GB = %d bits\n", GB)
	fmt.Printf("1TB = %d bits\n", TB)
	fmt.Printf("1PB = %d bits\n", PB)
	fmt.Printf("1EB = %d bits\n", uint64(EB))
	//fmt.Printf("1ZB = %d bits\n", ZB)    // overflows uint64
	//fmt.Printf("1YB = %d bits\n", YB)    // overflows uint64
	fmt.Printf("1ZB / 1TB = %d \n", ZB/TB)
	fmt.Printf("1YB / 1TB = %d \n", YB/TB)

}
