
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "Assignment_three/inventory"
)

func readString(reader *bufio.Reader, prompt string) string {
    fmt.Print(prompt)
    str, _ := reader.ReadString('\n')
    return strings.TrimSpace(str)
}

func readFloat(reader *bufio.Reader, prompt string) (float64, error) {
    str := readString(reader, prompt)
    return strconv.ParseFloat(str, 64)
}

func readInt(reader *bufio.Reader, prompt string) (int, error) {
    str := readString(reader, prompt)
    return strconv.Atoi(str)
}

func showMenu() {
    fmt.Println("\n=== Inventory Management System ===")
    fmt.Println("1. Add New Product")
    fmt.Println("2. Update Stock")
    fmt.Println("3. Search Product by ID")
    fmt.Println("4. Search Product by Name")
    fmt.Println("5. Display All Products")
    fmt.Println("6. Display Products by Price (Ascending)")
    fmt.Println("7. Display Products by Stock (Ascending)")
    fmt.Println("8. Exit")
    fmt.Print("Enter your choice: ")
}

func addProduct(reader *bufio.Reader, inv *inventory.Inventory) {
    fmt.Println("\n--- Add New Product ---")
    
    id, err := readInt(reader, "Enter Product ID: ")
    if err != nil {
        fmt.Println("Invalid ID format")
        return
    }

    name := readString(reader, "Enter Product Name: ")
    if name == "" {
        fmt.Println("Name cannot be empty")
        return
    }

    price, err := readFloat(reader, "Enter Product Price: ")
    if err != nil {
        fmt.Println("Invalid price format")
        return
    }

    stock, err := readInt(reader, "Enter Initial Stock: ")
    if err != nil {
        fmt.Println("Invalid stock format")
        return
    }

    product := inventory.NewProduct(id, name, price, stock)
    err = inv.AddProduct(product)
    if err != nil {
        fmt.Printf("Error adding product: %v\n", err)
        return
    }

    fmt.Println("Product added successfully!")
}

func displayProducts(products []*inventory.Product) {
    if len(products) == 0 {
        fmt.Println("No products found")
        return
    }

    fmt.Println("\nProduct List:")
    fmt.Printf("%-6s | %-20s | %-10s | %-10s\n", "ID", "Name", "Price", "Stock")
    fmt.Println(strings.Repeat("-", 52))
    
    for _, p := range products {
        fmt.Printf("%-6d | %-20s | $%-9.2f | %-10d\n", 
            p.GetID(), p.GetName(), p.GetPrice(), p.GetStock())
    }
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    inv := inventory.NewInventory()

    for {
        showMenu()
        choice, err := readInt(reader, "")
        if err != nil {
            fmt.Println("Invalid input. Please try again.")
            continue
        }

        switch choice {
        case 1:
            addProduct(reader, inv)

        case 2:
            id, err := readInt(reader, "Enter Product ID: ")
            if err != nil {
                fmt.Println("Invalid ID format")
                continue
            }

            quantity, err := readInt(reader, "Enter New Stock Quantity: ")
            if err != nil {
                fmt.Println("Invalid quantity format")
                continue
            }

            err = inv.UpdateStock(id, quantity)
            if err != nil {
                fmt.Printf("Error: %v\n", err)
            } else {
                fmt.Println("Stock updated successfully!")
            }

        case 3:
            id, err := readInt(reader, "Enter Product ID to search: ")
            if err != nil {
                fmt.Println("Invalid ID format")
                continue
            }

            product, err := inv.SearchByID(id)
            if err != nil {
                fmt.Printf("Error: %v\n", err)
            } else {
                displayProducts([]*inventory.Product{product})
            }

        case 4:
            name := readString(reader, "Enter Product Name to search: ")
            products := inv.SearchByName(name)
            displayProducts(products)

        case 5:
            displayProducts(inv.GetAllProducts())

        case 6:
            displayProducts(inv.GetProductsSortedByPrice())

        case 7:
            displayProducts(inv.GetProductsSortedByStock())

        case 8:
            fmt.Println("Thank you for using the Inventory Management System. Goodbye!")
            return

        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}