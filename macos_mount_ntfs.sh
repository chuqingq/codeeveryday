Mac 挂载NTFS移动硬盘进行读写操作 Readonly file system

1. diskutil info /Volumes/YOUR_NTFS_DISK_NAME 

找到 Device Node

Device Node:              /dev/disk1s1

2. hdiutil eject /Volumes/YOUR_NTFS_DISK_NAME

"disk1" unmounted.
"disk1" ejected.

弹出你的硬盘

3. 创建一个目录，稍后将mount到这个目录 

sudo mkdir /Volumes/MYHD

4. 将NTFS硬盘 挂载 mount 到mac

sudo mount_ntfs -o rw,nobrowse /dev/disk1s1 /Volumes/MYHD/

