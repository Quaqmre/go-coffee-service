package coffee

import "sync"

//Coffeer interface for methods
type Coffeer interface {
	Add(Coffe)
	AddReview(Coffe)
	List() []Coffe
	ListReview() []Coffe
	GetAmberName() string
}

//Coffe strore all coffee property for single one
type Coffe struct {
	Name           string
	ProductName    string
	ProductDate    string
	ConsuptionDate string
	Rate           int
}

//Amber for a customer see all coffee
type amber struct {
	Name           string
	Quantity       int
	coffeeList       []Coffe
	coffeeReviewList []Coffe
	defaultBarista Coffeer
}
type barista struct {
}

//SetAmberName change specspic AmberName
func SetAmberName(s string) {
	item := globals()
	item.Name = s
}

//SetBarista hiring to the new barista :D
func SetBarista(e Coffeer) {
	item := globals()
	item.defaultBarista = e
}

var (
	mu   sync.RWMutex
	ambr = newDefaultAmber()
)

func newDefaultAmber() amber {
	return amber{Name: "Kardesler", Quantity: 0, coffeeList: []Coffe{}, defaultBarista: &barista{}}
}
func globals() *amber {
	mu.RLock()
	defer mu.RUnlock()
	return &ambr
}

var _ Coffeer = (*barista)(nil)

func (br *barista) Add(b Coffe) {
	item := globals()
	item.coffeeList = append(ambr.coffeeList, b)
}
func (br *barista) AddReview(b Coffe) {
	item := globals()
	item.coffeeReviewList = append(item.coffeeReviewList, b)
}
func (br *barista) List() []Coffe {
	item := globals()
	clone := make([]Coffe, len(item.coffeeList))
	copy(clone, item.coffeeList)
	return clone
}
func (br *barista) ListReview() []Coffe {
	item := globals()
	clone := make([]Coffe, len(item.coffeeReviewList))
	copy(clone, item.coffeeReviewList)
	return clone
}
func (br *barista) GetAmberName() string {
	item := globals()
	return item.Name
}
