package main // Chamando o pagote main

import "runtime/debug"

func f3() {

	// usar um pacote do Go debug.PrintStack()
	debug.Stack()
}