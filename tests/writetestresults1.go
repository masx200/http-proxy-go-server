package tests

import (
	"bufio"
	"io"
	"os"
)

// WriteTestResultsToFile 写入测试结果到文件
func WriteTestResultsToFile(results []string, file *os.File) error {
	// 写入到测试记录.log

	var err error
	// 移动到文件末尾
	_, err = file.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)

	// 写入分隔符
	_, err = writer.WriteString("\n\n###\n\n")
	if err != nil {
		return err
	}

	// 写入测试结果
	for _, line := range results {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
