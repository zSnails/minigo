.PHONY: grammar
grammar:
	antlr4 -Dlanguage=Go -o grammar Minigo.g4 -package grammar -visitor -listener
