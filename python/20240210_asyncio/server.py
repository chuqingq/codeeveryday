import asyncio

async def handle_client(reader, writer):
    while True:
        data = await reader.read(1024)
        if not data or data == b'quit':
            break
        writer.write(data)
        await writer.drain()

    writer.close()

async def run_server():
    server = await asyncio.start_server(handle_client, 'localhost', 8888)
    async with server:
        await server.serve_forever()

try:
    asyncio.run(run_server())
except KeyboardInterrupt:
    pass

