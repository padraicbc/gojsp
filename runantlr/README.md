# Recreate the `base` parser using `JavaScriptLexer.g4` and `JavaScriptParser.g4`

This post processes the outputted go files which have a few errors relating to naming.
You will need `java` installed. You can replace the lexer and parser files with your own variations.
These are going to continually change so don't expect consistency. 
The grammar files have fixes based on [these suggestions](https://github.com/antlr/grammars-v4/tree/master/javascript/javascript/Go) so you'll need to do the same if you use the original grammar files from [here](https://github.com/antlr/grammars-v4/tree/master/javascript/javascript).

### To Run:
  
* Using these grammar files:

    `go run path_to/runantlr/main.go  -p output_path/whatever`

* Pass your own grammar files:
 
    `go run path_to/runantlr/main.go  -p output_path/whatever -lx path/JavaScriptLexer.g4 -p4 path/JavaScriptParser.g4` 

`p` is the output path and the base will become the package name so `/home/foo/jp` would create the base in `/home/foo/jp` and the package would be  `package jp`




