$(function() {
	setTimeout(function() {
		console.log("123");
	    console.log($('#my').attr("href"));
	    $('#my').attr('href', 'http://www.huawei.com');
	}, 0);// 设置成0，则打印结果为huawei；设置为1，则打印结果为baidu。因此说明：phantomjs不会等脚本都执行完成才执行open的回调。
});