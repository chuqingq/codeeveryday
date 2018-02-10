var weapi = require('wechat-api');

var config = {
  token: 'chuqq',
  appid: 'wx0288bf03ed5da89b',
  appsecret: 'd4624c36b6795d1d99dcf0547af5443d'
};

var api = new weapi(config.appid, config.appsecret);

api.sendText('oLo4_wcfnRNtxh0psNCBQaWY4wHg', 'hello world', function(err, data, res) {
	console.log('err: ', err, ', data: ', data, ', res: ', res);
}

var data = {
	time: {
	  "value":'2017-03-15',
	  "color":"#173177"
	},
	content: {
	  "value":'提醒中午吃饭',
	  "color":"#173177"
	},
};
api.sendTemplate('oLo4_wcfnRNtxh0psNCBQaWY4wHg', 'pf58G7IhZXMrk63XS9OrdwFrhhoIa81zpg74X96AGRg', 'http://weixin.qq.com/download', data, function(err, data, res) {console.log(err, data, res)})
