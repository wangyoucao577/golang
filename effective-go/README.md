# Effective Go
我的 [Effective Go](https://golang.org/doc/effective_go.html) 学习笔记, 主要是新知识点的highlight.     

## The blank identifier
[The blank identifier](https://golang.org/doc/effective_go.html#blank), i.e. `_`.      
除了通常的忽略不需要的变量外, 另外两种特殊用法:    

- [Import for side effect](https://golang.org/doc/effective_go.html#blank_import): 导入不需要显式使用的包, 以获得其附属功能.    

例如    
```go
import _ "net/http/pprof"
```    
导入了`net/http/pprof`包. 虽然未显式地使用这个包, 但是由于`import`会调用`net/http/pprof`的`init()`, 而其中注册了用于debug的HTTP handlers, 于是获得了相应的debug功能.    

- [Interface checks](https://golang.org/doc/effective_go.html#blank_implements): 通过`_`检查类型是否实现了某个约定必须要实现的接口, 将未实现的错误在编译时暴露出来.    

例如 [json.RawMessage](https://golang.org/pkg/encoding/json/#RawMessage) 类型约定必须实现 [json.Marshaler](https://golang.org/pkg/encoding/json/#Marshaler) 和 [json.Unmarshaler](https://golang.org/pkg/encoding/json/#Unmarshaler)接口, 通过如下定义可实现编译时检查, 详见[代码](https://golang.org/src/encoding/json/stream.go?s=6715:6737#L275).        
```go
var _ Marshaler = (*RawMessage)(nil)
var _ Unmarshaler = (*RawMessage)(nil)
```

## Embedding
[Embedding](https://golang.org/doc/effective_go.html#embedding).     
`Interface`的Embedding比较容易理解, 重点是`struct`的Embedding提供了方法的类似继承的功能(要求匿名嵌入).     

- 例如示例中的`bufio.ReadWriter`的`struct`    
其匿名嵌入了`bufio.Reader`和`bufio.Writer`两个`struct`. 由于`bufio.Reader`和`bufio.Writer`实现了`io.ReadWriter`接口, 于是`bufio.ReadWriter`便也实现了`io.ReadWriter`接口, 不再需要显式实现.    
```go
// ReadWriter stores pointers to a Reader and a Writer.
// It implements io.ReadWriter.
type ReadWriter struct {
    *Reader  // *bufio.Reader
    *Writer  // *bufio.Writer
}
```    

- 而若嵌入时带上成员名, 则没有此特性     
如下示例, `ReadWriter`若需要实现`io.Reader`接口, 需要写如下的传递代码, 比较啰嗦.     
```go
type ReadWriter struct {
    reader *Reader
    writer *Writer
}

func (rw *ReadWriter) Read(p []byte) (n int, err error) {
    return rw.reader.Read(p)
}
```

- my test code 
```bash
$ cd embedding 
$ go build 
$ ./embedding
type main.readerNoFieldName satisfied io.Reader
type main.readerWithFieldName didn't satisfy io.Reader
```

