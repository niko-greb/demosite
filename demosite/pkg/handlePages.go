package pkg

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type Urls struct {
	Id uint16
	Url, ShortUrl, Comment string

}

var urlPack = []Urls{}
var showUrl = Urls{}
/** БЛОК СТРАНИЦ САЙТА **/
func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("mysql", "test:test@tcp(10.0.0.246)/urls?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM `urls` LIMIT 4")
	if err != nil {
		panic(err)
	}

	urlPack = []Urls{}
	for res.Next() {
		var shorty Urls
		err = res.Scan(&shorty.Id, &shorty.Url, &shorty.ShortUrl ,&shorty.Comment)
		if err != nil {
			panic(err)
		}
		urlPack = append(urlPack, shorty)
		//fmt.Println(fmt.Sprintf("Short url: %s ", shorty.ShortUrl))
	}

	t.ExecuteTemplate(w, "index", urlPack)
}

func contacts(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/contacts.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "contacts", nil)
}

func cutter(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/cutter.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "cutter", nil)
}

func showUrls(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/urls.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	db, err := sql.Open("mysql", "test:test@tcp(10.0.0.246)/urls?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query(fmt.Sprintf("SELECT * FROM `urls` WHERE id = '%s'", vars["id"]))
	if err != nil {
		panic(err)
	}

	showUrl = Urls{}
	for res.Next() {
		var shorty Urls
		err = res.Scan(&shorty.Id, &shorty.Url, &shorty.ShortUrl ,&shorty.Comment)
		if err != nil {
			panic(err)
		}
		showUrl = shorty

	}

	t.ExecuteTemplate(w, "urls", showUrl)
}

func HandleFunc() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/contacts/", contacts).Methods("GET")
	rtr.HandleFunc("/cutter", cutter).Methods("GET")
	rtr.HandleFunc("/cutresult", CutResult).Methods("POST")
	rtr.HandleFunc("/urls/{id:[0-9]+}", showUrls).Methods("GET")
	http.Handle("/", rtr)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":8080", nil)
}

