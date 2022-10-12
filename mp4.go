package ExtractAudio

import (
	"github.com/zhangyiming748/ExtractAudio/log"
	"github.com/zhangyiming748/replace"
	"os"
	"os/exec"
	"strings"
)

type File struct {
	in  string // 输入文件绝对路径
	out string // 输出文件绝对路径
}

func Mp4(dir string) {
	//var files []File
	oldFiles:=getFiles(dir, "mp4")
	for _, Name := range oldFiles {
		oldName:=strings.Join([]string{dir,Name},"/")
		newName := strings.Replace(oldName, ".mp4", ".m4a", -1)
		command(oldName,newName)

	}
}
func getFiles(dir, pattern string) []string {
	files, _ := os.ReadDir(dir)
	var aim []string
	types := strings.Split(pattern, ";") //"wmv;rm"
	for _, f := range files {
		//fmt.Println(f.Name())
		if l := strings.Split(f.Name(), ".")[0]; len(l) != 0 {
			//log.Info.Printf("有效的文件:%v\n", f.Name())
			for _, v := range types {
				if strings.HasSuffix(f.Name(), v) {
					log.Debug.Printf("有效的目标文件:%v\n", f.Name())
					//absPath := strings.Join([]string{dir, f.Name()}, "/")
					//log.Printf("目标文件的绝对路径:%v\n", absPath)
					aim = append(aim, f.Name())
				}
			}
		}
	}
	return aim
}

func command(src, dst string) {
	in := src
	out := dst
	cmd := exec.Command("ffmpeg", "-i", in, "-vn", "-y", "-acodec", "copy", out)
	//ffmpeg -i 3.mp4 -vn -y -acodec copy 3.m4a
	log.Debug.Printf("生成的命令是:%s", cmd)
	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Debug.Printf("cmd.StdoutPipe产生的错误:%v", err)
	}
	if err = cmd.Start(); err != nil {
		log.Debug.Printf("cmd.Run产生的错误:%v", err)
	}
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		t:=replace.Replace(string(tmp))
		log.Info.Println(t)
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		log.Debug.Panicln("命令执行中有错误产生", err)
	}
}
