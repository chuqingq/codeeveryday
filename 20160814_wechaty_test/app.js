const Wechaty = require('wechaty')
const bot = new Wechaty()

bot
.on('scan', ({url, code}) => console.log(`Scan QrCode to login: ${code}\n${url}`))
.on('login',         user => console.log(`User ${user} logined`))
.on('message',    message => {
    (!bot.self(message) && setTimeout(function() {
        console.log(`Message: ${message}`)
        const from = message.get('from')
        console.log(`from: ${from}`)
        message.set('to', from)
        bot.send(message)
    }, 5000))
    // (!bot.self(message) && bot.reply(message, message.getContentString()))
})
.init()
