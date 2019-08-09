package slasterisk

import (
	"io"
	"os"

	"github.com/sunicy/go-lame"
)

func CheckFilePath(filePath string) (exists bool, err error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return MakeNeededDirs(filePath)
	}
	return true, err
	//os.MkdirAll
}

func MakeNeededDirs(path string) (exists bool, err error) {
	if err := os.MkdirAll(path, 0755); err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func WavToMp3(wavFile string, mp3File string) {
	// open files
	wavFile, _ := os.OpenFile(wavFileName, os.O_RDONLY, 0555)
	mp3File, _ := os.OpenFile(mp3FileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	defer mp3File.Close()

	// parsing wav info
	// NOTE: reader position moves even if it is not a wav file
	wavHdr, err := lame.ReadWavHeader(wavFile)
	if err != nil {
		panic("not a wav file, err=" + err.Error())
	}

	wr, _ := lame.NewWriter(mp3File)
	wr.EncodeOptions = wavHdr.ToEncodeOptions()
	io.Copy(wr, wavFile) // wavFile's pos has been changed!
	wr.Close()
}
