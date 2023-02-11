# Translator

接入阿里云机器翻译的命令行工具。

## 构建
```bash
& git clone https://github.com/hanelalo/translator.git
& cd translator
& go mod vendor
& go build
```
> 也可以通过 `go build -o <name>` 指定可执行文件的名称。
## 命令
```
  -config string
    	配置文件 (default "/opt/translator/config.yaml")
  -scene string
    	翻译场景，详见阿里云文档 (default "general")
  -source string
    	源语言, 详细语言代码请看阿里云文档 https://help.aliyun.com/document_detail/215387.html?spm=a2c4g.11186623.0.0.78382e50NX8PG3 (default "en")
  -target string
    	目标语言, 详细语言代码请看阿里云文档 https://help.aliyun.com/document_detail/215387.html?spm=a2c4g.11186623.0.0.78382e50NX8PG3 (default "zh")
```
## 配置文件
配置文件地址，也可以通过 `TRANSLATOR_CONFIG_LOCATION` 环境变量进行配置，如果配置了该环境变量，将忽略 `-config` 参数。

```yaml
regionId: cn-hangzhou
accessKeyId: <accessKeyId>
accessKeySecret: <accessKeySecret>
```
* regionId 
  
  区域 id，可以不修改。

* accessKeyId 
  
  阿里云访问授权 id。

* accessKeySecret 
  
  阿里云访问授权 secret。