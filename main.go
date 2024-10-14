package main

import (
	"fmt"
	"html/template"
	"im/controller"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/user/login", controller.UserLogin)
	http.HandleFunc("/user/register", controller.UserRegister)

	http.Handle("/asset/", http.FileServer(http.Dir(".")))

	RegisterView()

	http.ListenAndServe(":8080", nil)
}

func RegisterTemplate() {
	GlobTemplate := template.New("root")
	GlobTemplate, err := GlobTemplate.Parse("view/**/*")
	if err != nil {
		log.Fatal(err)
	}

	for _, tpl := range GlobTemplate.Templates() {
		pattern := tpl.Name()
		fmt.Println(pattern)
		http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			GlobTemplate.ExecuteTemplate(w, pattern, nil)
		})
		fmt.Println("register=>" + pattern)
	}
}

func RegisterView() {
	tpl, err := template.ParseGlob("view/**/*")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range tpl.Templates() {
		tplname := v.Name()
		if strings.HasSuffix(tplname, ".shtml") {
			fmt.Println("HandleFunc " + v.Name())
			http.HandleFunc(tplname, func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("parse " + v.Name() + "==" + tplname)
				err := tpl.ExecuteTemplate(w, tplname, nil)
				if err != nil {
					log.Fatal(err.Error())
				}
			})
		}
	}
}
