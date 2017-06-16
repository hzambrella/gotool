package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//sudo ./crawl -img_dir=data/pic

const UrlPrefix string = "http://www.aitaotu.com"

var (
	ch1     chan string
	ch2     chan string
	ch3     chan int
	img_dir string
	log_dir string
	keyword string
)

//初始化变量
func init() {
	flag.StringVar(&img_dir, "img_dir", "", "where is images to save")//定义string命令行参数
	log_dir="/var/tmp"
	keyword="Hello"
//	flag.StringVar(&log_dir, "log_dir", "/var/tmp", "where is log to save")
//	flag.StringVar(&keyword, "kw", "Hello", "search for special keyword")

	ch1 = make(chan string, 20)
	ch2 = make(chan string, 1000)
	ch3 = make(chan int, 1000)

	logpath := path.Join(log_dir, "crawl.log")//路径连接
	logfile, err := os.OpenFile(logpath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)//参数  不存在就创建|读写|追加写 
	if err != nil {
		log.Printf("create log %q ERR %s", logpath, err)
		os.Exit(1)
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)//设置flag,这个很有用。|是种操作符
	log.SetOutput(logfile)//设置日志文件路径
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()//解析flag
	if img_dir == "" || keyword == "" {
		flag.PrintDefaults()//命令行说明
		os.Exit(1)
		return
	}

	//检查目录是否存在
	img_dir = path.Join(img_dir, keyword)
	file, err := os.Stat(img_dir)
	if err != nil || !file.IsDir() {
		dir_err := os.Mkdir(img_dir, os.ModePerm)
		if dir_err != nil {
			fmt.Printf("create dir %q failed\n", img_dir)
			os.Exit(1)//程序结束
		}
	}

	go getListUrl()
	go parseListUrl()
	go downloadImage()

	count := 0
	for num := range ch3 {
		count = count + num
		fmt.Println("count:", count)
	}
	fmt.Println("crawl end")
}

func getListUrl() {
	docUrl := fmt.Sprintf("%s/search/%s/", UrlPrefix, keyword)
	doc, err := goquery.NewDocument(docUrl)//获得HTML
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}

	doc.Find(".picbox").Each(func(i int, s *goquery.Selection) {
		text, _ := s.Find("a").Attr("href")
		list_url := UrlPrefix + text
		ch1 <- list_url
	})
}

//根据模块和总数据列出所有的图片页面
func parseListUrl() {
	suffix := ".html"
	for list_url := range ch1 {
		page_count := getPageCount(list_url)
		prefix := strings.TrimRight(list_url, suffix)
		for i := 1; i <= page_count; i++ {
			img_list_url := prefix + "_" + strconv.Itoa(i) + suffix
			ch2 <- img_list_url
		}
	}
}

//获取总页数
func getPageCount(list_url string) (count int) {
	count = 0
	doc, _ := goquery.NewDocument(list_url)
	doc.Find(".pages ul li").Each(func(i int, s *goquery.Selection) {
		text := s.Find("a").Text()
		if text == "末页" {
			last_page_url, _ := s.Find("a").Attr("href")
			prefix := strings.Trim(last_page_url, ".html")
			index := strings.Index(prefix, "_")
			last_page_num := prefix[index+1:]
			page_num, _ := strconv.Atoi(last_page_num)
			count = page_num
		}
	})
	return count
}

//解析图片url
func downloadImage() {
	for img_list_url := range ch2 {
		doc, _ := goquery.NewDocument(img_list_url)
		doc.Find("#big-pic p a").Each(func(i int, s *goquery.Selection) {
			img_url, _ := s.Find("img").Attr("src")
			go func() {
				saveImages(img_url)
			}()
		})
	}
}

//下载图片
func saveImages(img_url string) {
	log.Printf("Get %s", img_url)
	u, err := url.Parse(img_url)
	if err != nil {
		log.Println("parse url failed:", img_url, err)
		return
	}

	tmp := strings.TrimLeft(u.Path, "/")
	tmp = strings.ToLower(strings.Replace(tmp, "/", "-", -1))
	filename := path.Join(img_dir, tmp)

	if checkExists(filename) {
		log.Printf("Exists %s", filename)
		return
	}

	response, err := http.Get(img_url)
	if err != nil {
		log.Println("get img_url failed:", err)
		return
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("read data failed:", img_url, err)
		return
	}

	image, err := os.Create(filename)
	if err != nil {
		log.Println("create file failed:", filename, err)
		return
	}

	ch3 <- 1
	defer image.Close()
	image.Write(data)
}

func checkExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
