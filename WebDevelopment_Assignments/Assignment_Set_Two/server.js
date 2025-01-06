const express = require('express');
const sqlite3 = require('sqlite3').verbose();
const path = require('path');
const app = express();


app.use(express.json());
app.use(express.static('public'));


const db = new sqlite3.Database('expenses.db', (err) => {
    if (err) console.error(err.message);
    console.log('Connected to the expenses database.');
});


db.run(`
    CREATE TABLE IF NOT EXISTS expenses (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        amount DECIMAL(10,2) NOT NULL,
        description TEXT,
        category TEXT NOT NULL,
        date TEXT DEFAULT CURRENT_TIMESTAMP
    )
`);


app.get('/api/expenses', (req, res) => {
    db.all('SELECT * FROM expenses ORDER BY date DESC', [], (err, rows) => {
        if (err) {
            res.status(500).json({ error: err.message });
            return;
        }
        res.json(rows);
    });
});

app.post('/api/expenses', (req, res) => {
    const { amount, description, category } = req.body;
    db.run(
        'INSERT INTO expenses (amount, description, category) VALUES (?, ?, ?)',
        [amount, description, category],
        function(err) {
            if (err) {
                res.status(500).json({ error: err.message });
                return;
            }
            res.json({
                id: this.lastID,
                amount,
                description,
                category
            });
        }
    );
});

app.delete('/api/expenses/:id', (req, res) => {
    db.run('DELETE FROM expenses WHERE id = ?', req.params.id, function(err) {
        if (err) {
            res.status(500).json({ error: err.message });
            return;
        }
        res.json({ message: "deleted", changes: this.changes });
    });
});

app.get('/api/summary', (req, res) => {
    db.all(`
        SELECT category, SUM(amount) as total
        FROM expenses
        GROUP BY category
    `, [], (err, rows) => {
        if (err) {
            res.status(500).json({ error: err.message });
            return;
        }
        res.json(rows);
    });
});

const port = 3000;
app.listen(port, () => {
    console.log(`Server running at http://localhost:${port}`);
});