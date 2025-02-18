# 豆瓣租房信息爬取

> 目前个人已测试深圳可用，理论上其它城市也可用，只要修改配置文件中修改`WebSite`的地址到对应的城市的豆瓣租房小组地址即可
> 如：成都的租房小组地址是 [https://www.douban.com/group/CDzufang/discussion](https://www.douban.com/group/CDzufang/discussion) ,把配置文件中换成对应url即可

## 更新
- v1.1.0: `2025-02-17` 接入coze deepseek,用于最后识别是否符合要求 

## CookieCloud支持
本项目使用了[CookieCloud](https://github.com/easychen/CookieCloud)，可点去自己搭建CookieCloud，也可使用一些[公共的CookieCloud服务](https://github.com/easychen/CookieCloud#%E7%AC%AC%E4%B8%89%E6%96%B9)

## 如何运行项目
### 从docker-compose启动(推荐)
1. 将docker-compose.yml拷贝到本地，并确保他们在相同目录下(请确保已安装`docker`和`docker-compose`)  
```yml
version: '3.7'
services:
  douban:
    image: yazzyk/douban:latest
    container_name: douban
    volumes:
      - ${DOUBAN_ROOT}/config.toml:/app/config.toml
      - ${DOUBAN_ROOT}/nutsDB:/app/nutsDB
      - ${DOUBAN_ROOT}/logs:/app/logs
    ports:
      - '5050:5050'
```
以上的${DOUBAN_ROOT}为配置文件和日志文件的存放目录，可以自行修改
2. 运行
```bash
docker-compose up -d
```

### 从docker启动
```bash
docker pull yazzyk/douban:latest
docker run -d --name douban -v ${DOUBAN_ROOT}/config.toml:/app/config.toml -v ${DOUBAN_ROOT}/logs:/app/logs yazzyk/douban:latest
```
以上的${DOUBAN_ROOT}为配置文件和日志文件的存放目录，可以自行修改

### 自行编译运行
#### 1. 下载并安装 Golang  
请参考: [https://go.dev/doc/install](https://go.dev/doc/install)

#### 2. 下载代码
```bash
git clone github.com/yazzyk/douban-rent-room
cd douban-rent-room
go mod tidy
```

#### 3. 修改配置文件
根据个人情况修改`config.toml`

#### 4. 运行代码
```bash
go run main.go 
```

### 从Realease下载运行
1. 前往[Release](https://github.com/Yazzyk/douban-rent-room/releases)下载对应平台的压缩包
2. 根据个人需要修改配置文件`config.toml`(配置文件和可执行文件要在同一个目录下)
3. 运行程序

## 本地构建
```bash
make
```

## docker构建
> 构建依赖于makefile的构建结果
```bash
docker build -t <IMAGE_NAME> .
```
or 
```bash
docker-compose -f docker-compose.build.yml build
```
