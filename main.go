package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/fofapro/fofa-go/fofa"
	"github.com/spaolacci/murmur3"
	"gopkg.in/yaml.v2"
	"hash"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type apiconfig struct {
	Email string `yaml:"Email,omitempty"`
	Apikey string `yaml:"Apikey,omitempty"`
}

type FofaResult struct {
	Error   bool       `json:"error"`
	Mode    string     `json:"mode"`
	Page    int        `json:"page"`
	Query   string     `json:"query"`
	Size    int        `json:"size"`
	Results [][]string `json:"results"`
}

var (
	SearchKeyword = flag.String("k","","example  -k title=\"百度\"\nexample  -k domain=\"baidu.com\"\nexample  -k 'domain=\"baidu.com\" && city=\"Nanjing\"'\n......\nAnd Support Fofa Other Syntax")
	SearchFile    = flag.String("f","","example  -f target.txt" )
	OutputFile    = flag.String("o","","example -o result.csv")
	IconHashCount = flag.String("i","","example -i https://www.baidu.com/favicon.ico")
)

func main()  {
	flag.Parse()
	if (*SearchKeyword =="" && *OutputFile =="" && *SearchFile == "" && *IconHashCount =="")||(*SearchKeyword !="" && *OutputFile !="" && *SearchFile != "" && *IconHashCount !=""){
		flag.Usage()
	}
	if (*SearchKeyword !="" && *OutputFile != ""){
		QueryFofa()
	}
	if (*SearchKeyword !="" && *OutputFile == "") {
		*OutputFile ="result.csv"
		QueryFofa()
	}
	if (*SearchFile !="" && *OutputFile!="" ){
		FofaReadfile()
	}
	if (*SearchFile !="" && *OutputFile =="" ){
		*OutputFile = "result.csv"
		FofaReadfile()
	}
	if *IconHashCount != "" {
		PrintResult(Mmh3Hash32(IconHash(*IconHashCount)))
	}
}

func QueryFofa()  {
	myConfig:=GetConfig("config.yaml")
	email := myConfig.Email
	key := myConfig.Apikey
	//fmt.Println("登录账号为：",email)
	clt := fofa.NewFofaClient([]byte(email), []byte(key))
	if clt == nil {
		fmt.Printf("create fofa client\n")
		return
	}
	var searcccc string
	searcccc = *SearchKeyword
	searcccc = searcccc + " && (is_honeypot=false && is_fraud=false)"
	fmt.Println(searcccc)
	result, err := clt.QueryAsJSON(1, []byte(searcccc))
	if err != nil {
		fmt.Printf("%v\n", err.Error())
	}
	fmt.Println("正在查询:",searcccc)
	//fmt.Printf("%s\n", result)
	Allresult :=FofaResult{}
	json.Unmarshal(result,&Allresult)
	fmt.Printf("%s\n", Allresult.Results)
	if len(Allresult.Results) > 0 {
		arraryTocsv(Allresult.Results)
	}
	//fmt.Println(Allresult.Size)
	if Allresult.Size >10{
		fofapage1 := Allresult.Size/88
		fofapage := fofapage1+1
		//fmt.Println(fofapage)
		for i:=2;i<fofapage;i++{
			result, err = clt.QueryAsJSON(uint(i), []byte(*SearchKeyword))
			if err != nil {
				fmt.Printf("%v\n", err.Error())
			}
			json.Unmarshal(result,&Allresult)
			arraryTocsv(Allresult.Results)
			fmt.Printf("%s\n", Allresult.Results)
		}
	}
}

func FofaReadfile()  {
	fi, err := os.Open(*SearchFile)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		*SearchKeyword=string(a)
		QueryFofa()
	}

}

func arraryTocsv(temp [][] string)  {
	tempLen := len(temp)
	_, err := os.Stat(*OutputFile)
	if err == nil{
		fd,_:=os.OpenFile(*OutputFile,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
		for i:=0;i<tempLen;i++ {
			domaintemp := temp[i][1]
			a:=HttpFormat(domaintemp)
			titletemp := temp[i][4]
			titletemp = strings.Replace(titletemp,","," ",-1)
			titletemp = strings.Replace(titletemp,"\n"," ",-1)
			titletemp = strings.Replace(titletemp,"        ","",-1)
			rrr := (temp[i][0]+","+a+","+temp[i][2]+","+temp[i][3]+","+titletemp+","+temp[i][5]+","+temp[i][6]+"\n")
			fd.Write([]byte(rrr))
		}
		return
	}
	if os.IsNotExist(err){
		//fmt.Println("File not exist")
		f, err := os.Create(*OutputFile)
		defer f.Close()
		if err != nil {
			// 创建文件失败处理
		} else {
			content := "Dmoain,Host,Ip,Port,Title,Country,City"+"\n"
			_, err = f.Write([]byte(content))
			if err != nil {
				// 写入失败处理
			}
			fd,_:=os.OpenFile(*OutputFile,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
			for i:=0;i<tempLen;i++ {
				domaintemp := temp[i][1]
				a:=HttpFormat(domaintemp)
				titletemp := temp[i][4]
				titletemp = strings.Replace(titletemp,","," ",-1)
				titletemp = strings.Replace(titletemp,"\n"," ",-1)
				titletemp = strings.Replace(titletemp,"        ","",-1)
				rrr := (temp[i][0]+","+a+","+temp[i][2]+","+temp[i][3]+","+titletemp+","+temp[i][5]+","+temp[i][6]+"\n")
				fd.Write([]byte(rrr))
			}
		}
		return
	}
	fmt.Println("File error")
	return

}

func GetConfig(filename string) apiconfig {
	var configfile apiconfig
	configinfo,err:=ioutil.ReadFile(filename)
	if err!=nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}
	err=yaml.Unmarshal(configinfo,&configfile)
	if err!=nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return configfile
}

func HttpFormat(formatstr3 string) string{
	if strings.Contains(formatstr3,"https://"){
		return formatstr3
	}else {
		formatstr3 = "http://" +formatstr3
		return formatstr3
	}
}

func IconHash(formatstr4 string) []byte {
	resp, err := http.Get(formatstr4)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	formatstr4 = base64.StdEncoding.EncodeToString(body)
	var buffer bytes.Buffer
	for i := 0; i < len(formatstr4); i++ {
		ch := formatstr4[i]
		buffer.WriteByte(ch)
		if (i+1)%76 == 0 {
			buffer.WriteByte('\n')
		}
	}
	buffer.WriteByte('\n')
	return buffer.Bytes()
}
func Mmh3Hash32(raw []byte) string {
	var h32 hash.Hash32 = murmur3.New32()
	h32.Write([]byte(raw))
	return fmt.Sprintf("%d", int32(h32.Sum32()))
}
func PrintResult(result string) {
	fmt.Printf("icon_hash=\"%s\"\n", result)
}
