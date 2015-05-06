var jsdom = require("jsdom");
var exec = require('child_process').exec;

jsdom.env({
	url: "http://open.baidu.com/special/time/",
	scripts: ["D:\\jquery-1.9.1.min.js"],
	features: {
		ProcessExternalResources: ["script"]
	},
	done: function (errors, window) {
		var $ = window.$;
		var curtime = $('#time').text();
		console.log('get time: ' + curtime);

		// 更新系统时间
		exec('time ' + curtime, function(err, stdout, stderr) {
			console.log('change time ' + curtime + ': ' + (err ? err : 'success'));
			process.exit();
		});
	}
});
