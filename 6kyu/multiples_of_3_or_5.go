package kata

func Multiple3And5(number int) int {
	listOfMultiples := []int{}
	setOfMultiples := map[int]bool{}
	for i := 1; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			listOfMultiples = append(listOfMultiples, i)
		}
		for _, v := range listOfMultiples {
			setOfMultiples[v] = true
		}
	}
	if len(setOfMultiples) > 0 {
		sum := 0
		for k := range setOfMultiples {
			sum += k
		}
		return sum
	}
	return 0
}
