package main

func main() {
	e := NewApp()

	e.Logger.Fatal(e.Start(":3000"))
}
