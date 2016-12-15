# kman

本工具主要用来收集系统的内存、CPU、磁盘,端口信息。由于我们以往通过敲命令的方式获取上述信息时，展示的方式对用户并不友好，
而且展示的信息不全等，于是kman可以作为一种辅助手段。


### 如何使用

- 构建静态文件

```
$ make build

```

- 构建完成后，生成 `kman` 静态连接文件

```
$ ./kman -h


NAME:
   kman - useful for getting infomation in our system

USAGE:
   kman [global options] command [command options] [arguments...]

VERSION:
   0.5.0

COMMANDS:
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
     --mem		display the system memory info
   --cpu		display the system cpu info
   --disk		display the system disk info
   --pname 		find the process with name
   --port 		find the process with portl; eg: --port 8080 or --port all
   --tw			get the tcp timewait info
   --help, -h		show help
   --version, -v	print the version
```


### 获取系统的内存信息

```
$ ./kman -mem

+--------+-----------+--------+--------+
| TOTAL  | AVAILABLE |  FREE  |  USED  |
+--------+-----------+--------+--------+
| 2.1 GB | 1.9 GB    | 1.7 GB | 373 MB |
+--------+-----------+--------+--------+
```

### 获取系统的CPU信息

```
$ ./kman -cpu

+------+------+--------+--------+-----------------+------+--------+
| CPU  | USER | SYSTEM |  IDLE  | IDLE PERCENTAGE | NICE | IOWAIT |
+------+------+--------+--------+-----------------+------+--------+
| cpu0 | 6.86 |   11.6 | 817.17 | 97.790888%      |    0 |   0.47 |
| cpu1 |  6.5 |  11.24 |    822 | 97.887441%      |    0 |   0.19 |
+------+------+--------+--------+-----------------+------+--------+
```

### 获取系统的磁盘使用情况

```
$ ./kman -disk

+-------+-------+--------+-------------+
| TOTAL | FREE  |  USED  | USEDPERCENT |
+-------+-------+--------+-------------+
| 90 GB | 77 GB | 8.9 GB | 9.855309%   |
+-------+-------+--------+-------------+
```

### 根据进程名称获取所属进程的信息

ps : 进程名称为模糊搜索

```
$ ./kman -pname docker

+------+-----------------+--------+
| PID  |      NAME       | PARENT |
+------+-----------------+--------+
| 1110 | dockerd         |      1 |
| 1197 | docker-containe |   1110 |
+------+-----------------+--------+
```

### 根据端口号获取所属进程的信息

```
$ ./kman -port 22

+------+------+------+--------+
| PID  | NAME | PORT | STATUS |
+------+------+------+--------+
| 1269 | sshd |   22 | LISTEN |
+------+------+------+--------+
```

```
可以通过 $ ./kman -port all 获取所有端口信息
```


### 根据TIME_WAIT情况

```
$ ./kman -tw

+--------+------------+----------+------------+-------------+
| LISTEN | FIN WAIT 2 | LAST ACK | CLOSE WAIT | ESTABLISHED |
+--------+------------+----------+------------+-------------+
|      5 |          5 |        1 |          5 |          28 |
+--------+------------+----------+------------+-------------+
```