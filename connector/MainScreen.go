package connector

import (
	"net/http"
	"text/template"
)

func MainScreenHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/MainScreen.html")
	if err != nil {
		http.Error(w, "ошибка загрузки фронта", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
