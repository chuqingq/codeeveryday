var page = require('webpage').create();

page.open('http://www.baidu.com', function(status) {
    console.log('status='+status);

    console.log('document='+document);
    console.log('window='+window);

    // var recognition = new webkitSpeechRecognition();// 卡在这里
    // console.log(recognition);

    var SpeechRecognition = window.SpeechRecognition || 
			window.mozSpeechRecognition || 
			window.webkitSpeechRecognition || 
			window.msSpeechRecognition || 
			window.oSpeechRecognition;
    console.log(SpeechRecognition);// undefined
});
