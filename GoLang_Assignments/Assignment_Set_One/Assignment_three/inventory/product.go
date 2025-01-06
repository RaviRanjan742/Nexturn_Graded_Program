package inventory

type Product struct {
    id    int
    name  string
    price float64
    stock int
}

func NewProduct(id int, name string, price float64, stock int) *Product {
    return &Product{
        id:    id,
        name:  name,
        price: price,
        stock: stock,
    }
}

func (p *Product) GetID() int {
    return p.id
}

func (p *Product) GetName() string {
    return p.name
}

func (p *Product) GetPrice() float64 {
    return p.price
}

func (p *Product) GetStock() int {
    return p.stock
}

func (p *Product) SetStock(quantity int) {
    p.stock = quantity
}
