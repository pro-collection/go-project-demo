package demos

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type UserGroup struct {
	Name       string `json:"name"`
	Server     string `json:"server"`
	Port       int    `json:"port"`
	Password   string `json:"password"`
	Cipher     string `json:"cipher"`
	Key        string `json:"key"`
	Keygen     int    `json:"key_gen"`
	UDPTimeout int    `json:"time_out"`
}

type Config struct {
	UserGroups []*UserGroup `json:"user_groups"`
	UDPTimeout int          `json:"time_out"`
}

/*
 * @msg: 读取json 文件数据，按照参数 v 的格式解析
 * @param: fileName
 * @return:
 */
func LoadJsonData(fileName string, v interface{}) (*Config, error) {
	// 重新获取绝对路径
	filePathString, _ := filepath.Abs(fileName)

	// 打开文件
	file, err := os.OpenFile(filePathString, os.O_RDONLY, 0644)

	var MConfig = &Config{}

	if err != nil {
		if err != nil {
			fmt.Println("Error opening file:", err)
			return MConfig, err
		}
	}

	defer file.Close()

	// 读取文件为 []byte
	dataJson, err := io.ReadAll(file)

	json.Unmarshal(dataJson, MConfig)

	return MConfig, err
}
