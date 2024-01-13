package leetcode

// @leet start
func merge(nums1 []int, m int, nums2 []int, n int) {
	// NOTE: First attempt, ignore

	// nums1 = [1,2,3,4,0,0,0,], m = 4
	//                  ^
	// nums2 = [2,5,6], n = 3
	//              ^
	// [1,2,2,3,4,6]

	// nums1 = [1,2,3,0,0,0,0], m = 3
	//                ^
	// nums2 = [3,5,6,7], n = 4
	//          ^
	// [1,2,2,3,5,6,7]

	// nums1 = [1,2,3,0,0,0], m = 3
	//                ^ i = 3
	// nums2 = [3,5,6], n = 3
	//                ^ j = 3
	// [1,2,2,3,5,6]
	//              ^ i = 6

	// nums1 = [0], m = 0
	//          ^ i = 0
	// nums2 = [1], n = 1
	//            ^ j = 1
	// [1]
	//    ^ i = 1

	// nums1 = [4,5,6,0,0,0], m = 3
	//          ^ i = 0
	// nums2 = [4,2,3], n = 3
	//          ^ j = 0
	// [1,5,6,0,0,0]
	//  ^ i = 0
	// NOTE: Swapping into nums2 doesn't work in this case ^

	// Iterate backwards with 3 pointers. i starts at the end of nums1 (length
	// m), j starts at the end of nums2 (length n), and k starts at the end of
	// the result (still nums1 but length m + n). This takes advantage of the
	// extra space at the end of nums1 and allows us to proceed with merge sort
	// logic as per usual and not worry about clobbering the values in the
	// beginning of nums1.

	// nums1 = [1,2,3,0,0,0], m = 3
	//          ^ i = 0
	// nums2 = [2,5,6], n = 3
	//        ^ j = -1
	// [1,2,2,3,5,6]
	//  ^ k = 0

	for i, j, k := m-1, n-1, m+n-1; k >= 0; k-- {
		switch {
		case j < 0:
			return
		case i < 0:
			nums1[k] = nums2[j]
			j--
		case nums1[i] >= nums2[j]:
			nums1[k] = nums1[i]
			i--
		default:
			nums1[k] = nums2[j]
			j--
		}
	}
}

// @leet end

