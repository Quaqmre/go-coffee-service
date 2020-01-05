package coffie

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	g := globals()
	g.defaultBarista.Add(Coffie{Name: "Filtre", ProductName: "StarWars", Rate: 60})
	g.defaultBarista.Add(Coffie{Name: "Türk", ProductName: "Mehmet Efendi", Rate: 80})
	g.defaultBarista.Add(Coffie{Name: "Brazilian", ProductName: "Kava", Rate: 20})
	fmt.Println(g.defaultBarista.List())
}
