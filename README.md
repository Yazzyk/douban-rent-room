# 豆瓣深圳租房信息爬取

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
根据个人情况修改config.toml

#### 4. 运行代码
```bash
go run main.go 
```

## 本地构建
```bash
make
```

## docker构建
```bash
docker build -t <IMAGE_NAME> .
```

