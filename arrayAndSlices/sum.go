package arrayandslices

func SumSlice(nums []int) (res int) {
	for _, v := range nums {
		res += v
	}
	return
}

func SumAll(numsToSum ...[]int) []int {
	lenOfNums := len(numsToSum)
	sums := make([]int, lenOfNums)
	for i, v := range numsToSum {
		sums[i] = SumSlice(v)
	}
	return sums
}

func TailAll(numToTail ...[]int) []int {
	var tails []int
	for _, v := range numToTail {
		if len(v) == 0 {
			tails = append(tails, 0)
		} else {
			tails = append(tails, v[len(v)-1])
		}
	}
	return tails
}
