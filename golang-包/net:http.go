http包提供了HTTP客户端和服务端的实现。
Get、Head、Post和PostForm函数发出HTTP/ HTTPS请求。

Get、Head、Post和PostForm函数发出HTTP/ HTTPS请求。
  resp, err := http.Get("http://example.com/")
...
  resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
...
  resp, err := http.PostForm("http://example.com/form",
  url.Values{"key": {"Value"}, "id": {"123"}})

程序在使用完回复后必须关闭回复的主体。
	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

要管理HTTP客户端的头域、重定向策略和其他设置，创建一个Client：
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}
	resp, err := client.Get("http://example.com")
	// ...
	req, err := http.NewRequest("GET", "http://example.com", nil)
	// ...
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	// ...

要管理代理、TLS配置、keep-alive、压缩和其他设置，创建一个Transport：
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{RootCAs: pool},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://example.com")
Client和Transport类型都可以安全的被多个go程同时使用。出于效率考虑，应该一次建立、尽量重用。

ListenAndServe使用指定的监听地址和处理器启动一个HTTP服务端。处理器参数通常是nil，
这表示采用包变量DefaultServeMux作为处理器。
Handle和HandleFunc函数可以向DefaultServeMux添加处理器。

	http.Handle("/foo", fooHandler)
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))


