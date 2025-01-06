package employee

import (
    "errors"
    "fmt"
)

type Employee struct {
    ID         int
    Name       string
    Age        int
    Department string
}

func NewEmployee(id int, name string, age int, department string) (*Employee, error) {
    if age < MinAge {
        return nil, errors.New("employee must be at least 18 years old")
    }

    if department != DepartmentIT && department != DepartmentHR {
        return nil, errors.New("invalid department")
    }

    return &Employee{
        ID:         id,
        Name:       name,
        Age:        age,
        Department: department,
    }, nil
}

func (e *Employee) String() string {
    return fmt.Sprintf("ID: %d, Name: %s, Age: %d, Department: %s", 
        e.ID, e.Name, e.Age, e.Department)
}