package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Biodata struct {
	Nama   string
	Alamat string
	Umur   int
	Email  string
}

var biodata = map[string]Biodata{
	"olivia@corp.com":        {"Olivia", "Otista, Jakarta Timur", 30, "olivia@corp.com"},
	"dian@gr.com":            {"Dian", "Kalapa Gading, Jakarta Utara", 25, "Dian@gr.com"},
	"jeremy.first@media.com": {"Jeremy First", "Menteng, Jakarta Pusat", 45, "jeremy.first@media.com"},
	"anaz@media.com":         {"Anaztasia", "Ampera, Jakarta Selatan", 28, "anaz@media.com"},
	"Truncat@sport.com":      {"Truncat", "Meruya, Jakarta Barat", 21, "Truncat@sport.com"},
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl, err := template.ParseFiles("login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tpl.Execute(w, biodata)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		email := r.FormValue("email")
		isFound := false
		for _, data := range biodata {
			if data.Email == email {
				isFound = true
				break
			}
		}
		if isFound {
			http.Redirect(w, r, "/success?email="+email, http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/failed", http.StatusSeeOther)
		}

	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

}

func loginSuccess(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	bio := biodata[email]
	tpl, err := template.ParseFiles("success.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, bio)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func loginFailed(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "failed.html")
}

func main() {
	http.HandleFunc("/", loginPage)
	http.HandleFunc("/success", loginSuccess)
	http.HandleFunc("/failed", loginFailed)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
