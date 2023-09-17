package handleFile

import (
	"encoding/json"
	"fmt"
	"go-project-demo/packages/pro3_proxypool/pkg/getter"
	"go-project-demo/packages/pro3_proxypool/pkg/models"
	"os"
	"sync"
)

func FindFile(filename string) (file *os.File, err error) {
	if _, err = os.Stat(filename); os.IsNotExist(err) {
		// 文件不存在，创建文件
		file, err = os.Create(filename)
		if err != nil {
			fmt.Println("创建文件失败:", err)
			return nil, err
		}

		fmt.Println("文件创建成功")
	} else {
		// 文件存在，直接使用
		// 读写且清空源文件
		file, err = os.OpenFile(filename, os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("打开文件失败:", err)
			return nil, err
		}

		fmt.Println("文件已存在，可以直接使用")
	}

	return
}

// WriteFileWithNetWork .
// 写入文件
func WriteFileWithNetWork(file *os.File) {
	var ipList []*models.IP

	var functionList = []GetIPFunction{
		getter.IP89,
		//getter.KDL,
		//getter.PLPSSL,
		getter.IP3306,
		//getter.PZZQZ,
	}

	var wg sync.WaitGroup
	for _, function := range functionList {
		wg.Add(1)
		go func(f GetIPFunction) {
			temp := f()

			for _, ip := range temp {
				ipList = append(ipList, ip)
			}

			wg.Done()
		}(function)

	}

	wg.Wait()

	jsonData, _ := json.Marshal(ipList)

	// 文件阶段
	err := file.Truncate(0)

	// 光标移动
	_, err = file.Seek(0, 0)

	_, err = file.WriteString(string(jsonData) + "\n")
	if err != nil {
		fmt.Println("写入文件失败:", err)
		return
	}
}

func WriteToLocal(writePath string, ipList *[]*models.IP) {
	file, err := FindFile(writePath)
	if err != nil {
		return
	}

	jsonData, _ := json.Marshal(ipList)

	err = file.Truncate(0)
	_, err = file.Seek(0, 0)

	_, err = file.WriteString(string(jsonData) + "\n")
	if err != nil {
		fmt.Println("写入文件失败:", err)
		return
	}
}

// ReadFile
// 读取文件， 返回 []byte
func ReadFile(file *os.File) ([]byte, error) {

	// 获取文件大小
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("获取文件信息失败:", err)
		return nil, err
	}
	fileSize := fileInfo.Size()

	buffer := make([]byte, fileSize)
	// 读取文件内容
	_, err = file.Read(buffer)

	if err != nil {
		fmt.Println("读取文件失败:", err)
		return nil, err
	}

	return buffer, nil
}

// FilterGetUsedIpList
// 过滤 ip , 只留下可使用的 ip 地址信息
func FilterGetUsedIpList(fileContent []byte) []*models.IP {
	var ipList []*models.IP
	err := json.Unmarshal(fileContent, &ipList)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}

	ipChan := make(chan *models.IP, len(ipList))

	// 校验 ip 是否好用
	fmt.Printf("长度: %d \n", len(ipList))

	var wg sync.WaitGroup

	for _, ip := range ipList {
		wg.Add(1)

		go getUsedIP(ip, ipChan, &wg)
	}

	wg.Wait()
	close(ipChan)

	var list []*models.IP

	for ip := range ipChan {
		list = append(list, ip)
	}

	return list
}

func GetIP() {
	file, err := FindFile("source_ip.json")
	defer file.Close()
	if err != nil {
		return
	}

	// 网络获取 ip 地址核心方法
	WriteFileWithNetWork(file)

	// 重新读取一遍文档
	file, err = FindFile("source_ip.json")
	fileContent, err := ReadFile(file)

	list := FilterGetUsedIpList(fileContent)
	fmt.Println("ip 处理结束")
	for _, ip := range list {
		fmt.Println("ip: ", ip)
	}

	//重新写入文件
	WriteToLocal("ip.json", &list)
}
