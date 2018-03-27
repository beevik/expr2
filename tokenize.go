package main

import (
	"log"
	"math/big"
	"unicode/utf8"
)

const eof = 0

type fullRune struct {
	val rune
	buf []byte
}

type exprLex struct {
	line    []byte
	text    []byte
	current fullRune
	empty   bool
}

func newExprLex(line []byte) *exprLex {
	return &exprLex{line: line, empty: true}
}

func (l *exprLex) advance() (r fullRune) {
	if len(l.line) > 0 {
		var size int
		r.val, size = utf8.DecodeRune(l.line)
		r.buf, l.line = l.line[:size], l.line[size:]
		if r.val == utf8.RuneError {
			log.Print("invalid utf8")
			r = l.advance()
		}
	}
	return
}

func (l *exprLex) getc() rune {
	if l.current.val != eof {
		l.text = append(l.text, l.current.buf...)
	}
	l.current = l.advance()
	return l.current.val
}

func (l *exprLex) initc() rune {
	if l.empty {
		l.empty = false
		return l.getc()
	}
	return l.current.val
}

func (l *exprLex) clearBuf() {
	l.text = l.text[:0]
}

func (l *exprLex) Error(e string) {
	log.Fatal(e)
}

func (l *exprLex) Lex(lval *exprSymType) int {
	c := l.initc()

yystate0:

	l.clearBuf()

	goto yystart1

	goto yystate1 // silence unused label error
yystate1:
	c = l.getc()
yystart1:
	switch {
	default:
		goto yyabort
	case c == '.':
		goto yystate3
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate2
	case c >= '0' && c <= '9':
		goto yystate8
	}

yystate2:
	c = l.getc()
	switch {
	default:
		goto yyrule1
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate2
	}

yystate3:
	c = l.getc()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate4
	}

yystate4:
	c = l.getc()
	switch {
	default:
		goto yyrule2
	case c == 'E' || c == 'e':
		goto yystate5
	case c >= '0' && c <= '9':
		goto yystate4
	}

yystate5:
	c = l.getc()
	switch {
	default:
		goto yyabort
	case c == '+' || c == '-':
		goto yystate6
	case c >= '0' && c <= '9':
		goto yystate7
	}

yystate6:
	c = l.getc()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate7
	}

yystate7:
	c = l.getc()
	switch {
	default:
		goto yyrule2
	case c >= '0' && c <= '9':
		goto yystate7
	}

yystate8:
	c = l.getc()
	switch {
	default:
		goto yyrule2
	case c == '.':
		goto yystate4
	case c == 'E' || c == 'e':
		goto yystate5
	case c >= '0' && c <= '9':
		goto yystate8
	}

yyrule1: // [ \t\r\n]+

	goto yystate0
yyrule2: // {float}
	{

		lval.num = &big.Rat{}
		_, ok := lval.num.SetString(string(l.text))
		if !ok {
			log.Printf("bad number %q", l.text)
			return 0
		}
		return NUM
	}
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	l.empty = true
	return int(c)
}
