%{
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
%}

%yyc c
%yyn c = l.getc()

D           [0-9]+
E           [eE][-+]?{D}
float       {D}"."{D}?{E}?|{D}{E}?|"."{D}{E}?

%%
    l.clearBuf()

[ \t\r\n]+

{float}
    lval.num = &big.Rat{}
    _, ok := lval.num.SetString(string(l.text))
    if !ok {
        log.Printf("bad number %q", l.text)
        return 0
    }
    return NUM

%%
    l.empty = true
    return int(c)
}