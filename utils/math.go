package utils

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Lcm(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	result := numbers[0]
	for _, num := range numbers[1:] {
		result = (result * num) / gcd(result, num)
	}
	return result
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Digits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return digits
}
