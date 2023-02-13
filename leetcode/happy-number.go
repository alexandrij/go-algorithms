package algorithms

func IsHappy(n int) bool {
	seen := map[int]bool{}
	for n != 1 {
		if _, ok := seen[n]; ok {
			break
		}
		seen[n] = true

		var sum int
		for n > 0 {
			d := n % 10
			sum += d * d
			n = (n - d) / 10
		}
		n = sum
	}
	return n == 1
}
