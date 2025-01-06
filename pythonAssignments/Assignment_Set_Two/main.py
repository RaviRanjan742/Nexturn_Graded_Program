
from book_management import BookManager
from customer_management import CustomerManager
from sales_management import SalesManager

class BookMartSystem:
    def __init__(self):
        self.book_manager = BookManager()
        self.customer_manager = CustomerManager()
        self.sales_manager = SalesManager(self.book_manager, self.customer_manager)

    def book_menu(self):
        print("\nBook Management:")
        print("1. Add Book")
        print("2. View Books")
        print("3. Search Book")
        print("4. Back to Main Menu")
        
        try:
            choice = int(input("Enter your choice: "))
            
            if choice == 1:
                title = input("Enter title: ")
                author = input("Enter author: ")
                price = float(input("Enter price: "))
                quantity = int(input("Enter quantity: "))
                if self.book_manager.add_book(title, author, price, quantity):
                    print("Book added successfully!")
                    
            elif choice == 2:
                books = self.book_manager.view_books()
                if not books:
                    print("No books in inventory")
                for book in books:
                    print("\n" + book.display())
                    
            elif choice == 3:
                title = input("Enter book title to search: ")
                book = self.book_manager.search_book(title)
                if book:
                    print("\n" + book.display())
                else:
                    print("Book not found")
                    
            elif choice == 4:
                return
            
            else:
                print("Invalid choice!")
                
        except ValueError as e:
            print(f"Invalid input: {e}")

    def customer_menu(self):
        print("\nCustomer Management:")
        print("1. Add Customer")
        print("2. View Customers")
        print("3. Back to Main Menu")
        
        try:
            choice = int(input("Enter your choice: "))
            
            if choice == 1:
                name = input("Enter name: ")
                email = input("Enter email: ")
                phone = input("Enter phone: ")
                if self.customer_manager.add_customer(name, email, phone):
                    print("Customer added successfully!")
                    
            elif choice == 2:
                customers = self.customer_manager.view_customers()
                if not customers:
                    print("No customers registered")
                for customer in customers:
                    print("\n" + customer.display())
                    
            elif choice == 3:
                return
            
            else:
                print("Invalid choice!")
                
        except ValueError as e:
            print(f"Invalid input: {e}")

    def sales_menu(self):
        print("\nSales Management:")
        print("1. Sell Book")
        print("2. View Sales")
        print("3. Back to Main Menu")
        
        try:
            choice = int(input("Enter your choice: "))
            
            if choice == 1:
                name = input("Enter customer name: ")
                email = input("Enter customer email: ")
                phone = input("Enter customer phone: ")
                book_title = input("Enter book title: ")
                quantity = int(input("Enter quantity: "))
                if self.sales_manager.sell_book(name, email, phone, book_title, quantity):
                    print("Sale completed successfully!")
                    
            elif choice == 2:
                sales = self.sales_manager.view_sales()
                if not sales:
                    print("No sales recorded")
                for sale in sales:
                    print("\n" + sale.display())
                    
            elif choice == 3:
                return
            
            else:
                print("Invalid choice!")
                
        except ValueError as e:
            print(f"Invalid input: {e}")

    def run(self):
        while True:
            print("\nWelcome to BookMart!")
            print("1. Book Management")
            print("2. Customer Management")
            print("3. Sales Management")
            print("4. Exit")
            
            try:
                choice = int(input("Enter your choice: "))
                
                if choice == 1:
                    self.book_menu()
                elif choice == 2:
                    self.customer_menu()
                elif choice == 3:
                    self.sales_menu()
                elif choice == 4:
                    print("Thank you for using BookMart!")
                    break
                else:
                    print("Invalid choice!")
                    
            except ValueError as e:
                print(f"Invalid input: {e}")
            except Exception as e:
                print(f"An error occurred: {e}")

if __name__ == "__main__":
    system = BookMartSystem()
    system.run()