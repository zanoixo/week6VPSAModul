package main

func main() {
	var studenti = make(map[string]Student)
	studenti["63230001"] = Student{"Student1", "priimek1", []int{6, 5, 7}}
	studenti["63230002"] = Student{"Student2", "priimek2", []int{6, 6, 6}}
	studenti["63230003"] = Student{"Student3", "priimek3", []int{10, 10, 10}}

	dodajOceno(studenti, "63230001", 10)

	povprecje(studenti, "63230001")

	izpisRedovalnice(studenti)
	izpisiKoncniUspeh(studenti)

}
