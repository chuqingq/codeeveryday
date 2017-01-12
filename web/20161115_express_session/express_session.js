// npm init
// npm install express
// npm install express-session

var express = require('express')
var session = require('express-session')

var app = express()

// Use the session middleware
app.use(session({ secret: 'keyboard cat', cookie: { maxAge: 5000 }}))

// Access the session as req.session
app.get('/', function(req, res, next) {
  var sess = req.session
  if (sess.views) {
    sess.views++
    res.setHeader('Content-Type', 'text/html')
    res.write('<p>views: ' + sess.views + '</p>')
    res.write('<p>expires in: ' + (sess.cookie.maxAge / 1000) + 's</p>')
    res.end()
  } else {
    sess.views = 1
    res.end('welcome to the session demo. refresh!')
  }
})

app.listen(3000)
console.log('Listening on port 3000')

// result: sess.cookie.maxAge表示5秒内如果不访问页面，则超时；如果访问，则重新计时5秒。
