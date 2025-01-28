package clipboard_test

import (
	"fmt"

	. "github.com/kawana77b/go-clip/internal/clipboard"
)

func Example() {
	WriteAll("日本語")
	text, _ := ReadAll()
	fmt.Println(text)

	// Output:
	// 日本語
}
