升级beego框架
go get -u github.com/astaxie/beego

安装bee
go get -u github.com/beego/bee
安装完之后，bee 可执行文件默认存放在 $GOPATH/bin 里面，
所以您需要把 $GOPATH/bin 添加到您的环境变量中

Usage:

    bee command [arguments]

The commands are:

    version     show the bee & beego version
    migrate     run database migrations
    api         create an api application base on beego framework
    bale        packs non-Go files to Go source files    
    new         create an application base on beego framework
    run         run the app which can hot compile
    pack        compress an beego project
    fix         Fixes your application by making it compatible with newer versions of Beego
    dlv         Start a debugging session using Delve
    dockerize   Generates a Dockerfile for your Beego application
    generate    Source code generator
    hprose      Creates an RPC application based on Hprose and Beego frameworks
    pack        Compresses a Beego application into a single file
    rs          Run customized scripts
    run         Run the application by starting a local development server
    server      serving static content over HTTP on port

bee new myproject
    创建了 src/myproject/
        conf/           app.conf
        routers/        router.go
        controllers/    default.go
        models/
        views/          index.tpl
        static/
            js/
            css/
            img/
        tests/          default_test.go

bee api apipriject
    与 new相比 去掉了 static/ 和 views/
    创建了 src/apiproject

bee run 运行程序

bee pack
    src/apimyproject/apiproject.tar.gz
发布应用的时候打包，会把项目打包成 zip 包，这样我们部署的时候直接把打包之后的项目上传，解压就可以部署了：

bee version

bee migrate 
bee migrate rollback
bee migrate reset
bee migrate refresh
    默认参数
        -driver=mysql -conn="root:@tcp(127.0.0.1:3306)/test"

bee dockerize -image="library/golang:1.6.4" -expose=9000
    