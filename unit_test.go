package ExtractAudio

import "testing"

func TestExtrace(t *testing.T) {
	src:="/Users/zen/Github/ExtractAudio/example/钢琴"
	dst:="/Users/zen/Github/ExtractAudio/example/钢琴"
	Extrace(src,dst)
}
func TestGetFiles(t *testing.T) {
	src:="/Users/zen/Github/ExtractAudio/example/mp4"
	getFiles(src,"mp4")
}
func TestMp4(t *testing.T) {
	src:="/Users/zen/Github/ExtractAudio/example/mp4"
	Mp4(src)
}
