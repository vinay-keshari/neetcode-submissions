func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		if _, exists := count[num]; exists {
			count[num] += 1
		} else {
			count[num] = 1
		}
	}
	n := len(nums)+1 
	buckets := make([][]int, n)
	
	for key, value := range count {
		buckets[value] = append(buckets[value], key)
	}

	res := []int{}
	for i := (len(buckets) - 1); i>0 ; i--{
		for _, num := range buckets[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	} 

	return res
}
