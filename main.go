package main

import (
	"bytes"
	"embed"
	"flag"
	"log"
	"os"
	"regexp"
	"strings"
)

//go:embed javascript_parser_base.go_
var parserbase []byte

//go:embed javascript_lexer_base.go_
var lexerbase []byte
var base embed.FS

// fixes based on suggst here https://github.com/antlr/grammars-v4/tree/master/javascript/javascript/Go
// g4 files already have fixes
func main() {

	log.SetFlags(log.Llongfile)

	// 	// javascript_lexer.go
	// p *JavaScriptLexer => l *JavaScriptLexer
	// antlr.ParserBase => antlr.BaseParser
	path := flag.String("p", "", "-p path to antlr output")
	flag.Parse()
	*path = strings.Trim(*path, "/ ")

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
	jsLexerBytes = bytes.ReplaceAll(jsLexerBytes, []byte("package parser"), []byte("package "+*path))
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
		repl = bytes.ReplaceAll(repl, []byte("package parser"), []byte("package "+*path))
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
