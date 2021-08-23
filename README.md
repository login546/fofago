# fofago

Fofa批量查询工具，查询时自动拼接`&& (is_honeypot=false && is_fraud=false)`排除干扰资产。

~~~
  -f string
        example  -f target.txt
  -k string
        example  -k 'title="百度"''
        example  -k 'domain="baidu.com"'
        example  -k 'domain="baidu.com" && city="Nanjing"'
        ......
        And Support Fofa Other Syntax
  -o string
        example -o result.csv
~~~

首次使用需修改配置文件config.yaml

输入FOFA的email和key

## 使用

### 查询单条语句

~~~bash
./fofago -k 'domain="baidu.com"'         Mac下
fofago.exe -k 'domain="baidu.com"'       Windows下
~~~

![image](https://user-images.githubusercontent.com/38073810/130382551-0eaa0d10-fcf3-4aa9-819c-7c0cbe6ffa8b.png)

同时会在本地生成result.csv

如需指定输出文件名则增加`-o xxxx.csv`

![image](https://user-images.githubusercontent.com/38073810/130382622-8cf1f3ea-9bb0-4302-84c7-c438113dc8ed.png)

### 查询多条语句

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

![image](https://user-images.githubusercontent.com/38073810/130382724-a25e2dbf-aeba-4dea-b0f5-c61a58beafab.png)

![image](https://user-images.githubusercontent.com/38073810/130382759-1dc4353a-0f33-425f-923a-96de56596bfb.png)

### 输出

默认输出为result.csv，或使用以下命令自定义输出文件名

~~~bash
./fofago -f url.txt -o test.csv          Mac下
fofago.exe -f url.txt -o test1.csv       Windows下
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

