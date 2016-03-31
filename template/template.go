package template

const (
	TEMPLATE = `
{{define "rule" }}{{.Token0}}({{.Token1}}) ::= "<\n>{{end}} 

{{define "token"}} {{.Token0}} {{end}}

{{define "mayKey"}}<{{.Token0}}.{{.Token1}}>{{end}}

{{define "kleene"}}<{{.Token0}}.{{.Token1}}:{z|<{{.Token1}}(z)>}><\n>{{end}}

{{define "obMay"}}<if({{.Token0}}.{{.Token1}})>{{end}}

{{define "multiMay"}}<if( {{range .TList }}{{.Token0}}.{{.}}||{{end}} )> {{end}}

{{define "cbMay"}}<endif>{{end}}

{{define "mayBoolTok"}}<if({{.Token0}}.{{.Token1}})> {{.Token1}} <endif>{{end}}

{{define "endRule"}}"
{{end}}"
`

	CLASS_TEMPLATE = `{{define "class"}} 
package tmp;

public class {{.ClassName}} {
{{range .ClassIdent}} 
    {{.Type}} {{.Ident}};{{end}}
}
{{end}}`
)
