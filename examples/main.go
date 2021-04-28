package main

import (
	"fmt"
	"log"
	"reflect"
	"unsafe"

	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/gojsp/vast"
)

func main() {

	log.SetFlags(log.Llongfile)
	// coldHot()

	impexp()
	// fs()

	// arrow()
	toes5()
	// singleExp()
}

func visit(tree antlr.ParseTree, v *vast.Visitor) interface{} {

	return v.Visit(tree)

}

func refT(t vast.VNode) {
	s := reflect.ValueOf(t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		f = reflect.NewAt(f.Type(), unsafe.Pointer(f.UnsafeAddr())).Elem()
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
