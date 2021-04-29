package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/padraicbc/gojsp/vast"
)

func it() {
	b, err := ioutil.ReadFile("files/Iterators.js")
	if err != nil {
		log.Fatal(err)
	}
	code := string(b)

	v := vast.NewVisitor(code)
	// v.Debug = true
	// do whatever with errors
	go v.DefaultError()

	tree := v.Parser.Program()
	vp := visit(tree, v).(*vast.Program)
	for _, ch := range vp.Body {
		fmt.Println(ch.Code())
		fmt.Println(ch.Type(), "->", ch.FirstChild().Type(), "->", ch.FirstChild().FirstChild().Type())

		for _, ch2 := range vast.Children(ch.FirstChild().FirstChild()) {
			fmt.Println(ch2.Type(), "->", ch2.Code())
		}
	}

}
