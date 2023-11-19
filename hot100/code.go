package hot100

import "sort"

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	res := []int{}
	for i, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = i
		}
		if _, ok := m[target-num]; ok && m[target-num] != i {
			res = append(res, []int{m[target-num], i}...)
			return res
		}
	}
	return res
}
func groupAnagrams(strs []string) [][]string {
	mapStrs := map[string][]string{}
	for _, str := range strs {
		byteStr := []byte(str)
		sort.Slice(byteStr, func(i, j int) bool {
			return byteStr[i] > byteStr[j]
		})
		nstr := string(byteStr)
		mapStrs[nstr] = append(mapStrs[nstr], str)
	}
	res := [][]string{}
	for _, strings := range mapStrs {
		res = append(res, strings)
	}
	return res
}
