package hard

/*
805. Split Array With Same Average

In a given integer array A, we must move every element of A to either list B or list C. (B and C initially start empty.)

Return true if and only if after such a move, it is possible that the average value of B is equal to the average value of C, and B and C are both non-empty.

Example :
Input:
[1,2,3,4,5,6,7,8]
Output: true
Explanation: We can split the array into [1,4,5,8] and [2,3,6,7], and both of them have the average of 4.5.
Note:

The length of A will be in the range [1, 30].
A[i] will be in the range of [0, 10000].

 */

/* Approach #1: Meet in the Middle [Accepted]

First, let's get a sense of the condition that average(B) = average(C), where B, C are defined in the problem statement.

Say A (the input array) has N elements which sum to S, and B (one of the splitting sets) has K elements which sum to X.
Then the equation for average(B) = average(C) becomes X/K = S-X/N-K. This reduces to X(N-K) = (S-X)K which is X/K = S/N. That is, average(B) = average(A).
Now, we could delete average(A) from each element A[i] without changing our choice for B.
(A[i] -= mu, where mu = average(A)). This means we just want to choose a set B that sums to 0.

 */

func splitArraySameAverage(A []int) bool {


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
	sortfunc(A)

	var sum int = 0
	for i := 0; i < len(A); i++  {
		sum += A[i]
	}
	average := sum / len(A)

	var index int = 0
	for i := 0; i < len(A); i++  {
		if average >= A[i] {
			index = i
			break
		}
	}

	B := make([]int, 0, len(A))
	C := make([]int, 0, len(A))

	leftindex := index
	rightindex := index + 1
	for leftindex > -1 && rightindex < len(A) {

		if (A[leftindex] + A[rightindex]) / 2 == average {
			if len(B) > len(C) {
				C = append(C, A[leftindex], A[rightindex])
			}else {
				B = append(B, A[leftindex], A[rightindex])
			}
		}else {

		}

	}




}
