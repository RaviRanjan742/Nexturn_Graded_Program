package employee

import (
    "errors"
    "strings"
)

type EmployeeService struct {
    employees []*Employee
}

func NewEmployeeService() *EmployeeService {
    return &EmployeeService{
        employees: make([]*Employee, 0),
    }
}

func (s *EmployeeService) AddEmployee(emp *Employee) error {
    
    for _, e := range s.employees {
        if e.ID == emp.ID {
            return errors.New("employee ID already exists")
        }
    }

    s.employees = append(s.employees, emp)
    return nil
}

func (s *EmployeeService) SearchByID(id int) (*Employee, error) {
    for _, emp := range s.employees {
        if emp.ID == id {
            return emp, nil
        }
    }
    return nil, errors.New("employee not found")
}

func (s *EmployeeService) SearchByName(name string) (*Employee, error) {
    for _, emp := range s.employees {
        if strings.EqualFold(emp.Name, name) {
            return emp, nil
        }
    }
    return nil, errors.New("employee not found")
}

func (s *EmployeeService) ListByDepartment(department string) []*Employee {
    result := make([]*Employee, 0)
    for _, emp := range s.employees {
        if emp.Department == department {
            result = append(result, emp)
        }
    }
    return result
}

func (s *EmployeeService) CountByDepartment(department string) int {
    count := 0
    for _, emp := range s.employees {
        if emp.Department == department {
            count++
        }
    }
    return count
}
