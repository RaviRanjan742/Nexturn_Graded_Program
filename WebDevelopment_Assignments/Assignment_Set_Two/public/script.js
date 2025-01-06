let chart = null;


const categoryColors = {
    Food: '#FF6384',
    Travel: '#36A2EB',
    Shopping: '#FFCE56',
    Bills: '#4BC0C0',
    Entertainment: '#9966FF'
};


async function fetchExpenses() {
    try {
        const response = await fetch('/api/expenses');
        const expenses = await response.json();
        updateExpenseTable(expenses);
        updateChart();
    } catch (error) {
        console.error('Error fetching expenses:', error);
    }
}


document.getElementById('expenseForm').addEventListener('submit', async (e) => {
    e.preventDefault();
    
    const expense = {
        amount: parseFloat(document.getElementById('amount').value),
        description: document.getElementById('description').value,
        category: document.getElementById('category').value
    };

    try {
        await fetch('/api/expenses', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(expense)
        });

        
        e.target.reset();
        
        
        fetchExpenses();
    } catch (error) {
        console.error('Error adding expense:', error);
    }
});


async function deleteExpense(id) {
    try {
        await fetch(`/api/expenses/${id}`, {
            method: 'DELETE'
        });
        fetchExpenses();
    } catch (error) {
        console.error('Error deleting expense:', error);
    }
}


function updateExpenseTable(expenses) {
    const tbody = document.querySelector('#expenseTable tbody');
    tbody.innerHTML = '';

    expenses.forEach(expense => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>$${expense.amount.toFixed(2)}</td>
            <td>${expense.description}</td>
            <td>
                <span class="category-badge" style="background-color: ${categoryColors[expense.category]}">
                    ${expense.category}
                </span>
            </td>
            <td>
                <button class="delete-btn" onclick="deleteExpense(${expense.id})">Delete</button>
            </td>
        `;
        tbody.appendChild(row);
    });
}


async function updateChart() {
    try {
        const response = await fetch('/api/summary');
        const data = await response.json();

        const ctx = document.getElementById('chart-container');
        
        if (chart) {
            chart.destroy();
        }

        chart = new Chart(ctx, {
            type: 'pie',
            data: {
                labels: data.map(item => item.category),
                datasets: [{
                    data: data.map(item => item.total),
                    backgroundColor: Object.values(categoryColors)
                }]
            },
            options: {
                responsive: true,
                maintainAspectRatio: false,
                plugins: {
                    legend: {
                        position: 'right'
                    }
                }
            }
        });
    } catch (error) {
        console.error('Error updating chart:', error);
    }
}


fetchExpenses();