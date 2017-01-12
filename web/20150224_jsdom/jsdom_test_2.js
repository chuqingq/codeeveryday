var jsdom = ;
var jquery = require('fs').readFileSync("./jquery.js", "utf-8");

require("jsdom").env({
    url: 'http://localhost:6060/pkg/',
    src: [jquery],
    // scripts: ['./jquery.js'],
    done: function (errors, window) {
        console.log("there have been", window.$("a").length, "nodejs releases!");
        // console.log("there have been", jquery(window, "chuqq").length, "nodejs releases!");
    }
});