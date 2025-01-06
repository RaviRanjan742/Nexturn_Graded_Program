const express = require('express');
const sqlite3 = require('sqlite3').verbose();
const bcrypt = require('bcrypt');
const jwt = require('jsonwebtoken');
const cors = require('cors');
const app = express();


app.use(express.json());
app.use(cors());
app.use(express.static('public'));


const db = new sqlite3.Database('agile.db', (err) => {
    if (err) console.error(err.message);
    console.log('Connected to the agile database.');
});


db.serialize(() => {
    
    db.run(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT UNIQUE NOT NULL,
        password TEXT NOT NULL,
        email TEXT UNIQUE NOT NULL
    )`);

    
    db.run(`CREATE TABLE IF NOT EXISTS comments (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER,
        content TEXT NOT NULL,
        section TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY(user_id) REFERENCES users(id)
    )`);

    
    db.run(`CREATE TABLE IF NOT EXISTS forum_posts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER,
        title TEXT NOT NULL,
        content TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY(user_id) REFERENCES users(id)
    )`);
});


const authenticateToken = (req, res, next) => {
    const token = req.header('Authorization')?.split(' ')[1];
    if (!token) return res.sendStatus(401);

    jwt.verify(token, 'your_jwt_secret', (err, user) => {
        if (err) return res.sendStatus(403);
        req.user = user;
        next();
    });
};


app.post('/api/register', async (req, res) => {
    try {
        const { username, password, email } = req.body;
        const hashedPassword = await bcrypt.hash(password, 10);
        
        db.run('INSERT INTO users (username, password, email) VALUES (?, ?, ?)',
            [username, hashedPassword, email],
            function(err) {
                if (err) {
                    return res.status(400).json({ error: err.message });
                }
                res.status(201).json({ id: this.lastID });
            });
    } catch (error) {
        res.status(500).json({ error: error.message });
    }
});


app.post('/api/login', async (req, res) => {
    const { username, password } = req.body;
    
    db.get('SELECT * FROM users WHERE username = ?', [username], async (err, user) => {
        if (err) return res.status(500).json({ error: err.message });
        if (!user) return res.status(400).json({ error: 'User not found' });
        
        const valid = await bcrypt.compare(password, user.password);
        if (!valid) return res.status(400).json({ error: 'Invalid password' });
        
        const token = jwt.sign({ id: user.id, username }, 'your_jwt_secret');
        res.json({ token });
    });
});


app.get('/api/comments/:section', (req, res) => {
    db.all(`
        SELECT comments.*, users.username
        FROM comments
        JOIN users ON comments.user_id = users.id
        WHERE comments.section = ?
        ORDER BY comments.created_at DESC
    `, [req.params.section], (err, rows) => {
        if (err) return res.status(500).json({ error: err.message });
        res.json(rows);
    });
});


app.post('/api/comments', authenticateToken, (req, res) => {
    const { content, section } = req.body;
    db.run(
        'INSERT INTO comments (user_id, content, section) VALUES (?, ?, ?)',
        [req.user.id, content, section],
        function(err) {
            if (err) return res.status(500).json({ error: err.message });
            res.json({ id: this.lastID });
        }
    );
});


app.get('/api/forum', (req, res) => {
    db.all(`
        SELECT forum_posts.*, users.username
        FROM forum_posts
        JOIN users ON forum_posts.user_id = users.id
        ORDER BY forum_posts.created_at DESC
    `, [], (err, rows) => {
        if (err) return res.status(500).json({ error: err.message });
        res.json(rows);
    });
});


app.post('/api/forum', authenticateToken, (req, res) => {
    const { title, content } = req.body;
    db.run(
        'INSERT INTO forum_posts (user_id, title, content) VALUES (?, ?, ?)',
        [req.user.id, title, content],
        function(err) {
            if (err) return res.status(500).json({ error: err.message });
            res.json({ id: this.lastID });
        }
    );
});

const port = 3000;
app.listen(port, () => {
    console.log(`Server running at http://localhost:${port}`);
});