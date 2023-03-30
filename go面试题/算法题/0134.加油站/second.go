package main

func canCompleteCircuit4(gas []int, cost []int) int {
	cur, total := 0, 0
	index := 0
	for i := 0; i < len(gas); i++ {
		cur += gas[i] - cost[i]
		total += gas[i] - cost[i]
		if cur < 0 {
			index = i + 1
			cur = 0
		}
	}
	if total < 0 {
		return -1
	}
	return index
}
