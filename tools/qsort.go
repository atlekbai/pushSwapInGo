/*
* @Author: Tlekbai Ali
* @Date:   2020-01-04 15:34:40
* @Last Modified by:   Tlekbai Ali
* @Last Modified time: 2020-01-04 16:24:41
*/

package tools

func partition(array []int, low, high int) int {
	pivot := array[low]
	j := low
	i := low
	for i <= high {
		if array[i] <= pivot {
			array[i], array[j] = array[j], array[i]
			j++
		}
		i++
	}
	array[low], array[j] = array[j], array[low]
	return j
}

/*
** Quick Sort algorithm implementation
 */
func Qsort(array []int, low, high int) {
	if low >= high {
		return
	}
	mid := partition(array, low, high)
	Qsort(array, low, mid-1)
	Qsort(array, mid+1, high)
}
