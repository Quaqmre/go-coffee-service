package coffee

import "testing"

func TestAmberName(t *testing.T) {
	item := globals()
	item.defaultBarista = &mockCoffe{}
	SetAmberName("asd")
	if item.Name != "asd" {
		t.Error("SetAmberName goes wrong")
	}
}
func TestSetBarista(t *testing.T) {
	mockedBarista := &mockCoffe{}
	SetBarista(mockedBarista)

	item := globals()
	item.defaultBarista.Add(Coffe{})

	if !mockedBarista.added {
		t.Error("doest setted barista with mock")
	}

}

type mockCoffe struct {
	added       bool
	addedReview bool
	listed      bool
}

var _ Coffeer = (*mockCoffe)(nil)

func (m *mockCoffe) Add(b Coffe) {
	m.added = true
}
func (m *mockCoffe) AddReview(b Coffe) {
	m.addedReview = true
}
func (m *mockCoffe) List() []Coffe {
	m.listed = true
	return nil
}
func (m *mockCoffe) ListReview() []Coffe {
	m.listed = true
	return nil
}
func (m *mockCoffe) GetAmberName() string {
	item := globals()
	return item.Name
}
