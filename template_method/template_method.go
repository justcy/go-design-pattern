package main

import "fmt"

type Downloader interface {
	Download(uri string)
}
type implement interface {
	download()
	save()
}
type Template struct {
	implement
	uri string
}
func NewTemplate(impl implement) *Template {
	return &Template{
		implement: impl,
	}
}
func (t *Template) Download(uri string) {
	t.uri = uri
	fmt.Print("prepare downloading\n")
	t.implement.download()
	t.implement.save()
	fmt.Print("finish downloading\n")
}

func (t *Template) save() {
	fmt.Print("default save\n")
}

type HTTPDownloader struct {
	*Template
}
func NewHTTPDownloader() Downloader {
	downloader := &HTTPDownloader{}
	template := NewTemplate(downloader)
	downloader.Template = template
	return downloader
}

func (d *HTTPDownloader) download() {
	fmt.Printf("download %s via http\n", d.uri)
}

func (*HTTPDownloader) save() {
	fmt.Printf("http save\n")
}

type FTPDownloader struct {
	*Template
}

func NewFTPDownloader() Downloader {
	downloader := &FTPDownloader{}
	template := NewTemplate(downloader)
	downloader.Template = template
	return downloader
}

func (d *FTPDownloader) download() {
	fmt.Printf("download %s via ftp\n", d.uri)
}
