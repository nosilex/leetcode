/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

	Input: strs = ["flower","flow","flight"]
	Output: "fl"

Example 2:

	Input: strs = ["dog","racecar","car"]
	Output: ""
	Explanation: There is no common prefix among the input strings.
*/

package leetcode

func longestCommonPrefix(strs []string) string {
    if len(strs) == 1 {
        return strs[0]
    }

    var out string
    for k, v := range strs {
        if k == 0 {
            out = v
            continue
        }
        if len(out) > len(v) {
            out = out[:len(v)]
        }
        for i := 0; i < len(v); i++ {
            if i > len(out) - 1 {
                break
            }
            if v[i] != out[i] {
                out = out[:i]
                break
            }
        }
    }

    return out
}