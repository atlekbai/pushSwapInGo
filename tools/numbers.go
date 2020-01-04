/*
* @Author: Tlekbai Ali
* @Date:   2020-01-04 14:34:35
* @Last Modified by:   Tlekbai Ali
* @Last Modified time: 2020-01-04 15:58:17
*/

package tools

/*
** Function returns maximum of given numbers
 */
func Max(numbers ...int) int {
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

/*
** Function returns minimum of given numbers
 */
func Min(numbers ...int) int {
	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}

/*
** Function checks for duplicates in sorted array
 */
func CheckDuplicates(array []int) bool {
	for i := 1; i < len(array); i++ {
		if array[i-1] == array[i] {
			return false
		}
	}
	return true
}

/*
** Function returns true if given string is a valid number.
** Otherwise false
 */
func CheckNumber(str string) bool {
	minus := false
	if len(str) == 0 {
		return false
	}
	if str[0] == '-' && len(str) == 1 {
		return false
	}
	i := 0
	for str[i] == '-' {
		if minus == true && str[i] == '-' {
			return false
		}
		if str[i] == '-' {
			minus = true
		}
	}
	for str[i] >= '0' && str[i] <= '9' {
		i++
	}
	if i != len(str) {
		return false
	}
	return true
}