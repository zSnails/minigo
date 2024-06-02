grammar: Minigo.g4
	antlr4 -Dlanguage=Go -o grammar Minigo.g4 -package grammar -visitor -listener
