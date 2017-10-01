
$ lscpu
Architecture:          x86_64
CPU 运行模式：    32-bit, 64-bit
Byte Order:            Little Endian
CPU(s):                4
On-line CPU(s) list:   0-3
每个核的线程数：2
每个座的核数：  2
Socket(s):             1
NUMA 节点：         1
厂商 ID：           GenuineIntel
CPU 系列：          6
型号：              37
Model name:            Intel(R) Core(TM) i5 CPU       M 520  @ 2.40GHz
步进：              2
CPU MHz：             2400.000
CPU max MHz:           2400.0000
CPU min MHz:           1199.0000
BogoMIPS:              4788.23
虚拟化：           VT-x
L1d 缓存：          32K
L1i 缓存：          32K
L2 缓存：           256K
L3 缓存：           3072K
NUMA node0 CPU(s):     0-3
Flags:                 fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush dts acpi mmx fxsr sse sse2 ss ht tm pbe syscall nx rdtscp lm constant_tsc arch_perfmon pebs bts rep_good nopl xtopology nonstop_tsc aperfmperf pni pclmulqdq dtes64 monitor ds_cpl vmx smx est tm2 ssse3 cx16 xtpr pdcm sse4_1 sse4_2 popcnt aes lahf_lm tpr_shadow vnmi flexpriority ept vpid dtherm ida arat



$ go test -bench AtomicAdd -benchtime 10s
goos: linux
goarch: amd64
BenchmarkAtomicAdd-4   	2000000000	         6.56 ns/op
PASS
ok  	_/home/chuqq/work/codeeveryday/golang/20171001_go_atomic_bench	13.803s

