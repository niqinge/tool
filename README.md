## 通用工具项目, 不受限某个场景

#### 使用方式
- 1.download
```text
    $: cd ~/src/github.com
    $: mkdir niqinge
    $: cd niqinge
    $: git clone git@github.com:niqinge/utils.git
```
- 2.引包
```text
    $: go mod vendor
```

#### 功能介绍

- 日志包
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

- mysql链接
```text
 
```