
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "Assignment_two/banking"
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
    fmt.Println("\n=== Bank Transaction System ===")
    fmt.Println("1. Create New Account")
    fmt.Println("2. Deposit Money")
    fmt.Println("3. Withdraw Money")
    fmt.Println("4. Check Balance")
    fmt.Println("5. View Transaction History")
    fmt.Println("6. List All Accounts")
    fmt.Println("7. Exit")
    fmt.Print("Enter your choice: ")
}

func createAccount(reader *bufio.Reader, bank *banking.Bank) {
    fmt.Println("\n--- Create New Account ---")
    
    id, err := readInt(reader, "Enter Account ID: ")
    if err != nil {
        fmt.Println("Invalid ID format")
        return
    }

    name := readString(reader, "Enter Account Holder Name: ")
    if name == "" {
        fmt.Println("Name cannot be empty")
        return
    }

    initialBalance, err := readFloat(reader, "Enter Initial Balance: ")
    if err != nil {
        fmt.Println("Invalid amount format")
        return
    }

    account := banking.NewAccount(id, name, initialBalance)
    err = bank.AddAccount(account)
    if err != nil {
        fmt.Printf("Error creating account: %v\n", err)
        return
    }

    fmt.Println("Account created successfully!")
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    bank := banking.NewBank()

    for {
        showMenu()
        choice, err := readInt(reader, "")
        if err != nil {
            fmt.Println("Invalid input. Please try again.")
            continue
        }

        switch choice {
        case 1:
            createAccount(reader, bank)

        case 2:
            id, err := readInt(reader, "Enter Account ID: ")
            if err != nil {
                fmt.Println("Invalid ID format")
                continue
            }

            amount, err := readFloat(reader, "Enter Deposit Amount: ")
            if err != nil {
                fmt.Println("Invalid amount format")
                continue
            }

            if err := bank.Deposit(id, amount); err != nil {
                fmt.Printf("Error: %v\n", err)
            } else {
                fmt.Println("Deposit successful!")
            }

        case 3:
            id, err := readInt(reader, "Enter Account ID: ")
            if err != nil {
                fmt.Println("Invalid ID format")
                continue
            }

            amount, err := readFloat(reader, "Enter Withdrawal Amount: ")
            if err != nil {
                fmt.Println("Invalid amount format")
                continue
            }

            if err := bank.Withdraw(id, amount); err != nil {
                fmt.Printf("Error: %v\n", err)
            } else {
                fmt.Println("Withdrawal successful!")
            }

        case 4:
            id, err := readInt(reader, "Enter Account ID: ")
            if err != nil {
                fmt.Println("Invalid ID format")
                continue
            }

            if balance, err := bank.GetBalance(id); err != nil {
                fmt.Printf("Error: %v\n", err)
            } else {
                fmt.Printf("Current Balance: $%.2f\n", balance)
            }

        case 5:
            id, err := readInt(reader, "Enter Account ID: ")
            if err != nil {
                fmt.Println("Invalid ID format")
                continue
            }

            if history, err := bank.GetTransactionHistory(id); err != nil {
                fmt.Printf("Error: %v\n", err)
            } else {
                fmt.Println("\nTransaction History:")
                for _, transaction := range history {
                    fmt.Println(transaction)
                }
            }

        case 6:
            accounts := bank.ListAccounts()
            if len(accounts) == 0 {
                fmt.Println("No accounts found")
            } else {
                fmt.Println("\nAll Accounts:")
                for _, acc := range accounts {
                    fmt.Printf("ID: %d, Name: %s, Balance: $%.2f\n", 
                        acc.GetID(), acc.GetName(), acc.GetBalance())
                }
            }

        case 7:
            fmt.Println("Thank you for using our banking system. Goodbye!")
            return

        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}