package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ipfs/go-cid"
	mc "github.com/multiformats/go-multicodec"
	mh "github.com/multiformats/go-multihash"
	"github.com/urfave/cli/v2"
)

func calculateCID(c *cli.Context) error {
	filePath := c.String("file")
	fcid, err := calculateCIDv1(filePath)
	if err != nil {
		fmt.Printf("计算 CID 失败: %s\n", err)
		return err
	}
	fmt.Printf("文件 %s 的 CID v1 为: %s\n", filePath, fcid)
	return nil
}

func calculateCIDv1(path string) (string, error) {
	// 检查路径是文件还是目录
	info, err := os.Stat(path)
	if err != nil {
		fmt.Printf("获取文件/目录信息失败: %s\n", err)
		return "", err
	}

	var fileData []byte
	if info.IsDir() {
		// 处理目录
		fileData, err = processDirectory(path)
		if err != nil {
			fmt.Printf("处理目录失败: %s\n", err)
			return "", err
		}
	} else {
		// 处理文件
		fileData, err = os.ReadFile(path)
		if err != nil {
			fmt.Printf("读取文件失败: %s\n", err)
			return "", err
		}
	}

	// 创建 CID v1 的前缀
	pref := cid.Prefix{
		Version:  1,
		Codec:    uint64(mc.Raw),
		MhType:   mh.SHA2_256,
		MhLength: -1,
	}

	// 计算 CID
	c, err := pref.Sum(fileData)
	if err != nil {
		fmt.Printf("生成 CID 失败: %s\n", err)
		return "", err
	}

	return c.String(), nil
}

// processDirectory 处理目录，递归读取所有文件并计算哈希
func processDirectory(dirPath string) ([]byte, error) {
	var buf bytes.Buffer

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 跳过目录本身
		if info.IsDir() {
			return nil
		}

		// 读取文件内容
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// 写入文件路径和内容到缓冲区
		buf.WriteString(path)
		buf.Write(data)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
