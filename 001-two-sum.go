package leetcode

func twoSum(nums []int, target int) []int {
    var result []int
    if len(nums) == 2 && nums[0] + nums[1] == target {
        result = append(result, 0)
        result = append(result, 1)
        return result
    }
    for i := 0; i < len(nums); i++ {
        for j := 0; j < len(nums); j++ {
            if i == j {
                continue
            }
            if nums[i] + nums[j] == target {
                result = append(result, i)
                result = append(result, j)
                return result
            }
        }
    }

    return result
}