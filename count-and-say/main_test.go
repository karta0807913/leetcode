package main

func _countAndSay(n int) []byte {
	if n == 1 {
		return []byte{1 | 0x30}
	}
	say := _countAndSay(n - 1)
	var ans []byte
	x := say[0]
	count := 0
	for _, val := range say {
		if x != val {
			ans = append(ans, byte(count|0x30))
			ans = append(ans, x)
			count = 0
			x = val
		}
		count += 1
	}
	ans = append(ans, byte(count|0x30))
	ans = append(ans, x)
	return ans
}

func countAndSay(n int) string {
	return string(_countAndSay(n))
}
