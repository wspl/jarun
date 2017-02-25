package main

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Download(url string, dst string) {
	headRes, err := http.Head(url)
	if err != nil {
		panic(err)
	}
	defer headRes.Body.Close()

	size, err := strconv.Atoi(headRes.Header.Get("Content-Length"))
	if err != nil {
		panic(err)
	}

	out, err := os.Create(dst)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	done := make(chan int)
	go func() {
		gotSize := 0
		for {
			select {
			case <- done:
				println("100.0% - Completed!")
				return
			default:
				file, err := os.Open(dst)
				if err != nil {
					panic(err)
				}

				fi, err := file.Stat()
				if err != nil {
					panic(err)
				}

				doneSize := int(fi.Size())
				offsetByte := doneSize - gotSize
				gotSize = doneSize

				if size == 0 {
					size = 1
				}

				percent := float64(gotSize) / float64(size) * 100
				rate := float64(offsetByte) / 1024

				p := strconv.FormatFloat(percent, 'f', 1, 64) + "%"

				r := ""
				if rate < 1024 {
					r = strconv.FormatFloat(rate, 'f', 2, 64) + " KB/s"
				} else {
					r = strconv.FormatFloat(rate / 1024, 'f', 2, 64) + " MB/s"
				}

				d := ""
				if gotSize < 1024 * 1024 {
					d = strconv.FormatFloat(float64(gotSize) / 1024, 'f', 2, 64) + " KB"
				} else {
					d = strconv.FormatFloat(float64(gotSize) / 1024 / 1024, 'f', 2, 64) + " MB"
				}

				println(p, "-", d, "|", r)
			}
			time.Sleep(time.Second / 2)
		}
	}()

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	n, err := io.Copy(out, res.Body)
	if err != nil {
		panic(err)
	}

	done <- int(n)
}