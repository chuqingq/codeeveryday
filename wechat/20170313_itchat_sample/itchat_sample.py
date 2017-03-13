# pip install itchat

import itchat

@itchat.msg_register(itchat.content.TEXT)
def text_reply(msg):
    print('msg: ' + msg['Text'])
    return msg['Text']

itchat.auto_login()
itchat.send('Hello, filehelper2', toUserName='filehelper')
itchat.run()

