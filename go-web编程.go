HTTP 是一种无状态、由文本构成请求-响应协议，使用的是客户端-服务器计算模型
	无状态：每一次请求都是独立的
客户端被称作用户代理，服务器被称为Web服务器 HTTP客户端都是一个web浏览器

HTTP请求
	请求行（请求方法 URI HTTP版本）
	零个或多个请求首部header
	一个空行
	可选报文主体body
请求方法：
	HTTP 0.9只有GET一个方法
	HTTP 1.0添加了POST方法和HEAD方法
	HTTP 1.1 这个当前使用得最多的
		添加了PUT、DELETE、OPTIONS、TRACE、CONNECT 5个方法
		要求实现的只有 GET 和 HEAD方法
	GET--命令服务器返回指定的资源
	HEAD--与GET类似，只是不要求报文主体，通常用于不获取报文主体情况下，取得响应的首部
	POST--命令服务器将报文主体的内容传递给URI指定的资源，
		至于服务器具体会对这些数据执行什么动作则服务器自己决定
		HTML 2.0 通过添加HTML表单来实现 form标签的属性method设置请求方法(仅仅支持两个)：GET POST
	PUT--命令服务器将报文主体的内容设置为。URI指定的资源。
			如果URI指定的位置上已经有数据存在，报文主体内容就替代已有数据
			如果资源不存在，那么在URI指定的位置上新创建一个资源
	DELETE--命令服务器删除URI指定的资源
	TRACE--命令服务器返回请求本身	通过这个方法，客户端可以知道介于它和服务器之间的其他服务器是如何处理请求的
	OPTIONS--返回服务器支持的HTTP方法列表
	CONNECT--命令服务器与客户端建立一个网络连接。 通常用于设置SSL隧道已开启HTTPS功能
	PATCH--命令服务器使用报文主体中的数据对URI指定的资源进行修改

	安全的请求方法：GET、HEAD、OPTIONS 和 TRACE 不会对服务器的状态进行更改
	不安全的请求方法：POST、PUT、DELETE	能够对服务器的状态进行修改

	幂等的请求方法：安全的方法都是幂等的	 PUT 和 DELETE 也是幂等的 
	不幂等的请求方法：PATCH 和 POST 

请求首部：记录了与请求本身和客户端有关的信息	有若干个用冒号分隔的键值对，最后以回车CR和换行LF结尾
	HOST是HTTP 1.1 唯一强烈要求的首部
	Accept 	客户端在HTTP响应中能够接收的内容类型
	Accept-Charset	客户端在HTTP响应中能够接收的内容的字符集
	Authorization	发送基本的身份验证证书
	Cookie	客户端把服务器之前发送过来的所有cookie回传给服务器
	Content-Length	请求主体的字节长度
	Content-Type	请求主体 内容类型	
					默认 x-www-form-urlen-coded 
					上传文件	multipart/form-data
	Host 	服务器的名字和端口号 默认80
	Referrer	发送请求的页面所在的地址
	User-Agent	对发起请求的客户端进行描述

HTTP响应
	一个状态行 (状态码status-code 原因短语reason-phrase)
	零个或多个的响应首部
	一个空行
	一个可选的报文主体

状态码：
	1xx：情报状态码  服务器告知客户端：自己已经接收到了客户端发送的请求，并且已经进行了处理
	2xx：成功状态码  服务器告知客户端：自己已经接收到了客户端发送的请求，并且已经 成功地 进行了处理  
		常用 200 OK
	3xx：重定向状态码	服务器告知客户端：自己已经接收到了客户端发送的请求，并且已经成功地处理了请求，为了完成请求的动作，客户端还需要一些工作。
		这些状态码大多用于实现了URL重定向
		301 永久重定向	网址和显示内容都会更新
		302 临时重定向 容易发生URL劫持  保持原网址而内容确是新网址的内容
	4xx：客户端错误状态码	服务器告知客户端：发送的请求出现了问题
		常见： 404 Not Found
	5xx：服务器错误状态码	服务器告知客户端：自己由于某些原因无法正确的处理请求
		常见：500 Internal Server Error 
响应首部
	Allow	告知客户端：支持哪些请求方法
	Content-Length	响应主体的字节长度
	Content-Type	响应包含主体，这里说明响应主体的内容类型
	Date	格林尼治标准时间GMT 记录的当前时间
	Location	仅仅用于重定向时使用，告知客户端应该向哪个URL发送请求
	Server	返回响应的服务器的域名
	Set-Cookie	在客户端设置一个cookie，响应里可以设置多个 Set-Cookie首部
	WWW-Authenticate	服务器告知客户端：在Authorization请求首部中应该提供哪种类型的身份验证信息

URI： 
	URI统一资源标识符Uniform Resource Identifier：
		分为 URN统一资源名称Uniform Resource Name 
			和 URL统一资源定位符Uniform Resource Location 
	<方案名称(required)>:<分层部分(required)>[ ? <查询参数(optional)> ][ # <片段(optional)> ]
	方案名称scheme name：有很多  web服务通常使用HTTP
	分层部分：资源的识别信息 UR
		http://sausheong:password@www.example.com/docs/file?name=sausheong&location=singapore#summary
		https://www.baidu.com/
	使用URL编码对特殊符号进行转换（URL编码又称百分号编码） %+ASCII码： %20

HTTP/2  由SPDY/2协议改进而来
	1、HTTP 1.x 使用纯文本方式表示；HTTP/2 使用二进制
	2、HTTP 1.x 每次只能发送单个请求；
		HTTP/2 是完全多路复用的 
			多个请求和响应可以在同一时间里使用同一个连接
			还会对请求首部进行压缩以减少需要传送的数据量，并允许服务器将响应推送到客户端
	3、在HTTP方法和状态码等功能 HTTP/2与HTTP 1.x 保持了一致相同

Web应用：
	1、向发送请求的客户端返回HTML，客户端向用户展示渲染后HTML
	2、程序向客户端传送数据时必须使用HTTP协议
  可以这么描述
  	1、通过HTTP协议，以HTTP请求报文的形式获取客户端输入；
  	2、对HTTP请求报文进行处理，并执行必要的操作
  	3、生成HTML，并以HTTP响应报文的形式将其返回给客户端

Web应用可以分为： 处理器  模版引擎
MVC模型：
  	控制器是根据请求对模型进行修改
  		它是苗条的，它应该只包含路由代码和HTTP报文的解包和打包逻辑
  	模型用户表示底层的数据
  		它应该的丰满的，它包含应用的逻辑和数据
  	视图是用来生成HTML或者称为渲染
处理器
	1、接收和处理客户端发来的请求
	2、调用模版引擎生成HTML，然后把数据填充至响应报文当中
  处理器既是控制器controller也是模型model
  
模版引擎
	模版引擎使用数据和模版生成HTML，模版中可以包含HTML也可以不包含
	分类：
		静态模版 直接用数据替换占位符 无逻辑模版
		动态模版	在静态模版进行了扩展，包含了一些编程语言结构，如：条件控制语句、迭代语句和变量	
			如：JSP ASP ERB

go开发的Web应用：可扩展、模块化和可维护等特性

