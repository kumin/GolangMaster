package entities

type Properties struct {
	Color    string
	Price    float64
	Category string
}

type Product struct {
	ID         int64
	Name       string
	Properties *Properties
}
