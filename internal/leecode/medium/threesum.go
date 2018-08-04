package medium

/*
https://leetcode.com/problems/3sum/description/

Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]

 */

/*
超过时间，
去重和符合表达式的判断，做的都不理想
 */
func ThreeSum1(nums []int) [][]int {

	if len(nums) < 3 {
		return nil
	}

	abcSlice := make([][]int, 0, len(nums) * len(nums[1:]) * len(nums[2:]))

	count := len(nums)
	for i := 0; i < count - 2; i++ {
		a := nums[i]
		for j := i + 1; j < count - 1; j++ {
			b := nums[j]
			for k := j + 1; k < count; k++ {
				c := nums[k]
				if a + b + c == 0 {
					abcSlice = append(abcSlice, []int{a, b, c})
				}
			}
		}
	}

	var excludeindex map[int]int = make(map[int]int)
	for i := 0; i < len(abcSlice); i++ {
		triplets := abcSlice[i]
		for j := i + 1; j < len(abcSlice); j ++ {
			sectriplets := abcSlice[j]
			//if ((sectriplets[0] == triplets[0] || sectriplets[0] == triplets[1] || sectriplets[0] == triplets[2]) &&
			//	(sectriplets[1] == triplets[0] || sectriplets[1] == triplets[1] || sectriplets[1] == triplets[2]) &&
			//	(sectriplets[2] == triplets[0] || sectriplets[2] == triplets[1] || sectriplets[2] == triplets[2])) {
			if CompareTwoSlice(triplets, sectriplets) {
				if _, ok := excludeindex[j]; !ok {
					excludeindex[j] = j
				}
			}
		}
	}

	retabcSlice := make([][]int, 0, len(abcSlice))
	for i := 0; i < len(abcSlice); i++ {
		if _, ok := excludeindex[i]; !ok {
			retabcSlice = append(retabcSlice, abcSlice[i])
		}
	}

	return retabcSlice
}

func CompareTwoSlice(a []int, b []int) bool {

	var flag int = -1
	for i := 0; i < len(b); i++ {
		if a[0] == b[i] {
			flag = i
			break
		}
	}
	if flag == -1 {
		return false
	}

	for i := 0; i < len(b); i++ {
		if a[1] == b[i] && i != flag {
			return true
		}
	}

	return false
}

func ThreeSum2(nums []int) [][]int {
	ret := [][]int{}

	if nums != nil  && (len(nums) >= 3) {

		length := len(nums)
		aMap := make(map[int]int, length)
		bMap := make(map[int]int, length)
		aList := make([]int, 0, length)
		bList := make([]int, 0, length)

		for _, num := range nums {
			if num <= 0 {
				aMap[num] = aMap[num] + 1
			} else {
				bMap[num] = bMap[num] + 1
			}
		}

		for num, count := range aMap {
			aList = append(aList, num)
			if count >= 2 {
				aList = append(aList, num)
			}
		}

		for num, count := range bMap {
			bList = append(bList, num)
			if count >= 2 {
				bList = append(bList, num)
			}
		}

		for i := 0; i < len(aList) - 1; i++ {
			num1 := aList[i]
			if (i > 0 && num1 == aList[i - 1]) {
				continue
			}
			for j := i + 1; j < len(aList); j++ {
				num2 := aList[j]
				if j > i + 1 && num2 == aList[j - 1] {
					continue
				}
				if _, ok := bMap[-(num1 + num2)]; ok {
					ret = append(ret, []int{num1, num2, -(num1 + num2)})
				}
			}
		}

		for i := 0; i < len(bList) - 1; i++ {
			num2 := bList[i]
			if (i > 0 && num2 == bList[i - 1]) {
				continue
			}
			for j := i + 1; j < len(bList); j++ {
				num3 := bList[j]
				if j > i + 1 && num3 == bList[j - 1] {
					continue
				}
				if _, ok := aMap[-(num3 + num2)]; ok {
					ret = append(ret, []int{-(num3 + num2), num2, num3})
				}
			}
		}

		if aMap[0] >= 3 {
			ret = append(ret, []int{0, 0, 0})
		}
	}

	return ret
}

func ThreeSum(nums []int) [][]int {

	if len(nums) < 3 {
		return nil
	}

	abcSlice := make([][]int, 0)

	//1 对输入数组进行排序
	sortfunc := func(arr []int) {
		for i := 0; i < len(arr) - 1; i++ {
			for j := 0; j < len(arr) - 1 - i; j++ {
				if arr[j] > arr[j + 1] {
					arr[j], arr[j + 1] = arr[j + 1], arr[j]
				}
			}
		}
	}
	sortfunc(nums)

	//进行检索
	/*
	遍历数组中每一个元素，
	在遍历每一个元素的同时，在循环内部固定该元素作为可能元组(a+b+c=0)中的一个元素，对其后面元素进行查抄，查看是否能找出和这个元素相加等于0的其他两个元素
	 */
	for i := 0; i < len(nums) - 2; i ++ {
		//遍历元素，已大于0无需再往后遍历
		if nums[i] > 0 {
			break
		}

		//由于该数组是排序数组，遍历作为元组中的元素不能重复选取，直接跳过
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		sum := 0 - nums[i]
		beginIndex := i + 1
		endIndex := len(nums) - 1
		for endIndex > beginIndex {
			if nums[beginIndex] + nums[endIndex] == sum {
				abcSlice = append(abcSlice, []int{-sum, nums[beginIndex], nums[endIndex]})

				//对于后面两个元素的选取 也需要排重 选取过的元素不能重复选取，直接跳过
				for endIndex > beginIndex && nums[beginIndex] == nums[beginIndex + 1] {
					beginIndex++
				}
				for endIndex > beginIndex && nums[endIndex] == nums[endIndex - 1] {
					endIndex --
				}

				beginIndex++
				endIndex--
			} else if nums[beginIndex] + nums[endIndex] > sum {
				endIndex --
			} else {
				beginIndex++
			}
		}
	}

	return abcSlice
}


