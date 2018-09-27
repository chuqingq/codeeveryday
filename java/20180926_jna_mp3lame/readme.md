# 使用liblame把pcm流式压缩成mp3

## 安装libmp3lame-dev

    sudo apt install libmp3lame-dev

## 编译demo

代码中引用头文件：

    #include "lame/lame.h"

编译选项：

    gcc -o pcm2mp3 pcm2mp3.c -lmp3lame -Wall

执行可执行文件：

    ./pcm2mp3 7.5depingfangshiduoshao.pcm 1.mp3

## 原pcm播放

    vlc.exe  --demux=rawaud --rawaud-channels 1 --rawaud-samplerate 16000 7.5depingfangshiduoshao.pcm

## 验证流式OK

    dd if=1.mp3 bs=1k count=1 of=1.1.mp3

## JNA封装

运行：

    java -cp myartifact-0.0.1-SNAPSHOT.jar:jna-3.0.9.jar   com.chuqq.mygroup.myartifact.App
