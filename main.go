package main

func main() {
	// hello world
	const a = "Hello,World"
	println(a)

	// 奇数と偶数
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			println(i, "-偶数")
		} else if i%2 == 1 {
			println(i, "-奇数")
		}
	}
}
