import asyncio

async def tcp_echo_client(message):
    reader, writer = await asyncio.open_connection('127.0.0.1', 8888)
    print(f'Send: {message!r}')
    writer.write(message.encode())
    data = await reader.read(100)
    print(f'Received: {data.decode()!r}')
    print('Close the socket')
    writer.close()

message = 'Hello World!'

try:
    asyncio.run(tcp_echo_client(message))
except KeyboardInterrupt:
    pass

