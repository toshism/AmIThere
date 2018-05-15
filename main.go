package main

import (
	"fmt"
	"time"

	"gocv.io/x/gocv"
)

func capture_image(img *gocv.Mat) {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer webcam.Close()
	webcam.Read(img)
}

func am_i_there(img *gocv.Mat) bool {
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	classifier.Load("/home/tosh/bin/haarcascade_profileface.xml")

	rects := classifier.DetectMultiScale(*img)

	if len(rects) > 0 {
		return true
	} else {
		return false
	}
}

func main() {
	img := gocv.NewMat()
	defer img.Close()

	am_there := "no"
	for i := 0; i <= 5; i++ {
		capture_image(&img)
		if am_i_there(&img) {
			am_there = "yes"
			break
		}
		time.Sleep(5 * time.Second)
	}
	fmt.Print(am_there)
}
