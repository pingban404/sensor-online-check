# 传感器在线检查系统

## 项目简介

`sensor-online-check` 是一个用于监控传感器在线状态的系统。它定期检查传感器的状态，并记录每次检查的结果。

## 目录结构
```
├─check 检查传感器状态的脚本
├─config 配置文件目录
├─esclient ES客户端目录
├─query 查询传感器状态的脚本
└─utils 工具函数目录
```

## 环境依赖

- Go 1.18+
- Git

## 安装步骤

1. 克隆项目仓库：
```sh
 git clone https://github.com/your-repo/sensor-online-check.git
 cd sensor-online-check
```
2. 安装依赖
  ```sh
  go mod download
  ```
3. 配置环境变量（如果必要）
  ```sh
  export YOUR_ENV_VAR=value
  ```
### 运行项目
1. 编译项目
```sh
 go build -o sensor-check
 ```
2. 运行可执行文件
```sh
./sensor-check
```
3. 查看日志
```sh
tail -f app.log
```
   

