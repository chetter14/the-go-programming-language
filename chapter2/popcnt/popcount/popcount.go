// Counts the number of set bits in number 'x'
package popcount

var pc [256]byte

func init() {
	for i, _ := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	total_cnt := 0
	for i := 0; i < 8; i++ {
		total_cnt += int(pc[byte(x>>(i*8))])
	}
	return total_cnt
}
