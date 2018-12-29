// Copyright 2015 The Loadcat Authors. All rights reserved.

package ui

import (
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"strings"
)

func GetCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		fmt.Println(err.Error())
	}
	s = strings.Replace(s, "\\", "/", -1)
	s = strings.Replace(s, "\\\\", "/", -1)
	i := strings.LastIndex(s, "/")
	path := string(s[0 : i+1])
	return path
}

var (

	TplLayout = template.Must(template.New("layout.html").ParseFiles(GetCurrentPath() + "templates/layout.html"))

	TplBalancerList     = template.Must(template.Must(TplLayout.Clone()).ParseFiles(GetCurrentPath() + "templates/balancerList.html"))
	TplBalancerNewForm  = template.Must(template.Must(TplLayout.Clone()).ParseFiles(GetCurrentPath() + "templates/balancerNewForm.html"))
	TplBalancerView     = template.Must(template.Must(TplLayout.Clone()).ParseFiles(GetCurrentPath() + "templates/balancerView.html"))
	TplBalancerEditForm = template.Must(template.Must(TplLayout.Clone()).ParseFiles(GetCurrentPath() + "templates/balancerEditForm.html"))

	TplServerNewForm  = template.Must(template.Must(TplLayout.Clone()).ParseFiles(GetCurrentPath() + "templates/serverNewForm.html"))
	TplServerEditForm = template.Must(template.Must(TplLayout.Clone()).ParseFiles(GetCurrentPath() + "templates/serverEditForm.html"))
)
