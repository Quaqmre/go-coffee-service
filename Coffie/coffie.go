package coffie

import "sync"

//Coffieer interface for methods
type Coffieer interface {
	Add(Coffie)
	AddReview(Coffie)
	List() []Coffie
	ListReview() []Coffie
	GetAmberName() string
}

//Coffie strore all coffie property for single one
type Coffie struct {
	Name           string
	ProductName    string
	ProductDate    string
	ConsuptionDate string
	Rate           int
}

//Amber for a customer see all coffie
type amber struct {
	Name           string
	Quantity       int
	coffieList       []Coffie
	coffieReviewList []Coffie
	defaultBarista Coffieer
}
type barista struct {
}

//SetAmberName change specspic AmberName
func SetAmberName(s string) {
	item := globals()
	item.Name = s
}

//SetBarista hiring to the new barista :D
func SetBarista(e Coffieer) {
	item := globals()
	item.defaultBarista = e
}

var (
	mu   sync.RWMutex
	ambr = newDefaultAmber()
)

func newDefaultAmber() amber {
	return amber{Name: "Kardesler", Quantity: 0, coffieList: []Coffie{}, defaultBarista: &barista{}}
}
func globals() *amber {
	mu.RLock()
	defer mu.RUnlock()
	return &ambr
}

var _ Coffieer = (*barista)(nil)

func (br *barista) Add(b Coffie) {
	item := globals()
	item.coffieList = append(ambr.coffieList, b)
}
func (br *barista) AddReview(b Coffie) {
	item := globals()
	item.coffieReviewList = append(item.coffieReviewList, b)
}
func (br *barista) List() []Coffie {
	item := globals()
	clone := make([]Coffie, len(item.coffieList))
	copy(clone, item.coffieList)
	return clone
}
func (br *barista) ListReview() []Coffie {
	item := globals()
	clone := make([]Coffie, len(item.coffieReviewList))
	copy(clone, item.coffieReviewList)
	return clone
}
func (br *barista) GetAmberName() string {
	item := globals()
	return item.Name
}
