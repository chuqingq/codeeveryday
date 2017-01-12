var port1 = 20022;
var port2 = 20023;


var net = require('net');

var sock1, sock2;

connectSock1();

function connectSock1() {
	var server1 = net.createServer(function(sock) {
		console.log('sock1 connected');
		sock1 = sock;
		sock1.on('error', handleSock1Error);
		sock1.on('end', function() {
			console.log('sock1 end');
                        sock1.end();
			sock1 = undefined;
		});

		if (sock1 != undefined && sock2 != undefined) {
			sock1.pipe(sock2);
			sock2.pipe(sock1);
		}
	});

	server1.listen(port1);
}

function handleSock1Error() {
	console.log('sock1 error');
	sock1 = undefined;
}




connectSock2();

function connectSock2() {
	var server2 = net.createServer(function(sock) {
		console.log('sock2 connected');
		sock2 = sock;
		sock2.on('error', handleSock2Error);
		sock2.on('end', function() {
			console.log('sock2 end');
                        sock2.end();
			sock2 = undefined;
		});

		if (sock1 != undefined && sock2 != undefined) {
			sock1.pipe(sock2);
			sock2.pipe(sock1);
		}
	});

	server2.listen(port2);
}

function handleSock2Error() {
	console.log('sock2 error');
	sock2 = undefined;
}


