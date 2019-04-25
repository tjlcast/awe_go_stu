package main

func main() {
	var p = f()

}

func f() *int {
	v := 1
	return &v
}
