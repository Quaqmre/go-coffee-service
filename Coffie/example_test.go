package coffee

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	g := globals()
	g.defaultBarista.Add(Coffe{Name: "Filtre", ProductName: "StarWars", Rate: 60})
	g.defaultBarista.Add(Coffe{Name: "Türk", ProductName: "Mehmet Efendi", Rate: 80})
	g.defaultBarista.Add(Coffe{Name: "Brazilian", ProductName: "Kava", Rate: 20})
	fmt.Println(g.defaultBarista.List())
}
