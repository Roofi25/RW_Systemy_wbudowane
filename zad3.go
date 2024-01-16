package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

type Student struct {
	Imie     string
	Nazwisko string
	Indeks   int
	Email    string
}

type rms struct {
	N1 string `json:"Imie"`     // ta właściwość zostanie zwrócona jako Imie
	N2 string `json:"Nazwisko"` // ta właściwość zostanie zwrócona jako Nazwisko
	N3 int    `json:"Indeks"`   // ta właściwość zostanie zwrócona jako Indeks
	N4 string `json:"-"`        // ta właściwość nie zostanie zwrócona
}

func studentsToJSONFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // nagłówek JSON

	var studenciBezEmailu []rms

	for _, student := range studenci {
		studentBezEmailu := rms{
			N1: student.Imie,
			N2: student.Nazwisko,
			N3: student.Indeks,
		}
		studenciBezEmailu = append(studenciBezEmailu, studentBezEmailu)
	}

	data, _ := json.Marshal(studenciBezEmailu) // konwersja na JSON
	w.Write(data)                              // zwrócenie danych JSON
}

func stronaFunc(w http.ResponseWriter, r *http.Request) {
	// zwrócenie statycznej strony strona.html
	//http.ServeFile(w, r, "dane.html")
	tmpl, _ := template.ParseFiles("dane.html")
	tmpl.Execute(w, studenci)
}

var studenci []Student = []Student{
	{"Jan", "Kowalski", 12345, "test@test"},
	{"Marek", "Nowak", 30000, "to@tamto"},
	{"Anna", "Zdyb", 23232, "anna@zdyb"},
}

func minmax(tab []int) (int, int) {
	if len(tab) < 1 {
		return 0, 0
	}
	min, max := tab[0], tab[0]

	for i := 1; i < len(tab); i++ {
		if tab[i] < min {
			min = tab[i]
		} else if tab[i] > max {
			max = tab[i]
		}
	}
	return min, max
}

func losujFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	a, _ := strconv.Atoi(vars["a"])
	b, _ := strconv.Atoi(vars["b"])

	var liczbyLosowe []int

	for i := 0; i < a; i++ {
		liczbaLosowa := rand.Intn(b + 1)
		liczbyLosowe = append(liczbyLosowe, liczbaLosowa)
	}

	fmt.Fprintf(w, "<html><body>%v liczb wylosowanych z zakresu od 0 do %v: %v</body><html>", a, b, liczbyLosowe)
}

func pageFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mode := vars["mode"]
	id := vars["id"]

	fmt.Fprintf(w, "<html><body>Mode: %v <br>ID: %v</body></html>", mode, id)
}

func main() {
	numbers := []int{10, 2, 24, 13, 20}
	a, b := minmax(numbers)
	fmt.Println("Min: ", a, "Max: ", b)

	r := mux.NewRouter()
	r.HandleFunc("/losuj/{a:[0-9]+}/{b:[0-9]+}", losujFunc)
	r.HandleFunc("/page/{mode}/{id:[0-9]+}", pageFunc)

	r.HandleFunc("/test", stronaFunc)
	r.HandleFunc("/json", studentsToJSONFunc)
	http.ListenAndServe(":10000", r)
}
