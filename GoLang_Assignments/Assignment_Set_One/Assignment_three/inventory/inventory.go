package inventory

import (
    "fmt"
    "sort"
    "strings"
)

type Inventory struct {
    products []*Product
}

func NewInventory() *Inventory {
    return &Inventory{
        products: make([]*Product, 0),
    }
}

func (i *Inventory) AddProduct(product *Product) error {
    if product.stock < 0 {
        return fmt.Errorf("stock cannot be negative")
    }

    if product.price < 0 {
        return fmt.Errorf("price cannot be negative")
    }

    
    for _, p := range i.products {
        if p.id == product.id {
            return fmt.Errorf("product with ID %d already exists", product.id)
        }
    }

    i.products = append(i.products, product)
    return nil
}

func (i *Inventory) UpdateStock(id int, quantity int) error {
    if quantity < 0 {
        return fmt.Errorf("stock cannot be negative")
    }

    product, err := i.SearchByID(id)
    if err != nil {
        return err
    }

    product.SetStock(quantity)
    return nil
}

func (i *Inventory) SearchByID(id int) (*Product, error) {
    for _, product := range i.products {
        if product.id == id {
            return product, nil
        }
    }
    return nil, fmt.Errorf("product with ID %d not found", id)
}

func (i *Inventory) SearchByName(name string) []*Product {
    var results []*Product
    searchTerm := strings.ToLower(name)
    
    for _, product := range i.products {
        if strings.Contains(strings.ToLower(product.name), searchTerm) {
            results = append(results, product)
        }
    }
    return results
}

func (i *Inventory) GetAllProducts() []*Product {
    return i.products
}

func (i *Inventory) GetProductsSortedByPrice() []*Product {
    sorted := make([]*Product, len(i.products))
    copy(sorted, i.products)
    
    sort.Slice(sorted, func(i, j int) bool {
        return sorted[i].price < sorted[j].price
    })
    
    return sorted
}

func (i *Inventory) GetProductsSortedByStock() []*Product {
    sorted := make([]*Product, len(i.products))
    copy(sorted, i.products)
    
    sort.Slice(sorted, func(i, j int) bool {
        return sorted[i].stock < sorted[j].stock
    })
    
    return sorted
}