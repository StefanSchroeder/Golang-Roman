[![GoDoc](https://godoc.org/github.com/StefanSchroeder/Golang-Roman?status.png)](https://godoc.org/github.com/StefanSchroeder/Golang-Roman)
[![Build Status](https://travis-ci.org/StefanSchroeder/Golang-Roman.svg?branch=master)](https://travis-ci.org/StefanSchroeder/Golang-Roman)
[![Go Report Card](http://goreportcard.com/badge/StefanSchroeder/Golang-Roman)](http://goreportcard.com/report/StefanSchroeder/Golang-Roman)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Golang-Roman
============

A small Go-package for Roman Numbers

	package main
	import roman "github.com/StefanSchroeder/Golang-Roman"
	import "fmt"
	func main() {
		fmt.Printf("2+2=%v\n" , roman.Roman(2+2))
		// Prints 2+2=IV
		fmt.Printf("IX+IX=%v\n" , roman.Arabic("XVIII"))
		// Prints IX+IX=18
	}

Added Travis.

## Changelog:

1.0: Initial

2.0: Changed Roman() to return tupel (incl. error). Thanks to Jim Eagle for the contribution.

