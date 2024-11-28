# PandaX device sdk for go

本章节介绍了如何安装和配置设备 SDK，以及提供了相关例⼦来演示如何使⽤设备 SDK上报设备数据以及服务调⽤


⽀持MQTT 协议版本：3.1.1
golang version：1.20+

## SDK 功能列表


- 发送 raw/telemetry/attribute 类型的数据到 IoT Hub.
- 从 IoT Hub 接收 raw/attribute/command 类型的数据
- 支持 MQTT/MQTTs.
- 支持自动重连.

> 未来会新增断线后缓存数据的功能目前暂未实现

## SDK API 列表


|         API         | Function                                   |
| :------------------ | :----------------------------------------- |
| PublishRaw        | 推送原始数据到 IoT Hub |
| PublishTelemetry | 推送遥测数据到 IoT Hub|
| PublishAttribute  | 推送属性数据到 IoT Hub |
| SubscribeRaw   | 从 IoT Hub 订阅原始数据 |
| SubscribeAttribute   | 从 IoT Hub 订阅属性数据 |
| SubscribeCommand   | 从 IoT Hub 订阅命令数据 |
| Connect      | 连接 IoT Hub    |
| Close      | 断开 IoT Hub |
| NewClient      | 创建连接客户端，同时可以配置属性如 qos, ssl等|

### 安装使用 SDK

``` shell
go get -u github.com/PandaXGO/device-sdk-go
```

```go
import "github.com/PandaXGO/device-sdk-go"
```

### 快速使用:

```go
// 创建默认的 client
_brokerAddr = "tcp://127.0.0.1:1883"
cli := client.NewClient(_brokerAddr, _username, _pwd)()

// 连接到 IoT Hub
cli.Connect()

// 订阅原始信息
cli.SubscribeRaw(context.TODO(), rawTopicHandler)

// 推送遥测数据
cli.PublishTelemetry(ctx, v)

// 关闭 client
cli.Close()
```

```go
// 创建支持 ssl 的client
_brokerAddr := "ssl://127.0.0.1:1883"
cli := client.NewClient(_brokerAddr, _username, _pwd)(
        client.WithUseSSL(true))

```

### Client Configuration

|         Parameter   | Description        |           Default        |
| :------------------ | :------------------| :----------------------- |
|host |IoT Hub 地址| "" |
|username | 设备 Token | "" |
|password | 可为空 | "" |

> 上述参数为必须设置的参数, 如果你想要创建支持 ssl、qos 及其他属性的 client 可以在创建是附加属性就像这样 **_client.WithQoS(1)_**

## 例子
[helloworld](samples/helloworld.go)

