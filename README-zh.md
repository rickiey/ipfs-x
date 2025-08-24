# ipfs-x

使用ipfs协议，将文件上传到ipfs网络，并返回ipfs的hash值，用于分享文件。
并且和其存储系统交互，包括但不限于

+ Filecoin
+ Swarm
+ Pinata
+ Arweave
+ Walrus

在运行或构建本项目之前，请确保满足以下条件：

安装 Go（版本 >= 1.18）
安装 IPFS 并确保IPFS节点正在运行
（可选）配置其他存储系统的访问凭证（如Pinata API密钥、Walrus账户等）

## 基础功能

+ 本地文件获取CID
+ 本地文件上传到IPFS网络
+ 下载IPFS网络上的文件

## 交互功能

+ 上传文件到其他存储系统，附带IPFS的CID

### 上传到 walrus

1. 获取文件CID
2. 上传文件到walrus，携带metadata {cid: CID}