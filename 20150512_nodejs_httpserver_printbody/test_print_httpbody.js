var http = require('http');

http.createServer(function (req, res) {
	var resbody = '';
	req.on('data', function(chunk) {
		resbody += chunk;
	});
	req.on('end', function() {
		console.log('recv body:', resbody);
	});
	
	res.writeHead(200, {'Content-Type': 'text/plain'});
	res.end('Hello World\n');
}).listen(1337);

console.log('Server running at http://0.0.0.0:1337/');