package adaption

/*
   Takes the list of tokens per rule from the scanner
   Determines the kind of rules and calls the template for execution
*/

import (
	"github.com/templategeneration/genjson"
	tFiles "github.com/templategeneration/template"
	"github.com/templategeneration/utils"
	"strings"
	"text/template"
)

var fileName string
var destination string
var w = &util.OutputUtil{}

type Input struct {
	Token0 string
	Token1 string
	Token2 string
	TList  []string
}

type GenTemp struct {
	Gj *genjson.Encoding
}
type PojoClass struct {
	ClassName  string
	ClassIdent []Var
}

type Var struct {
	Type  string
	Ident string
}

func (g GenTemp) PrintFunc(tList []Token, dest string) {
	w.DestinationTemplate = dest
	for _, v := range tList {
		switch v.kind {
		case kleTok:
			g.printKleene(tList)
			return
		case lex:
			printMorphen(tList)
			return
		case lexRule:
			g.printLexer(tList)
			return
		}
	}
	g.printRule(tList)
}

func (g GenTemp) printKleene(tList []Token) {
	var i string
	p := &PojoClass{}
	var m = make(map[string]interface{})
	for _, t := range tList {
		switch t.kind {
		case ruleName:
			i = t.ident
			g.Gj.Add(t.ident, "", 1)
			tRule(t.ident, i)
			p.ClassName = strings.Title(t.ident)
			m[t.ident] = ""
		case endRule:
			tEndRule()
		case kleTok:
			tKleene(i, t.ident)
			p.ClassIdent = append(p.ClassIdent, Var{"ArrayList<" + strings.Title(t.ident) + ">", t.ident})
			g.Gj.Add(t.ident, i, 1)
		case oBrack:
			tToken(t.ident)
		case cBrack:
			tToken(t.ident)
		case plusTok:
			tmayToken(i, t.ident)
		case keyTok:
			tToken(t.ident)
		case mayKey:
			p.ClassIdent = append(p.ClassIdent, Var{"boolean", t.ident})
			tboolToken(i, t.ident)
		case mayOBrack:
			tmayToken(i, t.ident)
		case mayCBrack:
			tmayToken(i, t.ident)
		case plus | multi | div | minus | equal:
			tToken(t.ident)
		case eofTok:
			break
		case grammarName:
			fileName = t.ident
		default:
			tmayToken(i, t.ident)
			if i != t.ident {
				g.Gj.Add(t.ident, i, 0)
				p.ClassIdent = append(p.ClassIdent, Var{"String", t.ident})
			}
		}
	}
	genClassTemplate(*p)
}

func printMorphen(tList []Token) {
	var i string
	for j, t := range tList {
		switch t.kind {
		case ruleName:
			i = t.ident
			tRule(t.ident, i)
		case endRule:
			tEndRule()
		case tok:
			tToken(t.ident)
		case kleTok:
			tKleene(i, t.ident)
		case actTok:
			tmayToken(i, t.ident)
		case oBrack:
			tToken(t.ident)
		case cBrack:
			tToken(t.ident)
		case mayOBrack:
			tmayOBIdent(j, t.ident, tList)
		case plus:
			tToken(t.ident)
		case multi:
			tToken(t.ident)
		case div:
			tToken(t.ident)
		case minus:
			tToken(t.ident)
		case equal:
			tToken(t.ident)

		case grammarName:
			fileName = t.ident
			return
		case lex:
			tmayToken(i, "ID")
		default:
			tmayToken(i, t.ident)
		}
	}
}

func (g GenTemp) printLexer(tList []Token) {
	// needed later
}

func (g GenTemp) printRule(tList []Token) {
	var i string
	p := &PojoClass{}
	for j, t := range tList {
		switch t.kind {
		case ruleName: // rule :
			i = t.ident
			tRule(t.ident, i)
			p.ClassName = strings.Title(t.ident)
			g.Gj.Add(t.ident, "", 1)
		case endRule: // ;
			tEndRule()
		case kleTok:
			tKleene(i, t.ident)
		case keyTok:
			tToken(t.ident)
		case mayOBrack: // (
			tmayOBIdent(j, i, tList)
		case mayCBrack: // )
			tmayCB()
		case plus:
			tToken(t.ident)
		case multi:
			tToken(t.ident)
		case div:
			tToken(t.ident)
		case minus:
			tToken(t.ident)
		case equal:
			tToken(t.ident)
		case grammarName:
			fileName = t.ident
		default:
			tmayToken(i, t.ident)
			if i != t.ident {
				p.ClassIdent = append(p.ClassIdent, Var{"String", t.ident})
				g.Gj.Add(t.ident, i, 0)
			}
		}
	}
	genClassTemplate(*p)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func tRule(rule string, kleene string) {
	check(genTemplate("rule", Input{Token0: rule, Token1: kleene}))
}

func tToken(token string) {
	check(genTemplate("token", Input{Token0: token}))
}

func tEndRule() {
	check(genTemplate("endRule", Input{}))
}

func tmayToken(t0, t1 string) {
	check(genTemplate("mayKey", Input{Token0: t0, Token1: t1}))

}

func tKleene(r, k string) {
	check(genTemplate("kleene", Input{Token0: r, Token1: k}))
}

func tMayKey(r, k string) {
	check(genTemplate("mayKey", Input{Token0: r, Token1: k}))
}

func tboolToken(r, k string) {
	check(genTemplate("mayBoolTok", Input{Token0: r, Token1: k}))
}

func tmayOB(t0, t1 string) {
	check(genTemplate("obMay", Input{Token0: t0, Token1: t1}))
}

func tmayCB() {
	check(genTemplate("cbMay", Input{}))
}
func tmayOBIdent(j int, i string, tList []Token) {
	t := make([]string, 0)
	for i := j; tList[i].kind != mayCBrack; i++ {
		if tList[i].kind == mayKey {
			t = append(t, tList[i].ident)
		}
	}
	genMultiTemplate(i, t)
}

func genMultiTemplate(i string, t []string) {
	if len(t) > 1 {
		check(genTemplate("multiMay", Input{Token0: i, TList: t}))
	} else {
		check(genTemplate("obMay", Input{Token0: i, Token1: t[0]}))
	}
}

func genTemplate(trule string, input Input) error {
	tc, err := template.New("template").Parse(tFiles.TEMPLATE)
	check(err)
	err = tc.ExecuteTemplate(w, trule, input)
	return err
}

func genClassTemplate(pojo PojoClass) {
	tc, err := template.New("ClassTemp").Parse(tFiles.CLASS_TEMPLATE)
	check(err)
	err = tc.ExecuteTemplate(util.PrintJFile{pojo.ClassName}, "class", pojo)
}
