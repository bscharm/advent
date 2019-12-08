package one

func SumFuelRequirements(modules []int) int {
	total := 0
	for _, i := range modules {
		total += calculateFuelRequirement(i)
	}
	return total
}

func calculateFuelRequirement(i int) int {
	val := (i / 3) - 2

	if val <= 0 {
		return 0
	}

	val += calculateFuelRequirement(val)
	return val
}
