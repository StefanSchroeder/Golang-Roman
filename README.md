Golang-Roman
============

[![Build Status](https://travis-ci.org/StefanSchroeder/Golang-Roman.png?branch=master)](https://travis-ci.org/StefanSchroeder/Golang-Roman)


A small Go-package for Roman Numbers

	package main
	import "roman"
	import "fmt"
	func main() {
		fmt.Printf("2+2=%v\n" , roman.Roman(2+2))
		fmt.Printf("IX+IX=%v\n" , roman.Arabic("XVIII"))
	// fmt.Printf("is IX a Roman number?: %v\n" , roman.IsRoman("IX"))
	}

Added Travis.
