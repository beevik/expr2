//line parse.y:2
package main

import __yyfmt__ "fmt"

//line parse.y:3
import (
	"fmt"
	"math/big"
)

//line parse.y:12
type exprSymType struct {
	yys int
	num *big.Rat
}

const NUM = 57346
const UMINUS = 57347

var exprToknames = []string{
	"NUM",
	" +",
	" -",
	" *",
	" /",
	"UMINUS",
}
var exprStatenames = []string{}

const exprEofCode = 1
const exprErrCode = 2
const exprMaxDepth = 200

//line parse.y:63

//line yacctab:1
var exprExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const exprNprod = 9
const exprPrivate = 57344

var exprTokenNames []string
var exprStates []string

const exprLast = 26

var exprAct = []int{

	6, 7, 8, 9, 2, 1, 16, 0, 10, 11,
	0, 12, 13, 14, 15, 5, 0, 3, 0, 8,
	9, 4, 6, 7, 8, 9,
}
var exprPact = []int{

	11, -1000, 17, 11, 11, -1000, 11, 11, 11, 11,
	-1000, -5, 12, 12, -1000, -1000, -1000,
}
var exprPgo = []int{

	0, 4, 5,
}
var exprR1 = []int{

	0, 2, 1, 1, 1, 1, 1, 1, 1,
}
var exprR2 = []int{

	0, 1, 3, 3, 3, 3, 2, 3, 1,
}
var exprChk = []int{

	-1000, -2, -1, 6, 10, 4, 5, 6, 7, 8,
	-1, -1, -1, -1, -1, -1, 11,
}
var exprDef = []int{

	0, -2, 1, 0, 0, 8, 0, 0, 0, 0,
	6, 0, 2, 3, 4, 5, 7,
}
var exprTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	10, 11, 7, 5, 3, 6, 3, 8,
}
var exprTok2 = []int{

	2, 3, 4, 9,
}
var exprTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var exprDebug = 0

type exprLexer interface {
	Lex(lval *exprSymType) int
	Error(s string)
}

const exprFlag = -1000

func exprTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(exprToknames) {
		if exprToknames[c-4] != "" {
			return exprToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func exprStatname(s int) string {
	if s >= 0 && s < len(exprStatenames) {
		if exprStatenames[s] != "" {
			return exprStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func exprlex1(lex exprLexer, lval *exprSymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = exprTok1[0]
		goto out
	}
	if char < len(exprTok1) {
		c = exprTok1[char]
		goto out
	}
	if char >= exprPrivate {
		if char < exprPrivate+len(exprTok2) {
			c = exprTok2[char-exprPrivate]
			goto out
		}
	}
	for i := 0; i < len(exprTok3); i += 2 {
		c = exprTok3[i+0]
		if c == char {
			c = exprTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = exprTok2[1] /* unknown char */
	}
	if exprDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", exprTokname(c), uint(char))
	}
	return c
}

func exprParse(exprlex exprLexer) int {
	var exprn int
	var exprlval exprSymType
	var exprVAL exprSymType
	exprS := make([]exprSymType, exprMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	exprstate := 0
	exprchar := -1
	exprp := -1
	goto exprstack

ret0:
	return 0

ret1:
	return 1

exprstack:
	/* put a state and value onto the stack */
	if exprDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", exprTokname(exprchar), exprStatname(exprstate))
	}

	exprp++
	if exprp >= len(exprS) {
		nyys := make([]exprSymType, len(exprS)*2)
		copy(nyys, exprS)
		exprS = nyys
	}
	exprS[exprp] = exprVAL
	exprS[exprp].yys = exprstate

exprnewstate:
	exprn = exprPact[exprstate]
	if exprn <= exprFlag {
		goto exprdefault /* simple state */
	}
	if exprchar < 0 {
		exprchar = exprlex1(exprlex, &exprlval)
	}
	exprn += exprchar
	if exprn < 0 || exprn >= exprLast {
		goto exprdefault
	}
	exprn = exprAct[exprn]
	if exprChk[exprn] == exprchar { /* valid shift */
		exprchar = -1
		exprVAL = exprlval
		exprstate = exprn
		if Errflag > 0 {
			Errflag--
		}
		goto exprstack
	}

exprdefault:
	/* default state action */
	exprn = exprDef[exprstate]
	if exprn == -2 {
		if exprchar < 0 {
			exprchar = exprlex1(exprlex, &exprlval)
		}

		/* look through exception table */
		xi := 0
		for {
			if exprExca[xi+0] == -1 && exprExca[xi+1] == exprstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			exprn = exprExca[xi+0]
			if exprn < 0 || exprn == exprchar {
				break
			}
		}
		exprn = exprExca[xi+1]
		if exprn < 0 {
			goto ret0
		}
	}
	if exprn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			exprlex.Error("syntax error")
			Nerrs++
			if exprDebug >= 1 {
				__yyfmt__.Printf("%s", exprStatname(exprstate))
				__yyfmt__.Printf(" saw %s\n", exprTokname(exprchar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for exprp >= 0 {
				exprn = exprPact[exprS[exprp].yys] + exprErrCode
				if exprn >= 0 && exprn < exprLast {
					exprstate = exprAct[exprn] /* simulate a shift of "error" */
					if exprChk[exprstate] == exprErrCode {
						goto exprstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if exprDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", exprS[exprp].yys)
				}
				exprp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if exprDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", exprTokname(exprchar))
			}
			if exprchar == exprEofCode {
				goto ret1
			}
			exprchar = -1
			goto exprnewstate /* try again in the same state */
		}
	}

	/* reduction by production exprn */
	if exprDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", exprn, exprStatname(exprstate))
	}

	exprnt := exprn
	exprpt := exprp
	_ = exprpt // guard against "declared and not used"

	exprp -= exprR2[exprn]
	exprVAL = exprS[exprp+1]

	/* consult goto table to find next state */
	exprn = exprR1[exprn]
	exprg := exprPgo[exprn]
	exprj := exprg + exprS[exprp].yys + 1

	if exprj >= exprLast {
		exprstate = exprAct[exprg]
	} else {
		exprstate = exprAct[exprj]
		if exprChk[exprstate] != -exprn {
			exprstate = exprAct[exprg]
		}
	}
	// dummy call; replaced with literal code
	switch exprnt {

	case 1:
		//line parse.y:27
		{
			if exprS[exprpt-0].num.IsInt() {
				fmt.Println(exprS[exprpt-0].num.Num().String())
			} else {
				fmt.Println(exprS[exprpt-0].num.String())
			}
		}
	case 2:
		//line parse.y:37
		{
			exprVAL.num.Add(exprS[exprpt-2].num, exprS[exprpt-0].num)
		}
	case 3:
		//line parse.y:41
		{
			exprVAL.num.Sub(exprS[exprpt-2].num, exprS[exprpt-0].num)
		}
	case 4:
		//line parse.y:45
		{
			exprVAL.num.Mul(exprS[exprpt-2].num, exprS[exprpt-0].num)
		}
	case 5:
		//line parse.y:49
		{
			exprVAL.num.Quo(exprS[exprpt-2].num, exprS[exprpt-0].num)
		}
	case 6:
		//line parse.y:53
		{
			exprVAL.num = exprS[exprpt-0].num
			exprVAL.num.Neg(exprVAL.num)
		}
	case 7:
		//line parse.y:58
		{
			exprVAL.num = exprS[exprpt-1].num
		}
	case 8:
		exprVAL.num = exprS[exprpt-0].num
	}
	goto exprstack /* stack new state and value */
}
