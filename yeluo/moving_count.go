// 创建人：lg
// 创建时间：2020/9/21
package main

func movingCount(m int, n int, k int) int {
	flag := make([][]int, m)
	res := 0
	for i := 0; i < m; i++ {
		flag[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i >= 0 && flag[i][j] != 1) || (j >= 0 && flag[i][j] != 1) {
				tmp := getStep(i, j)
				if tmp <= k {
					res++
					flag[i][j] = 1
				}
			}
		}
	}
	return res
}

func getStep(a, b int) int {
	res := 0

	for a > 0 {
		res += a % 10
		a /= 10
	}

	for b > 0 {
		res += b % 10
		b /= 10
	}

	return res
}


func main() {
	println(movingCount(10,20,10))
}

