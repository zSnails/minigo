.PHONY: grammar
grammar:
	antlr4 -Dlanguage=Go -o grammar Minigo.g4 -package grammar -visitor -listener

.PHONY: qbe_build
qbe_build:
	qbe -o out.s file.ssa && cc out.s
