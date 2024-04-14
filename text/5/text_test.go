package main

import (
	"bytes"
	"testing"
	"text/template"
)

// начало решения

// var templateText = `{{if .Balance>=100 }} Balance ok {{ else }} {{if .Balance>=0 && .Balance<100 }} Пора пополнить {{ end }}`
var templateText = `{{if (ge .Balance 100)}}{{.Name}}, добрый день! Ваш баланс - {{.Balance}}₽. Все в порядке.{{else if (gt .Balance 0)}}{{.Name}}, добрый день! Ваш баланс - {{.Balance}}₽. Пора пополнить.{{else}}{{.Name}}, добрый день! Ваш баланс - {{.Balance}}₽. Доступ заблокирован.{{end}}`

//var templateText = `{{if (ge .Balance 100)}}Алиса, добрый день! Ваш баланс - {{.Balance}}Р. Все в порядке.{{end}}`

// конец решения

type User struct {
	Name    string
	Balance int
}

// renderToString рендерит данные по шаблону в строку
func renderToString(tpl *template.Template, data any) string {
	var buf bytes.Buffer
	tpl.Execute(&buf, data)
	return buf.String()
}

func Test(t *testing.T) {
	tpl := template.New("message")
	tpl = template.Must(tpl.Parse(templateText))

	user := User{"Алиса", -1}
	got := renderToString(tpl, user)

	const want = "Алиса, добрый день! Ваш баланс - 500₽. Все в порядке."
	if got != want {
		t.Errorf("%v: want '%v'; got '%v'", user, want, got)
	}
}
