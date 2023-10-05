package main

// import "fmt"
// func f(xp *int) {
// 	*xp = 100
// }

// 値を入れ替えるswap関数の実装
func swap(x int, y int) (int, int) {
	x, y = y, x
	return x, y
}

// 値を入れ替えるswap2関数の実装
func swap2(x, y *int) (int, int){
	x, y = y, x
	return *x, *y
}

func main() {
	// // hello world
	// const a = "Hello,World"
	// println(a)

	// // 奇数と偶数
	// for i := 1; i <= 100; i++ {
	// 	if i%2 == 0 {
	// 		println(i, "-偶数")
	// 	} else if i%2 == 1 {
	// 		println(i, "-奇数")
	// 	}
	// }

	// // 構造体
	// p := struct {
	// 	name string
	// 	age int
	// }{name: "suzuki masashi", age: 42}
	// // {name: "kizawa daisuke",age: 48}
	// println(p.name, p.age)

	// ns := [...]int{10, 20, 30, 40, 50}
	// // 要素にアクセス
	// println(ns[3]) // 添字は変数でもよい
	// // 長さ
	// println(len(ns))
	// // スライス演算
	// fmt.Println(ns[1:2])

	// 値を入れ替えるswap関数の実装
	n, m := swap(10, 20)
	println(n, m)

	// println(swapp(10,20))

	// 値を入れ替えるswap2関数
	x, y := 10, 20
	swap2(&x, &y)
	println(x, y)


}
