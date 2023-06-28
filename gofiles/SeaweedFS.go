package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/chrislusf/seaweedfs/weed"
)

const (
	masterURL = "http://localhost:9333"
)

func main() {
	// 创建SeaweedFS客户端
	client, err := weed.NewClient(masterURL)
	if err != nil {
		log.Fatal(err)
	}

	// 创建文件
	filePath := "example.txt"
	if err := createFile(client, filePath); err != nil {
		log.Fatal(err)
	}
	fmt.Println("文件创建成功！")

	// 读取文件内容
	content, err := readFile(client, filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("文件内容：", content)

	// 更新文件内容
	newContent := "这是更新后的内容"
	if err := updateFile(client, filePath, []byte(newContent)); err != nil {
		log.Fatal(err)
	}
	fmt.Println("文件更新成功！")

	// 读取更新后的文件内容
	content, err = readFile(client, filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("文件更新后的内容：", content)

	// 删除文件
	if err := deleteFile(client, filePath); err != nil {
		log.Fatal(err)
	}
	fmt.Println("文件删除成功！")
}

// 创建文件
func createFile(client *weed.Seaweed, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	_, fileId, err := client.SubmitFile(fileBytes, filepath.Base(filePath))
	if err != nil {
		return err
	}

	// 可选：将返回的fileId保存到数据库或其他地方

	return nil
}

// 读取文件内容
func readFile(client *weed.Seaweed, filePath string) (string, error) {
	fileId, err := getFileIdFromDB(filePath)
	if err != nil {
		return "", err
	}

	fileUrl := client.LookupFileId(fileId).Url
	response, err := client.Get(fileUrl)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	fileBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(fileBytes), nil
}

// 更新文件内容
func updateFile(client *weed.Seaweed, filePath string, newContent []byte) error {
	fileId, err := getFileIdFromDB(filePath)
	if err != nil {
		return err
	}

	fileUrl := client.LookupFileId(fileId).Url
	_, err = client.Upload(fileUrl, newContent, fileId, false)
	if err != nil {
		return err
	}

	return nil
}

// 删除文件
func deleteFile(client *weed.Seaweed, filePath string) error {
	fileId, err := getFileIdFromDB(filePath)
	if err != nil {
		return err
	}
	err = client.Delete(fileId)
	if err != nil {
		return err
	}
	// 可选：从数据库或其他地方删除fileId
	return nil
}

// 从数据库或其他地方获取文件的fileId
func getFileIdFromDB(filePath string) (string, error) {
	// 假设你已经实现了从数据库或其他地方获取fileId的逻辑
	// 这里只是一个示例，你可以根据实际情况进行实现
	fileId := "your-file-id"
	return fileId, nil
}
