package solution

import "strconv"

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    
    s := strconv.Itoa(x)
    l := len(s)-1
    for i := 0; i <= l; i++ {
        if s[i] != s[l-i]  {
            return false
        }
    }
    return true
}