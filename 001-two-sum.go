/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Example 1:

    Input: nums = [2,7,11,15], target = 9
    Output: [0,1]
    Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:

    Input: nums = [3,2,4], target = 6
    Output: [1,2]

Example 3:

    Input: nums = [3,3], target = 6
    Output: [0,1]
*/

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