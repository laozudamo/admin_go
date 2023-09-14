package controller

import (
	"admin_go/response"
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func HandleVideo(c *gin.Context) {
	var wg sync.WaitGroup
	wg.Add(1)

	var url string
	go func(w *sync.WaitGroup) {
		url = SplitVideo()
		wg.Done()
	}(&wg)
	wg.Wait()

	data := map[string]interface{}{
		"url": url,
	}
	response.Success(c, 200, nil, data)
}

func SplitVideo() string {
	t1 := time.Now().Second()
	inputPath := "./uploads/1021312.mp4"
	var out bytes.Buffer
	hashStr := fmt.Sprintf("%d", time.Now().Unix())
	outputPath := "./uploads/" + hashStr + ".mp4"
	cmdArguments := []string{"-ss", "00:00:15", "-t", "00:19:00", "-i", inputPath, "-vcodec", "copy", "-acodec", "copy", outputPath}
	// ffmpeg 切割视频
	cmd := exec.Command("ffmpeg", cmdArguments...)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		panic(err)
	} else {
		t2 := time.Now().Second() - t1
		fmt.Println("...... 转换成功.......")
		fmt.Println("t2:", t2)
		return hashStr
	}
}

func GetFps() {

}

func Command(command []string) {
	var out bytes.Buffer
	cmd := exec.Command("ffmpeg", command...)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("...... 转换成功.......")
	}
}

func translate(inputPath string, fileName string, ext string) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		var out bytes.Buffer
		fmt.Print("默认输出音频路径: ./audio\n")
		outputName := strings.Replace(fileName, ext, "mp3", 1)
		fmt.Println(outputName)
		outPutDir := "./audio"
		outputPath := outPutDir + "/" + outputName
		cmdArguments := []string{"-i", inputPath, outputPath}

		// ffmpeg -i input.mp4 output.mp3
		cmd := exec.Command("ffmpeg", cmdArguments...)
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("...... 转换成功.......")
		}
		wg.Done()
	}()
	fmt.Println("hello")
	wg.Wait()
}

func split(inputPath string, fileName string, ext string) {
	var wg sync.WaitGroup
	wg.Add(1)
	//now := time.Now().Format("2006-01-02 15:04:05")
	go func() {
		var out bytes.Buffer
		//outputName := strings.Replace("121212", ext, "mp4", 1)
		//fmt.Println(outputName)
		outPutDir := "./video"
		outputPath := outPutDir + "/" + "11121212" + ".mp4"
		cmdArguments := []string{"-ss", "00:00:15", "-t", "00:01:00", "-i", inputPath, "-vcodec", "copy", "-acodec", "copy", outputPath}
		// ffmpeg 切割视频
		cmd := exec.Command("ffmpeg", cmdArguments...)
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			panic(err)
		} else {
			fmt.Println("...... 转换成功.......")
		}
		wg.Done()
	}()

	wg.Wait()
}
