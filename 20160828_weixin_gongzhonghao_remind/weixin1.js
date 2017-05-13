var process = require('process');
var wechat = require('wechat');
var weapi = require('wechat-api');
var connect = require('express');
var app = connect();

process.on('uncaughtException', function(err) {
    console.log('uncaughtException:', err);
});

var config = {
  token: 'chuqq',
  appid: 'wx0288bf03ed5da89b',
  appsecret: 'd4624c36b6795d1d99dcf0547af5443d'
  //encodingAESKey: 'encodinAESKey'
};

var api = new weapi(config.appid, config.appsecret);

app.use(connect.query());
app.use('/wechat', wechat(config, wechat.text(function(message, req, res, next) {
  console.log('/wechat text:', message.Content);
  handleMsg(message.Content, message.FromUserName, res);
  return;
  /*
  // 回复屌丝(普通回复)
  res.reply('hehe');
  //你也可以这样回复text类型的信息
  res.reply({
    content: 'text object',
    type: 'text'
  });
  // 回复一段音乐
  res.reply({
    type: "music",
    content: {
      title: "来段音乐吧",
      description: "一无所有",
      musicUrl: "http://mp3.com/xx.mp3",
      hqMusicUrl: "http://mp3.com/xx.mp3"
      //thumbMediaId: "thisThumbMediaId"
    }
  });
  // 回复高富帅(图文回复)
  res.reply([
    {
      title: '你来我家接我吧',
      description: '这是女神与高富帅之间的对话',
      picurl: 'http://nodeapi.cloudfoundry.com/qrcode.jpg',
      url: 'http://nodeapi.cloudfoundry.com/'
    }
  ]);
  }*/
}).voice(function(message, req, res, next) {
  console.log("/wechat voice:", message.Recognition);
  handleMsg(message.Recognition, message.FromUserName, res);
})));


function handleMsg(content, user, res) {
  var opts = {
    "query":content,
    "city":"南京",
    "category": "remind"
  };

  api.semantic(user, opts, function(err, data, res2) {
    console.log('semantic result: ', data);
    if (err != null
      || data == undefined
      || data.semantic == undefined
      || data.semantic.details == undefined
      || data.semantic.details.datetime == undefined) {
        res.reply('error: ' + err + '.\n' + content);
        return;
    }
    var datetime = data.semantic.details.datetime.date + ' ' + data.semantic.details.datetime.time;
    // res.reply("定时成功: "+ datetime + "\n" + message.Recognition);

    // 定时
      var timeA = parseDate(datetime);
      if (!timeA) {
        res.reply('时间格式非法');
        return;
      }
      var timeout = timeA.getTime() - new Date().getTime();
      if (timeout <= 0) {
        res.reply('时间已过期');
        return;
      }
      setTimeout(function() {
        api.sendText(user, content, function(err, data, res) {
          console.log('err:', err);
          console.log('data:', data);
          console.log('res:', res);
        });
      }, timeout);
      res.reply('定时提醒成功：\n'+datetime+'\n'+content);
  });
}

/*  
  将String类型解析为Date类型.  
  parseDate('2006-1-1') return new Date(2006,0,1)  
  parseDate(' 2006-1-1 ') return new Date(2006,0,1)  
  parseDate('2006-1-1 15:14:16') return new Date(2006,0,1,15,14,16)  
  parseDate(' 2006-1-1 15:14:16 ') return new Date(2006,0,1,15,14,16);  
  parseDate('2006-1-1 15:14:16.254') return new Date(2006,0,1,15,14,16,254)  
  parseDate(' 2006-1-1 15:14:16.254 ') return new Date(2006,0,1,15,14,16,254)  
  parseDate('不正确的格式') retrun null  
*/
function parseDate(str){
  if(typeof str == 'string'){
    // var results = str.match(/^ *(\d{4})-(\d{1,2})-(\d{1,2}) */);   
    // if(results && results.length>3)   
    //   return new Date(parseInt(results[1]),parseInt(results[2]) -1,parseInt(results[3]));    
    var results = str.match(/^ *(\d{4})-(\d{1,2})-(\d{1,2}) +(\d{1,2}):(\d{1,2}):(\d{1,2}) */);
    if(results && results.length>6)
      return new Date(parseInt(results[1]),parseInt(results[2]) -1,parseInt(results[3]),parseInt(results[4]),parseInt(results[5]),parseInt(results[6]));
    // results = str.match(/^ *(\d{4})-(\d{1,2})-(\d{1,2}) +(\d{1,2}):(\d{1,2}):(\d{1,2})\.(\d{1,9}) */);   
    // if(results && results.length>7)   
    //   return new Date(parseInt(results[1]),parseInt(results[2]) -1,parseInt(results[3]),parseInt(results[4]),parseInt(results[5]),parseInt(results[6]),parseInt(results[7]));    
  }
  return null;
}


// 启动服务

var server = app.listen(80, function() {
  var host = server.address().address;
  var port = server.address().port;

  console.log('weixin app listening at http://%s:%s', host, port);
});

