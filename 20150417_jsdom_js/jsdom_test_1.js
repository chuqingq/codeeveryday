var jsdom = require("jsdom");

//jsdom.env(
//	"http://open.baidu.com/special/time/",
//	[],// , "http://www.baidu.com/js/aladdin/clock/clock.js" "http://code.jquery.com/jquery.js"
//	function (errors, window) {
//		console.log(window.document.getElementById("time").innerHtml);
//	}
//);
//jsdom.env({
//  url: "http://open.baidu.com/special/time/",
//  scripts: [],
//  done: function (errors, window) {
//	  console.log(window.baidu_time);
//    // console.log(window.document.getElementById("time"));
//  }
//});
jsdom.env({
  html: "<!DOCTYPE html><html><head><script>window.myvar = 1;</script></head><body><div id=\"time\">123</div></body></html>",//
  scripts: ["D:\\jquery-1.9.1.min.js"],//[path.resolve(__dirname, "../jquery-fixtures/jquery-1.6.2.js")],
  features: {
      FetchExternalResources: ["script"],
      ProcessExternalResources: ["script"],
      SkipExternalResources: false
    },
  done: function (errors, window) {
	  var $ = window.jQuery;
	  console.log(window.myvar);
	  console.log($('#time').text());
    // console.log(window.document.getElementById("time"));
  }
});