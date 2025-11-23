package redovalnica

import "fmt"

type Student struct {
	ime     string
	priimek string
	ocene   []int
}

func izpisiKoncniUspeh(studenti map[string]Student) {

	for studentId, student := range studenti {
		var ocena float64 = povprecje(studenti, studentId)

		if ocena < 6 {
			fmt.Println(student.ime, student.priimek+":", "povprečna ocena", ocena, "->", "Neuspešen študent")
		} else if ocena >= 6 && ocena < 9 {
			fmt.Println(student.ime, student.priimek+":", "povprečna ocena", ocena, "->", "Povprečen študent")
		} else {
			fmt.Println(student.ime, student.priimek+":", "povprečna ocena", ocena, "->", "Odličen študent!")
		}
	}
}

func izpisRedovalnice(studenti map[string]Student) {

	for studentId, student := range studenti {
		fmt.Println(studentId, "-", student.ime, student.priimek+":", student.ocene)
	}
}

func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {

	_, ok := studenti[vpisnaStevilka]

	if !ok {
		fmt.Println("Študent ne obstaja")
		return -1.0
	}

	var povprecje float64 = 0

	for _, ocena := range studenti[vpisnaStevilka].ocene {
		povprecje += float64(ocena)
	}

	povprecje /= float64(len(studenti[vpisnaStevilka].ocene))

	if povprecje < 6.0 {
		return 0.0
	}

	return povprecje
}

func dodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {

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
	noveOcene.ocene = append(noveOcene.ocene, ocena)
	studenti[vpisnaStevilka] = noveOcene
}
