package utils

import "C"
import (
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"unsafe"

	sc "golang.org/x/text/encoding/simplifiedchinese"
)

func CInt(i int64) C.int {
	return C.int(i)
}

func GoInt(ci C.int) int64 {
	return int64(ci)
}

func CStr2GoInt(str *C.char) int64 {
	return Str2Int(GoString(str))
}

func CBool(b bool) C.int {
	if b {
		return 1
	}
	return 0
}

func CString(str string) *C.char {
	gbstr, _ := sc.GB18030.NewEncoder().String(str)
	return C.CString(gbstr)
}

func GoString(str *C.char) string {
	utf8str, _ := sc.GB18030.NewDecoder().String(C.GoString(str))
	return utf8str
}

// utf8Togb18030编码 适配中文函数名
func GBString(str string) string {
	gbstr, _ := sc.GB18030.NewEncoder().String(str)
	return gbstr
}

func Str2ptr(s string) uintptr {
	return uintptr(unsafe.Pointer(CString(s)))
}

func Int2ptr(i int64) uintptr {
	return uintptr(i)
}

func Byte2ptr(b []byte) uintptr {
	//return uintptr(*((*int64)(unsafe.Pointer(&b))))
	return uintptr(unsafe.Pointer(&b))
}

func Bool2ptr(b bool) uintptr {
	if b == true {
		return uintptr(1)
	}
	return uintptr(0)
}

func Ptr2str(ptr uintptr) string {
	return GoString((*C.char)(unsafe.Pointer(ptr)))
}

func Ptr2bool(ptr uintptr) bool {
	i := int(ptr)
	if i == 1 {
		return true
	}
	return false
}

func Ptr2int(ptr uintptr) int {
	return int(ptr)
}

func _ptr2str(ptr uintptr) string {
	var b []byte
	for {
		sbyte := *((*byte)(unsafe.Pointer(ptr)))
		if sbyte == 0 {
			break
		}
		b = append(b, sbyte)
		ptr += 1
	}
	return string(b)
}

// GetCurrPath 获取当前程序路径
// 假设框架程序CleverQQ Air.exe运行在D:\CleverQQ, 则返回D:\CleverQQ
func GetCurrPath() (string, error) {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return path, nil
}

/**
 * 读取文件, 返回[]byte
 */
func File2Bytes(filename string) ([]byte, error) {
	// File
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// FileInfo:
	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// []byte
	data := make([]byte, stats.Size())
	_, err = file.Read(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Str2Int(str string) int64 {
	val, _ := strconv.ParseInt(str, 10, 64)
	return val
}

func Int2Str(val int64) string {
	str := strconv.FormatInt(val, 10)
	return str
}

func Float642Str(val float64, prec int) string {
	str := strconv.FormatFloat(val, 'f', prec, 64)
	return str
}

func EscapeEmoji(text string) string {
	data := []byte(text)
	ret := []byte{}
	skip := 0
	for i := range data {
		if skip > 1 {
			skip -= 1
			continue
		}
		if data[i] == byte(240) && data[i+1] == byte(159) {
			code := hex.EncodeToString(data[i : i+4])
			ret = append(ret, []byte(fmt.Sprintf("[emoji=%s]", code))...)
			skip = 4
		} else {
			ret = append(ret, data[i])
		}
	}
	return string(ret)
}

// UnescapeEmoji 将 [emoji=FFFFFFFF] 转化为 emoji的utf-8字符串
func UnescapeEmoji(text string) string {
	data := []byte(text)
	ret := []byte{}
	skip := 0
	for i := range data {
		if skip > 1 {
			skip -= 1
			continue
		}
		if i+7 < len(data) && bytes.Equal(data[i:i+7], []byte("[emoji=")) {
			end := bytes.IndexByte(data[i:], byte(93))
			if end == -1 {
				return text
			}
			code, _ := hex.DecodeString(string(data[i+7 : end+i]))
			ret = append(ret, code...)
			skip = end + 1
		} else {
			ret = append(ret, data[i])
		}
	}
	return string(ret)
}
