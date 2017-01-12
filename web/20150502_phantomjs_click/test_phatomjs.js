//console.log('Hello, world!');
//phantom.exit();

var page = require('webpage').create();

//page.open('http://www.baidu.com', function() {
//	page.render('baidu.png');
//	phantom.exit();
//});

//page.open('http://www.baidu.com', function() {
//	console.log('before evaluate');
//	page.includeJs("http://lib.sinaapp.com/js/jquery/1.9.1/jquery-1.9.1.min.js", function() {
//		page.evaluate(function() {
//			// $('a[href="/signin"]').get(0).click();
//			$('a[href="http://news.baidu.com"]').get(0).click()
//	    });
//	    // page.render('golangtc.png');
//	    page.render('baidu.png');
//	    phantom.exit();
//	});
//});

page.open('http://www.baidu.com', function() {
	console.log('1111');
	page.evaluate(function() {
		// $('a[href="http://news.baidu.com"]').get(0).click();
		$('#kw').val(123);
		$('#su').get(0).click();
	});
	setTimeout(function() {
		page.render('baidu.png');
		phantom.exit();
	}, 3000);
});