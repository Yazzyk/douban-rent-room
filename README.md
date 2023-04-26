# 豆瓣租房信息爬取

> 目前个人已测试深圳ok，理论上其它城市也ok，只要修改配置文件中修改`WebSite`的地址到对应的城市的豆瓣租房小组地址即可
> 如：成都的租房小组地址是 [https://www.douban.com/group/CDzufang/discussion](https://www.douban.com/group/CDzufang/discussion) ,把配置文件中换成对应url即可

## 如何运行项目
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
```bash
docker build -t <IMAGE_NAME> .
```

