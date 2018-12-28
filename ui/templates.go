// Copyright 2015 The Loadcat Authors. All rights reserved.

package ui

import (
	"html/template"
)

var (
	TplLayout = template.Must(template.New("layout.html").ParseFiles("./templates/layout.html"))

	TplBalancerList     = template.Must(template.Must(TplLayout.Clone()).ParseFiles("./templates/balancerList.html"))
	TplBalancerNewForm  = template.Must(template.Must(TplLayout.Clone()).ParseFiles("./templates/balancerNewForm.html"))
	TplBalancerView     = template.Must(template.Must(TplLayout.Clone()).ParseFiles("./templates/balancerView.html"))
	TplBalancerEditForm = template.Must(template.Must(TplLayout.Clone()).ParseFiles("./templates/balancerEditForm.html"))

	TplServerNewForm  = template.Must(template.Must(TplLayout.Clone()).ParseFiles("./templates/serverNewForm.html"))
	TplServerEditForm = template.Must(template.Must(TplLayout.Clone()).ParseFiles("./templates/serverEditForm.html"))
)
