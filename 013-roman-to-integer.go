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

func romanToInt(s string) int {
    mapsv := map[string]int {
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,
    }

    result := 0
    for i := 0; i < len(s); i++ {
        cc := string(s[i])
        vc, ok := mapsv[cc]
        if !ok {
            continue
        }

        if i+1 < len(s) {
            cn := string(s[i+1])
            vn, ok := mapsv[cn]; 
            if !ok {
                continue
            }

            if (cc == "I" && (cn == "V" || cn == "X")) || (cc == "X" && (cn == "L" || cn == "C")) || (cc == "C" && (cn == "D" || cn == "M")) {
                result += vn - vc
                i++
                continue
            }
        }

        result += vc
    }

    return result
}