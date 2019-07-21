package hard

/*
639. Decode Ways II

https://leetcode.com/problems/decode-ways-ii/description/

A message containing letters from A-Z is being encoded to numbers using the following mapping way:

'A' -> 1
'B' -> 2
...
'Z' -> 26
Beyond that, now the encoded string can also contain the character '*', which can be treated as one of the numbers from 1 to 9.

Given the encoded message containing digits and the character '*', return the total number of ways to decode it.

Also, since the answer may be very large, you should return the output mod 109 + 7.

Example 1:
Input: "*"
Output: 9
Explanation: The encoded message can be decoded to the string: "A", "B", "C", "D", "E", "F", "G", "H", "I".
Example 2:
Input: "1*"
Output: 9 + 9 = 18
Note:
The length of the input string will fit in range [1, 105].
The input string will only contain the character '*' and digits '0' - '9'.

 */

func NumDecodings(s string) int {
	return numDecodingsDynamicProgramming(s)
}


/*
From the solutions discussed above, we can observe that the number of decodings possible upto any index,
i, is dependent only on the characters upto the index ii and not on any of the characters following it.
This leads us to the idea that this problem can be solved by making use of Dynamic Programming.

We can also easily observe from the recursive solution that,
the number of decodings possible upto the index i can be determined easily if we know the number of decodings possible upto the index i-1 and i-2.
第i个元素有多少种情况 是根据第i-1,i-2个元素的情况决定的

Thus, we fill in the dp array in a forward manner.
dp[i] is used to store the number of decodings possible by considering the characters in the given string s upto the (i-1)^{th} index only(including it).
我们以向前的方式填充dp数组，dp[i]用于存储第i-1个字符情况的数量

The equations for filling this dp at any step again depend on the current character and the just preceding character.
These equations are similar to the ones used in the recursive solution.

The following animation illustrates the process of filling the dp for a simple example.
 */
/*
当前的元素计算为1个元素还是计算为两个元素 要根据前一个元素和前前一个元素来得出
（一开始做题目陷入的错误思路
一开始计算这道题目的时候，对于第一个元素是看成一个还是看成两个，要根据后一个元素，
但其实如果第二个元素看成一个还是两个，并没有去细想这个点，仔细想他其实是根据前一个来定的，而前一个是一个还是两个又是根据前前一个来定的
这样从后往前思考，只要给出一个起点就可以不断往后计算出来了，并且可以循环
）
而一开始的错误思路，是把当前元素和后面元素是否连起来看成是两种情况，在两种情况的前提下继续计算后面部分的结果，这种递归效率肯定是慢的
 */
func numDecodingsDynamicProgramming(s string) int {
	var m int64 = 1000000007
	var dp []int64 = make([]int64,len(s) + 1)
	dp[0] = 1
	if s[0] == '*' {
		dp[1] = 9
	} else if s[0] == '0' {
		dp[1] = 0
	}else {
		dp[1] = 1
	}

	for i := 1; i < len(s); i++ {
		if s[i] == '*' {
			//当前元素单独为一个当个编码情况
			dp[i + 1] = 9 * dp[i]

			//当前元素和前面元素一起组合成一个编码情况
			if s[i - 1] == '1' {
				dp[i + 1] = (dp[i + 1] + dp[i-1] * 9) % m
			}else if s[i - 1] == '2' {
				dp[i + 1] = (dp[i + 1] + dp[i-1] * 6) % m
			}else if s[i - 1] == '*' {
				dp[i + 1] = (dp[i + 1] + dp[i-1] * 15) % m
			}
		}else {
			if s[i] == '0' {
				dp[i+1] = 0
			}else {
				dp[i+1] = dp[i]
			}

			if s[i - 1] == '1' {
				dp[i + 1] = (dp[i + 1] + dp[i-1]) % m
			}else if s[i - 1] == '2' && s[i] <= '6' {
				dp[i + 1] = (dp[i + 1] + dp[i-1]) % m
			}else if s[i - 1] == '*' {
				if s[i] <= '6' {
					dp[i + 1] = (dp[i + 1] + dp[i-1] * 2) % m
				}else {
					dp[i + 1] = (dp[i + 1] + dp[i-1]) % m
				}
			}
		}
	}
	return int(dp[len(s)])
}

//start 开始下标
//end 结束下标
func numDecodingRecursive(s string, start, end int) int {

	//需要进行再次拆分
	if end - start > 1 {
		g1a := numDecodingResult(s, start, start)
		g1b := numDecodingRecursive(s, start + 1, end)

		g1ret := g1a * g1b
		if g1ret > 1000000000 + 7 {
			g1ret = g1ret % (1000000000 + 7)
		}

		g2a := numDecodingResult(s, start, start + 1)
		g2b := numDecodingRecursive(s, start + 2, end)

		g2ret := g2a * g2b
		if g2ret > 1000000000 + 7 {
			g2ret = g2ret % (1000000000 + 7)
		}

		ret := g1ret + g2ret
		if ret > 1000000000 + 7 {
			ret = ret % (1000000000 + 7)
		}
		return ret
	}else if end - start > 0 {
		g1a := numDecodingResult(s, start, start)
		g1b := numDecodingRecursive(s, start + 1, end)

		g1ret := g1a * g1b
		if g1ret > 1000000000 + 7 {
			g1ret = g1ret % (1000000000 + 7)
		}

		g2ret := numDecodingResult(s, start, start + 1)

		ret := g1ret + g2ret
		if ret > 1000000000 + 7 {
			ret = ret % (1000000000 + 7)
		}
		return ret
	}else if end - start == 0 {
		return numDecodingResult(s, start, start)
	}else {
		return 1
	}
}

func numDecodingResult(s string, start, end int) int {

	if end - start == 0 {
		//传入1个元素时
		if s[start] == '*' {
			return 9
		} else if s[start] == '0' {
			return 0
		}else {
			return 1
		}
	} else {
		//传入2个元素时
		//if end - start == 1
		if s[start] != '*' && s[end] != '*' {
			if s[start] == '1' || (s[start] == '2' && s[end] <= '6') {
				return 1
			} else {
				return 0
			}
		} else if s[start] == '*' && s[end] == '*' {
			return 15
		} else {
			if s[start] == '*' {
				if s[end] <= '6' {
					return 2
				} else {
					return 1
				}
			} else {
				//if s[end] == '*' {

				if s[start] == '1' {
					return 9
				} else if s[start] == '2' {
					return 6
				} else {
					return 0
				}
			}
		}
	}
}

