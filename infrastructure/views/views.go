package views

import (
  "html/template"
)

type templateParser struct {
	HTML string
}

func (tP *templateParser) Write(p []byte) (n int, err error) {
	tP.HTML += string(p)
	return len(p), nil
}

type staticTemplate struct {
  Title string
}

func Render(page string, data interface{}) (error, []byte){
	tp := &templateParser{}
	t, _ := template.ParseFiles("views/" + page + ".html")
	t.Execute(tp, data)
	return nil, []byte(tp.HTML)
}

func Static(page string) (error, []byte) {
  data := &staticTemplate{"Title"}
  return Render(page, data)
}

func Error(err error) []byte {
  e, v := Render("error", err)
  if e != nil {
    return []byte(err.Error());
  } 
  return v
}



