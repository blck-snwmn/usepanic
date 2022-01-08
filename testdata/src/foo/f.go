package foo

func panic(v interface{}) { // want "don't use `panic` in this package."
}

func do() {
	panic("test") // want "don't use `panic` in this package."
}
