package main

import (
	"log"
	"net/http"
	"os"
)

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return neuteredReaddirFile{f}, nil
}

type neuteredReaddirFile struct {
	http.File
}

func (f neuteredReaddirFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func main() {
	// 把以下這段不 comment out 後，就不會再產生 go的 default behaviour (which is directory listing)
	// http.Dir("./tmp/static")
	// 這部分設定了伺服器的根檔案系統為 ./tmp/static。

	// 行為：
	// http.Dir("./tmp/static") 告訴伺服器，當靜態檔案請求進來時，所有的文件都應該從 ./tmp/static 目錄中查找。
	// 例如，如果路徑是 example.txt，則伺服器會嘗試從 ./tmp/static/example.txt 中讀取檔案。
	fs := justFilesFilesystem{http.Dir("./tmp/static")}

	// 移除 URL 路徑的前綴，讓路徑變簡單
	err := http.ListenAndServe(":8080", http.StripPrefix("/", http.FileServer(fs))) 


	// // default golang behaviour: 使用 http.Dir 提供靜態檔案服務
	// // go to http://localhost:8080/, it will list example.txt
	// fs := http.FileServer(http.Dir("./tmp/static"))
	// err := http.ListenAndServe(":8080", http.StripPrefix("/", fs)) 

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}