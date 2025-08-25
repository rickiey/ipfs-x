package cmd

import (
	"fmt"
	"os"

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

func calculateCIDv1(filePath string) (string, error) {
	// 读取文件内容
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("读取文件失败: %s\n", err)
		return "", err
	}

	// 创建 CID v1 的前缀
	pref := cid.Prefix{
		Version:  1,              // CID v1
		Codec:    uint64(mc.Raw), // 使用 Raw 编码（可根据需求改为 DagProtobuf 等）
		MhType:   mh.SHA2_256,    // 使用 SHA2-256 哈希
		MhLength: -1,             // 默认长度
	}

	// 计算文件的 CID
	c, err := pref.Sum(fileData)
	if err != nil {
		fmt.Printf("生成 CID 失败: %s\n", err)
		return "", err
	}

	// 输出 CID
	fmt.Println("CID v1:", c.String())
	return c.String(), nil
}
