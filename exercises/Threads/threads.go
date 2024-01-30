package main








import (
	"fmt"
	"time"
)

func Foo(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, s)
	}
}

func main() {
	go Foo("1st Thread")
	go Foo("2nd thread")

	time.Sleep(time.Second)     // makes sure that main thread wont be stopped before other threads are start
}
