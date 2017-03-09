var fs = require('fs');

var b = fs.readFileSync('./mail.qq.com1.har.json');
var input = JSON.parse(b.toString()).log.entries;
console.log('input count: ', input.length);

// 结构化的请求
var output = [];

for (var i in input) {
    var req = input[i];
    // console.log('req: ', req);

    var isRefer = false;
    var referer = getRefererFromHeaders(req.request.headers);
    for (var j in output) {
        var old = output[j];
        // console.log('old: ', old);
        if (referer == old.request.url) {
            if (!old.children) {
                old.children = [];
            }
            old.children.push(req);
            isRefer = true;
            break;
        }
    }

    if (!isRefer) {
        output.push(req);
    }
}

console.log('output: ', output.length);

function getRefererFromHeaders(headers) {
    for (var i in headers) {
        var h = headers[i];
        if (h.name == 'referer' || h.name == 'Referer') {
            // console.log('referer: ', h.value);
            return h.value;
        }
    }
}

// 打印这31个referer
for (var i in output) {
    var req = output[i];
    console.log('referer: ', getRefererFromHeaders(req.headers));
}

// 保存文件
fs.writeFileSync('1.json', JSON.stringify(output));

// 全部打印