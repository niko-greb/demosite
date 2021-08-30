package pkg

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

//cutResult Принимаем данные из форм
func CutResult(w http.ResponseWriter, r *http.Request) {
	defer handleDbErr()

	cuturl := r.FormValue("cutterurl")
	comment := r.FormValue("comment")

	shortyFirst := FirstStr(cuturl)
	randUrl := RandomUrl()
	var shortyUrl string
	shortyUrl = shortyFirst+randUrl

	insertToDB(cuturl, shortyUrl , comment, w, r)

}

// FirstStr Выделяем у ссылки http:// или https://
func FirstStr(str string) string{
	var a []string
	for i , c := range str {
		if i == 4 && fmt.Sprintf("%c", c) != "s"{
			a = append(a, "://")
			break
		} else if i == 8 {
			break
		}
		a = append(a , fmt.Sprintf("%c", c))
	}
	b := strings.Join(a, "")
	return b
}


func handleDbErr() {
	if rec := recover(); rec != nil {
		fmt.Println("Recovering from panic:", rec)

	}

}

func insertToDB(cuturl, shortyUrl , comment string, w http.ResponseWriter, r *http.Request) {
	if cuturl == "" {
		fmt.Fprint(w, "Лазанья!")
	} else {
		db, err := sql.Open("mysql", "test:test@tcp(127.0.0.1:3306)/urls?parseTime=true")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		insert, err := db.Query(fmt.Sprintf("INSERT INTO urls (`url`, `short_url`, `comment`)" +
			" VALUES ('%s', '%s', '%s')", cuturl, shortyUrl , comment))
		if err != nil {
			fmt.Fprint(w, "Такой url адресс уже существует")
			panic("err")

		}
		defer insert.Close()

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}
