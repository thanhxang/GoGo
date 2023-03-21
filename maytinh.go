package main

import "fmt"

func nhap(a, b int) int {
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	return 0
}

func main() {
	var a int
	var b int
	var pheptinh int = 1
	fmt.Println("Nhập số a đi:")
	fmt.Scan(&a)
	fmt.Println("Nhập số b đi:")
	fmt.Scan(&b)
	fmt.Println("Chọn phép tính đi nè:")
	fmt.Println("1.Cộng")
	fmt.Println("2.Trừ")
	fmt.Println("3.Nhân")
	fmt.Println("4.Chia")
	fmt.Println("5.Out")
	fmt.Scan(&pheptinh)

	//for pheptinh := 1; pheptinh < 5; pheptinh++ {
	//	if pheptinh == 5 {
	//		fmt.Println("Chọn sai số rồi. Chọn từ 1-4 thôi")
	//		pheptinh = pheptinh
	//	} else if pheptinh == 1 {
	//		fmt.Println("a cong b = ", a+b)
	//	} else if pheptinh == 2 {
	//		fmt.Println("a tru b = ", a-b)
	//	} else if pheptinh == 3 {
	//		fmt.Println("a nhan b = ", a*b)
	//	} else if pheptinh == 4 {
	//		fmt.Println("a chia b = ", a/b)
	//	}

	//if pheptinh == 1 {
	//	fmt.Println("a cong b = ", a+b)
	//} else if pheptinh == 2 {
	//	fmt.Println("a tru b = ", a-b)
	//} else if pheptinh == 3 {
	//	fmt.Println("a nhan b = ", a*b)
	//} else if pheptinh == 4 {
	//	fmt.Println("a chia b = ", a/b)
	//} else if pheptinh == 5 {
	//	fmt.Println("Chọn sai số rồi. Chọn [1-4] thôi!"
	//}

	switch pheptinh {
	case 1:
		fmt.Println("a cộng b = ", a+b)
	case 2:
		fmt.Println("a trừ b = ", a-b)
	case 3:
		fmt.Println("a nhân b = ", a*b)
	case 4:
		fmt.Println("a chia b = ", a/b)
	case 5:
		fmt.Println("Hay muốn nhập lại A với B?")
	default:
		fmt.Println("Chọn sai số rồi. Chọn [1-4] thôi!")
	}

}
