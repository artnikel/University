package server

import (
	"net/http"
	"text/template"
	p "university/position"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "index", nil)
}

func insert(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/insert.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "insert", nil)
}

func save_adding(w http.ResponseWriter, r *http.Request) {
	(p.Student).Insert_db(p.Student{}, w, r)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func delete(w http.ResponseWriter, r *http.Request) {
	(p.Student).Select_db(p.Student{})
	t, err := template.ParseFiles("templates/delete.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "delete", p.Lists)
}

func save_deleting(w http.ResponseWriter, r *http.Request) {
	(p.Student).Delete_db(p.Student{}, w, r)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func seeall(w http.ResponseWriter, r *http.Request) {

	(p.Student).Select_db(p.Student{})
	t, err := template.ParseFiles("templates/seeall.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "seeall", p.Lists)

}

func search(w http.ResponseWriter, r *http.Request) {

	(p.Student).Search_db(p.Student{}, w, r)
	t, err := template.ParseFiles("templates/search.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "search", p.Lists)

}

func update(w http.ResponseWriter, r *http.Request) {
	(p.Student).Select_db(p.Student{})

	t, err := template.ParseFiles("templates/update.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "update", p.Lists)
}

func save_updating(w http.ResponseWriter, r *http.Request) {
	(p.Student).Update_db(p.Student{}, w, r)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func HandleFunction() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/insert", insert).Methods("GET")
	rtr.HandleFunc("/save_adding", save_adding).Methods("POST")
	rtr.HandleFunc("/delete", delete).Methods("GET")
	rtr.HandleFunc("/save_deleting", save_deleting).Methods("POST")
	rtr.HandleFunc("/seeall", seeall).Methods("GET")
	rtr.HandleFunc("/search", search).Methods("POST")
	rtr.HandleFunc("/update", update).Methods("GET")
	rtr.HandleFunc("/save_updating", save_updating).Methods("POST")
	http.Handle("/", rtr)

	http.ListenAndServe(":8000", nil)
}