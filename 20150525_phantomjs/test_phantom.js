var page = require('webpage').create();

page.open('http://localhost:8080/D%3A/temp/CodeEveryDay/test_phantomjs/data.html', function() {
	console.log('after open');
	page.includeJs("http://lib.sinaapp.com/js/jquery/1.9.1/jquery-1.9.1.min.js", function() {
		console.log('after includeJs');
		var href = page.evaluate(function() {
			return $('#my').attr('href');
		});
		setTimeout(function() {
			console.log('href=', href);
			phantom.exit();
		}, 1000);
	});
});