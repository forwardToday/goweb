package controllers

import (
	"html/template"
	"minusblog/minus"
)

type MinusController struct {
	minus.Controller
}

func (c *MinusController) Get() {
	t, _ := template.ParseFiles("./view/minus.gtpl")
	t.Execute(c.Ct.ResponseWriter, nil)
	// fmt.Fprintf(c.Ct.ResponseWriter, "召唤神龙")
}
