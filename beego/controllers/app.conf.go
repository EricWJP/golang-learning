appname = beepkg
httpaddr = "127.0.0.1"
httpport = 9090
runmode ="dev"
autorender = false
recoverpanic = false
viewspath = "myview"
// 也可以直接通过beego.BConfig.AppName="beepkg"这样来修改


mysqluser = "root"
mysqlpass = "rootpass"
mysqlurls = "127.0.0.1"
mysqldb   = "beego"

// 获取方式
beego.AppConfig.String("mysqluser")
beego.AppConfig.String("mysqlpass")
beego.AppConfig.String("mysqlurls")
beego.AppConfig.String("mysqldb")

AppConfig 的方法如下：
	Set(key, val string) error
	String(key string) string
	Strings(key string) []string
	Int(key string) (int, error)
	Int64(key string) (int64, error)
	Bool(key string) (bool, error)
	Float(key string) (float64, error)
	DefaultString(key string, defaultVal string) string
	DefaultStrings(key string, defaultVal []string)
	DefaultInt(key string, defaultVal int) int
	DefaultInt64(key string, defaultVal int64) int64
	DefaultBool(key string, defaultVal bool) bool
	DefaultFloat(key string, defaultVal float64) float64
	DIY(key string) (interface{}, error)
	GetSection(section string) (map[string]string, error)
	SaveConfigFile(filename string) error


// 不同环境配置
[dev]
httpport = 8080
[prod]
httpport = 8088
[test]
httpport = 8888


多个配置文件
app.conf
[dev]
httpport = 8080
[prod]
httpport = 8088
[test]
httpport = 8888

app2.conf
runmode ="dev"
autorender = false
recoverpanic = false
viewspath = "myview"

[dev]
httpport = 8080
[prod]
httpport = 8088
[test]
httpport = 8888




从环境变量配置获取
runmode  = "${ProRunMode||dev}"
httpport = "${ProPort||9090}"


BConfig
保存了所有 beego 里面的系统默认参数，你可以通过 beego.BConfig 来访问和修改底下的所有配置信息.

beego.LoadAppConfig("ini", "conf/app2.conf")
加载多个配置文件，调用多次就可以了
如果后面文件跟前面的key冲突，后面的覆盖前面的


App配置
AppName
	应用名称，默认是 beego。bee new 创建的项目名称
	beego.BConfig.AppName = "beego"
RunMode
	运行模式，可选值： prod、dev或者test。默认dev（开发模式）
	beego.BConfig.RunMode = "dev"
RouterCaseSensitive
	是否路由忽略大小写匹配，默认是 true，区分大小写
	beego.BConfig.RouterCaseSensitive = true
ServerName
	beego服务器默认 在请求的时候输出server为beego
	beego.BConfig.ServerName = "beego"
RecoverPanic	
	是否异常恢复，默认值为true，即当应用出现异常的情况，通过recover恢复回来，而不会导致应用异常退出
	beego.BConfig.RecoverPanic = true
CopyRequestBody
	是否允许在HTTP请求时，返回原始请求体数据字节，默认为 false （GET or HEAD or 上传文件除外）
	beego.BConfig.CopyRequestBody = false
EnableGzip
	是否开始gzip支持，默认false 不支持 gzip， 
	一旦开启gzip，那么在模版输出的内容会进行gzip或者在zlib压缩，根据用户的Accept-Encoding来判断
	beego.BConfig.EnableGzip = false
	Gzip允许用户自定义压缩级别、压缩长度阈值和针对请求类型压缩:
		压缩级别, gzipCompressLevel = 9,取值为 1~9,如果不设置为 1(最快压缩)
		压缩长度阈值, gzipMinLength = 256,当原始内容长度大于此阈值时才开启压缩,默认为 20B(ngnix默认长度)
		请求类型, includedMethods = get;post,针对哪些请求类型进行压缩,默认只针对 GET 请求压缩
MaxMemory
	文件上传默认内存缓存大小，默认值是 1 << 26 (64M)
	beego.BConfig.MaxMemory = 1 << 26
EnableErrorsShow
	是否显示系统错误信息，默认 true
	beego.BConfig.EnableErrorsShow = true
EnableErrorsRender
	是否将错误进行渲染，默认 true，即出错会提示友好的出错页面，
	对于API类型的应用可能需要讲该选项设为false以阻止在dev模式下不必要的模版渲染信息返回



Web配置
AutoRender
	是否模板自动渲染，默认值为 true，对于 API 类型的应用，应用需要把该选项设置为 false，不需要渲染模板。
	beego.BConfig.WebConfig.AutoRender = true
EnableDocs
	是否开启文档内置功能，默认是 false
	beego.BConfig.WebConfig.EnableDocs = true
FlashName
	Flash 数据设置时 Cookie 的名称，默认是 BEEGO_FLASH
	beego.BConfig.WebConfig.FlashName = "BEEGO_FLASH"
FlashSeperator
	Flash 数据的分隔符，默认是 BEEGOFLASH
	beego.BConfig.WebConfig.FlashSeparator = "BEEGOFLASH"
DirectoryIndex
	是否开启静态目录的列表显示，默认不显示目录，返回 403 错误。
	beego.BConfig.WebConfig.DirectoryIndex = false
StaticDir
	静态文件目录设置，默认是static
	可配置单个或多个目录:
		单个目录, StaticDir = download. 相当于 beego.SetStaticPath("/download","download")
		多个目录, StaticDir = download:down download2:down2. 
			相当于 beego.SetStaticPath("/download","down") 和 beego.SetStaticPath("/download2","down2")
	beego.BConfig.WebConfig.StaticDir
StaticExtensionsToGzip
	允许哪些后缀名的静态文件进行 gzip 压缩，默认支持 .css 和 .js
	beego.BConfig.WebConfig.StaticExtensionsToGzip = []string{".css", ".js"}

	等价 config 文件中
		StaticExtensionsToGzip = .css, .js
TemplateLeft
	模板左标签，默认值是{{。
	beego.BConfig.WebConfig.TemplateLeft="{{"
TemplateRight
	模板右标签，默认值是}}。
	beego.BConfig.WebConfig.TemplateRight="}}"
ViewsPath
	模板路径，默认值是 views。
	beego.BConfig.WebConfig.ViewsPath="views"
EnableXSRF
	是否开启 XSRF，默认为 false，不开启。
	beego.BConfig.WebConfig.EnableXSRF = false
XSRFKEY
	XSRF 的 key 信息，默认值是 beegoxsrf。 EnableXSRF＝true 才有效
	beego.BConfig.WebConfig.XSRFKEY = "beegoxsrf"
XSRFExpire
	XSRF 过期时间，默认值是 0，不过期。
	beego.BConfig.WebConfig.XSRFExpire = 0


监听配置
Graceful
	是否开启热升级，默认是 false，关闭热升级。
	beego.BConfig.Listen.Graceful=false
ServerTimeOut
	设置 HTTP 的超时时间，默认是 0，不超时。
	beego.BConfig.Listen.ServerTimeOut=0
ListenTCP4
	监听本地网络地址类型，默认是TCP6，可以通过设置为true设置为TCP4。
	beego.BConfig.Listen.ListenTCP4 = true
EnableHTTP
	是否启用 HTTP 监听，默认是 true。
	beego.BConfig.Listen.EnableHTTP = true
HTTPAddr
	应用监听地址，默认为空，监听所有的网卡 IP。
	beego.BConfig.Listen.HTTPAddr = ""
HTTPPort
	应用监听端口，默认为 8080。
	beego.BConfig.Listen.HTTPPort = 8080
EnableHTTPS
	是否启用 HTTPS，默认是 false 关闭。当需要启用时，先设置 EnableHTTPS = true，并设置 HTTPSCertFile 和 HTTPSKeyFile
	beego.BConfig.Listen.EnableHTTPS = false
HTTPSAddr
	应用监听地址，默认为空，监听所有的网卡 IP。
	beego.BConfig.Listen.HTTPSAddr = ""
HTTPSPort
	应用监听端口，默认为 10443
	beego.BConfig.Listen.HTTPSPort = 10443
HTTPSCertFile
	开启 HTTPS 后，ssl 证书路径，默认为空。
	beego.BConfig.Listen.HTTPSCertFile = "conf/ssl.crt"
HTTPSKeyFile
	开启 HTTPS 之后，SSL 证书 keyfile 的路径。
	beego.BConfig.Listen.HTTPSKeyFile = "conf/ssl.key"
EnableAdmin
	是否开启进程内监控模块，默认 false 关闭。
	beego.BConfig.Listen.EnableAdmin = false
AdminAddr
	监控程序监听的地址，默认值是 localhost 。
	beego.BConfig.Listen.AdminAddr = "localhost"
AdminPort
	监控程序监听的地址，默认值是 8088 。
	beego.BConfig.Listen.AdminPort = 8088
EnableFcgi
	是否启用 fastcgi ， 默认是 false。
	beego.BConfig.Listen.EnableFcgi = false
EnableStdIo
	通过fastcgi 标准I/O，启用 fastcgi 后才生效，默认 false。
	beego.BConfig.Listen.EnableStdIo = false


Session配置
SessionOn
	session 是否开启，默认是 false。
	beego.BConfig.WebConfig.Session.SessionOn = false
SessionProvider
	session 的引擎，默认是 memory，详细参见 session 模块。
	beego.BConfig.WebConfig.Session.SessionProvider = ""
SessionName
	存在客户端的 cookie 名称，默认值是 beegosessionID。
	beego.BConfig.WebConfig.Session.SessionName = "beegosessionID"
SessionGCMaxLifetime
	session 过期时间，默认值是 3600 秒。
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600
SessionProviderConfig
	配置信息，根据不同的引擎设置不同的配置信息，详细的配置请看下面的引擎设置，详细参见 session 模块
SessionCookieLifeTime
	session 默认存在客户端的 cookie 的时间，默认值是 3600 秒。
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 3600
SessionAutoSetCookie
	是否开启SetCookie, 默认值 true 开启。
	beego.BConfig.WebConfig.Session.SessionAutoSetCookie = true
SessionDomain
	session cookie 存储域名, 默认空。
	beego.BConfig.WebConfig.Session.SessionDomain = ""


Log配置
log详细配置，请参见 `logs 模块`。
AccessLogs
	是否输出日志到 Log，默认在 prod 模式下不会输出日志，默认为 false 不输出日志。此参数不支持配置文件配置。
	beego.BConfig.Log.AccessLogs = false
FileLineNum
	是否在日志里面显示文件名和输出日志行号，默认 true。此参数不支持配置文件配置。
	beego.BConfig.Log.FileLineNum = true
Outputs
	日志输出配置，参考 logs 模块，console file 等配置，此参数不支持配置文件配置。
	beego.BConfig.Log.Outputs = map[string]string{"console": ""}
 or
	beego.BConfig.Log.Outputs["console"] = ""










