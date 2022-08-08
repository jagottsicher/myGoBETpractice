package mathUtils

func TheSumOf(inputValues ...int) int {

	sum := 0

	for _, v := range inputValues {
		sum += v
	}

	return sum
}
