package other

func do() {
	panic("test") // want "don't use `panic` in this package."
}

func main() {
	panic("test") // want "don't use `panic` in this package."
}
