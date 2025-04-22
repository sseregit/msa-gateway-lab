const express = require('express');
const app = express();

app.use(express.json()); // JSON 바디 파싱

// ✅ /api/query?name=jang&age=34
app.get('/api/query', (req, res) => {
    const { name, age } = req.query;
    res.json({ message: `Received query for name=${name}, age=${age}` });
});

// ✅ /api/user/:id
app.get('/api/user/:id', (req, res) => {
    const { id } = req.params;
    res.json({ message: `Fetching user with id=${id}` });
});

// ✅ /api/create
app.post('/api/create', (req, res) => {
    const { name, age } = req.body;
    res.status(201).json({ message: `Created user ${name} aged ${age}` });
});

// ✅ /api/delete/:id
app.delete('/api/delete', (req, res) => {
    const { id } = req.body;
    res.json({ message: `Deleted user with id=${id}` });
});

// ✅ /api/update/:id
app.put('/api/update', (req, res) => {
    const { name, age, id } = req.body;
    res.json({ message: `Updated user ${id} to name=${name}, age=${age}` });
});

// ✅ 서버 실행
const PORT = 3000;
app.listen(PORT, () => {
    console.log(`Server listening on http://localhost:${PORT}`);
});