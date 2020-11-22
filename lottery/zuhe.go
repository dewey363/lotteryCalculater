package lottery

import (
	"sort"
	"strconv"
	"strings"
)

type Input struct {
	String string
	Pick   int
}

// 根据输入参数计算所有组合
func ZuHe(input []Input) []string {
	rArr := make([]*[][]int, len(input))
	for k, v := range input {
		strArr := strings.Split(v.String, ",")
		arr := make([]int, len(strArr))
		for k, v := range strArr {
			arr[k], _ = strconv.Atoi(v)
		}
		pick := v.Pick
		result := make([][]int, 0)
		index := make([]int, len(arr))
		for i := 0; i < pick; i++ {
			index[i] = 1
		}
		combine(arr, pick, &result, index)
		rArr[k] = &result
	}
	start := make([]int, len(rArr))
	end := make([]int, len(rArr))
	for k, v := range rArr {
		start[k] = 0
		end[k] = len(*v)
	}
	rIndex := make([][]int, 0)
	arrayJoin(start, end, &rIndex)
	var ret []string
	for _, v := range rIndex {
		var r []int
		for k, i := range v {
			r = append(r, (*rArr[k])[i]...)
		}
		// 添加排序
		sort.Ints(r)
		// 转成字符数组
		sr := make([]string, len(r))
		for k, v := range r {
			sr[k] = strconv.Itoa(v)
		}
		ret = append(ret, strings.Join(sr, ","))
	}
	return ret
}

// 从字符串数组中取N个数的所有结果
func combine(arr []int, pick int, result *[][]int, index []int) {
	var re []int
	for k, v := range index {
		if v == 1 {
			re = append(re, arr[k])
		}
	}
	*result = append(*result, re)
	for i, j := 0, 0; i < len(index)-1; i++ {
		if index[i] == 1 {
			j++
		}
		if index[i] == 1 && index[i+1] == 0 {
			index[i] = 0
			index[i+1] = 1
			for k := 0; k < i; k++ {
				if k < j-1 {
					index[k] = 1
				} else {
					index[k] = 0
				}
			}
			combine(arr, pick, result, index)
		}
	}
}

// 计算多个数据组合的索引下标结果
func arrayJoin(start []int, end []int, result *[][]int) {
	var temp = make([]int, len(start))
	copy(temp, start)
	for i := len(start) - 1; i > 0; i-- {
		if start[i] >= end[i] {
			start[i] = 0
			start[i-1]++
			arrayJoin(start, end, result)
		}
	}
	if start[0] >= end[0] {
		return
	}
	*result = append(*result, temp)
	start[len(start)-1]++
	arrayJoin(start, end, result)
}
