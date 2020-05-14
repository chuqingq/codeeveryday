import smtplib
from email.mime.text import MIMEText
from email.header import Header

server = 'smtp.qiye.aliyun.com'
port = 465
sender = 'xxx'
sender_pwd = ''
receivers = [
    '',
]

subject = 'subject 123'
content = '123 456'

message = MIMEText(content, 'plain', 'utf-8')
message['From'] = sender
message['To'] = ','.join(receivers)
# message['X-Priority'] = '1'
message['Subject'] = Header(subject, 'utf-8')

try:
    with smtplib.SMTP_SSL(server, port) as client:
        # client.set_debuglevel(1)
        client.login(sender, sender_pwd)
        # print('Sending email to %s...' % str(receivers))
        print('msg: %s' % message.as_string())
        client.sendmail(sender, receivers, message.as_string())
        # time.sleep(5)
except Exception as e:
    print('sendmail error:', repr(e))
