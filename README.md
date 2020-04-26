# beego_blog
```个人beego项目
go服务器安装 golang，配置 这里就略过了

get get -u github.com/astaxie/beego
bee仓库是我自己修改过的，在使用过程中bee有很多缺陷
之前经常使用rails框架，按照rails一些规范，对bee进行了一些完善 仅仅针对mysql(后面有时间再继续深入，定制空间还是很大的)
go get -u github.com/EricWJP/bee

该项目使用mysql需要 mysql包
go get -u github.com/go-sql-driver/mysql

1、部署 
  在本地创建了多个项目 其中包括beego_blog, 最终选择b_blog作为项目名
  在github上创建项目时，由于仓库太多了，所以为了容易区分 取名为 beego_blog
  git clone https://github.com/EricWJP/beego_blog.git
  mv beego_blog b_blog
2、修改 conf/app.conf 文件
   修改项目启动配置
      例如 启动production配置 runmode=prod
   配置数据库
3、配置nginx 这里nginx 与 beego项目在同一个服务器
   这里没有什么特别之处
   server块
    proxy_pass http://localhost:8080;
    这里http://localhost:8080 是 beego启动后的URL，端口号是8080（app.conf: httpport = 8080）
4、启动beego项目
    这里可以用shell脚本实现部署流程(代码库更新 版本管理 共享文件 等等，推荐参考Capistrano部署流程), shell脚本并没有放到这个仓库
    也可以 直接后台
         bee run & 
    也可以
        使用编译成功的b_blog可执行文件后台启动
    也可以
        把b_blog stastic/ views/ conf/ 这几个打包成一个目录，放到任意目录
        进入该目录 使用 ./b_blog 命令来启动服务
    当然还有其他 部署工具  大致步骤： 也可以使用shell脚本来实现
        本地
            bee pack 打包成压缩文件
            scp 上传压缩文件到服务器
        服务器
            解压上传的压缩文件
            Supervisor管理beego进程，重启应用程序
        
   