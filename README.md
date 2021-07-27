## 📗 目录结构

```
├── api             (api层)
│   └── v1          (v1版本接口)
├── conf            (配置文)
├── configs         (配置包)
├── core            (核心文件)
├── docs            (swagger文档目录)
├── global          (全局对象)
├── initialize      (初始化)
│   └── internal    (初始化内部函数)
├── middleware      (中间件层)
├── model           (模型层) 
│   ├── request     (入参结构体)                
│   └── response    (出参结构体)
├── router          (路由层)               
├── service         (service层)              
└── utils           (工具包)
```

## 使用如下命令下载swag
```
go get -u github.com/swaggo/swag/cmd/swag
```

### 生成API文档
```
swag init
```
## 开启pprof,在配置文件中设置
```
system:
  pprof: true
```

### 通过浏览器访问
```
http://127.0.0.1:8800/debug/pprof/
```

### 火焰图
```
# 执行命令后,会在浏览器打开一个窗口
go tool pprof -http=:1234 http://localhost:8800/debug/pprof/goroutine
# 简单解释
# -http 表示使用交互式web接口查看获取的性能信息,指定可用的端口即可
# debug/pprof/需要查看的指标 (allocs,block,goroutine,heap...)
# go tool pprof http://localhost:8800/debug/pprof/goroutine?second=20
# 采集协程数据并持续20S
# 常用的命令有top,tree,web,list等
```
## 开启prometheus,在配置文件中设置
```
system:
  promhttp: true
```