package main

import "fmt"

func _main () {
	var isOnline bool
	fmt.Printf("b = %v\n", isOnline);
	update(&isOnline)
	fmt.Printf("b = %v\n", isOnline);
}

func update(b *bool) {
	*b = true
}
