// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

// Injectors from wire.go:

func inject(foo *Foo) *Bar {
	bar := NewBar(foo)
	return bar
}
