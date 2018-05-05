# [Go] my_gopl_test  
Sample and exercise codes from learning  [The Go Programming Language](http://gopl.io). 

## 实验平台
- Linux: `Ubuntu 16.04.4 LTS`
    - `Kernel 4.13.0-36-generic`
    - `go version go1.9.4 linux/amd64`


## 学习笔记
### 基本点
- `Go`是一门编译型语言，且仅支持静态编译链接（不支持动态链接）
- `Go`原生支持`unicode`
- `Go`的编译没有警告, 要么pass, 要么error
    - `Go`语言不允许 unused local variable, 否则会报编译错误
    - `import`但未被引用的包, 会导致编译错误
- `Go`的代码通过`package`组织(类似于其他语言的modules 或 libraries)
    - 一个`package`由位于单个目录下的一个或多个.go生成
    - 通常目录的名字即为`package`的名字
    - 而`package xxx`中的`xxx`为命名空间(引用名)
    - 按照惯例是导入路径的最后一段与命名空间一致
- 每个源文件都应
    - 首先以 `package xxx` 开始，以定义此文件属于哪个package.
    - 然后`import xxx` 导入所需要链接的`package`(`import`必须在`package`之后)
    - 再然后才是此文件中的代码实现
- `package main`，以及 `func main`
    - `package main`定义了独立的可执行程序
    - `func main`则定义了程序的入口函数
- 注释： `//`
- `Go`中 函数和包级别(package level entities)的变量/函数可以任意顺序声明, 并不影响其调用
- 变量的几种声明/初始化形式(声明即初始化)
    - `var s1, s2 string` (声明2个`string`变量, 隐式初始化为`""`字符串. 若类型为`int`, 则隐式初始化为`0`)
	- `s1,s2 := "",""` (声明2个变量s1 s2, 以空字符串`""`初始化, 于是即推导出其为`string`类型. 此种方式术语叫做 short variable declaration)
	- `var s1 = ""` (冗余的 `var`, 使用的较少)
	- `var s1 string = ""` (冗余的`var` 和`string`, 用的也比较少)
- Go中的循环语句仅有for一种, 其有几种形式
    - `for initialization, condition, post { `
        - 此种形式下, initialization/condition/post 均可以省略, 左大括号"{" 必须在 post 的同一行
	- `for index, arg := range os.Args[1:] {`
        - 此种区间遍历形式, 提供 索引 和 值 两个参数. 若不需要其中某个, 经常是不需要 index, 可以用 blank identifier "_" 即下划线来代替(Go语言不允许unused local variable). 如 `for _, arg := range os.Args[1:] {`






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
- gofmt: 格式化源代码，**强制无参数的命令来统一go的代码格式**, 默认行为为将diff的内容写到stdout，而要直接格式化源文件本身的话，加上 `-w` 选项   
    - `gofmt -l -w .`   
- goimports: 根据代码需要自动地添加或删除import   
- go doc: 在cmd中看go的文档   

### 杂项
- Go中的区间索引：
    - 左闭右开原则. 即区间包括第一个索引元素, 不包括最后一个。（比如 a = [1, 2, 3, 4, 5], a[0:3]=[1,2,3], 即包含左边第一个元素a[0], 但不包含右侧的索引元素a[3]. ）
    - 区间索引的左、右参数分别可以省略。左参数省略则默认为0, 右参数省略则为len(a)
- 名字的作用域
    - 函数内部定义的名字，只在函数内部有效
    - 函数外部定义的名字(包级名字), 在整个包的所有文件中都可以访问
        - 包级名字，若首字母大写(包括函数名和变量名)，那么就是导出的名字，即可以被外部的包访问
        - 包级名字，若首字母小写(包括函数名和变量名)，那么就是属于包内部的名字，可以在包的所有文件中访问
    - 包本身的名字，一般总是用小写
- 命名风格
    - 倾向于不要太长的名字
    - 倾向于驼峰命名法(优先大小写分隔，而不是下划线分隔)
- 四种声明
    - `var`: 变量声明
    - `const`: 常量声明
    - `type`: 类型声明
    - `func`: 函数声明
- 变量的默认零值    
(零值初始化机制可以确保每个声明的变量总是有一个良好定义的值, 因此在`Go`中不存在未初始化的变量)
    - `bool`: `false`
    - 数值型: `0`
    - `string`: `""`
    - 接口或引用类型(包括`slice/map/chan/func`):`nil`
    - 指针: `nil`  
- 简短变量声明(例如 `a:= 0`)
    - 简短变量声明`:=`是声明语句, 而`=`是赋值语句
    - 简短变量声明语句中必须至少要声明一个新的变量, 否则会编译失败(已声明过的变量在简短变量声明语句中等价于赋值)
- 指针
    - `C`风格, 即通过`&`取地址, 通过`*`取值, 类型为`*T`
- 变量的生命周期与内存分配
    - `Go`中作用域的概念与`C/C++`中的作用域的概念不同
        - `Go`中的声明语句的作用域是指源代码中可以有效使用这个名字的范围
        - `Go`中的作用域是编译时概念, 而生命周期是运行时概念
            - 前者是变量名字的可见范围, 后者是实际变量的可引用时间段
        - 注意`for/if/switch`等的隐式词法域
            - 比如`if`条件中声明的变量, 在`else block`中也可见
    - 变量的回收由`Go`垃圾回收器负责, 而是否可以垃圾回收的唯一标志为变量是否仍然可达
        - 所以局部变量的地址返回后也依然有效, 因为依然可达
    - `new`在`Go`中只是一个内建的函数
        - `new`函数返回的是指针
        - 通过`new`函数创建变量和直接声明的变量没有什么本质区别(后者需要一个临时变量, 然后才能取地址)
        - `new`和变量是在栈上分配还是堆上分配没有任何关系
        - 通常情况下`new`的使用较少
    - 包级变量的生命周期和整个程序的运行一致
    - 局部变量的声明周期, 从声明开始, 到不可达结束
    - 变量的内存分配
        - 由`Go`编译器自动决定是在栈上分配还是堆上分配
        - 对于局部变量来说
            - 从函数中逃逸的变量, 必须在堆上分配
            - 不从函数中逃逸的变量, 由编译器自动决定在栈上分配还是堆上分配
- 命令行参数处理
    - `Go`中一般通过`flag`包, 类似于`python`中的`argparse`
- 包的初始化
    - 依次初始化(会按照变量的初始化依赖顺序)
    - 包中复杂变量的初始化, 可以通过特殊的`init()`函数来进行(进入`main()`之前自动被调用)
    - 每个包只会被初始化一次, 不会重复初始化
    - `package main`最后被初始化


## Reference Links 
- http://gopl.io
- http://github.com/golang-china/gopl-zh
- http://bitbucket.org/golang-china/gopl-zh