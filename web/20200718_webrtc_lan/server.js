var app = require('express')();
var server = require('http').createServer(app);
var io = require('socket.io').listen(server);
app.get('/', (req, res) => {
	res.sendFile(__dirname+'/index.html')
})


io.on('connection', socket => {
	console.log('websocket connected')
	socket.on('candidate', candidate => {
		console.log('websocket recv candidate: ', candidate)
		socket.broadcast.emit('candidate', candidate)
	})
	socket.on('offer', offer => {
		console.log('websocket recv offer: ', offer)
		socket.broadcast.emit('offer', offer)
	})
	socket.on('answer', answer => {
		console.log('websocket recv answer: ', answer)
		socket.broadcast.emit('answer', answer)
	})

})



var port = 8000
server.listen(port, () => console.log('listen at', port))


