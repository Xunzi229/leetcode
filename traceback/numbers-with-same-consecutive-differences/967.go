package numbers_with_same_consecutive_differences

func numsSameConsecDiff(n int, k int) []int {
	if n <= 0 {
		return nil
	}

	res := make([]int, 0)

	// sn 剩余长度, num 目前数
	var backtrace = func(sn, num int) {}

	backtrace = func(sn, num int) {
		if sn == 0 {
			res = append(res, num)
			return
		}

		for i := 0; i <= 9; i++ {
			if sn == n && i == 0 {
				continue
			}

			if sn == n {
				num = i
				sn--
				backtrace(sn, i)
				sn++
				num = 0
				continue
			}

			if (num%10)-i == k || i-(num%10) == k {
				num = num*10 + i
				sn--
				backtrace(sn, num)
				sn++
				num = (num - i) / 10
			}
		}
	}
	backtrace(n, 0)
	return res
}
