package main

import (
	"fmt"
	"html/template"
	"text/template"
)

type Temps struct{
	notemp *template.Template
	indx *template.Template
	helo *template.Template
}

func notemp() *template.Template{
	src := "<html><body><h1>No Template</h1></body></html>"
	tmp, _ :=template.New("index").Parse(src)
}

func main(){

}



// type Template struct {
// 	name string
// 	*parse.Tree
// 	*common
// 	leftDelim  string
// 	rightDelim string
// }

// func New(name string) *Template {
// 	t := &Template{
// 		name: name,
// 	}
// 	t.init()
// 	return t
// }

// func (t *Template) init() {
// 	if t.common == nil {
// 		c := new(common)
// 		c.tmpl = make(map[string]*Template)
// 		c.parseFuncs = make(FuncMap)
// 		c.execFuncs = make(map[string]reflect.Value)
// 		t.common = c
// 	}
// }

// func (t *Template) Parse(text string) (*Template, error) {
// 	t.init()
// 	t.muFuncs.RLock()
// 	trees, err := parse.Parse(t.name, text, t.leftDelim, t.rightDelim, t.parseFuncs, builtins())
// 	t.muFuncs.RUnlock()
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Add the newly parsed trees, including the one for t, into our common structure.
// 	for name, tree := range trees {
// 		if _, err := t.AddParseTree(name, tree); err != nil {
// 			return nil, err
// 		}
// 	}
// 	return t, nil
// }
