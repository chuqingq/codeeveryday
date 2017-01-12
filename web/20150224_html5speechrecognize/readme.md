## 使用html5的speech接口

test_html5_speech.html 报错如下：

    SpeechRecognitionError {message: "", error: "network", clipboardData: undefined, path: NodeList[0], cancelBubble: false

可能是google网络不通导致的。修改hosts之后能够成功访问google。之后使用范例成功。
