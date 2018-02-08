const http = require('http');
const port = 3000;

const handler = (req, res) => {
    res.writeHead(200, { 'Content-Type': 'text/html' });
    res.end('<h1>Hello</h1>');
};

const server = http.createServer(handler)

server.listen(port, 'localhost', err => err ? console.error(e) : console.log('Started server'))