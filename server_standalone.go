// +build !appengine

package main

func main() {
	e.Logger.Fatal(e.Start(":8080"))
}

