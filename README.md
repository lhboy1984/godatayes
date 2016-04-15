# GoDataYes

GoDataYes是使用go语言实现的访问通联数据的工具包

使用统一接口GetData(api-argument)

相关的参数和返回值可参照api.go，该文件通过抓取通联文档直接生成
# Installation
go get https://github.com/waditu/tushare.git
# Dependencies
"github.com/PuerkitoBio/goquery"

"golang.org/x/net/html"
#Quick Start
fmt.Println(godatayes.GetData(godatayes.NewArgument("getSecID", map[string]interface{}{
		"ticker": "000001",
	})))