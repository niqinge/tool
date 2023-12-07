## 通用工具项目, 不受限某个场景

#### 使用方式
- 1.download
```text
    $: cd ~/go/src/github.com
    $: mkdir niqinge
    $: cd niqinge
    $: git clone git@github.com:niqinge/utils.git
```
- 2.引包
```text
    $: go mod tidy
```

### 功能介绍

#### 日志包log
```text
    说明:日志包使用的是uber开源日志框架, 详细请查看https://github.com/uber-go/zap.
    使用:
        1.在服务启动时调用logger包InitLogger(project string)方法进行初始化, 其中project可以是路径或者是你的服务名都可以.
        2.初始化之后, 既可以在服务的任何地方使用类似下面代码进行日志打印:
            Info: log.Info("Info", zap.String("test", "Info"))
            Warn: log.Warn("Warn", zap.String("test", "Warn"))
            Debug:log.Debug("Debug", zap.String("test", "Debug"))
            Error:log.Error("Error", zap.String("test", "Error"))
```

#### mysql
- mysql链接 
- 数据库迁移


#### apollo分布式配置中心客户端
- 说明: 需要安装携程apollo开源框架[服务端](https://github.com/ctripcorp/apollo) 
- 客户端链接及取值

#### 发邮件
- 需要配置发送人的账户信息

#### nsq消息队列(TODO)

#### reflects 反射
- 反射出结构体的方法，并执行指定方法后，返回其结果值

#### JWT
- JWT 生成, 解析

#### mycontext
- 唯一的requestId
- 基于已有的requestId 生成context
- 基于已有的context 获取requestId
