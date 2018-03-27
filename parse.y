%{

package main

import (
	"fmt"
	"math/big"
)

%}

%union {
	num *big.Rat
}

%type	<num>	expr
%token	<num>	NUM

%left '+' '-'
%left '*' '/'
%nonassoc UMINUS

%%

top:
	expr
	{
		if $1.IsInt() {
			fmt.Println($1.Num().String())
		} else {
			fmt.Println($1.String())
		}
	}

expr:
	expr '+' expr
	{
		$$.Add($1, $3)
	}
|	expr '-' expr
	{
		$$.Sub($1, $3)
	}
|	expr '*' expr
	{
		$$.Mul($1, $3)
	}
|	expr '/' expr
	{
		$$.Quo($1, $3)
	}
|	'-' expr 				%prec UMINUS
	{
		$$ = $2
		$$.Neg($$)
	}
|	'(' expr ')'
	{
		$$ = $2
	}
|	NUM

%%
