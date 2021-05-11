package query

// Paginate
type Paginate struct {
	Count   int
	Offset  int
	MaxPage int
	Page    int
	Data    []interface{}
}
