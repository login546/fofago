# fofago
Fofa批量查询工具

~~~
  -f string
        example  -f target.txt
  -k string
        example  -k title="百度"
        example  -k domain="baidu.com"
        example  -k 'domain="baidu.com" && city="Nanjing"'
        ......
        And Support Fofa Other Syntax
  -o string
        example -o result.csv
~~~



## 编译

### Mac

~~~bash
GOARCH=amd64 GOOS=darwin go build -o fofago-darwin main.go
~~~

### Linux

~~~bash
GOARCH=amd64 GOOS=linux go build -o fofago-linux main.go
~~~

### Windows

~~~bash
GOARCH=amd64 GOOS=windows go build -o fofago-windows.exe main.go
~~~



## 使用

####查询单条语句

~~~bash
./fofago -k domain="baidu.com"         Mac下
fofago.exe -k domain="baidu.com"       Windows下
~~~

#### 查询多条语句

新建一个文本，文本格式每行仅限一条查询语句，如下url.txt

~~~
app="佑友-佑友防火墙"
title="幻阵"
title="360新天擎"
title="登录_威思客"
"ClusterEngine" && title=="TSCEV4.0 login"
~~~

使用以下命令

~~~bash
./fofago -f url.txt          Mac下
fofago.exe -f url.txt        Windows下
~~~



#### 输出

默认输出为result.csv，或使用以下命令

~~~bash
./fofago -f url.txt -o test.csv          Mac下
fofago.exe -f url.txt -o test1.csv       Windows下
~~~

