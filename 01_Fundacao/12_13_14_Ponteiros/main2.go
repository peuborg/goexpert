package main

func main() {
	var1 := 5
	var2 := 10
	println(soma(&var1, &var2))
	println(var1)
	println(var2)
}

func soma(a, b *int) int {
	*a = 20
	*b = 30
	return *a + *b
}
