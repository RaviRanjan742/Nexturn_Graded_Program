
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "Assignment_Set_One/employee"
)

func readString(reader *bufio.Reader, prompt string) string {
    fmt.Print(prompt)
    str, _ := reader.ReadString('\n')
    return strings.TrimSpace(str)
}

func readInt(reader *bufio.Reader, prompt string) (int, error) {
    str := readString(reader, prompt)
    return strconv.Atoi(str)
}

func showMenu() {
    fmt.Println("\n=== Employee Management System ===")
    fmt.Println("1. Add Employee")
    fmt.Println("2. Search Employee by ID")
    fmt.Println("3. Search Employee by Name")
    fmt.Println("4. List Employees by Department")
    fmt.Println("5. Count Employees by Department")
    fmt.Println("6. Exit")
    fmt.Print("Enter your choice: ")
}

func addEmployee(reader *bufio.Reader, empService *employee.EmployeeService) {
    fmt.Println("\n--- Add New Employee ---")
    
    
    id, err := readInt(reader, "Enter Employee ID: ")
    if err != nil {
        fmt.Println("Invalid ID format")
        return
    }

    
    name := readString(reader, "Enter Employee Name: ")
    if name == "" {
        fmt.Println("Name cannot be empty")
        return
    }

    
    age, err := readInt(reader, "Enter Employee Age: ")
    if err != nil {
        fmt.Println("Invalid age format")
        return
    }

    
    fmt.Printf("Available Departments: %s, %s\n", employee.DepartmentIT, employee.DepartmentHR)
    department := readString(reader, "Enter Employee Department: ")

    
    emp, err := employee.NewEmployee(id, name, age, department)
    if err != nil {
        fmt.Printf("Error creating employee: %v\n", err)
        return
    }

    
    err = empService.AddEmployee(emp)
    if err != nil {
        fmt.Printf("Error adding employee: %v\n", err)
        return
    }

    fmt.Println("Employee added successfully!")
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    empService := employee.NewEmployeeService()

    for {
        showMenu()
        choice, err := readInt(reader, "")
        if err != nil {
            fmt.Println("Invalid input. Please try again.")
            continue
        }

        switch choice {
        case 1:
            addEmployee(reader, empService)

        case 2:
            id, err := readInt(reader, "Enter Employee ID to search: ")
            if err != nil {
                fmt.Println("Invalid ID format")
                continue
            }
            if emp, err := empService.SearchByID(id); err != nil {
                fmt.Printf("Error: %v\n", err)
            } else {
                fmt.Printf("Found: %s\n", emp.String())
            }

        case 3:
            name := readString(reader, "Enter Employee Name to search: ")
            if emp, err := empService.SearchByName(name); err != nil {
                fmt.Printf("Error: %v\n", err)
            } else {
                fmt.Printf("Found: %s\n", emp.String())
            }

        case 4:
            fmt.Printf("Available Departments: %s, %s\n", employee.DepartmentIT, employee.DepartmentHR)
            dept := readString(reader, "Enter Department: ")
            employees := empService.ListByDepartment(dept)
            if len(employees) == 0 {
                fmt.Printf("No employees found in %s department\n", dept)
            } else {
                fmt.Printf("\nEmployees in %s department:\n", dept)
                for _, emp := range employees {
                    fmt.Println(emp.String())
                }
            }

        case 5:
            fmt.Printf("Available Departments: %s, %s\n", employee.DepartmentIT, employee.DepartmentHR)
            dept := readString(reader, "Enter Department: ")
            count := empService.CountByDepartment(dept)
            fmt.Printf("Number of employees in %s: %d\n", dept, count)

        case 6:
            fmt.Println("Goodbye!")
            return

        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}