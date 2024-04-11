#!/usr/bin/env python3

import signal
import sys
import asyncio
import logging


logging.basicConfig(
    filename='transtream_wrapper.log',
    level=logging.DEBUG,
    format="%(asctime)s - %(levelname)s - %(message)s",
)


transtream_proc = None


async def transtream_inout():
    """处理输入输出"""
    logging.debug("start transtream_inout...")
    # asyncio.get_event_loop().add_reader(sys.stdin, input_handler)
    # 等待transtream进程启动
    asyncio.sleep(1)
    while True:
        input_text = await asyncio.get_event_loop().run_in_executor(None, input)
        if transtream_proc:
            transtream_proc.stdin.write(input_text.encode())
            await transtream_proc.stdin.drain()
            logging.debug(f"transtream_proc.stdin.write({input_text})")


# 启动transtream
async def transtream():
    """启动transtream"""
    logging.debug("start transtream...")
    proc = await asyncio.create_subprocess_shell(
        "./do_transtream",
        stdin=asyncio.subprocess.PIPE,
        stdout=asyncio.subprocess.PIPE,
        stderr=asyncio.subprocess.PIPE,
    )
    global transtream_proc
    transtream_proc = proc
    # 处理标准输出
    task_stdout = asyncio.create_task(handle_transtream_stdout(proc.stdout))
    # 处理标准错误（日志）
    task_stderr = asyncio.create_task(handle_transtream_stderr(proc.stderr))
    # 传入标准输入
    # proc.stdin.writelines([b'{"a":456,"b":7,"c":8}\n'])
    # await proc.stdin.drain()
    # 等待任务完成
    await asyncio.gather(proc.wait(), task_stdout, task_stderr)
    logging.debug("transtream finished.")
    asyncio.get_event_loop().stop()


async def handle_transtream_stdout(stream):
    """处理标准输出"""
    # logging.debug("handle_transtream_stdout")
    while True:
        line = await stream.readline()
        # logging.debug(f"transtream stdout: {line}")
        if not line:
            break
        # logging.debug("transtream stdout: " + line.decode().strip())
        # 打印出来给父进程使用
        print(line.decode())


async def handle_transtream_stderr(stream):
    """处理标准错误"""
    while True:
        line = await stream.readline()
        if not line:
            break
        logging.debug("transtream stderr: " + line.decode().strip())


async def main():
    """主函数"""
    # 启动transtream
    task_transtream = asyncio.create_task(transtream())
    # 处理输入输出
    task_inout = asyncio.create_task(transtream_inout())
    # 等待任务完成
    await asyncio.gather(task_inout, task_transtream)


def signal_handler(sig, frame):
    print("Ctrl+C received. Exiting...")
    sys.exit(0)


# 注册Ctrl+C信号处理器
signal.signal(signal.SIGINT, signal_handler)

if __name__ == "__main__":
    asyncio.run(main())
