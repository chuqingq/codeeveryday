var process = require('process');

var wechat = require('wechat');
var weapi = require('wechat-api');

var mongodb = require('mongodb');
var connect = require('express');
var fecha = require('fecha');

var app = connect();

process.on('uncaughtException', function(err) {
    console.log('uncaughtException:', err);
});

var config = {
  token: 'chuqq',
  appid: 'wx0288bf03ed5da89b',
  appsecret: 'd4624c36b6795d1d99dcf0547af5443d'
};

var api = new weapi(config.appid, config.appsecret);

var OAuth = require('wechat-oauth');
var oauth = new OAuth(config.appid, config.appsecret);

var collection;

// 打开mongo，并删除已超期的任务
mongodb.MongoClient.connect('mongodb://localhost:27017/weixin', function(err, db) {
  if (err) {
    console.log('mongodb.MongoClient.connect error:', err);
    return;
  }
  console.log("MongoClient.connect success");

  collection = db.collection('remind');
  if (!collection) {
    console.log('db.collection("remind") error: collection is null');
    return;
  }

  // 每10秒，遍历mongo中过期的数据，并发送提醒
  setInterval(function() {
    // console.log('setInterval')
    collection.find({time: {$lt: new Date()}, ishandled: false}).toArray(function(err, docs) {
      if (err) {
        console.log('collection.find.toArray error:', err);
        return;
      }
      // console.log('docs:', docs);
      docs.forEach(function(doc) {
        console.log('send doc:', doc);
        var data = {
            time: {value: fecha.format(doc.time, 'YYYY-MM-DD HH:mm:ss'), color: '#173177'},
            content: {value: doc.content, color: '#173177'}
        };
        // api.sendText(doc.user, doc.content, function(err, data, res) {
        api.sendTemplate(doc.user, 'n3lQoXJNPH01DuVLkkeBajv0BIpJXUKAWQUSIbLYWHA', 'http://www.baidu.com', data, function(err, data, res) {
          if (err) {
            console.log('api.sendTemplate error: ', err, data, res);
          }
          collection.updateOne({_id: mongodb.ObjectId(doc._id)}, {'$set': {ishandled: true}}, function(err) {
            console.log('delete doc:', doc);
          });
        });
      });
    });
  }, 10*1000);
});

app.use(connect.query());
app.use('/wechat', wechat(config, wechat.text(function(message, req, res, next) {
  console.log('/wechat text:', message.Content);
  handleMsg(message.FromUserName, message.Content, res);
}).voice(function(message, req, res, next) {
  console.log("/wechat voice:", message.Recognition);
  handleMsg(message.FromUserName, message.Recognition, res);
}).event(function(message, req, res, next) {
  console.log('/wechat event: ', message);
  var msg = '';
  if (message.EventKey == 'SETREMIND_10') {
    msg = '10分钟后提醒我';
  } else if (message.EventKey == 'SETREMIND_30') {
    msg = '30分钟后提醒我';
  } else if (message.EventKey == 'SETREMIND_60') {
    msg = '60分钟后提醒我';
  } else {
    msg = '不支持的event: ' + message.EventKey;
  }
  handleMsg(message.FromUserName, msg, res);
})));

// text和voice均调此接口处理remind消息
function handleMsg(user, content, res) {
  var opts = {
    "query": content,
    "city": "南京",
    "category": "remind"
  };

  api.semantic(user, opts, function(err, data, res2) {
    if (err) {
      console.log('api.semantic error: ', err, data, res2);
      // return;
    }
    console.log('semantic result: ', data);

    var datetime;
    if (err != null
      || data == undefined
      || data.semantic == undefined
      || data.semantic.details == undefined
      || data.semantic.details.datetime == undefined) {
        console.log('semantic error: ' + err + '.\n' + content);
        var array = content.split(' ', 3);
        if (array.length == 3) {
            datetime = array[0] + ' ' + array[1];
        } else if (array.length == 2) {
            datetime = array[0];
        } else {
            res.reply('未识别时间\n' + fecha.format(new Date(), 'YYYY-MM-DD HH:mm:ss'));
            return;
        }
    } else {
        console.log('datetime: ', data.semantic.details.datetime);
        datetime = data.semantic.details.datetime.date + ' ' + data.semantic.details.datetime.time;
        console.log('semantic datetime: ' + datetime);
    }
    var datetime2 = fecha.parse(datetime, 'YYYY-MM-DD HH:mm:ss');
    if (datetime2.toString() == 'Invalid Date') {
      res.reply('Time is invalid format\n' + fecha.format(new Date(), 'YYYY-MM-DD HH:mm:ss'));
      return;
    }

    if (datetime2 < new Date()) {
      res.reply('Time is overdue\n' + datetime + ' before ' + fecha.format(new Date(), 'YYYY-MM-DD HH:mm:ss'));
      return;
    }

    // 插入数据库中
    var doc = {time: datetime2, user: user, content: content, ishandled: false};
    collection.insert(doc, function(err) {
      if (err) {
        console.log('Server error: ', err);
        res.reply('Server error: ' + err);
        return;
      }

      console.log('insert doc:', doc, '::', err);
      res.reply('Set timer success!\nTime: '+datetime+'\nContent: '+content);
    });
  });
}

//app.use('/', function(req, res) {
//  res.end('Hello World!');
//});
app.use('/remind/new', function(req, res) {
  console.log('/remind/new:'+req.query.code);
  oauth.getAccessToken(req.query.code, function(err, result) {
    if (err) {
      console.log("getUser error: " + err);
    }
    console.log('getUser result: ' + JSON.stringify(result));
  });
  console.log('query: '+ JSON.stringify(req.query));
  res.end('/remind/new');
});
app.use('/remind/save', function(req, res) {
  console.log('/remind/save:');
  res.end('/remind/save');
});
app.use('/remind/get', function(req, res) {
  console.log('/remind/get:');
  res.end('/remind/get');
});


// 启动服务
var server = app.listen(80, function() {
  var host = server.address().address;
  var port = server.address().port;

  console.log('weixin app listening at http://%s:%s', host, port);
});

