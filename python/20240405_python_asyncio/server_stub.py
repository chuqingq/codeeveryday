#!/usr/bin/env python

import signal
import sys
import asyncio


async def loop_streaming():
    """循环推流"""
    while True:
        print("start streaming...")
        await streaming()


async def streaming():
    """一次推流"""
    proc = await asyncio.create_subprocess_shell("python stream.py")
    await proc.wait()


# 启动transtream
async def transtream():
    """启动transtream"""
    print("start transtream...")
    proc = await asyncio.create_subprocess_shell(
        "python transtream.py",
        stdin=asyncio.subprocess.PIPE,
        stdout=asyncio.subprocess.PIPE,
        stderr=asyncio.subprocess.PIPE,
    )
    # 处理标准输出
    task_stdout = asyncio.create_task(handle_transtream_stdout(proc.stdout))
    # 处理标准错误（日志）
    task_stderr = asyncio.create_task(handle_transtream_stderr(proc.stderr))
    # 传入标准输入
    proc.stdin.writelines([b'{"a":456,"b":7,"c":8}\n'])
    # await proc.stdin.drain()
    # 等待任务完成
    await asyncio.gather(proc.wait(), task_stdout, task_stderr)


async def handle_transtream_stdout(stream):
    """处理标准输出"""
    print("handle_transtream_stdout")
    while True:
        line = await stream.readline()
        print(f'transtream stdout: {line}')
        if not line:
            break
        print('transtream stdout: '+line.decode().strip())


async def handle_transtream_stderr(stream):
    """处理标准错误"""
    while True:
        line = await stream.readline()
        if not line:
            break
        print('transtream stderr: '+line.decode().strip())


async def main():
    """主函数"""
    # 循环推流
    task_streaming = asyncio.create_task(loop_streaming())
    # 启动transtream
    task_transtream = asyncio.create_task(transtream())
    # 等待任务完成
    await asyncio.gather(task_streaming, task_transtream)


def signal_handler(sig, frame):
    print("Ctrl+C received. Exiting...")
    sys.exit(0)


# 注册Ctrl+C信号处理器
signal.signal(signal.SIGINT, signal_handler)

if __name__ == "__main__":
    asyncio.run(main())
