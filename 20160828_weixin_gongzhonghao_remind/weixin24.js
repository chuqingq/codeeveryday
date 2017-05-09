var process = require('process');

var wechat = require('wechat');
var weapi = require('wechat-api');

var mongodb = require('mongodb');
var connect = require('express');
var fecha = require('fecha');

var fs = require('fs')
  , Log = require('log')
  , log = new Log('debug', fs.createWriteStream('my.log'));

var config = require('./package.json').config;

var app = connect();

process.on('uncaughtException', function(err) {
    log.error('uncaughtException:', err);
});

var api = new weapi(config.appid, config.appsecret);

// 设置菜单
api.createMenu(config.menu, function(err, res) {
  if (err) {
    log.error('菜单设置失败：', err);
  } else {
    log.info('菜单设置成功');
  }
});

var OAuth = require('wechat-oauth');
var oauth = new OAuth(config.appid, config.appsecret);

var collection;

// 打开mongo，并删除已超期的任务
mongodb.MongoClient.connect('mongodb://localhost:27017/weixin', function(err, db) {
  if (err) {
    log.error('mongodb.MongoClient.connect error:', err);
    return;
  }
  log.info("MongoClient.connect success");

  collection = db.collection('remind');
  if (!collection) {
    log.error('db.collection("remind") error: collection is null');
    return;
  }

  // 每10秒，遍历mongo中过期的数据，并发送提醒
  setInterval(function() {
    // console.log('setInterval')
    collection.find({time: {$lt: new Date()}, ishandled: false}).toArray(function(err, docs) {
      if (err) {
        log.error('collection.find.toArray error:', err);
        return;
      }
      // console.log('docs:', docs);
      docs.forEach(function(doc) {
        log.info('send doc:', doc);
        var data = {
            time: {value: fecha.format(doc.time, 'YYYY-MM-DD HH:mm:ss'), color: '#173177'},
            content: {value: doc.content, color: '#173177'}
        };
        // api.sendText(doc.user, doc.content, function(err, data, res) {
        api.sendTemplate(doc.user, 'BCN7n5qe41QjX3-b13Z7ZBqfoGpE1UpePga8uY2CHEE', 'http://www.baidu.com', data, function(err, data, res) {
          if (err) {
            log.error('api.sendTemplate error: ', err, data, res);
          }
          collection.updateOne({_id: mongodb.ObjectId(doc._id)}, {'$set': {ishandled: true}}, function(err) {
            log.debug('delete doc:', doc);
          });
        });
      });
    });
  }, 10*1000);
});

app.use(connect.query());
app.use(config.urlprefix, wechat(config, wechat.text(function(message, req, res, next) {
  log.debug('/wechat text:', message.Content);
  handleMsg(message.FromUserName, message.Content, res);
}).voice(function(message, req, res, next) {
  log.debug("/wechat voice:", message.Recognition);
  handleMsg(message.FromUserName, message.Recognition, res);
}).event(function(message, req, res, next) {
  log.debug('/wechat event: ', message);
  if (message.Event != 'CLICK') {
    log.debug('not expected event: ', message.Event);
    return;
  }
  var match = message.EventKey.match(/^SETREMIND_([0-9]+)$/)
  if (!res) {
    log.warn('非预期的eventkey: ', message.EventKey);
    res.reply('非预期的eventkey: ' + message.EventKey)
    return;
  }
  var msg = match[1] + '分钟后提醒我';
  log.debug('msg: ' + msg);
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
      log.error('api.semantic error: ', err, data, res2);
      // return;
    }
    // log.debug('semantic result: ', data);

    var datetime;
    if (err != null
      || data == undefined
      || data.semantic == undefined
      || data.semantic.details == undefined
      || data.semantic.details.datetime == undefined) {
        log.error('semantic error: ' + err + '.\n' + content);
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
        // log.debug('datetime: ', data.semantic.details.datetime);
        datetime = data.semantic.details.datetime.date + ' ' + data.semantic.details.datetime.time;
        // log.debug('semantic datetime: ' + datetime);
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
        log.error('collection.insert error: ', err);
        res.reply('Server error: ' + err);
        return;
      }

      log.debug('insert doc:', doc, '::', err);
      res.reply('提醒设置成功！\n时间：'+datetime+'\n内容：'+content);
    });
  });
}

//app.use('/', function(req, res) {
//  res.end('Hello World!');
//});
app.use('/remind/new', function(req, res) {
  log.debug('/remind/new:'+req.query.code);
  oauth.getAccessToken(req.query.code, function(err, result) {
    if (err) {
      log.error("getUser error: " + err);
    }
    log.debug('getUser result: ' + JSON.stringify(result));
  });
  log.debug('query: '+ JSON.stringify(req.query));
  res.end('/remind/new');
});
app.use('/remind/save', function(req, res) {
  log.debug('/remind/save:');
  res.end('/remind/save');
});
app.use('/remind/get', function(req, res) {
  log.debug('/remind/get:');
  res.end('/remind/get');
});


// 启动服务
var server = app.listen(config.listenport, function() {
  var host = server.address().address;
  var port = server.address().port;

  log.info('weixin app listening at http://%s:%s', host, port);
});

