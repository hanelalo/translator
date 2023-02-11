package main

import (
	"flag"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alimt"
	"github.com/goccy/go-yaml"
	"os"
	"strings"
)

func main() {
	source := flag.String("source", "en", "源语言, 详细语言代码请看阿里云文档 https://help.aliyun.com/document_detail/215387.html?spm=a2c4g.11186623.0.0.78382e50NX8PG3")
	target := flag.String("target", "zh", "目标语言, 详细语言代码请看阿里云文档 https://help.aliyun.com/document_detail/215387.html?spm=a2c4g.11186623.0.0.78382e50NX8PG3")
	scene := flag.String("scene", "general", "翻译场景，详见阿里云文档")
	configFile := flag.String("config", "/opt/translator/config.yaml", "配置文件")
	flag.Parse()
	configFileLocation := getConfigFileLocation(*configFile)
	config, err := parseConfig(configFileLocation)
	if err != nil {
		return
	}
	content := flag.Arg(0)
	if content == "" {
		fmt.Println("请输入要翻译的内容")
		return
	}
	content = strings.ReplaceAll(content, "\n", "")
	content = strings.ReplaceAll(content, "\r", "")
	// 创建ecsClient实例
	alimtClient, err := alimt.NewClientWithAccessKey(
		config.RegionId,        // 地域ID
		config.AccessKeyId,     // 您的Access Key ID
		config.AccessKeySecret) // 您的Access Key Secret
	if err != nil {
		// 异常处理
		panic(err)
	}
	// 创建API请求并设置参数
	request := alimt.CreateTranslateECommerceRequest()
	// 等价于 request.PageSize = "10"
	request.Method = "POST"          //设置请求
	request.FormatType = "text"      //翻译文本的格式
	request.SourceLanguage = *source //源语言
	request.SourceText = content     //原文
	request.TargetLanguage = *target //目标语言
	request.Scene = *scene           //目标语言
	// 发起请求并处理异常
	response, err := alimtClient.TranslateECommerce(request)
	if err != nil {
		// 异常处理
		panic(err)
	}
	fmt.Println(response.Data.Translated)
}

func getConfigFileLocation(locationFromCommand string) string {
	locationFromEnv := os.Getenv("TRANSLATOR_CONFIG_LOCATION")
	if locationFromEnv == "" {
		return locationFromCommand
	}
	return locationFromEnv
}

func parseConfig(configFile string) (ClientConfig, error) {
	config := ClientConfig{}
	yml, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Printf("配置文件读取异常, %s", err)
		return config, err
	}
	err = yaml.Unmarshal(yml, &config)
	if err != nil {
		fmt.Printf("配置文件解析异常, %s", err)
		return config, err
	}
	return config, nil
}

type ClientConfig struct {
	RegionId        string `yaml:"regionId"`
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
}
