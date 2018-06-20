package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func main() {
	// 如何具体分配到相应的函数来处理请求呢？conn首先会解析request:c.readRequest(),
	// 然后获取相应的handler:handler := c.server.Handler，也就是我们刚才在调用函数ListenAndServe时候的第二个参数，
	// 传递是nil，也就是为空，那么默认获取handler = DefaultServeMux,那么这个变量用来做什么的呢？
	// 对，这个变量就是一个路由器，它用来匹配url跳转到其相应的handle函数，那么这个我们有设置过吗?有，
	// 我们调用的代码里面第一句不是调用了http.HandleFunc("/", sayhelloName)嘛。
	// 这个作用就是注册了请求/的路由规则，当请求uri为"/"，路由就会转到函数sayhelloName，DefaultServeMux会调用ServeHTTP方法，
	// 这个方法内部其实就是调用sayhelloName本身，最后通过写入response的信息反馈到客户端。
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
