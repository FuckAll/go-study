package algorithms

import (
	"testing"
)

// 寻找素数：只有1和它本身能够被他所除的叫做素数

// 查找出[2,n)中有多少个素数
func countPrimes(n int) int {
	var count int
	for i := 2; i < n; i++ {
		//if !isPrimes(i) {
		if !isPrimesNG(i) {
			continue
		}
		count++
	}
	return count
}

func isPrimes(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//
//12 = 2 × 6
//12 = 3 × 4
//12 = sqrt(12) × sqrt(12)
//12 = 4 × 3
//12 = 6 × 2
// 只需要sqrt(12)的时候就可以判断是否是素数了，不需要后面
func isPrimesNG(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func TestCountPrimes(t *testing.T) {
	primes := countPrimes(3)
	if primes != 1 {
		t.Errorf("want 1 but get:%v", primes)
		return
	}

	primes = countPrimes(4)
	if primes != 2 {
		t.Errorf("want 2 but get:%v", primes)
		return
	}

	primes = countPrimes(10)
	if primes != 4 {
		t.Errorf("want 4 but get:%v", primes)
	}
}

// 更进阶的一种写法
func countPrimesNG(n int) int {
	isPrimes := make([]bool, n)
	for i := range isPrimes {
		isPrimes[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrimes[i] {
			for j := i * 2; j < n; j += i {
				isPrimes[j] = false
			}
		}
	}
	var count int
	for i := 2; i < n; i++ {
		if isPrimes[i] {
			count++
		}
	}
	return count
}

func TestCountPrimesNG(t *testing.T) {
	primes := countPrimesNG(3)
	if primes != 1 {
		t.Errorf("want 1 but get:%v", primes)
		return
	}

	primes = countPrimesNG(4)
	if primes != 2 {
		t.Errorf("want 2 but get:%v", primes)
		return
	}

	primes = countPrimesNG(10)
	if primes != 4 {
		t.Errorf("want 4 but get:%v", primes)
	}
}

func countPrimesNGNG(n int) int {
	isPrimes := make([]bool, n)
	for i := range isPrimes {
		isPrimes[i] = true
	}

	for i := 2; i*i < n; i++ {
		if isPrimes[i] {
			for j := i * i; j < n; j += i {
				isPrimes[j] = false
			}
		}
	}

	var count int
	for i := 2; i < n; i++ {
		if isPrimes[i] {
			count++
		}
	}

	return count

}

func TestCountPrimesNGNG(t *testing.T) {
	primes := countPrimesNGNG(3)
	if primes != 1 {
		t.Errorf("want 1 but get:%v", primes)
		return
	}
	primes = countPrimesNGNG(4)
	if primes != 2 {
		t.Errorf("want 2 but get:%v", primes)
		return
	}
	primes = countPrimesNGNG(10)
	if primes != 4 {
		t.Errorf("want 4 but get:%v", primes)
	}
}
