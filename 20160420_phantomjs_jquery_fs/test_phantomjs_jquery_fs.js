var page = require('webpage').create();
var fs = require('fs');

page.open('http://www.my285.com/wuxia/huangyi/xunqinji/423.htm', function() {
	console.log('enter open');

    page.includeJs("http://lib.sinaapp.com/js/jquery/1.9.1/jquery-1.9.1.min.js", function() {
        console.log('enter includeJs');

        var content = page.evaluate(function() {
            console.log('enter evaluate');

            return $('td[colspan="2"]').text();
        });

        fs.write('xqj1.txt', content, 'w');

        phantom.exit();
    });
	
    console.log('after includeJs');
});