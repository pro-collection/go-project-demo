package cmd

import (
	"github.com/spf13/cobra"
	"go-project-demo/src/pro1/internal/word"
	"log"
	"strings"
)

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscoreToUpperCamelCase
	ModeUnderscoreToLowerCamelCase
	ModeCamelCaseToUnderscore
)

var str string
var mode int8

var desc = strings.Join([]string{
	"支持的各种但是转换格式如下： ",
	"1： 全部单词转为大写",
	"2： 全部单词转为小写",
	"3： 下划线单词转大写驼峰单词",
	"4： 下划线单词转小写驼峰单词（首单词字母小写）",
	"5： 驼峰单词转下划线单词",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
			break
		case ModeLower:
			content = word.ToLower(str)
			break
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
			break
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
			break
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
			break
		default:
			log.Fatalln("站不支持该格式转换， 请执行 help word 查看帮助文档")
			break
		}

		log.Printf("输出结果：%s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}
