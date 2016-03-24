package adaption

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

const eof = -1

type Token struct {
	ident string
	kind  keyword
}

type Adaption struct {
	input       string
	pos         int
	back        int
	state       adaFn
	token       string
	tList       []Token
	destination string
	gT          GenTemp
}

type keyword int

var key = map[string]keyword{
	"grammar": grammarKey,
	"@header": injectKey,
	"//":      commentKey,
	"skip":    skipKey,
	"header":  headerKey,
	"EOF":     eofTok,
}

var op = map[string]keyword{
	"+": plus,
	"*": multi,
	"/": div,
	"-": minus,
	"=": equal,
}

const (
	ruleName keyword = iota
	mulKey
	aloKey
	tokKey
	mayKey
	plusTok // 5
	keyTok
	mayTok
	kleTok
	actTok
	eofTok // 10
	multTok
	tok
	aloTok
	oBrack
	cBrack // 15
	mayOBrack
	mayCBrack
	grammarName
	err
	lexRule // 20
	lex
	regularExpre
	endRule
	plus
	mult
	minus
	multi
	div
	equal
)

const (
	grammarKey keyword = iota
	injectKey          // 1
	commentKey
	skipKey
	headerKey // 4
)

type adaFn func(*Adaption) adaFn

func (a *Adaption) run() {
	for a.state = scanning; a.state != nil; {
		a.state = a.state(a)
	}
}

func scanning(a *Adaption) adaFn {
	switch r := a.peek(); {
	case r == '/' && a.peek() == '/':
		return skipComment
	case isWS(r) || isNL(r):
		a.next()
		return scanning
	case r == ';':
		a.add(";", endRule)
		a.gT.PrintFunc(a.tList, a.destination)
		a.tList = make([]Token, 0)
		a.next()
		return scanning
	case r == ':':
		return resetRule
	case r == '\'':
		return scan
	case r == '(':
		a.add("(", mayOBrack)
		a.next()
		return scanning
	case r == ')':
		a.add(")", mayCBrack)
		a.next()
		return scanning
	case r == -1:
		return nil
	case r <= unicode.MaxASCII && unicode.IsPrint(r):
		return scan
	default:
		if a.next() == eof {
			return nil
		}
	}
	return a.perror("not rec. token " + a.token)
}

func scan(a *Adaption) adaFn {
	switch r := a.peek(); {
	case unicode.IsUpper(r):
		return scanLexer
	case r != '\'':
		return scanToken
	case r == '\'':
		return scanKeyword
	default:
		return a.perror("scan not detect" + a.token)
	}
}

func scanToken(a *Adaption) adaFn {
	if i, ok := key[a.token]; ok {
		switch i {
		case grammarKey:
			return scanGrammar
		case injectKey:
			return scanInjection
		case headerKey:
			return scanInjection
		case eofTok:
			a.add(a.token, eofTok)
		default:
			panic("not defined but in key")
		}
		return scanning
	}

	switch r := a.peek(); {
	case r == '?':
		a.next()
		a.add(a.token, mayTok)
	case r == '*':
		a.add(a.token, kleTok)
	case r == '+':
		a.add(a.token, plusTok)
	case r == '{':
		a.add(a.token, oBrack)
	case r == '}':
		a.add(a.token, cBrack)
	case r == '(':
		a.add(a.token, oBrack)
	case r == ')':
		a.add(a.token, mayKey)
		a.detMulti()
	case r == ' ':
		a.add(a.token, tok)
	case r == eof:
		return nil
	default:
		a.token += string(r)
		a.next()
		return scanToken
	}
	a.next()
	if a.token != "" {
		panic("not Token not rec")
	}
	return scanning
}

func scanKeyword(a *Adaption) adaFn {
	switch r := a.peek(); {
	case r == ' ':
		a.add(a.token, keyTok)
	case r == '?':
		a.add(a.token, mayKey)
	case r == '*':
		a.add(a.token, mulKey)
	case r == '+':
		a.add(a.token, aloKey)
	case r == '{':
		a.add("{", oBrack)
	case r == '}':
		a.add("}", cBrack)
	case r == '\'':
		a.next()
		return scanKeyword
	case r == eof:
		return nil
	default:
		a.token += string(r)
		a.next()
		return scanKeyword
	}
	if a.token != "" {
		a.perror("Keyword not rec " + a.token)
	}
	return scanning
}

func scanLexer(a *Adaption) adaFn {
	switch r := a.peek(); {
	case r == ' ' || r == ';' || r == ':':
		if a.token == "EOF" {
			a.add(a.token, eofTok)
		}
		a.add(a.token, lex)
	default:
		a.token += string(r)
		a.next()
		return scanLexer
	}
	return scanning
}

func skipInjection(a *Adaption) adaFn {
	r := a.next()
	if r == '}' {
		a.token = ""
		return scanning
	}
	return skipInjection
}

func scanRegular(a *Adaption) adaFn {
	r := a.peek()
	if isNL(r) || r == eof || r == ';' {
		a.add(a.token, regularExpre)
		return scanning
	}
	a.next()
	a.token += string(r)
	return scanRegular
}

func skipComment(a *Adaption) adaFn {
	r := a.next()
	if isNL(r) {
		a.token = ""
		return scanning
	}
	return skipComment
}

func (a *Adaption) detMulti() keyword {
	switch r := a.next(); {
	case r == '?':
		return mayTok
	case r == '*':
		return multTok
	case r == '+':
		return aloTok
	case r == ' ':
		return tok
	case r == ')':
		a.add(string(r), mayCBrack)
	default:
		a.perror("no mult. detect " + a.token)
	}
	return err
}

func (a *Adaption) next() rune {
	r, l := utf8.DecodeRuneInString(a.input[a.pos:])
	a.pos += l
	a.back = l
	if int(a.pos) >= len(a.input) {
		return eof
	}
	return r
}

func (a *Adaption) peek() rune {
	r := a.next()
	a.pos -= a.back
	return r
}

func resetRule(a *Adaption) adaFn {
	r := a.tList[len(a.tList)-1].kind
	if r == lex {
		a.tList[len(a.tList)-1].kind = lexRule
		return scanRegular
	} else {
		a.tList[len(a.tList)-1].kind = ruleName
	}
	a.next()
	return scanning
}
func isWS(r rune) bool {
	return r == ' ' || r == '\t'
}

func isNL(r rune) bool {
	return r == '\n' || r == '\r'
}

func scanGrammar(a *Adaption) adaFn {
	a.token = ""
	for r := a.peek(); r != ';'; {
		a.token += string(r)
		r = a.next()
	}
	a.add(strings.Trim(a.token, " "),
		grammarName)
	return scanning
}

func scanInjection(a *Adaption) adaFn {
	r := a.next()
	if r == '}' {
		return scanning
	}
	a.token = ""
	return scanInjection
}

func scanComment(a *Adaption) adaFn {
	r := a.next()
	if isNL(r) {
		return scanning
	}
	return scanComment
}

func (a *Adaption) perror(er string) adaFn {
	a.add(er, err)
	panic("ERROR: " + er)
	return nil
}

func RunAdaption(input *[]byte, desti *string, gt *GenTemp) *[]Token {
	adapt := &Adaption{
		input:       string(*input),
		tList:       make([]Token, 0),
		pos:         0,
		destination: *desti,
		gT:          *gt}
	adapt.run()
	return &adapt.tList

}

func (a *Adaption) add(i string, k keyword) {
	if i != "" {
		a.tList = append(a.tList, Token{ident: i, kind: k})
		a.token = ""
		return
	} else {
		var i = a.pos + a.back
		t := a.input[a.pos:i]
		if i, ok := op[t]; ok {
			switch i {
			case plus:
				a.tList = append(a.tList, Token{ident: t, kind: plus})
			case multi:
				a.tList = append(a.tList, Token{ident: t, kind: multi})
			}
			a.next()
		}
	}
}
