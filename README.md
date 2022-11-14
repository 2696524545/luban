### 多集群管理平台


#### 环境依赖
```
nodejs 16.x
golang 1.18.x
kubernetes 1.24.x
vue 2.x
element-ui 2.x
mysql5.7
```

#### client-go
```
go get k8s.io/client-go@v0.24.6
```


#### 目录结构说明
后端
```
micheng@michengdeMacBook-Pro demo % tree  .
.
├── README.md
├── api             接口
│   ├── routers     路由
│   │   ├── k8s.go
│   │   └── user.go
│   └── types       请求/返回 结构体
│       └── types.go
├── cmd
│   ├── config      配置文件结构体
│   │   └── config.go
│   ├── main.go     主程序
│   └── options     options方法
│       └── options.go
├── controller       业务逻辑
│   ├── client.go
│   ├── k8s.go
│   └── user.go
├── etc
│   └── config.yaml  配置文件
├── go.mod
├── go.sum
└── pkg
    ├── core
    │   └── k8s
    ├── model        数据库模型
    │   ├── k8s_cluster.go
    │   └── user.go
    ├── types
    │   └── types.go
    └── utils         工具
        ├── mysql.go
        └── viper.go

15 directories, 19 files

```
前端
```
micheng@michengdeMacBook-Pro luban-demo % tree -I node_modules 
.
├── README.md
├── babel.config.js
├── jsconfig.json
├── package-lock.json
├── package.json
├── public
│   ├── favicon.ico
│   └── index.html
├── src
│   ├── App.vue
│   ├── api           // 接口
│   │   └── test.js
│   ├── assets        // 静态资源
│   │   └── logo.png
│   ├── components   // 页面组件
│   │   ├── HelloWorld.vue
│   │   └── One.vue
│   ├── main.js
│   ├── plugin
│   │   ├── store   // 状态
│   │   │   └── index.js
│   │   └── utils   // 工具
│   │       └── request.js
│   ├── router      // 路由
│   └── view        // 页面
└── vue.config.js

10 directories, 16 files

```
