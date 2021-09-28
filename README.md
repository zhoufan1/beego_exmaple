<<<<<<< HEAD
# beego_exmaple
beego exmaple
=======
# beego_example API 项目

> beego 用例
> 用例功能：  
> - db使用
> - logging 配置  
> - 全局配置  

- 开发运行  
```bee run -downdoc=true -gendoc=true```

- 生成文档  
```bee generate docs```

- 交叉打包  
  - Mac 编译 Linux 执行文件
```GOOS=linux GOARCH=amd64 go build```

  - Windows 编译 Linux 执行文件  
    ```
    SET CGO_ENABLED=0
    SET GOOS=linux
    SET GOARCH=amd64
    go build 
    ```
>>>>>>> dev
