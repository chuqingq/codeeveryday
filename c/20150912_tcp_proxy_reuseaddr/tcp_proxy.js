var localHost = '127.0.0.1';
var localPort = 22;

var remoteHost = '127.0.0.1';
var remotePort = 20022;


var net = require('net');

var sock1, sock2;

connectSock1();

function connectSock1() {
	sock1 = net.connect(remotePort, remoteHost, function() {
		console.log('sock1 connected');
		sock1.on('end', function() {
			sock1.end();
			sock1 = undefined;
			setTimeout(connectSock1, 3000);
			console.log('sock1 end');
		});

		if (sock1 != undefined && sock2 != undefined) {
			sock1.pipe(sock2);
			sock2.pipe(sock1);
		}
	});
	sock1.on('error', handleSock1Error);
}

function handleSock1Error() {
	// sock1.end();
	sock1 = undefined;
	setTimeout(connectSock1, 3000);
	console.log('sock1 error');
}


connectSock2();

function connectSock2() {
	sock2 = net.connect(localPort, localHost, function() {
		console.log('sock2 connected');
		sock2.on('end', function() {
			sock2.end();
			sock2 = undefined;
			setTimeout(connectSock2, 3000);
			console.log('sock2 end');
		});

		if (sock1 != undefined && sock2 != undefined) {
			sock1.pipe(sock2);
			sock2.pipe(sock1);
		}
	});
	sock2.on('error', handleSock2Error);
}

function handleSock2Error() {
	sock2.end();
	sock2 = undefined;
	setTimeout(connectSock2, 3000);
	console.log('sock2 error');
}

