package ExtractAudio

import (
	"encoding/json"
	"fmt"
	"github.com/zhangyiming748/ExtractAudio/log"
	"github.com/zhangyiming748/replace"
	"os"
	"strings"
)

type Info struct {
	video string
	audio string
	title string
}
type entry struct {
	MediaType                  int    `json:"media_type"`
	HasDashAudio               bool   `json:"has_dash_audio"`
	IsCompleted                bool   `json:"is_completed"`
	TotalBytes                 int    `json:"total_bytes"`
	DownloadedBytes            int    `json:"downloaded_bytes"`
	Title                      string `json:"title"`
	TypeTag                    string `json:"type_tag"`
	Cover                      string `json:"cover"`
	VideoQuality               int    `json:"video_quality"`
	PreferedVideoQuality       int    `json:"prefered_video_quality"`
	GuessedTotalBytes          int    `json:"guessed_total_bytes"`
	TotalTimeMilli             int    `json:"total_time_milli"`
	DanmakuCount               int    `json:"danmaku_count"`
	TimeUpdateStamp            int64  `json:"time_update_stamp"`
	TimeCreateStamp            int64  `json:"time_create_stamp"`
	CanPlayInAdvance           bool   `json:"can_play_in_advance"`
	InterruptTransformTempFile bool   `json:"interrupt_transform_temp_file"`
	QualityPithyDescription    string `json:"quality_pithy_description"`
	QualitySuperscript         string `json:"quality_superscript"`
	CacheVersionCode           int    `json:"cache_version_code"`
	PreferredAudioQuality      int    `json:"preferred_audio_quality"`
	AudioQuality               int    `json:"audio_quality"`
	Avid                       int    `json:"avid"`
	Spid                       int    `json:"spid"`
	SeasionId                  int    `json:"seasion_id"`
	Bvid                       string `json:"bvid"`
	OwnerId                    int    `json:"owner_id"`
	OwnerName                  string `json:"owner_name"`
	OwnerAvatar                string `json:"owner_avatar"`
	PageData                   struct {
		Cid              int    `json:"cid"`
		Page             int    `json:"page"`
		From             string `json:"from"`
		Part             string `json:"part"`
		Link             string `json:"link"`
		RichVid          string `json:"rich_vid"`
		Vid              string `json:"vid"`
		HasAlias         bool   `json:"has_alias"`
		Weblink          string `json:"weblink"`
		Offsite          string `json:"offsite"`
		Tid              int    `json:"tid"`
		Width            int    `json:"width"`
		Height           int    `json:"height"`
		Rotate           int    `json:"rotate"`
		DownloadTitle    string `json:"download_title"`
		DownloadSubtitle string `json:"download_subtitle"`
	} `json:"page_data"`
}

func Extrace(src string) {
	var infos []Info
	head := getDir(src)
	log.Info.Printf("给定的目标文件下全部文件夹:%s\n", head)
	for _, first := range head {
		second := strings.Join([]string{src, first}, "/")
		if strings.Contains(second, "DS_store") {
			continue
		}
		log.Info.Printf("拼接第一级文件名:%s\n", second)
		entry := strings.Join([]string{second, "entry.json"}, "/")
		log.Info.Printf("拼接entry文件名:%s\n", entry)
		e := readEntry(entry)
		random := getDir(second)[0]
		audio := strings.Join([]string{second, random, "audio.m4s"}, "/")
		log.Info.Printf("拼接audio文件名:%s\n", audio)
		video := strings.Join([]string{second, random, "video.m4s"}, "/")
		log.Info.Printf("拼接video文件名:%s\n", video)
		info := &Info{
			video: video,
			audio: audio,
			title: e.PageData.Part,
		}
		infos = append(infos, *info)

	}
	for _, value := range infos {
		log.Info.Printf("%+v\n", value)
	}
	extrace_help(infos)
}

func extrace_help(infos []Info) {
	for _, info := range infos {
		oldName := info.audio
		newName := strings.Join([]string{info.title, "mp3"}, ".")
		log.Debug.Printf("文件:%v重命名为:%v\n", oldName, newName)
		os.Rename(oldName, newName)
	}
}

func getDir(pwd string) (partname []string) {
	//获取文件或目录相关信息
	fileInfoList, err := os.ReadDir(pwd)
	if err != nil {
		log.Debug.Panicln(err)
	}
	for i := range fileInfoList {
		partname = append(partname, fileInfoList[i].Name())
	}
	return partname
}

func readEntry(dir string) (e entry) {
	bytes, err := os.ReadFile(dir)
	if err != nil {
		fmt.Println("读取json文件失败", err)
		return
	}

	err = json.Unmarshal(bytes, &e)
	if err != nil {
		fmt.Println("解析数据失败", err)
		return
	}
	log.Info.Printf("获取到的partname:%s\n", e.PageData.Part)
	log.Info.Printf("获取到的title:%s\n", e.Title)
	e.PageData.Part = replace.Replace(e.PageData.Part)
	e.Title = replace.Replace(e.Title)
	log.Info.Printf("替换后的partname:%s\n", e.PageData.Part)
	log.Info.Printf("替换后的title:%s\n", e.Title)
	return e
}
