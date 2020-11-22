package lottery

import (
	"sort"
	"strconv"
	"strings"
)

// 大乐透计算项目结果
type LotteryResult struct {
	Name  string
	Count int
}

// 取得所有计算项目的名称
func GetResultName() []string {
	var names []string
	for _, r := range ljArr {
		names = append(names, r.key)
	}
	return names
}

// 大乐透计算项目
type lotteryJudge struct {
	key  string                                          // 计算项目名称
	cond string                                          // 判断条件
	fn   func(this lotteryJudge, s string) LotteryResult // 判断函数
}

func contains(lj lotteryJudge, str string) LotteryResult {
	lr := LotteryResult{
		Name:  lj.key,
		Count: 0,
	}
	for _, s := range strings.Split(str, ",") {
		if strings.Contains(lj.cond, " "+s+" ") {
			lr.Count++
		}
	}
	return lr
}

func largeThen18(lj lotteryJudge, str string) LotteryResult {
	lr := LotteryResult{
		Name:  lj.key,
		Count: 0,
	}
	for _, s := range strings.Split(str, ",") {
		if num, err := strconv.Atoi(s); err == nil && num >= 18 {
			lr.Count++
		}
	}
	return lr
}

func odd(lj lotteryJudge, str string) LotteryResult {
	lr := LotteryResult{
		Name:  lj.key,
		Count: 0,
	}
	for _, s := range strings.Split(str, ",") {
		if num, err := strconv.Atoi(s); err == nil && num%2 == 1 {
			lr.Count++
		}
	}
	return lr
}

func ac(lj lotteryJudge, str string) LotteryResult {
	lr := LotteryResult{
		Name:  lj.key,
		Count: 0,
	}
	var nums []int
	for _, s := range strings.Split(str, ",") {
		if num, err := strconv.Atoi(s); err == nil {
			nums = append(nums, num)
		}
	}

	var dValue []int
	for k, v := range nums {
		for _, v1 := range nums[k+1:] {
			dValue = append(dValue, v1-v)
		}
	}

	lr.Count = int(removeDuplicate(dValue) - 4)
	return lr
}

func removeDuplicate(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	sort.Ints(nums)

	i, j := 0, 1
	for ; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

var ljArr []lotteryJudge

func init() {
	ljArr = []lotteryJudge{
		{
			key: "AC值",
			fn:  ac,
		},
		{
			key: "大数",
			fn:  largeThen18,
		},
		{
			key: "奇数",
			fn:  odd,
		},
		{
			key:  "质数",
			cond: " 1 2 3 5 7 11 13 17 19 23 29 31 ",
			fn:   contains,
		},
		{
			key:  "边界",
			cond: " 1 2 3 5 7 11 13 17 19 23 29 31 ",
			fn:   contains,
		},
		{
			key:  "0路",
			cond: " 3 6 9 12 15 18 21 24 27 30 33 ",
			fn:   contains,
		},
		{
			key:  "1路",
			cond: " 1 4 7 10 13 16 19 22 25 28 31 34 ",
			fn:   contains,
		},
		{
			key:  "2路",
			cond: " 2 5 8 11 14 17 20 23 26 29 32 35 ",
			fn:   contains,
		},
		{
			key:  "1区",
			cond: " 1 2 3 4 5 ",
			fn:   contains,
		},
		{
			key:  "2区",
			cond: " 6 7 8 9 10 ",
			fn:   contains,
		},
		{
			key:  "3区",
			cond: " 11 12 13 14 15 ",
			fn:   contains,
		},

		{
			key:  "4区",
			cond: " 16 17 18 19 20 ",
			fn:   contains,
		},
		{
			key:  "5区",
			cond: " 21 22 23 24 25 ",
			fn:   contains,
		},
		{
			key:  "6区",
			cond: " 26 27 28 29 30 ",
			fn:   contains,
		},
		{
			key:  "7区",
			cond: " 31 32 33 34 35 ",
			fn:   contains,
		},
		{
			key:  "1行",
			cond: " 1 6 11 16 21 26 31 ",
			fn:   contains,
		},
		{
			key:  "2行",
			cond: " 2 7 12 17 22 27 32 ",
			fn:   contains,
		},
		{
			key:  "3行",
			cond: " 3 8 13 18 23 28 33 ",
			fn:   contains,
		},

		{
			key:  "4行",
			cond: " 4 9 14 19 24 29 34 ",
			fn:   contains,
		},
		{
			key:  "5行",
			cond: " 5 10 15 20 25 30 35 ",
			fn:   contains,
		},
	}
}

func Compute(str string) []LotteryResult {
	var lrArr []LotteryResult
	for _, lj := range ljArr {
		lrArr = append(lrArr, lj.fn(lj, str))
	}
	return lrArr
}
