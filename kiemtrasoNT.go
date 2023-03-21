package main

import "fmt"

func kiemtra(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Print("Nhập số lượng phần tử có trong mảng: ")
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for _, x := range a {
		if kiemtra(x) {
			fmt.Println(x, "là số nguyên tố")
		}
		fmt.Println(x, "không phải là số nguyên tố")
	}
}
