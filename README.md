# boxAPI0417

## 属性
![](https://img.shields.io/badge/status-build%20-green.svg)
![](https://img.shields.io/badge/design-Concurrent%20%5B%E5%B9%B6%E5%8F%91%5D-red.svg)
![](https://img.shields.io/badge/language-Golang-blue.svg)

## 目录

- boxAPI0417
  - src
    - initilization1
      - regist.go（注册box <- 上位机）
    - getConfig3
      - getServerUrl.go(获取上位机地址和box网络配置)
      - configNetwork.go（从底层配置网络）
    - heartbeat2
      - heartbeat.go（box心跳）
      - getSyncTask.go（获取同步任务）
      - ackForSyncTask.go(上报同步任务结果)
    - rpcFacetrack(数据推送)
      - rpcClient
        - main.go
      - rpcServer
        - main.go
      - rpcFacetrack.go（数据推送处理函数）
    - StdMsgForm(go标准的上位机通用信息格式的数据结构)
    - StdJsonrpc(go标准Jsonrpc以及考斯的部分数据结构)
    - utils
      - getIP.go(获取配置文件中的IP地址)
      - reboot.go(重启box)
      - UUID.go(box的唯一标示符的设定)
  - sundry
    - client-interface(考斯比特大陆人脸对接数据结构 JSON)
    - server-interface(考斯比特大陆人脸对接数据结构 JSON)
    - ...
  - test(临时验证)


## 参考
[Golang标准库文档](https://studygolang.com/pkgdoc)  

[Golang基础-进阶-实战编程 bilibili](https://www.bilibili.com/video/av36489007/?p=25)

[jsonrpc http简书](https://www.jianshu.com/p/74ac2439afb2)