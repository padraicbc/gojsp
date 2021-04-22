package main

import (
	"bytes"
	"embed"
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

//go:embed javascript_parser_base.go_
var parserbase []byte

//go:embed javascript_lexer_base.go_
var lexerbase []byte
var base embed.FS

// fixes based on suggested here https://github.com/antlr/grammars-v4/tree/master/javascript/javascript/Go
// g4 files already have fixes
func main() {

	log.SetFlags(log.Llongfile)
	_, fname, _, _ := (runtime.Caller(0))

	mnpth := filepath.Dir(fname) + "/"
	// 	// javascript_lexer.go
	// p *JavaScriptLexer => l *JavaScriptLexer
	// antlr.ParserBase => antlr.BaseParser
	l4 := flag.String("l4", mnpth+"JavaScriptLexer.g4", "-l4 path to lexerut")

	p4 := flag.String("p4", mnpth+"JavaScriptParser.g4", "-p path to parser")
	path := flag.String("p", "", "-p path to antlr output")
	flag.Parse()
	*path = strings.Trim(*path, "/ ")

	antlr4 := []string{"-jar", mnpth + "antlr-4.9.2-complete.jar", "-Dlanguage=Go",
		"-visitor", "-no-listener", *l4, *p4, "-Xexact-output-dir", "-o", *path,
		"-package", filepath.Base(*path)}
	cmd := exec.Command("/usr/bin/java", antlr4...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	jsLexerBytes, err := os.ReadFile(*path + "/javascript_lexer.go")
	if err != nil {
		log.Println(err)
		return
	}
	// use l for listener
	ptol := regexp.MustCompile(`p\s+\*JavaScriptLexer\b`)
	// correct name
	bp := regexp.MustCompile(`\bantlr.ParserBase\b`)
	// remove pointer *
	jsp := regexp.MustCompile(`\*JavaScriptLexerBase\b`)

	jsLexerBytes = jsp.ReplaceAll(jsLexerBytes, []byte(`JavaScriptLexerBase`))

	jsLexerBytes = ptol.ReplaceAll(jsLexerBytes, []byte("l *JavaScriptLexer"))
	jsLexerBytes = bp.ReplaceAll(jsLexerBytes, []byte(`antlr.BaseParser`))
	fo, err := os.CreateTemp(".", "")
	if err != nil {
		log.Println(err)
		return
	}

	if _, err = fo.Write(jsLexerBytes); err != nil {
		log.Println(err)
		return
	}
	if err = os.Rename(fo.Name(), *path+"/javascript_lexer.go"); err != nil {

		log.Println(err)
		return

	}

	if err = os.WriteFile(*path+"/javascript_parser_base.go",
		bytes.ReplaceAll(parserbase, []byte("package parser"),
			[]byte("package "+*path)), os.ModePerm); err != nil {

		log.Println(err)
		return

	}
	if err = os.WriteFile(*path+"/javascript_lexer_base.go",
		bytes.ReplaceAll(lexerbase, []byte("package parser"),
			[]byte("package "+*path)), os.ModePerm); err != nil {

		log.Println(err)
		return

	}

	var files = []string{
		"javascriptparser_visitor.go",
		"javascriptparser_base_visitor.go",
		"javascript_lexer.go",
		"javascript_parser.go",
	}

	for _, fl := range files {
		log.Println(*path + "/" + fl)
		f, err := os.ReadFile(*path + "/" + fl)
		if err != nil {
			log.Println(err)
			return
		}
		repl := bytes.ReplaceAll(f,
			[]byte(`"github.com/antlr/antlr4/runtime/Go/antlr"`),
			[]byte(`antlr "github.com/padraicbc/antlr4"`))
		fo, err := os.CreateTemp(".", "")
		if err != nil {
			log.Println(err)
			return
		}
		fo.Write(repl)
		if err = os.Rename(fo.Name(), *path+"/"+fl); err != nil {

			log.Println(err)
			return

		}

	}

}
