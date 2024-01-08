package main

import "fmt"

func main() {
	msgs := []string{"Hello, World!", "こんにちは，世界",
		"Olá Mundo", "안녕하세요, 월드!",
		"Bonjour le monde!", "你好世界！"}
	for _, m := range msgs {
		fmt.Println(m)
	}
}
