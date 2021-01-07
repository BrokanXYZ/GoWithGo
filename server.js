const express = require('express');
const path = require('path');
const app = express();

// prod
app.use(express.static(path.join(__dirname, 'build')));

// local
app.use(express.static(path.join(__dirname, 'public')));

app.get('/', function (req, res) {
  res.sendFile(path.join(__dirname, 'build', 'index.html'));
});

const port = process.env.PORT || 8080;

app.listen(port);
console.log(`Listening on ${port}`);