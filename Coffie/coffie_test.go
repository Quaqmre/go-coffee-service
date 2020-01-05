package coffie

import "testing"

func TestAmberName(t *testing.T) {
	item := globals()
	item.defaultBarista = &mockCoffie{}
	SetAmberName("asd")
	if item.Name != "asd" {
		t.Error("SetAmberName goes wrong")
	}
}
func TestSetBarista(t *testing.T) {
	mockedBarista := &mockCoffie{}
	SetBarista(mockedBarista)

	item := globals()
	item.defaultBarista.Add(Coffie{})

	if !mockedBarista.added {
		t.Error("doest setted barista with mock")
	}

}

type mockCoffie struct {
	added       bool
	addedReview bool
	listed      bool
}

var _ Coffieer = (*mockCoffie)(nil)

func (m *mockCoffie) Add(b Coffie) {
	m.added = true
}
func (m *mockCoffie) AddReview(b Coffie) {
	m.addedReview = true
}
func (m *mockCoffie) List() []Coffie {
	m.listed = true
	return nil
}
func (m *mockCoffie) ListReview() []Coffie {
	m.listed = true
	return nil
}
func (m *mockCoffie) GetAmberName() string {
	item := globals()
	return item.Name
}
