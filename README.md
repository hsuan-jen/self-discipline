## 📗 目录结构

```
├── api             (api层)
│   └── v1          (v1版本接口)
├── config          (配置包)
├── configs         (配置文件及常量)
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
├── source          (source层)                
└── utils           (工具包)
```

# 使用如下命令下载swag
go get -u github.com/swaggo/swag/cmd/swag
```

#### 生成API文档

````
swag init
````