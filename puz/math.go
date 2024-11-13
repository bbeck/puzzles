package puz

import "math"

// Abs returns the absolute value of the provided integer or float.
func Abs[T Integer | Float](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

// Sign returns the sign of the provided number.
func Sign[T Integer](n T) T {
	if n == 0 {
		return 0
	}
	if n < 0 {
		return -n / n
	}
	return n / n
}

// Sum adds together each of the provided elements.
func Sum[T Number](elems ...T) T {
	var sum T
	for _, elem := range elems {
		sum += elem
	}
	return sum
}

// Product multiplies together each of the provided elements.
// If no elements are passed in as arguments then 1 is returned.
func Product[T Number](elems ...T) T {
	var product T = 1
	for _, elem := range elems {
		product *= elem
	}
	return product
}

// Min returns the smallest element of the provided elements.
func Min[T Ordered](elems ...T) T {
	min := elems[0]
	for _, elem := range elems[1:] {
		if elem < min {
			min = elem
		}
	}
	return min
}

// Max returns the largest element of the provided elements.
func Max[T Ordered](elems ...T) T {
	max := elems[0]
	for _, elem := range elems[1:] {
		if elem > max {
			max = elem
		}
	}
	return max
}

func Pow[T Number, U Unsigned](base T, exp U) T {
	var result T = 1
	for {
		if exp&1 == 1 {
			result *= base
		}
		exp >>= 1

		if exp == 0 {
			break
		}

		base *= base
	}

	return result
}

func ModPow[T Integer](base, exp, mod T) T {
	base = base % mod

	var result T = 1
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		exp >>= 1
		base = (base * base) % mod
	}

	return result
}

// Digits returns the individual digits of the provided number in order from
// left to right.  If the provided number is negative, then the sign will be
// ignored.
func Digits[T Integer](n T) []T {
	n = Abs(n)
	if n < 10 {
		return []T{n}
	}

	length := 1 + int(math.Floor(math.Log10(float64(n))))
	digits := make([]T, length)
	for i := len(digits) - 1; n > 0; i-- {
		digits[i] = n % 10
		n /= 10
	}
	return digits
}

// JoinDigits interprets the provided digits as a number.
func JoinDigits[T Integer](ds []T) T {
	var n T
	for _, d := range ds {
		n = n*10 + d
	}
	return n
}

// ChineseRemainderTheorem solves a system of congruences for x.
//
//	x = a_i (mod n_i) for all i
//
// See: https://shainer.github.io/crypto/math/2017/10/22/chinese-remainder-theorem.html
func ChineseRemainderTheorem(as, ns []int) int {
	N := 1
	for _, n := range ns {
		N *= n
	}

	result := 0
	for i := 0; i < len(as); i++ {
		ai, ni := as[i], ns[i]
		_, _, si := ExtendedEuclid(ni, N/ni)
		result += ai * si * (N / ni)
	}

	for result < 0 {
		result = result + N
	}
	return result % N
}

// ExtendedEuclid computes the GCD of x and y as well as coefficients a and b
// such that:
//
//	a*x + b*y = gcd(a, b).
func ExtendedEuclid(x, y int) (int, int, int) {
	x0, x1, y0, y1 := 1, 0, 0, 1

	var q int
	for y > 0 {
		q, x, y = x/y, y, x%y
		x0, x1 = x1, x0-q*x1
		y0, y1 = y1, y0-q*y1
	}

	return q, x0, y0
}
