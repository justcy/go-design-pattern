package main

func main() {
	http := NewHTTPDownloader()
	http.Download("http://example.com/abc.zip")
	ftp := NewFTPDownloader()
	ftp.Download("ftp://example.com/123.zip")
}
