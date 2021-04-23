## Very early attempt at creating an [ANTLR](https://www.antlr.org/) visitor for javascript files using GO. Going to change like Irish weather so don't expect things to stay the same any time soon. 


The `base` is the result of running what is described [here](https://github.com/padraicbc/gojsp/tree/master/runantlr#readme)

The [vast](https://github.com/padraicbc/gojsp/tree/master/vast) folder is the implementation so far toward creating an ast which will allow easy manipulation/translation of the original source.
Thre are a few example functions in the exampls folder including how to parse and [manipulate](https://github.com/padraicbc/gojsp/blob/master/examples/singleexpress.go) a single expression. 
