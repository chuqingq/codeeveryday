#!/usr/bin/env python

import signal
import sys
import asyncio
import logging


logging.basicConfig(
    # filename="transtream_wrapper.log",
    level=logging.DEBUG,
    format="%(asctime)s - %(levelname)s - %(message)s",
)


# 执行transtream，直到进程退出
async def transtream():
    logging.debug("start transtream...")
    proc = await asyncio.create_subprocess_shell(
        "./do_transtream",
        stdin=asyncio.subprocess.PIPE,
        stdout=asyncio.subprocess.PIPE,
        # stderr=asyncio.subprocess.PIPE,
    )
    await asyncio.gather(handle_stdout(proc), handle_stdin(proc), proc.wait())


# 处理标准输入
async def handle_stdin(proc):
    logging.debug("start handle_transtream_stdin..")
    # 创建一个 asyncio.StreamReader 对象
    stdin_reader = asyncio.StreamReader()
    # 创建一个 asyncio.StreamReaderProtocol 对象，并将其绑定到标准输入流上
    reader_protocol = asyncio.StreamReaderProtocol(stdin_reader)
    await asyncio.get_event_loop().connect_read_pipe(lambda: reader_protocol, sys.stdin)
    # 从标准输入流读取数据
    while True:
        try:
            # 读取一行输入
            line = await stdin_reader.read(1024)
            if not line:
                logging.info("input EOF")
                break  # 如果没有更多输入，则退出循环
            # 处理输入数据
            proc.stdin.write(line)
            await proc.stdin.drain()
        except KeyboardInterrupt:
            # 捕获键盘中断（Ctrl+C），退出循环
            break
    # 结束进程，即关闭进程的标准输入
    proc.stdin.close()
    await proc.stdin.wait_closed()
    logging.info("end handle_transtream_stdin")


# 处理标准输出
async def handle_stdout(proc):
    """处理标准输出"""
    logging.debug("handle_transtream_stdout")
    while True:
        line = await proc.stdout.read(1024)
        if not line:
            break
        logging.debug("transtream stdout: " + line.decode().strip())
        # 打印出来给父进程使用
        print(line.decode().strip())
    logging.info("end handle_transtream_stdout")


# 处理标准错误
async def handle_transtream_stderr(proc):
    """处理标准错误"""
    while True:
        line = await proc.stderr.read(1024)
        if not line:
            break
        logging.error("transtream stderr: " + line.decode().strip())


def signal_handler(sig, frame):
    print("Ctrl+C received. Exiting...")
    sys.exit(0)


# 注册Ctrl+C信号处理器
signal.signal(signal.SIGINT, signal_handler)

if __name__ == "__main__":
    asyncio.run(transtream())
