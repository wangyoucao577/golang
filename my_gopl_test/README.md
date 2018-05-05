# [Go] my_gopl_test  
Sample and exercise codes from learning  [The Go Programming Language](http://gopl.io). 

## 实验平台
- Linux: `Ubuntu 16.04.4 LTS`
    - `Kernel 4.13.0-36-generic`
    - `go version go1.9.4 linux/amd64`


## 学习笔记
### 基本点
- Go是一门编译型语言，且仅支持静态编译链接（不支持动态链接）
- Go原生支持unicode
- Go的编译没有警告, 要么pass, 要么error
- Go的代码通过 package 组织(类似于其他语言的modules 或 libraries)，一个package由位于单个目录下的一个或多个.go生成，目录的名字即为package的名字。
- 每个源文件都应
    - 首先以 package xxx 开始，以定义此文件属于哪个package.
    - 然后import xxx 导入所需要链接的 package (import必须在package之后)
    - 再然后才是此文件中的 代码实现
- package main，以及 func main
    - package main定义了独立的可执行程序
    - func main则定义了程序的入口函数

### Go的代码组织 
(摘自[How to Write Go Code](https://golang.org/doc/code.html)的Code origanization章节原文, 非常好的入门文档, 强烈推荐!!)  
- Overview
    - Go programmers typically keep all their Go code in a single workspace.
    - A workspace contains many version control repositories (managed by Git, for example).
    - Each repository contains one or more packages.
    - Each package consists of one or more Go source files in a single directory.
    - The path to a package's directory determines its import path.
    - Note that this differs from other programming environments in which every project has a separate workspace and workspaces are closely tied to version control repositories.

- Workspaces
    - A workspace is a directory hierarchy with three directories at its root:
        - src contains Go source files,
        - pkg contains package objects, and
        - bin contains executable commands.
    - The go tool builds source packages and installs the resulting binaries to the pkg and bin directories.

    - The src subdirectory typically contains multiple version control repositories (such as for Git or Mercurial) that track the development of one or more source packages.

    - To give you an idea of how a workspace looks in practice, here's an example:
```
bin/    
    hello                          # command executable    
    outyet                         # command executable    
pkg/    
    linux_amd64/    
        github.com/golang/example/    
            stringutil.a           # package object    
src/    
    github.com/golang/example/    
        .git/                      # Git repository     metadata    
	hello/    
	    hello.go               # command source    
	outyet/    
	    main.go                # command source    
	    main_test.go           # test source    
	stringutil/    
	    reverse.go             # package source    
	    reverse_test.go        # test source    
    golang.org/x/image/    
        .git/                      # Git repository     metadata    
	bmp/    
	    reader.go              # package source    
	    writer.go              # package source    
    ... (many more repositories and packages omitted) ...   
```    




### Go的常用工具命令    
Go提供了一系列的工具命令，都可以通过一个单独的go命令调用    
- go run：编译一个或多个 .go, 链接库文件, 并运行最终生成的可执行文件 (不会保留可执行文件)   
- go build: 编译一个或多个 .go, 链接库文件, 生成可执行程序或package   
- go install: 编译一个或多个 .go, 链接库文件, 生成可执行程序或package, 并将其对应的安装到 bin/pkg 目录下供执行或其他程序链接   
- gofmt: 格式化源代码，强制无参数的命令来统一go的代码格式, 默认行为为将diff的内容写到stdout，而要直接格式化源文件本身的话，加上 `-w` 选项   
    - `gofmt -l -w .`   
- goimports: 根据代码需要自动地添加或删除import   
- go doc: 在cmd中看go的文档   



### Reference Links 
- http://gopl.io
- http://github.com/golang-china/gopl-zh
- http://bitbucket.org/golang-china/gopl-zh