package main

import (
	"github.com/zanoixo/week6VPSAModul/redovalnica"
)

func main() {
	var studenti = make(map[string]redovalnica.Student)
	studenti["63230001"] = redovalnica.Student{Ime: "Student1", Priimek: "priimek1", Ocene: []int{6, 5, 7}}
	studenti["63230002"] = redovalnica.Student{Ime: "Student2", Priimek: "priimek2", Ocene: []int{6, 6, 6}}
	studenti["63230003"] = redovalnica.Student{Ime: "Student3", Priimek: "priimek3", Ocene: []int{10, 10, 10}}

	redovalnica.DodajOceno(studenti, "63230001", 10)

	redovalnica.IzpisRedovalnice(studenti)
	redovalnica.IzpisiKoncniUspeh(studenti)

}
