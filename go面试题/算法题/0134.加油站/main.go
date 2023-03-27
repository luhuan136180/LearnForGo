package main

import "math"

//暴力解法
func canCompleteCircuit(gas []int, cost []int) int {
	rest := 0  //剩余油量
	index := 0 //当前位置
	for i := 0; i < len(cost); i++ {
		rest = gas[i] - cost[i]      //计算走完下一步所剩与油量
		index = (i + 1) % len(cost)  //计算下一步到达的索引下标
		for rest > 0 && index != i { //当油量有剩余，且没有回到出发点，则一直遍历
			rest += gas[index] - cost[index]
			index = (index + 1) % len(cost)
		}
		if rest >= 0 && index == i {
			return i
		}

	}
	return -1
}

//全局最优思想
//	情况⼀：如果gas的总和⼩于cost总和，那么⽆论从哪⾥出发，⼀定是跑不了⼀圈的
//	情况⼆：rest[i] = gas[i]-cost[i]为⼀天剩下的油，i从0开始计算累加到最后⼀站，如
//果累加没有出现负数，说明从0出发，油就没有断过，那么0就是起点。
//	情况三：如果累加的最⼩值是负数，汽车就要从⾮0节点出发，从后向前，看哪个节
//点能这个负数填平，能把这个负数填平的节点就是出发节点

func canCompleteCircuit2(gas []int, cost []int) int {
	curSum, min := 0, math.MaxInt64
	for i := 0; i < len(gas); i++ {
		rest := gas[i] - cost[i]
		curSum += rest
		if curSum < min {
			min = curSum
		}
	}
	if curSum < 0 { //1
		return -1
	}
	if min >= 0 { //2
		return 0
	}
	for i := len(gas) - 1; i >= 0; i-- { //3
		rest := gas[i] - cost[i]
		min += rest
		if min >= 0 {
			return i
		}
	}
	return -1
}

//局部最优思想
func canCompleteCircuit3(gas []int, cost []int) int {
	curSum, totalSum, start := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		curSum += gas[i] - cost[i]
		totalSum += gas[i] - cost[i]
		if curSum < 0 { // 当前累加rest[i]和 curSum⼀旦⼩于0
			start = i + 1 // 起始位置更新为i+1
			curSum = 0    // curSum从0开始
		}
	}
	if totalSum < 0 {
		return -1
	}
	return start
}
