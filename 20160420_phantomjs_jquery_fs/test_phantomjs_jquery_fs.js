var webpage = require('webpage');
var fs = require('fs');

var start = 1;
var stop = 730;

function generateUrl(index) {
    var indexStr = '' + index;
    while (indexStr.length < 3) {
        indexStr = '0' + indexStr;
    }
    return indexStr;
}

function capture(index) {
    if (index > stop) {
        console.log("exit");
        phantom.exit();
        return;
    }

    console.log('capture:', index);
    var indexStr = generateUrl(index);
    var page = webpage.create();
    page.open('http://www.my285.com/wuxia/huangyi/xunqinji/'+indexStr+'.htm', function() {
        console.log('enter open');

        page.includeJs("http://lib.sinaapp.com/js/jquery/1.9.1/jquery-1.9.1.min.js", function() {
            console.log('enter includeJs');

            var content = page.evaluate(function() {
                console.log('enter evaluate');

                return $('td[colspan="2"]').text();
            });
            // console.log('after evaluate', content);

            fs.write('xqj.txt', content+'\n\n', 'a');
            page.close();
            capture(index+1);
            return;
        });
        
        console.log('after includeJs');
    });  
}

capture(start);