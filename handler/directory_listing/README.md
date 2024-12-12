This is to demo directory listing behaviour and how to prevent it.

Run:
```
go run directory_listing.go
```
  
Go to `http://localhost:8080/example.txt` on browser.  
This will map to `tmp/static/example.txt`.  
So, if there is a file: tmp/static/example.txt, it will show the content of the file.  
If there is no such file, it will show 404, rather than list out the name of the file.

如果想看 golang  原有的不好的 directory listing behaviour, 則不 comment out 這段：
// // default golang behaviour: 使用 http.Dir 提供靜態檔案服務
	// // go to http://localhost:8080/, it will list example.txt
	// fs := http.FileServer(http.Dir("./tmp/static"))
	// err := http.ListenAndServe(":8080", http.StripPrefix("/", fs)) 