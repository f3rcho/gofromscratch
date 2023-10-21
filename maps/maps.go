package maps

import "fmt"

func ShowMaps() {
	countries := make(map[string]string)

	countries["mexico"] = "CDMX"
	countries["colombia"] = "Bogota"
	countries["peru"] = "Lima"
	// fmt.Println(countries)
	// fmt.Println(countries["mexico"])

	camp := map[string]int{
		"Barcelona": 39,
		"Madrid":    28,
		"Sevilla":   30,
		"Valencia":  29,
	}

	// fmt.Println(camp)

	for team, points := range camp {
		fmt.Printf("Equipo %s, tiene %d puntos\n", team, points)
	}

	delete(camp, "Madrid")
	// fmt.Println(camp)

	points, ok := camp["Juventus"]
	fmt.Printf("El puntaje capturado es %d, el equipo existe = %t \n", points, ok)
}
