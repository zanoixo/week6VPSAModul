// package redovalnica provides some simple methods to print final grades of students.

package redovalnica

import "fmt"

type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

// IzpisKoncniUspeh takes a map of a String which is a student number and a Student and prints how good student grades are
func IzpisiKoncniUspeh(studenti map[string]Student) {

	for studentId, student := range studenti {
		var ocena float64 = povprecje(studenti, studentId)

		if ocena < 6 {
			fmt.Println(student.Ime, student.Priimek+":", "povprečna ocena", ocena, "->", "Neuspešen študent")
		} else if ocena >= 6 && ocena < 9 {
			fmt.Println(student.Ime, student.Priimek+":", "povprečna ocena", ocena, "->", "Povprečen študent")
		} else {
			fmt.Println(student.Ime, student.Priimek+":", "povprečna ocena", ocena, "->", "Odličen študent!")
		}
	}
}

// IzpisRedovalnice takes a map of a String which is a student number and a Student and prints out student grades
func IzpisRedovalnice(studenti map[string]Student) {

	for studentId, student := range studenti {
		fmt.Println(studentId, "-", student.Ime, student.Priimek+":", student.Ocene)
	}
}

func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {

	_, ok := studenti[vpisnaStevilka]

	if !ok {
		fmt.Println("Študent ne obstaja")
		return -1.0
	}

	var povprecje float64 = 0

	for _, ocena := range studenti[vpisnaStevilka].Ocene {
		povprecje += float64(ocena)
	}

	povprecje /= float64(len(studenti[vpisnaStevilka].Ocene))

	if povprecje < 6.0 {
		return 0.0
	}

	return povprecje
}

// DodajOceno takes a map of a String which is a student number and a Student, a String which is a student number and and int represanting the grade 5-10 and adds it to the map
func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {

	if ocena > 10 || ocena < 1 {
		fmt.Println("Napačna ocena")
		return
	}

	_, ok := studenti[vpisnaStevilka]

	if !ok {
		fmt.Println("Študent ne obstaja")
		return
	}

	var noveOcene = studenti[vpisnaStevilka]
	noveOcene.Ocene = append(noveOcene.Ocene, ocena)
	studenti[vpisnaStevilka] = noveOcene
}
