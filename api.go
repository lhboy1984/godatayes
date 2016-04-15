package godatayes

var apis map[string]map[string]interface{}

func init() {
	apis = map[string]map[string]interface{}{
		//宏观行业
		//特色大数据
		//说明包含主题活跃周期数据。输入主题代码或名称，获取主题的活跃时间等信息，同时可输入是否最新活跃期，获取主题最新的活跃周期。(注：1、主题活跃周期数据自2013/1/1始；2、新闻量在某段时间内达到活跃阈值的主题即为活跃主题；3、数据按日更新。)
		"getThemesPeriod": map[string]interface{}{
			"url": "/api/subject/getThemesPeriod.json",
			"args": map[string]string{
				"isLatest":  "int16",  //是否最新活跃期，1为最新，0为非最新
				"themeID":   "int64",  //主题ID，允许多值输入，为空则输出所有
				"themeName": "string", //主题名称，支持模糊查询
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"themeID":    "int64",  //主题ID
				"beginDate":  "date",   //主题开始日期
				"themeName":  "string", //主题名称
				"endDate":    "date",   //主题结束日期
				"isLatest":   "int16",  //是否最新活跃期，1为最新，0为非最新
				"insertTime": "date",   //入库时间
				"updateTime": "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明包含雪球社交统计数据，输入一个或多个证券交易代码、统计起止日期，获取该证券一段时间内每天的雪球帖子数量、帖子占比(%)。(注：数据自2014/1/1始，按日更新。)
		"getSocialDataXQ": map[string]interface{}{
			"url": "/api/subject/getSocialDataXQ.json",
			"args": map[string]string{
				"beginDate": "date",   //统计开始日期，默认为30天前
				"endDate":   "date",   //统计结束日期，默认为当天
				"ticker":    "string", //证券交易代码，可多值输入
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"ticker":         "string", //证券交易代码
				"statisticsDate": "date",   //统计日期
				"postNum":        "int64",  //雪球帖子数量
				"postPercent":    "double", //雪球帖子占比(%)
				"insertTime":     "date",   //入库时间
				"updateTime":     "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明包含证券相关的新闻情感指数数据，输入一个或多个证券交易代码、起止日期，获取该证券一段时间内的新闻情感指数(即当天证券关联新闻的情感均值)。（注：1、2014/1/1起新闻来源众多、指数统计有效，2013年及之前的网站来源不全、数据波动大，数据自2004/10/28始；2、新闻量的统计口径为经算法处理后证券关联到的所有常规新闻；3、数据按日更新。)
		"getNewsSentimentIndex": map[string]interface{}{
			"url": "/api/subject/getNewsSentimentIndex.json",
			"args": map[string]string{
				"beginDate":    "date",   //查询开始日期，默认为一个月前
				"endDate":      "date",   //查询结束日期，默认为当天
				"secID":        "string", //证券内部编码，可通过交易代码和交易市场在getSecID获取到，可多值输入，如："000001.XSHE"、"000001.XSHE,600001.XSHG"
				"exchangeCD":   "string", //证券交易所代码(通联自编)。可选：XSHG、XSHE、XHKG。XSHG表示上海证券交易所，XSHE表示深圳证券交易所，XHKG表示香港交易所。可多值输入
				"secShortName": "string", //证券简称，支持模糊查询
				"ticker":       "string", //证券交易代码，可多值输入，如："000001"、"000001,600001"
				"field":        "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":           "string", //证券内部编码(通联自编)
				"exchangeCD":      "string", //证券交易所代码(通联自编)
				"ticker":          "string", //证券交易代码
				"newsPublishDate": "date",   //新闻发布日期
				"exchangeName":    "string", //证券交易所
				"secShortName":    "string", //证券简称
				"sentimentIndex":  "double", //新闻情感指数，证券当天关联新闻的总体看法(当天的证券新闻情感均值)，正数表示看涨、负数表示看跌、0为中性，绝对值越高情感越强烈
				"insertTime":      "date",   //入库时间
				"updateTime":      "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明获取一段时间内主题删除的关联证券数据，输入主题代码或名称、查询起止时间，可以获取该时间段内主题删除的关联证券信息，包括证券代码、证券简称、证券交易场所，同时返回关联开始时间、关联结束时间、关联具体描述、数据入库及更新时间。(注：1、主题与证券的关联自2013/12/28始、2014年12月起关联数据完整；2、数据按日更新。)
		"getThemesTickersDelete": map[string]interface{}{
			"url": "/api/subject/getThemesTickersDelete.json",
			"args": map[string]string{
				"themeID":   "int64",  //主题ID，允许多值输入，如:"1,2,3,4"
				"themeName": "string", //主题名称，支持模糊查询
				"beginDate": "date",   //查询开始日期，默认为当天
				"endDate":   "date",   //查询结束日期，默认为当天
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"themeID":      "int64",  //主题ID
				"themeName":    "string", //主题名称
				"secID":        "string", //证券内部编码(通联自编)
				"exchangeCD":   "string", //证券交易场所代码(通联自编)
				"ticker":       "string", //证券交易代码
				"beginDate":    "date",   //关联开始日期
				"secShortName": "string", //证券简称
				"exchangeName": "string", //证券交易所
				"endDate":      "date",   //关联结束日期，日期为空表示关联至今依旧存在
				"insertTime":   "date",   //入库时间
				"updateTime":   "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明包含主题的热度数据。输入主题代码或名称、同时可输入起止日期，获取一段时间内主题每天的新闻数量、主题热度(即主题每天新闻数量占当日所有主题新闻总量的百分比(%))。(注：数据自2014/1/1始，每天更新)
		"getThemesHeat": map[string]interface{}{
			"url": "/api/subject/getThemesHeat.json",
			"args": map[string]string{
				"themeID":   "int64",  //主题ID，允许多值输入，如:"1,2,3,4"
				"themeName": "string", //主题名称，支持模糊查询
				"beginDate": "date",   //查询开始日期，默认为30天前
				"endDate":   "date",   //查询结束日期，默认为当天
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"themeID":        "int64",  //主题ID
				"statisticsDate": "date",   //统计日期
				"themeName":      "string", //主题名称
				"newsNum":        "double", //当天主题新闻数量，即出现主题关键词的新闻数量
				"newsNumPercent": "double", //当天主题新闻占比(%)，当天该主题关联的新闻数量占所有主题新闻总数的百分比(%)
				"insertTime":     "date",   //入库时间
				"updateTime":     "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明包含证券关联的主题数据，主题源自申万行业。输入证券交易所代码、证券交易代码或简称，可以获取关联的主题等信息，包括证券代码、证券简称、证券交易场所，同时返回三个维度的关联分数、关联开始时间、关联结束时间、关联具体描述、数据入库及更新时间，同时可输入查询起止时间，以获取证券在该时间段内关联到的主题信息。(注：1、源自行业的主题与证券的关联自2014/12/26始；2、数据按日更新、同时刷新关联状态。)
		"getSectorThemesByTickers": map[string]interface{}{
			"url": "/api/subject/getSectorThemesByTickers.json",
			"args": map[string]string{
				"secID":        "string", //证券内部编码，可通过交易代码和交易市场在getSecID获取到，可多值输入，如："000001.XSHE"、"000001.XSHE,600001.XSHG"
				"secShortName": "string", //证券简称，支持模糊查询，如"浦发"、"浦发银行"
				"ticker":       "string", //证券交易代码，允许多值输入，用","分隔，如"000001"、"000001,600001"。
				"beginDate":    "date",   //查询开始日期，默认为当天
				"endDate":      "date",   //查询结束日期，默认为当天
				"exchangeCD":   "string", //证券交易所代码(通联自编)。可选：XSHG、XSHE、XHKG、XNYS、NOBB、XNAS。XSHG表示上海证券交易所，XSHE表示深圳证券交易所，XHKG表示香港交易所，XNYS表示纽约证券交易所，NOBB表示三板市场，XNYS表示美国纳斯达克证券交易所。可多值输入
				"field":        "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":                 "string", //证券内部编码(通联自编)
				"exchangeCD":            "string", //证券交易所代码(通联自编)
				"ticker":                "string", //证券交易代码
				"beginDate":             "date",   //关联开始日期
				"themeID":               "int64",  //主题ID
				"exchangeName":          "string", //证券交易所
				"secShortName":          "string", //证券简称
				"themeName":             "string", //主题名称
				"endDate":               "date",   //关联结束日期，日期为空表示关联至今依旧存在
				"returnScore":           "double", //收益关联程度，主题和证券在短期收益上的相似度，取值范围[0，1]，值越大表示关联度越高
				"textContributionScore": "double", //文本贡献关联度，主题和证券在新闻文本中的相似度，取值范围[0，1]，值越大表示关联度越高
				"industryScore":         "double", //行业关联度，主题和证券在行业分布上的相似度的关联程度，取值范围[0，1]，值越大表示关联度越高
				"insertTime":            "date",   //入库时间
				"updateTime":            "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明获取一段时间内新增(开始)的活跃主题数据，输入的时间参数在主题活跃周期的起始时间列进行查询。输入查询起止时间、是否最新活跃期、主题来源，可以获取该时间段内开始活跃的主题信息，包括主题ID、主题名称、主题开始时间、主题结束时间、是否最新活跃期、数据入库及更新时间。(注：1、主题活跃周期数据自2013/1/1始；2、数据按日更新。)
		"getActiveThemesInsert": map[string]interface{}{
			"url": "/api/subject/getActiveThemesInsert.json",
			"args": map[string]string{
				"beginDate":   "date",   //查询开始日期，默认为当天
				"endDate":     "date",   //查询结束日期，默认为当天
				"isLatest":    "Int16",  //是否最新活跃期，1为最新、输入该值表示主题最新活跃期的开始时间在所查询时间段内，0为非最新，默认为1
				"themeSource": "Int16",  //主题来源，1为源自网络主题，2为自动挖掘生成的主题，3为通联人工自编主题，5为申万行业主题
				"field":       "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"themeID":     "int64",  //主题ID
				"themeName":   "string", //主题名称
				"themeSource": "Int16",  //主题来源
				"beginDate":   "date",   //主题开始日期
				"endDate":     "date",   //主题结束日期，结束日期为当天表示主题当天依旧活跃(第二天会重新刷新结束时间)
				"isLatest":    "int16",  //是否最新活跃期，1为最新，0为非最新
				"insertTime":  "date",   //入库时间
				"updateTime":  "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明获取新闻发布来源列表，输入参数为空、直接获取通联新闻数据的爬取来源列表。
		"getNewsPublishSite": map[string]interface{}{
			"url": "/api/subject/getNewsPublishSite.json",
			"args": map[string]string{
				"field": "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"newsPublishSite": "string", //新闻发布来源，即新闻的实际最终来源
			},
			"indices": map[string]string{},
		},
		//说明包含新闻关联的公司数据，同时可获取针对不同公司的新闻情感数据。输入新闻ID，获取相关的公司信息，如：公司代码、公司全称，同时返回新闻标题、发布时间、入库时间信息。其中，公司代码可继续通过证券编码及基本上市信息(getSecID)查找公司相关的证券。(注：1、自2014/1/1起新闻来源众多、新闻量日均4万左右，2013年及之前的网站来源少、新闻数据量少；2、数据实时更新。)
		"getCompanyByNews": map[string]interface{}{
			"url": "/api/subject/getCompanyByNews.json",
			"args": map[string]string{
				"newsID": "int64",  //新闻ID，可多值输入，可由getNewsInfoByTime(获取一天某段时间内的新闻信息)这个API中获取新闻ID
				"field":  "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"newsPublishTime": "date",   //新闻发布时间
				"partyID":         "int64",  //公司代码(通联自编)，可通过证券编码及基本上市信息(getSecID)继续查找公司相关的证券
				"newsID":          "int64",  //新闻ID
				"newsTitle":       "string", //新闻标题
				"partyFullName":   "string", //公司全称
				"relatedScore":    "double", //，
				"sentiment":       "int16",  //公司新闻情感，新闻针对特定公司的情感分类，1为正面、0为中性、-1为负面
				"sentimentScore":  "double", //，利用人工准备的训练数据构建情感识别模型
				"newsInsertTime":  "date",   //新闻入库时间
			},
			"indices": map[string]string{},
		},
		//说明包含新闻关联的证券数据，同时可获取针对不同证券的新闻情感数据。输入新闻ID，获取相关的证券信息，如：证券代码、证券简称、证券交易场所，同时返回新闻标题、发布来源、发布时间、入库时间等新闻相关信息。每天更新。(注：1、自2014/1/1起新闻来源众多、新闻量日均4万左右，2013年及之前的网站来源少、新闻数据量少；2、数据实时更新。)
		"getTickersByNews": map[string]interface{}{
			"url": "/api/subject/getTickersByNews.json",
			"args": map[string]string{
				"newsID": "int64",  //新闻ID，可多值输入，可由getNewsInfoByTime(获取一天某段时间内的新闻信息)这个API中获取新闻ID
				"field":  "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"newsPublishTime": "date",   //新闻发布时间
				"secID":           "string", //证券内部编码(通联自编)
				"exchangeCD":      "string", //证券交易所代码
				"ticker":          "string", //证券交易代码
				"newsID":          "int64",  //新闻ID
				"newsTitle":       "string", //新闻标题
				"secShortName":    "string", //证券简称
				"exchangeName":    "string", //证券交易所
				"relatedScore":    "double", //该接口只展示显著关联结果
				"sentiment":       "int16",  //证券所属公司新闻情感，新闻针对特定证券所属公司的情感分类，1为正面、0为中性、-1为负面
				"sentimentScore":  "double", //，利用人工准备的训练数据构建情感识别模型
				"newsPublishSite": "string", //新闻发布来源，即新闻的实际最终来源
				"newsInsertTime":  "date",   //新闻入库时间
			},
			"indices": map[string]string{},
		},
		//说明包含新闻全文等信息。输入新闻ID，获取新闻全文相关字段，如：新闻ID、标题、摘要、正文、来源链接、初始来源、作者、发布来源、发布时间、入库时间等。(注：1、自2014/1/1起新闻来源众多、新闻量日均4万左右，2013年及之前的网站来源少、新闻数据量少；2、数据实时更新。)
		"getNewsContent": map[string]interface{}{
			"url": "/api/subject/getNewsContent.json",
			"args": map[string]string{
				"newsID": "int64",  //新闻ID，可多值输入
				"field":  "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"newsPublishTime":  "date",   //新闻发布时间
				"newsID":           "int64",  //新闻ID
				"newsTitle":        "text",   //新闻标题
				"newsSummary":      "text",   //(源自算法处理结果)
				"newsBody":         "text",   //新闻正文
				"newsURL":          "string", //新闻链接
				"newsOriginSource": "string", //新闻初始来源，即新闻原始出处
				"newsAuthor":       "text",   //新闻作者
				"newsPublishSite":  "string", //新闻发布来源，即新闻的实际最终来源
				"newsInsertTime":   "date",   //新闻入库时间
			},
			"indices": map[string]string{},
		},
		//说明根据公告分类获取相应公告信息，输入一个或多个公告分类，可以获取所查询证券相关的公告信息，包括公告ID、公告名称、证券交易场所、证券交易所对公告的原始分类、公告发布时间、公告所属分类、公告分类入库时间、更新时间。(注：公告分类数据自2009/1/5始，按日更新)
		"getReportByCategory": map[string]interface{}{
			"url": "/api/subject/getReportByCategory.json",
			"args": map[string]string{
				"Category":  "string", //公告自动分类，分类包括：季度报告-业绩报告、中期报告-业绩报告、年度报告-业绩报告、业绩预报-业绩报告、业绩快报-业绩报告、招股说明书-首发公告、发行公告-首发公告、路演公告-首发公告、招股意向书-首发公告、定价公告-首发公告、受限股上市-股份变动、配股上市-股份变动、公积金转增-股份变动、持股变动-股权变动、股权变动-股权变动、股权收购-股权变动、要约收购-股权变动、吸收合并-股权变动、减资公告-股权变动、增发预案-增发公告、增发获准-增发公告、增发招股意向书-增发公告、发行公告-增发公告、上市公告-增发公告、激励计划-股权激励、对象名单-股权激励、进展公告-股权激励、配股预案-配股公告、配股说明书-配股公告、配股提示公告-配股公告、配股获准公告-配股公告、其他配股事项-配股公告、风险警示-上市资格、暂停上市-上市资格、恢复上市-上市资格、终止上市-上市资格、关联交易、资产重组、诉讼仲裁、高管变动、债券相关、政府补贴、违纪违规、担保事项、项目投资、股价异动、重大合同、重大损失、其他公告。分类单值输入。
				"beginDate": "date",   //查询开始日期，默认为30天前
				"endDate":   "date",   //查询结束日期，默认为当天
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"reportID":    "int64",  //公告代码
				"publishDate": "date",   //公告发布时间
				"ticker":      "string", //证券交易代码
				"year":        "int32",  //公告年份
				"category":    "string", //公告分类结果
				"title":       "string", //公告标题
				"site":        "string", //证券交易场所，包括"SH"、"SZ"。"SH"指上海证券交易所、"SZ"指深圳证券交易所
				"reportType":  "string", //证券交易所对公告的原始分类
				"insertTime":  "date",   //自动分类数据入库时间
				"updateTime":  "date",   //自动分类数据更新时间
			},
			"indices": map[string]string{},
		},
		//说明根据公告ID获取公告原始内容数据，输入公告ID，获取公告原文等信息，包括公告ID、公告名称、证券交易场所、证券交易所对公告的原始分类、公告发布时间、公告具体内容、公告链接、公告入库时间。(注：公告数据自2000/1/8始，按日更新)
		"getReportContentByID": map[string]interface{}{
			"url": "/api/subject/getReportContentByID.json",
			"args": map[string]string{
				"reportID": "int64",  //公告代码
				"field":    "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"publishDate": "date",   //公告发布时间
				"reportID":    "int64",  //公告代码
				"year":        "int32",  //公告年份
				"tiker":       "string", //证券交易代码
				"title":       "string", //公告标题
				"site":        "string", //证券交易场所，包括"SH"、"SZ"。"SH"指上海证券交易所、"SZ"指深圳证券交易所
				"reportType":  "string", //证券交易所对公告的原始分类
				"txtContent":  "text",   //公告内容
				"URL":         "string", //公告原文链接
				"insertTime":  "date",   //公告入库时间
			},
			"indices": map[string]string{},
		},
		//说明包含所有主题基本信息。输入主题代码或名称、主题来源，可以获取主题相关信息，包括主题ID、主题名称、主题描述、主题来源、当天是否活跃、主题插入时间、主题更新时间等。(注：1、主题基期自2011/4/16始；2、数据按日更新主题活跃状态。)
		"getThemesContent": map[string]interface{}{
			"url": "/api/subject/getThemesContent.json",
			"args": map[string]string{
				"isMain":      "int16",  //是否聚类后的主题，对查询当天的活跃主题进行合并，1为合并后的主要主题，0为非主要主题，输入1可查询当天合并后的活跃主题(主题相似度达到阈值归为一类)，每天刷新。
				"themeID":     "int64",  //主题ID，允许多值输入，为空则输出所有
				"themeName":   "string", //主题名称，支持模糊查询
				"themeSource": "int16",  //源自网络主题：从云财经、新浪财经、东方财富等网站爬取得到的。
				"field":       "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"themeID":          "int64",  //主题ID
				"themeName":        "string", //主题名称
				"themeDescription": "string", //主题描述
				"themeSource":      "int16",  //主题来源
				"themeBaseDate":    "date",   //主题基期，计算主题收益等指数时的开始时间
				"isActive":         "int16",  //当天是否活跃，1为活跃，0为不活跃，每天刷新
				"isMain":           "int16",  //是否聚类后的主题
				"insertTime":       "date",   //入库时间
				"updateTime":       "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明包含主题关联的证券数据。输入主题代码或名称，可以获取主题关联的证券等信息，包括证券代码、证券简称、证券交易场所，同时返回三个维度的关联分数、关联开始时间、关联结束时间、关联具体描述、数据入库及更新时间，同时可输入查询起止时间，以获取主题在该时间段内关联的证券信息。(注：1、主题与证券的关联自2013/12/28始、2014年12月起关联数据完整；2、数据按日更新、同时刷新关联状态。)
		"getTickersByThemes": map[string]interface{}{
			"url": "/api/subject/getTickersByThemes.json",
			"args": map[string]string{
				"themeID":   "int64",  //主题ID，允许多值输入，如:"1,2,3,4"
				"themeName": "string", //主题名称，支持模糊查询
				"beginDate": "date",   //查询开始日期，默认为当天
				"endDate":   "date",   //查询结束日期，默认为当天
				"isNew":     "int16",  //关联至今是否存在，1为存在，0为关联已结束、返回的各关联度为0，默认查询全部
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"themeID":               "int64",  //主题ID
				"themeName":             "string", //主题名称
				"secID":                 "string", //证券内部编码(通联自编)
				"exchangeCD":            "string", //证券交易场所代码(通联自编)
				"ticker":                "string", //证券交易代码
				"beginDate":             "date",   //关联开始日期
				"secShortName":          "string", //证券简称
				"exchangeName":          "string", //证券交易所
				"endDate":               "date",   //关联结束日期，日期为空表示关联至今依旧存在
				"isNew":                 "date",   //关联至今是否存在
				"returnScore":           "double", //收益关联程度，主题和证券在短期收益上的相似度，取值范围[0，1]，值越大表示关联度越高
				"textContributionScore": "double", //文本贡献关联度，主题和证券在新闻文本中的相似度，取值范围[0，1]，值越大表示关联度越高
				"industryScore":         "double", //行业关联度，主题和证券在行业分布上的相似度，取值范围[0，1]，值越大表示关联度越高
				"insertTime":            "date",   //入库时间
				"updateTime":            "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明获取一段时间内主题新增的关联证券数据，输入主题代码或名称、查询起止时间，可以获取该时间段内主题新增的关联证券信息，包括证券代码、证券简称、证券交易场所，同时返回三个维度的关联分数、关联开始时间、关联结束时间、关联具体描述、数据入库及更新时间。(注：1、主题与证券的关联自2013/12/28始、2014年12月起关联数据完整；2、数据按日更新。)
		"getThemesTickersInsert": map[string]interface{}{
			"url": "/api/subject/getThemesTickersInsert.json",
			"args": map[string]string{
				"themeID":   "int64",  //主题ID，允许多值输入，如:"1,2,3,4"
				"themeName": "string", //主题名称，支持模糊查询
				"beginDate": "date",   //查询开始日期，默认为当天
				"endDate":   "date",   //查询结束日期，默认为当天
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"themeID":               "int64",  //主题ID
				"themeName":             "string", //主题名称
				"secID":                 "string", //证券内部编码(通联自编)
				"exchangeCD":            "string", //证券交易场所代码(通联自编)
				"ticker":                "string", //证券交易代码
				"beginDate":             "date",   //关联开始日期
				"secShortName":          "string", //证券简称
				"exchangeName":          "string", //证券交易所
				"endDate":               "date",   //关联结束日期，日期为空表示关联至今依旧存在
				"isNew":                 "date",   //关联至今是否存在
				"returnScore":           "double", //收益关联程度，主题和证券在短期收益上的相似度，取值范围[0，1]，值越大表示关联度越高
				"textContributionScore": "double", //文本贡献关联度，主题和证券在新闻文本中的相似度，取值范围[0，1]，值越大表示关联度越高
				"industryScore":         "double", //行业关联度，主题和证券在行业分布上的相似度，取值范围[0，1]，值越大表示关联度越高
				"insertTime":            "date",   //入库时间
				"updateTime":            "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明包含按单个统计日期获取的雪球社交数据，输入一个统计日期，获取当天雪球帖子涉及的所有证券、各证券雪球帖子数量、帖子占比(%)。(注：数据自2014/1/1始，按日更新。)
		"getSocialDataXQByDate": map[string]interface{}{
			"url": "/api/subject/getSocialDataXQByDate.json",
			"args": map[string]string{
				"statisticsDate": "date",   //统计日期，单值输入
				"field":          "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"ticker":         "string", //证券交易代码
				"statisticsDate": "date",   //统计日期
				"postNum":        "int64",  //雪球帖子数量
				"postPercent":    "double", //雪球帖子占比(%)
				"insertTime":     "date",   //入库时间
				"updateTime":     "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明包含证券相关的新闻热度指数数据，输入一个或多个证券交易代码、起止日期，获取该证券一段时间内的新闻热度指数(即证券当天关联新闻数量占当天新闻总量的百分比(%))。每天更新。（注：1、2014/1/1起新闻来源众多、指数统计有效，2013年及之前的网站来源不全、数据波动大，数据自2004/10/28始；2、新闻量的统计口径为经算法处理后证券关联到的所有常规新闻；3、数据按日更新。)
		"getNewsHeatIndex": map[string]interface{}{
			"url": "/api/subject/getNewsHeatIndex.json",
			"args": map[string]string{
				"beginDate":    "date",   //查询开始日期，默认为一个月前
				"endDate":      "date",   //查询结束日期，默认为当天
				"secID":        "string", //证券内部编码，可通过交易代码和交易市场在getSecID获取到，可多值输入，如："000001.XSHE"、"000001.XSHE,600001.XSHG"
				"exchangeCD":   "string", //证券交易所代码(通联自编)。可选：XSHG、XSHE、XHKG。XSHG表示上海证券交易所，XSHE表示深圳证券交易所，XHKG表示香港交易所。可多值输入
				"secShortName": "string", //证券简称，支持模糊查询
				"ticker":       "string", //证券交易代码，可多值输入，如："000001"、"000001,600001"
				"field":        "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":           "string", //证券内部编码(通联自编)
				"exchangeCD":      "string", //证券交易所代码(通联自编)
				"ticker":          "string", //证券交易代码
				"newsPublishDate": "date",   //新闻发布日期
				"exchangeName":    "string", //证券交易所
				"secShortName":    "string", //证券简称
				"heatIndex":       "double", //新闻热度指数，证券当天关联新闻数量占当天关联新闻总量的百分比(%）
				"insertTime":      "date",   //入库时间
				"updateTime":      "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明获取某天某一段时间内入库的新闻基本信息。输入新闻入库的日期、起止时间，获取该时间段内新入库的新闻相关信息，如：新闻ID、标题、摘要、初始来源、作者、发布来源、发布时间、新闻入库时间等。(注：1、自2014/1/1起新闻来源众多、新闻量日均4万左右，2013年及之前的网站来源少、新闻数据量少；2、数据实时更新。)
		"getNewsInfoByInsertTime": map[string]interface{}{
			"url": "/api/subject/getNewsInfoByInsertTime.json",
			"args": map[string]string{
				"newsInsertDate": "date",   //新闻入库日期，单值输入，如：20150101
				"beginTime":      "string", //查询新闻入库开始时间，格式为HH:MM，如 09:30
				"endTime":        "string", //查询新闻入库结束时间，格式为HH:MM，如 09:30
				"field":          "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"newsInsertTime":   "date",   //新闻入库时间
				"newsID":           "int64",  //新闻ID
				"newsTitle":        "text",   //新闻标题
				"newsSummary":      "text",   //(源自算法处理结果)
				"newsOriginSource": "string", //新闻初始来源，即新闻原始出处
				"newsAuthor":       "text",   //新闻作者
				"newsPublishSite":  "string", //新闻发布来源，即新闻的实际最终来源
				"newsPublishTime":  "date",   //新闻发布时间
			},
			"indices": map[string]string{},
		},
		//说明根据证券代码获取相应公告分类结果，输入一个或多个证券交易代码，可以获取所查询证券相关的公告信息，包括公告ID、公告名称、证券交易场所、证券交易所对公告的原始分类、公告分类结果、公告分类入库时间、更新时间。(注：公告分类数据自2009/1/5始，按日更新)
		"getReportByTicker": map[string]interface{}{
			"url": "/api/subject/getReportByTicker.json",
			"args": map[string]string{
				"ticker":    "string", //证券交易代码，可输入一个或多个，用","分隔，如"000001"、"000001,600001"。（可多值输入）
				"beginDate": "date",   //查询开始日期，默认为30天前
				"endDate":   "date",   //查询结束日期，默认为当天
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"reportID":    "int64",  //公告代码
				"publishDate": "date",   //公告发布时间
				"tiker":       "string", //证券交易代码
				"year":        "int32",  //公告年份
				"title":       "string", //公告标题
				"site":        "string", //证券交易场所，包括"SH"、"SZ"。"SH"指上海证券交易所、"SZ"指深圳证券交易所
				"reportType":  "string", //证券交易所对公告的原始分类
				"Category":    "string", //公告自动分类
				"insertTime":  "date",   //自动分类数据入库时间
				"updateTime":  "date",   //自动分类数据更新时间
			},
			"indices": map[string]string{},
		},
		//说明根据发布时间获取新闻关联的主题数据，原API(根据发布时间获取新闻关联的主题数据-getThemesByNewsTime)的升级版。输入新闻发布的起止时间，可以获取相关的主题信息，如：主题ID、主题名称，同时返回新闻标题、新闻发布时间、关联数据入库时间、更新时间等。(注：1、自2014/1/1起新闻来源众多、新闻量日均4万左右，2013年及之前的网站来源少、新闻数据量少；2、数据实时更新；3、关联数据入库起始时间为2015/06/17。)
		"getThemesByNewsTime2": map[string]interface{}{
			"url": "/api/subject/getThemesByNewsTime2.json",
			"args": map[string]string{
				"publishBeginTime": "string", //查询新闻发布开始时间，格式为：YYYY-MM-DD HH:MM:SS 、YYYY-MM-DDTHH:mm:SS或 YYYYMMDDHHMMSS，如：2015-02-26 18:05:24、2015-02-26T18:05:24或20150226180524
				"publishEndTime":   "string", //查询新闻发布结束时间，格式为：YYYY-MM-DD HH:MM:SS 、YYYY-MM-DDTHH:mm:SS或 YYYYMMDDHHMMSS，如：2015-02-26 18:05:24、2015-02-26T18:05:24或20150226180524
				"field":            "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"newsPublishTime": "date",   //新闻发布时间
				"newsID":          "int64",  //新闻ID
				"themeID":         "int64",  //主题ID
				"newsTitle":       "text",   //新闻标题
				"newsPublishSite": "string", //新闻发布来源，即新闻的实际最终来源
				"themeName":       "string", //主题名称
				"insertTime":      "date",   //入库时间
				"updateTime":      "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明获取某天活跃的主题数据。输入一个日期，获取在该日期活跃的主题。(注：1、主题活跃周期数据自2013/1/1始；2、新闻量在某段时间内达到活跃阈值的主题即为活跃主题；3、数据按日更新。)
		"getActiveThemes": map[string]interface{}{
			"url": "/api/subject/getActiveThemes.json",
			"args": map[string]string{
				"date":  "date",   //查询日期，输入一个日期获取当天的活跃主题
				"field": "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"themeID":    "int64",  //主题ID
				"beginDate":  "date",   //主题开始日期
				"themeName":  "string", //主题名称
				"endDate":    "date",   //主题结束日期
				"isLatest":   "int16",  //是否最新活跃期，1为最新，0为非最新
				"insertTime": "date",   //入库时间
				"updateTime": "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明包含证券关联的主题数据，主题源自网络。输入证券交易所代码、证券交易代码或简称，可以获取关联的主题等信息，包括证券代码、证券简称、证券交易场所，同时返回三个维度的关联分数、关联开始时间、关联结束时间、关联具体描述、数据入库及更新时间，同时可输入查询起止时间，以获取证券在该时间段内关联到的主题信息。(注：1、源自网络的主题与证券的关联自2013/12/28始、2014年12月起关联数据完整；2、数据按日更新。)
		"getWebThemesByTickers": map[string]interface{}{
			"url": "/api/subject/getWebThemesByTickers.json",
			"args": map[string]string{
				"secID":        "string", //证券内部编码，可通过交易代码和交易市场在getSecID获取到，可多值输入，如："000001.XSHE"、"000001.XSHE,600001.XSHG"
				"secShortName": "string", //证券简称，支持模糊查询，如"浦发"、"浦发银行"
				"ticker":       "string", //证券交易代码，允许多值输入，用","分隔，如"000001"、"000001,600001"。
				"beginDate":    "date",   //查询开始日期，默认为当天
				"endDate":      "date",   //查询结束日期，默认为当天
				"exchangeCD":   "string", //证券交易所代码(通联自编)。可选：XSHG、XSHE、XHKG、XNYS、NOBB、XNAS。XSHG表示上海证券交易所，XSHE表示深圳证券交易所，XHKG表示香港交易所，XNYS表示纽约证券交易所，NOBB表示三板市场，XNYS表示美国纳斯达克证券交易所。可多值输入
				"field":        "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":                 "string", //证券内部编码(通联自编)
				"exchangeCD":            "string", //证券交易所代码(通联自编)
				"ticker":                "string", //证券交易代码
				"beginDate":             "date",   //关联开始日期
				"themeID":               "int64",  //主题ID
				"exchangeName":          "string", //证券交易所
				"secShortName":          "string", //证券简称
				"themeName":             "string", //主题名称
				"endDate":               "date",   //关联结束日期，日期为空表示关联至今依旧存在
				"returnScore":           "double", //收益关联程度，主题和证券在短期收益上的相似度，取值范围[0，1]，值越大表示关联度越高
				"textContributionScore": "double", //文本贡献关联度，主题和证券在新闻文本中的相似度，取值范围[0，1]，值越大表示关联度越高
				"industryScore":         "double", //行业关联度，主题和证券在行业分布上的相似度，取值范围[0，1]，值越大表示关联度越高
				"insertTime":            "date",   //入库时间
				"updateTime":            "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明获取新闻关联的主题数据，该API以获取新闻关联的主题(getThemesByNews)为基础、进行过滤优化。输入新闻ID，可以获取相关的主题信息，如：主题ID、主题名称，同时返回新闻标题、新闻发布时间、关联数据插入时间、更新时间等。(注：1、自2014/1/1起新闻来源众多、新闻量日均4万左右，2013年及之前的网站来源少、新闻数据量少；2、数据实时更新；3、关联数据入库起始时间为2015/04/07。)
		"getThemesByNewsLF": map[string]interface{}{
			"url": "/api/subject/getThemesByNewsLF.json",
			"args": map[string]string{
				"newsID":     "int64",  //新闻ID，可多值输入，可由getNewsInfoByTime(获取一天某段时间内的新闻信息)这个API中获取新闻ID
				"insertDate": "date",   //关联数据入库日期，单值输入
				"beginTime":  "string", //查询关联数据入库开始时间，格式为HH:MM，如 09:30
				"endTime":    "string", //查询关联数据入库结束时间，格式为HH:MM，如 09:30
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"newsID":          "int64",  //新闻ID
				"themeID":         "int64",  //主题ID
				"newsTitle":       "text",   //新闻标题
				"newsPublishTime": "date",   //新闻发布时间
				"newsPublishSite": "string", //新闻发布来源
				"themeName":       "string", //主题名称
				"insertTime":      "date",   //入库时间
				"updateTime":      "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明根据发布时间获取新闻关联的主题数据，只包含与公司相关的新闻。输入新闻发布的起止时间，可以获取相关的主题信息，如：主题ID、主题名称，同时返回新闻标题、新闻发布时间、关联数据入库时间、更新时间等。(注：1、自2014/1/1起新闻来源众多、新闻量日均4万左右，2013年及之前的网站来源少、新闻数据量少；2、数据实时更新；3、关联数据入库起始时间为2015/04/07。)
		"getThemesByNewsTimeCompanyRel": map[string]interface{}{
			"url": "/api/subject/getThemesByNewsTimeCompanyRel.json",
			"args": map[string]string{
				"publishBeginTime": "string", //查询新闻发布开始时间，格式为：YYYY-MM-DD HH:MM:SS 、YYYY-MM-DDTHH:mm:SS或 YYYYMMDDHHMMSS，如：2015-02-26 18:05:24、2015-02-26T18:05:24或20150226180524
				"publishEndTime":   "string", //查询新闻发布结束时间，格式为：YYYY-MM-DD HH:MM:SS 、YYYY-MM-DDTHH:mm:SS或 YYYYMMDDHHMMSS，如：2015-02-26 18:05:24、2015-02-26T18:05:24或20150226180524
				"field":            "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"newsID":          "int64",  //新闻ID
				"themeID":         "int64",  //主题ID
				"newsTitle":       "text",   //新闻标题
				"newsPublishSite": "string", //新闻发布来源，即新闻的实际最终来源
				"newsPublishTime": "date",   //新闻发布时间
				"themeName":       "string", //主题名称
				"insertTime":      "date",   //入库时间
				"updateTime":      "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明获取新闻关联的主题数据，该API以获取新闻关联的主题(优化后)(getThemesByNewsLF)为基础、再次进行过滤优化，是所有获取新闻关联的主题API中最严格的优化结果、数据量也最少。输入新闻ID，可以获取相关的主题信息，如：主题ID、主题名称，同时返回新闻标题、新闻发布时间、关联数据插入时间、更新时间等。(注：1、自2014/1/1起新闻来源众多、新闻量日均4万左右，2013年及之前的网站来源少、新闻数据量少；2、数据实时更新；3、关联数据入库起始时间为2015/04/07。)
		"getThemesByNewsMF": map[string]interface{}{
			"url": "/api/subject/getThemesByNewsMF.json",
			"args": map[string]string{
				"newsID":     "int64",  //新闻ID，可多值输入，可由getNewsInfoByTime(获取一天某段时间内的新闻信息)这个API中获取新闻ID
				"insertDate": "date",   //关联数据入库日期，单值输入
				"beginTime":  "string", //查询关联数据入库开始时间，格式为HH:MM，如 09:30
				"endTime":    "string", //查询关联数据入库结束时间，格式为HH:MM，如 09:30
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"newsID":          "int64",  //新闻ID
				"themeID":         "int64",  //主题ID
				"newsTitle":       "text",   //新闻标题
				"newsPublishSite": "string", //新闻发布来源
				"newsPublishTime": "date",   //新闻发布时间
				"themeName":       "string", //主题名称
				"insertTime":      "date",   //入库时间
				"updateTime":      "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明分新闻源获取一天某段时间内的新闻信息，该API为未经去重处理的新闻数据。可先在获取新闻发布来源列表(getNewsPublishSite)API中获取通联的新闻源列表。输入新闻发布的日期、起止时间、新闻发布来源，获取该时间段内的新闻相关信息，如：新闻ID、标题、摘要、初始来源、作者、发布来源、发布时间、入库时间等。(注：1、自2014/1/1起新闻来源众多、新闻量日均4万左右，2013年及之前的网站来源少、新闻数据量少；2、数据实时更新。)
		"getNewsInfoByTimeAndSite": map[string]interface{}{
			"url": "/api/subject/getNewsInfoByTimeAndSite.json",
			"args": map[string]string{
				"newsPublishDate": "date",   //新闻发布日期，单值输入，输入格式“YYYYMMDD”
				"beginTime":       "string", //查询新闻发布开始时间，格式为HH:MM，如 09:30
				"endTime":         "string", //查询新闻发布结束时间，格式为HH:MM，如 09:30
				"newsPublishSite": "string", //新闻发布来源，即新闻的实际最终来源,可在获取新闻发布来源列表(getNewsPublishSite)API中获取通联的新闻源列表
				"pageTag":         "string", //新闻源页面标签，爬取的新闻源的新闻分类标签，目前只适用于newsPublishSite='和讯网'的新闻，所涉及的类别为：互联网金融,保险,信托,债券,公益,养老金,创投,商学院,商旅,国债期货,基金,外汇,奢侈品,房产,收藏,新闻,期货,汽车,现货,理财,科技,税务,股指期货,股票,舆情,评论,读书,银行,黄金。
				"field":           "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"newsPublishTime":  "date",   //新闻发布时间
				"newsID":           "int64",  //新闻ID
				"newsTitle":        "text",   //新闻标题
				"newsSummary":      "text",   //新闻摘要
				"newsOriginSource": "string", //新闻初始来源，即新闻原始出处
				"newsAuthor":       "text",   //新闻作者
				"newsPublishSite":  "string", //新闻发布来源，即新闻的实际最终来源
				"newsInsertTime":   "date",   //新闻入库时间
				"pageTag":          "string", //新闻源页面标签，爬取的新闻源的新闻分类标签，目前只适用于newsPublishSite='和讯网'的新闻
			},
			"indices": map[string]string{},
		},
		//说明包含新闻基本信息。输入新闻ID，获取新闻基本信息，如：新闻ID、标题、摘要、初始来源、作者、发布来源、发布时间、入库时间等。(注：1、自2014/1/1起新闻来源众多、新闻量日均4万左右，2013年及之前的网站来源少、每天新闻数据量少；2、数据实时更新。)
		"getNewsInfo": map[string]interface{}{
			"url": "/api/subject/getNewsInfo.json",
			"args": map[string]string{
				"newsID": "int64",  //新闻ID，可多值输入
				"field":  "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"newsPublishTime":  "date",   //新闻发布时间
				"newsID":           "int64",  //新闻ID
				"newsTitle":        "string", //新闻标题
				"newsSummary":      "string", //(源自算法处理结果)
				"newsOriginSource": "string", //新闻初始来源，即新闻原始出处
				"newsAuthor":       "text",   //新闻作者
				"newsPublishSite":  "string", //新闻发布来源，即新闻的实际最终来源
				"newsInsertTime":   "date",   //新闻入库时间
			},
			"indices": map[string]string{},
		},
		//说明包含证券相关的新闻数据，同时可获取针对不同证券的新闻情感数据。输入证券代码或简称、查询的新闻发布起止时间，同时可输入证券交易所代码，获取相关新闻数据，如：新闻ID、新闻标题、发布来源、发布时间、入库时间等。(注：1、自2014/1/1起新闻来源众多、新闻量日均4万左右，2013年及之前的网站来源少、新闻数据量少；2、数据实时更新。)
		"getNewsByTickers": map[string]interface{}{
			"url": "/api/subject/getNewsByTickers.json",
			"args": map[string]string{
				"secID":        "string", //证券内部编码，可通过交易代码和交易市场在getSecID获取到，可多值输入，如："000001.XSHE"、"000001.XSHE,600001.XSHG"
				"secShortName": "string", //证券简称，支持模糊查询
				"ticker":       "string", //证券交易代码，可多值输入，如："000001"、"000001,600001"
				"beginDate":    "date",   //查询新闻发布开始日期，默认为30天前
				"endDate":      "date",   //查询新闻发布结束日期，默认为当天
				"exchangeCD":   "string", //证券交易所代码(通联自编)。可选：XSHG、XSHE、XHKG、XNYS。XSHG表示上海证券交易所，XSHE表示深圳证券交易所，XHKG表示香港交易所，XNYS表示纽约证券交易所。可多值输入
				"field":        "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":           "string", //证券内部编码(通联自编)
				"exchangeCD":      "string", //证券交易所代码
				"ticker":          "string", //证券交易代码
				"newsPublishTime": "date",   //新闻发布时间
				"exchangeName":    "string", //证券交易所
				"secShortName":    "string", //证券简称
				"newsID":          "int64",  //新闻ID
				"newsTitle":       "string", //新闻标题
				"relatedScore":    "double", //，
				"sentiment":       "int16",  //证券所属公司新闻情感，新闻针对特定证券所属公司的情感分类，1为正面、0为中性、-1为负面
				"sentimentScore":  "double", //；该分数的计算结合了2种方法：1、机器学习(通用的深度学习+SVM算法，利用人工准备的训练数据构建情感识别模型)，2、词典（基于通联内部整理的情感词典）
				"newsPublishSite": "string", //新闻发布来源，即新闻的实际最终来源
				"newsInsertTime":  "date",   //新闻入库时间
			},
			"indices": map[string]string{},
		},
		//说明获取某天某一段时间内入库的新闻全文等信息。输入新闻入库的日期、起止时间，获取该时间段内新入库的新闻全文等信息，如：新闻ID、标题、摘要、正文、来源链接、初始来源、作者、发布来源、发布时间、新闻入库时间等。(注：1、自2014/1/1起新闻来源众多、新闻量日均4万左右，2013年及之前的网站来源少、新闻数据量少；2、数据实时更新。)
		"getNewsContentByInsertTime": map[string]interface{}{
			"url": "/api/subject/getNewsContentByInsertTime.json",
			"args": map[string]string{
				"newsInsertDate": "date",   //新闻入库日期，单值输入，如：20150101
				"beginTime":      "string", //查询新闻入库开始时间，格式为HH:MM，如 09:30
				"endTime":        "string", //查询新闻入库结束时间，格式为HH:MM，如 09:30
				"field":          "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"newsInsertTime":   "date",   //新闻入库时间
				"newsID":           "int64",  //新闻ID
				"newsTitle":        "text",   //新闻标题
				"newsSummary":      "text",   //(源自算法处理结果)
				"newsBody":         "text",   //新闻正文
				"newsURL":          "string", //新闻链接
				"newsOriginSource": "string", //新闻初始来源，即新闻原始出处
				"newsAuthor":       "text",   //新闻作者
				"newsPublishSite":  "string", //新闻发布来源，即新闻的实际最终来源
				"newsPublishTime":  "date",   //新闻发布时间
			},
			"indices": map[string]string{},
		},
		//说明获取当天活跃主题聚类对应关系数据。输入聚类后的主要主题代码或名称，可以获取同一类别的主题相关信息，包括主题ID、主题名称、主题插入时间、主题更新时间等。(注：1、可先在主题基本信息(getThemesContent)这个API中获取当天聚类后的主题；2、可输入isMain=0，在返回的数据中剔除主题自身的对应；3、数据每天刷新，只能获取当天数据。)
		"getThemesCluster": map[string]interface{}{
			"url": "/api/subject/getThemesCluster.json",
			"args": map[string]string{
				"isMain":    "int16",  //是否聚类后的主题，同类主题是否为聚类后的主要主题，1为主要主题，0为非主要主题，输入0则返回主要主题与非主要主题之间的对应，为空则输出同一类别下的所有主题
				"themeID":   "int64",  //主题ID，允许多值输入，为空则输出所有
				"themeName": "string", //主题名称，支持模糊查询
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"themeID":      "int64",  //主题ID
				"themeRelID":   "int64",  //同类主题ID
				"themeName":    "string", //主题名称
				"themeRelName": "string", //同类主题名称
				"isMain":       "int16",  //是否聚类后的主题，是指返回的同类主题是否为聚类后的主要主题，1为主要主题，0为非主要主题
				"insertTime":   "date",   //入库时间，即同类主题的入库时间
				"updateTime":   "date",   //更新时间，即同类主题的入库时间
			},
			"indices": map[string]string{},
		},
		//说明获取与某主题相似的其他主题数据。输入主题代码或名称，可以获取相似的主题信息，包括相似主题代码、相似主题名称、主题文本的相似度、主题关联证券的相似度。数据按日更新。
		"getThemesSimilarity": map[string]interface{}{
			"url": "/api/subject/getThemesSimilarity.json",
			"args": map[string]string{
				"themeID":   "int64",  //主题ID，允许多值输入，如:"1,2,3,4"
				"themeName": "string", //主题名称，支持模糊查询
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"themeID":         "int64",  //主题ID
				"themeSimID":      "int64",  //相似主题ID
				"themeName":       "string", //主题名称
				"themeSimName":    "string", //相似主题名称
				"textSimScore":    "double", //主题相似度，即两个主题文本内容相似度，取值范围[0，1]，值越大相似度越高
				"tickersSimScore": "double", //主题关联证券相似度，即两个主题关联到的证券的相似度，取值范围[0，1]，值越大相似度越高
				"insertTime":      "date",   //入库时间
				"updateTime":      "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明获取新闻关联的主题数据。输入新闻ID，可以获取相关的主题信息，如：主题ID、主题名称，同时返回新闻标题、新闻发布时间、关联数据插入时间、更新时间等。(注：1、自2014/1/1起新闻来源众多、新闻量日均4万左右，2013年及之前的网站来源少、新闻数据量少；2、数据实时更新；3、关联数据入库起始时间为2015/04/07。)
		"getThemesByNews": map[string]interface{}{
			"url": "/api/subject/getThemesByNews.json",
			"args": map[string]string{
				"newsID":     "int64",  //新闻ID，可多值输入，可由getNewsInfoByTime(获取一天某段时间内的新闻信息)这个API中获取新闻ID
				"insertDate": "date",   //关联数据入库日期，单值输入
				"beginTime":  "string", //查询关联数据入库开始时间，格式为HH:MM，如 09:30
				"endTime":    "string", //查询关联数据入库结束时间，格式为HH:MM，如 09:30
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"newsID":          "int64",  //新闻ID
				"newsTitle":       "text",   //新闻标题
				"newsPublishTime": "date",   //新闻发布时间
				"newsPublishSite": "string", //新闻发布来源
				"themeID":         "int64",  //主题ID
				"themeName":       "string", //主题名称
				"insertTime":      "date",   //入库时间
				"updateTime":      "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明根据发布时间获取新闻关联的主题数据。输入新闻发布的起止时间，可以获取相关的主题信息，如：主题ID、主题名称，同时返回新闻标题、新闻发布时间、关联数据入库时间、更新时间等。(注：1、自2014/1/1起新闻来源众多、新闻量日均4万左右，2013年及之前的网站来源少、新闻数据量少；2、数据实时更新；3、关联数据入库起始时间为2015/04/07。)
		"getThemesByNewsTime": map[string]interface{}{
			"url": "/api/subject/getThemesByNewsTime.json",
			"args": map[string]string{
				"publishBeginTime": "string", //查询新闻发布开始时间，格式为：YYYY-MM-DD HH:MM:SS 、YYYY-MM-DDTHH:mm:SS或 YYYYMMDDHHMMSS，如：2015-02-26 18:05:24、2015-02-26T18:05:24或20150226180524
				"publishEndTime":   "string", //查询新闻发布结束时间，格式为：YYYY-MM-DD HH:MM:SS 、YYYY-MM-DDTHH:mm:SS或 YYYYMMDDHHMMSS，如：2015-02-26 18:05:24、2015-02-26T18:05:24或20150226180524
				"field":            "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"newsID":          "int64",  //新闻ID
				"themeID":         "int64",  //主题ID
				"newsTitle":       "text",   //新闻标题
				"newsPublishSite": "string", //新闻发布来源，即新闻的实际最终来源
				"newsPublishTime": "date",   //新闻发布时间
				"themeName":       "string", //主题名称
				"insertTime":      "date",   //入库时间
				"updateTime":      "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明包含证券在股吧社交中的热度统计数据，输入一个或多个证券交易代码、统计起止日期，该证券在一段时间内每天相关的股吧帖子数量、帖子占比(%)。(注：数据自2014/1/1始，按日更新。)
		"getSocialDataGuba": map[string]interface{}{
			"url": "/api/subject/getSocialDataGuba.json",
			"args": map[string]string{
				"beginDate": "date",   //统计开始日期，默认为30天前
				"endDate":   "date",   //统计结束日期，默认为当天
				"ticker":    "string", //证券交易代码，可多值输入
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"ticker":         "string", //证券交易代码
				"statisticsDate": "date",   //统计日期
				"postNum":        "int64",  //股吧帖子数量
				"postPercent":    "double", //股吧帖子占比(%)
				"insertTime":     "date",   //入库时间
				"updateTime":     "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明包含主题在股吧社交中的热度统计数据，输入一个或多个主题代码、统计起止日期，获取该主题在一段时间内每天相关的股吧帖子数量、帖子占比(%)。(注：数据自2014/1/1始，按日更新。)
		"getSocialThemeDataGuba": map[string]interface{}{
			"url": "/api/subject/getSocialThemeDataGuba.json",
			"args": map[string]string{
				"beginDate": "date",   //统计开始日期，默认为30天前
				"endDate":   "date",   //统计结束日期，默认为当天
				"themeID":   "int64",  //主题代码，可多值输入
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"themeID":        "int64",  //主题代码
				"statisticsDate": "date",   //统计日期
				"postNum":        "int64",  //股吧帖子数量
				"postPercent":    "double", //股吧帖子占比(%)
				"insertTime":     "date",   //入库时间
				"updateTime":     "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明包含新闻全文信息，输入新闻ID，获取新闻全文、新闻链接等。(注：1、自2014/1/1起新闻来源众多、新闻量日均4万左右，2013年及之前的网站来源少、新闻数据量少；2、数据实时更新。)
		"getNewsBody": map[string]interface{}{
			"url": "/api/subject/getNewsBody.json",
			"args": map[string]string{
				"newsID": "int64",  //新闻ID，可多值输入
				"field":  "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"newsID":   "int64",  //新闻ID
				"newsBody": "text",   //新闻正文
				"newsURL":  "string", //新闻链接
			},
			"indices": map[string]string{},
		},
		//说明获取某天某一段时间内的新闻基本信息。输入新闻发布的日期、起止时间，获取该时间段内的新闻相关信息，如：新闻ID、标题、摘要、初始来源、作者、发布来源、发布时间、入库时间等。(注：1、自2014/1/1起新闻来源众多、新闻量日均4万左右，2013年及之前的网站来源少、新闻数据量少；2、数据实时更新。)
		"getNewsInfoByTime": map[string]interface{}{
			"url": "/api/subject/getNewsInfoByTime.json",
			"args": map[string]string{
				"newsPublishDate": "date",   //新闻发布日期，单值输入，如：20150101
				"beginTime":       "string", //查询新闻发布开始时间，格式为HH:MM，如 09:30
				"endTime":         "string", //查询新闻发布结束时间，格式为HH:MM，如 09:30
				"field":           "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"newsPublishTime":  "date",   //新闻发布时间
				"newsID":           "int64",  //新闻ID
				"newsTitle":        "string", //新闻标题
				"newsSummary":      "string", //新闻摘要(源自算法处理结果)
				"newsOriginSource": "string", //新闻初始来源，即新闻原始出处
				"newsAuthor":       "text",   //新闻作者
				"newsPublishSite":  "string", //新闻发布来源，即新闻的实际最终来源
				"newsInsertTime":   "date",   //新闻入库时间
			},
			"indices": map[string]string{},
		},
		//说明包含按单只证券代码获取的雪球社交数据，输入一个证券交易代码，获取该证券每天的雪球帖子数量、及帖子占比(%)。(注：数据自2014/1/1始，按日更新。)
		"getSocialDataXQByTicker": map[string]interface{}{
			"url": "/api/subject/getSocialDataXQByTicker.json",
			"args": map[string]string{
				"ticker": "string", //证券交易代码，单值输入
				"field":  "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"statisticsDate": "date",   //统计日期
				"ticker":         "string", //证券交易代码
				"postNum":        "int64",  //雪球帖子数量
				"postPercent":    "double", //雪球帖子占比(%)
				"insertTime":     "date",   //入库时间
				"updateTime":     "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明获取新闻关联的主题数据，只包含与公司相关的新闻。输入新闻ID，可以获取相关的主题信息，如：主题ID、主题名称，同时返回新闻标题、新闻发布时间、关联数据插入时间、更新时间等。(注：1、自2014/1/1起新闻来源众多、新闻量日均4万左右，2013年及之前的网站来源少、新闻数据量少；2、数据实时更新；3、关联数据入库起始时间为2015/04/07。)
		"getThemesByNewsCompanyRel": map[string]interface{}{
			"url": "/api/subject/getThemesByNewsCompanyRel.json",
			"args": map[string]string{
				"newsID":     "int64",  //新闻ID，可多值输入，可由getNewsInfoByTime(获取一天某段时间内的新闻信息)这个API中获取新闻ID
				"insertDate": "date",   //关联数据入库日期，单值输入
				"beginTime":  "string", //查询关联数据入库开始时间，格式为HH:MM，如 09:30
				"endTime":    "string", //查询关联数据入库结束时间，格式为HH:MM，如 09:30
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"newsID":          "int64",  //新闻ID
				"newsTitle":       "text",   //新闻标题
				"newsPublishTime": "date",   //新闻发布时间
				"newsPublishSite": "string", //新闻发布来源
				"themeID":         "int64",  //主题ID
				"themeName":       "string", //主题名称
				"insertTime":      "date",   //入库时间
				"updateTime":      "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明根据发布时间获取新闻关联的主题数据，该API以获取新闻关联的主题(getThemesByNewsTime)为基础、进行过滤优化。输入新闻发布的起止时间，可以获取相关的主题信息，如：主题ID、主题名称，同时返回新闻标题、新闻发布时间、关联数据入库时间、更新时间等。(注：1、自2014/1/1起新闻来源众多、新闻量日均4万左右，2013年及之前的网站来源少、新闻数据量少；2、数据实时更新；3、关联数据入库起始时间为2015/04/07。)
		"getThemesByNewsTimeLF": map[string]interface{}{
			"url": "/api/subject/getThemesByNewsTimeLF.json",
			"args": map[string]string{
				"publishBeginTime": "string", //查询新闻发布开始时间，格式为：YYYY-MM-DD HH:MM:SS 、YYYY-MM-DDTHH:mm:SS或 YYYYMMDDHHMMSS，如：2015-02-26 18:05:24、2015-02-26T18:05:24或20150226180524
				"publishEndTime":   "string", //查询新闻发布结束时间，格式为：YYYY-MM-DD HH:MM:SS 、YYYY-MM-DDTHH:mm:SS或 YYYYMMDDHHMMSS，如：2015-02-26 18:05:24、2015-02-26T18:05:24或20150226180524
				"field":            "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"newsID":          "int64",  //新闻ID
				"themeID":         "int64",  //主题ID
				"newsTitle":       "text",   //新闻标题
				"newsPublishSite": "string", //新闻发布来源，即新闻的实际最终来源
				"newsPublishTime": "date",   //新闻发布时间
				"themeName":       "string", //主题名称
				"insertTime":      "date",   //入库时间
				"updateTime":      "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明获取一段时间内删除(退出)的活跃主题数据，输入的时间参数在主题活跃周期的结束时间列进行查询。输入查询起止时间、是否最新活跃期、主题来源，可以获取该时间段内停止活跃的主题信息，包括主题ID、主题名称、主题开始时间、主题结束时间、是否最新活跃期、数据入库及更新时间。(注：1、主题活跃周期数据自2013/1/1始；2、数据按日更新。3、查询当天无活跃主题被删除、需等第二天9:00之后获取前一天停止活跃的主题数据。)
		"getActiveThemesDelete": map[string]interface{}{
			"url": "/api/subject/getActiveThemesDelete.json",
			"args": map[string]string{
				"beginDate":   "date",   //查询开始日期，默认为当天
				"endDate":     "date",   //查询结束日期，默认为当天
				"isLatest":    "Int16",  //是否最新活跃期，1为最新、输入该值表示主题最新活跃期的结束时间在所查询时间段内，0为非最新，默认为1
				"themeSource": "Int16",  //主题来源，1为源自网络主题，2为自动挖掘生成的主题，3为通联人工自编主题，5为申万行业主题
				"field":       "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"themeID":     "int64",  //主题ID
				"themeName":   "string", //主题名称
				"themeSource": "Int16",  //主题来源
				"beginDate":   "date",   //主题开始日期
				"endDate":     "date",   //主题结束日期，结束日期为当天表示主题当天依旧活跃(第二天会重新刷新结束时间)
				"isLatest":    "int16",  //是否最新活跃期，1为最新，0为非最新
				"insertTime":  "date",   //入库时间
				"updateTime":  "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明获取一段时间内新入库的主题数据。输入查询起止时间，可以获取该时间段内新入库的主题信息，包括主题ID、主题名称、主题描述、主题来源、当天是否活跃、主题插入时间、主题更新时间等。(注：1、主题基期自2011/4/16始；2、数据按日更新主题活跃状态。)
		"getThemesInsertDB": map[string]interface{}{
			"url": "/api/subject/getThemesInsertDB.json",
			"args": map[string]string{
				"beginDate":   "date",   //查询主题入库开始日期
				"endDate":     "date",   //查询主题入库结束日期
				"themeSource": "int16",  //主题来源，1为源自网络主题，2为自动挖掘生成的主题，3为通联人工自编主题，5为申万行业主题
				"field":       "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"themeID":          "int64",  //主题ID
				"insertTime":       "date",   //入库时间
				"themeName":        "string", //主题名称
				"themeDescription": "string", //主题描述
				"themeSource":      "int16",  //主题来源
				"themeBaseDate":    "date",   //主题基期，计算主题收益等指数时的开始时间
				"isActive":         "int16",  //当天是否活跃，1为活跃，0为不活跃，每天刷新
				"isMain":           "int16",  //是否聚类后的主题
				"updateTime":       "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明根据发布时间获取新闻关联的主题数据，该API以获取新闻关联的主题(优化后)(getThemesByNewsTimeLF)为基础、再次进行过滤优化，是所有获取新闻关联的主题API中最严格的优化结果、数据量也最少。输入新闻发布的起止时间，可以获取相关的主题信息，如：主题ID、主题名称，同时返回新闻标题、新闻发布时间、关联数据入库时间、更新时间等。(注：1、自2014/1/1起新闻来源众多、新闻量日均4万左右，2013年及之前的网站来源少、新闻数据量少；2、数据实时更新；3、关联数据入库起始时间为2015/04/07。)
		"getThemesByNewsTimeMF": map[string]interface{}{
			"url": "/api/subject/getThemesByNewsTimeMF.json",
			"args": map[string]string{
				"publishBeginTime": "string", //查询新闻发布开始时间，格式为：YYYY-MM-DD HH:MM:SS 、YYYY-MM-DDTHH:mm:SS或 YYYYMMDDHHMMSS，如：2015-02-26 18:05:24、2015-02-26T18:05:24或20150226180524
				"publishEndTime":   "string", //查询新闻发布结束时间，格式为：YYYY-MM-DD HH:MM:SS 、YYYY-MM-DDTHH:mm:SS或 YYYYMMDDHHMMSS，如：2015-02-26 18:05:24、2015-02-26T18:05:24或20150226180524
				"field":            "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"newsID":          "int64",  //新闻ID
				"themeID":         "int64",  //主题ID
				"newsTitle":       "text",   //新闻标题
				"newsPublishSite": "string", //新闻发布来源，即新闻的实际最终来源
				"newsPublishTime": "date",   //新闻发布时间
				"themeName":       "string", //主题名称
				"insertTime":      "date",   //入库时间
				"updateTime":      "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明获取新闻关联的主题数据，原API(获取新闻关联的主题数据-getThemesByNews)的升级版。输入新闻ID或新闻与主题的关联数据入库起止时间，可以获取相关的主题信息，如：主题ID、主题名称，同时返回新闻标题、新闻发布时间、关联数据入库时间、更新时间等。(注：1、自2014/1/1起新闻来源众多、新闻量日均4万左右，2013年及之前的网站来源少、新闻数据量少；2、数据实时更新；3、关联数据入库起始时间为2015/06/17)
		"getThemesByNews2": map[string]interface{}{
			"url": "/api/subject/getThemesByNews2.json",
			"args": map[string]string{
				"insertBeginTime": "string", //查询关联数据入库开始时间，格式为：YYYY-MM-DD HH:MM:SS 、YYYY-MM-DDTHH:mm:SS或 YYYYMMDDHHMMSS，如：2015-02-26 18:05:24、2015-02-26T18:05:24或20150226180524
				"insertEndTime":   "string", //查询关联数据入库结束时间，格式为：YYYY-MM-DD HH:MM:SS 、YYYY-MM-DDTHH:mm:SS或 YYYYMMDDHHMMSS，如：2015-02-26 18:05:24、2015-02-26T18:05:24或20150226180524
				"newsID":          "int64",  //新闻ID，可多值输入，可由getNewsInfoByTime(获取一天某段时间内的新闻信息)这个API中获取新闻ID
				"field":           "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"insertTime":      "date",   //入库时间
				"newsID":          "int64",  //新闻ID
				"themeID":         "int64",  //主题ID
				"newsTitle":       "text",   //新闻标题
				"newsPublishTime": "date",   //新闻发布时间
				"newsPublishSite": "string", //新闻发布来源，即新闻的实际最终来源
				"themeName":       "string", //主题名称
				"updateTime":      "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//说明包含公司关联的新闻数据，同时可获取针对不同公司的新闻情感数据。输入公司代码、查询的新闻发布起止时间，获取相关的新闻信息，如：新闻ID、新闻标题、发布来源、发布时间、新闻入库时间等。(注：1、自2014/1/1起新闻来源众多、新闻量日均4万左右，2013年及之前的网站来源少、新闻数据量少；2、数据实时更新。)
		"getNewsByCompany": map[string]interface{}{
			"url": "/api/subject/getNewsByCompany.json",
			"args": map[string]string{
				"partyID":   "int64",  //公司代码(通联自编)，可先通过证券编码及基本上市信息(getSecID)这个API查找公司代码
				"beginDate": "date",   //查询新闻发布开始日期，默认为30天前
				"endDate":   "date",   //查询新闻发布结束日期，默认为当天
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"partyID":         "int64",  //公司代码(通联自编)，可通过证券编码及基本上市信息(getSecID)继续查找公司相关的证券
				"newsPublishTime": "date",   //新闻发布时间
				"partyFullName":   "string", //公司全称
				"newsID":          "int64",  //新闻ID
				"newsTitle":       "string", //新闻标题
				"relatedScore":    "double", //公司新闻关联度，取值范围[0，1]，数值越大、关联度越高，分数>0.2319表示关联显著，该接口只展示显著关联结果
				"sentiment":       "int16",  //公司新闻情感，新闻针对特定公司的情感分类，1为正面、0为中性、-1为负面
				"sentimentScore":  "double", //，利用人工准备的训练数据构建情感识别模型
				"newsPublishSite": "string", //新闻发布来源，即新闻的实际最终来源
				"newsInsertTime":  "date",   //新闻入库时间
			},
			"indices": map[string]string{},
		},
		//说明获取某天某一段时间内的新闻全文等信息。输入新闻发布的日期、起止时间，获取该时间段内的新闻全文等信息，如：新闻ID、标题、摘要、正文、来源链接、初始来源、作者、发布来源、发布时间、入库时间等。(注：1、自2014/1/1起新闻来源众多、新闻量日均4万左右，2013年及之前的网站来源少、新闻数据量少；2、数据实时更新。)
		"getNewsContentByTime": map[string]interface{}{
			"url": "/api/subject/getNewsContentByTime.json",
			"args": map[string]string{
				"newsPublishDate": "date",   //新闻发布日期，单值输入，如：20150101
				"beginTime":       "string", //查询新闻发布开始时间，格式为HH:MM，如 09:30
				"endTime":         "string", //查询新闻发布结束时间，格式为HH:MM，如 09:30
				"field":           "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"newsPublishTime":  "date",   //新闻发布时间
				"newsID":           "int64",  //新闻ID
				"newsTitle":        "text",   //新闻标题
				"newsSummary":      "text",   //新闻摘要
				"newsBody":         "text",   //新闻正文
				"newsURL":          "string", //新闻链接
				"newsOriginSource": "string", //新闻初始来源，即新闻原始出处
				"newsAuthor":       "text",   //新闻作者
				"newsPublishSite":  "string", //新闻发布来源，即新闻的实际最终来源
				"newsInsertTime":   "date",   //新闻入库时间
			},
			"indices": map[string]string{},
		},
		//说明根据证券代码获取公告内容，输入一个或多个证券交易代码，可以获取所查询证券相关的公告信息，包括公告ID、公告名称、证券交易场所、证券交易所对公告的原始分类、公告发布时间、公告具体内容、公告链接、公告入库时间。(注：公告数据自2000/1/8始，按日更新)
		"getReportContent": map[string]interface{}{
			"url": "/api/subject/getReportContent.json",
			"args": map[string]string{
				"ticker":    "string", //证券交易代码，可输入一个或多个，用","分隔，如"000001"、"000001,600001"。（可多值输入）
				"beginDate": "date",   //查询开始日期，默认为30天前
				"endDate":   "date",   //查询结束日期，默认为当天
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"reportID":    "int64",  //公告代码
				"publishDate": "date",   //公告发布时间
				"tiker":       "string", //证券交易代码
				"year":        "int32",  //公告年份
				"title":       "string", //公告标题
				"site":        "string", //证券交易场所，包括"SH"、"SZ"。"SH"指上海证券交易所、"SZ"指深圳证券交易所
				"reportType":  "string", //证券交易所对公告的原始分类
				"txtContent":  "string", //公告内容
				"url":         "string", //公告原文链接
				"insertTime":  "date",   //自动分类数据入库时间
				"updateTime":  "date",   //自动分类数据更新时间
			},
			"indices": map[string]string{},
		},
		//说明包含证券关联的主题数据。输入证券交易所代码、证券交易代码或简称，可以获取关联的主题等信息，包括证券代码、证券简称、证券交易场所，同时返回三个维度的关联分数、关联开始时间、关联结束时间、关联具体描述、数据入库及更新时间，同时可输入查询起止时间，以获取证券在该时间段内关联到的主题信息。(注：1、主题与证券的关联自2013/12/28始、2014年12月起关联数据完整；2、数据按日更新。)
		"getThemesByTickers": map[string]interface{}{
			"url": "/api/subject/getThemesByTickers.json",
			"args": map[string]string{
				"secID":        "string", //证券内部编码，可通过交易代码和交易市场在getSecID获取到，可多值输入，如："000001.XSHE"、"000001.XSHE,600001.XSHG"
				"secShortName": "string", //证券简称，支持模糊查询，如"浦发"、"浦发银行"
				"ticker":       "string", //证券交易代码，允许多值输入，用","分隔，如"000001"、"000001,600001"。
				"beginDate":    "date",   //查询开始日期，默认为当天
				"endDate":      "date",   //查询结束日期，默认为当天
				"exchangeCD":   "string", //证券交易所代码(通联自编)。可选：XSHG、XSHE、XHKG、XNYS、NOBB、XNAS。XSHG表示上海证券交易所，XSHE表示深圳证券交易所，XHKG表示香港交易所，XNYS表示纽约证券交易所，NOBB表示三板市场，XNYS表示美国纳斯达克证券交易所。可多值输入
				"field":        "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":                 "string", //证券内部编码(通联自编)
				"exchangeCD":            "string", //证券交易所代码(通联自编)
				"ticker":                "string", //证券交易代码
				"beginDate":             "date",   //关联开始日期
				"themeID":               "int64",  //主题ID
				"exchangeName":          "string", //证券交易所
				"secShortName":          "string", //证券简称
				"themeName":             "string", //主题名称
				"endDate":               "date",   //关联结束日期，日期为空表示关联至今依旧存在
				"returnScore":           "double", //收益关联程度，主题和证券在短期收益上的相似度，取值范围[0，1]，值越大表示关联度越高
				"textContributionScore": "double", //文本贡献关联度，主题和证券在新闻文本中的相似度，取值范围[0，1]，值越大表示关联度越高
				"industryScore":         "double", //行业关联度，主题和证券在行业分布上的相似度，取值范围[0，1]，值越大表示关联度越高
				"insertTime":            "date",   //入库时间
				"updateTime":            "date",   //更新时间
			},
			"indices": map[string]string{},
		},
		//沪深股票信息
		//说明获取股票每日回报率的基本信息，包含交易当天的上市状态、日行情以及除权除息事项的基本数据。
		"getEquRetud": map[string]interface{}{
			"url": "/api/equity/getEquRetud.json",
			"args": map[string]string{
				"secID":                   "string", //证券内部编码，可通过交易代码和交易市场在getSecurityID获取到。（可多值输入）
				"ticker":                  "string", //股票交易代码，如'000001'（可多值输入）
				"beginDate":               "date",   //交易日期，可以选择某一固定的交易日或者某一区间的交易日
				"dailyReturnNoReinvLower": "double", //回报率，可查询回报率大于等于或者小于等于某一数值的记录
				"dailyReturnNoReinvUpper": "double", //回报率，可查询回报率大于等于或者小于等于某一数值的记录
				"dailyReturnReinvLower":   "double", //回报率，可查询回报率大于等于或者小于等于某一数值的记录
				"dailyReturnReinvUpper":   "double", //回报率，可查询回报率大于等于或者小于等于某一数值的记录
				"endDate":                 "date",   //交易日期，可以选择某一固定的交易日或者某一区间的交易日
				"isChgPctl":               "int16",  //交易日当天是否有涨跌幅限制，0——设置涨跌幅限制；1——IPO(包括借壳重组上市)不设置涨跌幅限制；2——股改完成恢复交易首日不设置涨跌幅限制；3——交易所规定的交易当天不设置涨跌幅限制的其他情况，默认为0。（可多值输入）
				"listStatusCD":            "string", //上市状态，选择交易日是否停牌，可选状态有L——上市，S——暂停，默认为L。（可多值输入）
				"field":                   "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":              "string", //证券内部ID
				"tradeDate":          "date",   //交易日
				"listStatusCD":       "string", //上市状态
				"dailyReturnReinv":   "double", //日回报率(考虑现金红利再投资)
				"dailyReturnNoReinv": "double", //日回报率(不考虑现金红利再投资)
				"ticker":             "string", //交易代码
				"exchangeCD":         "string", //交易市场代码
				"secShortName":       "string", //证券简称
				"currencyCD":         "string", //货币种类
				"isChgPctl":          "int16",  //是否涨跌幅限制
			},
			"indices": map[string]string{},
		},
		//说明沪深融资融券每日交易明细信息
		"getFstDetail": map[string]interface{}{
			"url": "/api/equity/getFstDetail.json",
			"args": map[string]string{
				"secID":     "string", //证券内部编码，可通过交易代码和交易市场在getSecID获取到。（可多值输入）
				"ticker":    "string", //交易代码，如'000001'（可多值输入）
				"beginDate": "date",   //起始日期
				"endDate":   "date",   //截止日期
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"tradeDate":    "date",   //交易日期
				"ticker":       "string", //交易代码
				"secID":        "string", //证券内部ID
				"assetClass":   "string", //证券类型，E表示股票，F表示基金。
				"exchangeCD":   "string", //交易市场
				"secShortName": "string", //证券简称
				"currencyCD":   "string", //交易货币代码
				"finVal":       "double", //本日融资余额
				"finBuyVal":    "double", //本日融资买入额
				"finRefundVal": "double", //本日融资偿还额
				"secVol":       "double", //本日融券余量
				"secSellVol":   "double", //本日融券卖出量
				"secRefundVol": "double", //本日融券偿还量
				"secVal":       "double", //本日融券余量金额
				"tradeVal":     "double", //本日融资融券余额
			},
			"indices": map[string]string{},
		},
		//说明获取上市公司股本结构及历次股本变动数据。
		"getEquShare": map[string]interface{}{
			"url": "/api/equity/getEquShare.json",
			"args": map[string]string{
				"secID":     "string", //证券ID，证券统一编码，可通过交易代码在getSecID获取到。（可多值输入）
				"ticker":    "string", //回购交易代码，如'204001'（可多值输入）
				"beginDate": "date",   //起始日期，默认为一年前当前日期
				"endDate":   "date",   //截止日期，默认为当前日期
				"partyID":   "string", //机构ID，法人统一编码，可通过法人名称在getPartyID获取到。（可多值输入）
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"changeDate":         "date",   //变更日期
				"secID":              "string", //证券ID
				"exchangeCD":         "string", //交易市场
				"ticker":             "string", //交易代码
				"secShortName":       "string", //证券简称
				"partyID":            "int64",  //机构ID
				"totalShares":        "double", //总股本
				"sharesA":            "double", //A股
				"sharesB":            "double", //B股
				"floatA":             "double", //流通A股
				"nonrestfloatA":      "double", //无限售流通A股
				"floatB":             "double", //流通B股
				"restShares":         "double", //有限售条件股份合计
				"nonrestFloatShares": "double", //无限售流通股份合计
			},
			"indices": map[string]string{},
		},
		//说明通过输入股票ID（通联编制）或股票交易代码（支持多值输入，最大支持50只），选择查询开始日期与结束日期，获取股票在一段时间ST标记信息。
		"getSecST": map[string]interface{}{
			"url": "/api/equity/getSecST.json",
			"args": map[string]string{
				"secID":     "string", //证券内部编码，可通过交易代码和交易市场在getSecID获取到。
				"ticker":    "string", //股票交易代码，如'000001'
				"beginDate": "date",   //查询开始日期，输入格式“YYYYMMDD”
				"endDate":   "date",   //查询结束日期，输入格式“YYYYMMDD”
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"ticker":        "string", //交易代码
				"tradeDate":     "date",   //日期
				"secID":         "string", //证券ID
				"exchangeCD":    "string", //交易市场
				"tradeAbbrName": "string", //证券简称
				"STflg":         "string", //ST标记，S*ST-公司经营连续三年亏损，退市预警+还没有完成股改;*ST-公司经营连续三年亏损，退市预警;ST-公司经营连续二年亏损，特别处理;SST-公司经营连续二年亏损，特别处理+还没有完成股改;S-还没有完成股改
			},
			"indices": map[string]string{},
		},
		//说明查询获取沪港通合资格股票名单，及历史调入、调出信息。
		"getEquSHHKCons": map[string]interface{}{
			"url": "/api/equity/getEquSHHKCons.json",
			"args": map[string]string{
				"consTypeCD": "string", //输入沪港通类别：01-港股通
				"intoDate":   "date",   //输入日期，可以获取这一天沪港通合资格股票名单，输入格式“YYYYMMDD”
				"isNew":      "int",    //是否最新：1-是，0-否
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"ticker":       "string", //交易代码
				"consTypeCD":   "string", //沪港通类别
				"secID":        "string", //证券ID
				"secShortName": "string", //证券简称
				"secFullName":  "string", //证券全称
				"intoDate":     "date",   //调入日期
				"outDate":      "date",   //调出日期
				"isNew":        "int",    //是否最新
			},
			"indices": map[string]string{},
		},
		//说明输入证券ID或股票交易代码，获取股票所属行业分类
		"getEquIndustry": map[string]interface{}{
			"url": "/api/equity/getEquIndustry.json",
			"args": map[string]string{
				"secID":             "string", //证券内部编码，可通过交易代码和交易市场在getSecurityID获取到。（可多值输入）
				"ticker":            "string", //股票交易代码，如"000001"（可多值输入）
				"intoDate":          "string", //输入日期，可以获取这一天股票的行业信息
				"industry":          "string", //行业分类标准，模糊查询，如输入“申万”或"申万行业"
				"industryVersionCD": "string", //010301-证监会行业V2012、
				"field":             "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":             "string", //证券ID
				"industry":          "string", //行业分类标准
				"industryID1":       "string", //一级行业编码
				"industryID2":       "string", //二级行业编码
				"industryID3":       "string", //三级行业编码
				"IndustryID4":       "string", //四级行业编码
				"ticker":            "string", //交易代码
				"exchangeCD":        "string", //交易市场
				"secShortName":      "string", //证券简称
				"secFullName":       "string", //证券全称
				"partyID":           "int64",  //机构ID
				"industryVersionCD": "string", //行业分类标准编码
				"industryID":        "string", //行业分类编码
				"industrySymbol":    "string", //行业分类编码，行业编制机构发布的行业编码
				"intoDate":          "date",   //成分纳入日期
				"outDate":           "date",   //成分剔除日期
				"isNew":             "int16",  //是否最新
				"industryName1":     "string", //一级行业
				"industryName2":     "string", //二级行业
				"industryName3":     "string", //三级行业
				"IndustryName4":     "string", //四级行业
			},
			"indices": map[string]string{},
		},
		//说明获取股票首次公开发行上市的基本信息，包含股票首次公开发行的进程及发行结果。
		"getEquIPO": map[string]interface{}{
			"url": "/api/equity/getEquIPO.json",
			"args": map[string]string{
				"secID":          "string", //证券内部编码，可通过交易代码和交易市场在getSecurityID获取到。（可多值输入）
				"ticker":         "string", //股票交易代码，如'000001'（可多值输入）
				"eventProcessCD": "int64",  //（可多值输入）
				"field":          "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":                 "string", //证券内部ID
				"eventProcessCD":        "int64",  //发行进程
				"ticker":                "string", //交易代码
				"exchangeCD":            "string", //交易市场
				"secShortName":          "string", //证券简称
				"publishDate":           "date",   //公告日期
				"onlineIssueDate":       "date",   //网上发行日期
				"listDate":              "date",   //上市日期
				"issuePrice":            "double", //发行价格
				"currencyCD":            "string", //货币种类
				"issueShares":           "double", //股票发行量
				"newIssueShares":        "double", //新股发行量
				"transShares":           "string", //老股转让数量
				"newIssueRaiseCap":      "double", //新股募集资金
				"OldShareRaiseCap":      "double", //老股募集资金
				"issuePE":               "double", //发行市盈率
				"onlineIssueLottoRatio": "double", //中签率
				"firstDayOpenPrice":     "double", //上市首日开盘价
				"firstDayClosePrice":    "double", //上市首日收盘价
			},
			"indices": map[string]string{},
		},
		//说明获取股票股权分置改革的基本信息，包含股改进程、股改实施方案以及流通股的变动情况。
		"getEquRef": map[string]interface{}{
			"url": "/api/equity/getEquRef.json",
			"args": map[string]string{
				"secID":          "string", //证券内部编码，可通过交易代码和交易市场在getSecurityID获取到。（可多值输入）
				"ticker":         "string", //股票交易代码，如'000001'（可多值输入）
				"beginDate":      "date",   //股权分置改革方案公告日期，可以选择某一固定日期或者某一区间日期
				"endDate":        "date",   //股权分置改革方案公告日期，可以选择某一固定日期或者某一区间日期
				"eventProcessCD": "int64",  //股权分置改革方案实施进程，可选：7——停止实施，3——股东大会否决，6——方案实施，默认为6。（可多值输入）
				"field":          "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":              "string", //证券内部ID
				"publishDate":        "date",   //公告日期
				"ticker":             "string", //交易代码
				"exchangeCD":         "string", //交易市场
				"secShortName":       "string", //证券简称
				"eventProcessCD":     "int64",  //事件进程
				"perShareDivRatio":   "double", //每股送股比例
				"perShareTransRatio": "double", //每股转增股比例
				"perCashDiv":         "double", //每股派现
				"currencyCD":         "string", //货币种类
				"recordDate":         "date",   //股权登记日
				"afFirstTradeDate":   "date",   //股改后交易首日
				"bfShares":           "double", //股改前总股本
				"afShares":           "double", //股改后总股本
				"bfTradeShares":      "double", //股改前流通股股数
				"bfTradeSharesRatio": "double", //股改前流通股比例
			},
			"indices": map[string]string{},
		},
		//说明获取股票进行股本拆细或者缩股的基本信息。
		"getEquSplits": map[string]interface{}{
			"url": "/api/equity/getEquSplits.json",
			"args": map[string]string{
				"secID":     "string", //证券内部编码，可通过交易代码和交易市场在getSecurityID获取到。（可多值输入）
				"ticker":    "string", //股票交易代码，如'000001'（可多值输入）
				"beginDate": "date",   //股本分拆方案公告日期，可以选择某一固定日期或者某一区间日期
				"endDate":   "date",   //股本分拆方案公告日期，可以选择某一固定日期或者某一区间日期
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":          "string", //内部编码
				"publishDate":    "date",   //公告日期
				"ticker":         "string", //证券代码
				"exchangeCD":     "string", //交易市场
				"secshortName":   "string", //证券简称
				"splitsRatio":    "double", //拆股率
				"splitsBaseDate": "date",   //基准日
				"reTradeDate":    "date",   //拆股后首个交易日
				"sharesBfSplits": "double", //拆股前总股本
				"sharesAfSplits": "double", //拆股后总股本
			},
			"indices": map[string]string{},
		},
		//说明获取上海、深圳交易所公布的每个交易日的融资融券交易汇总的信息，包括成交量、成交金额。本交易日可获取前一交易日的数据。
		"getFstTotal": map[string]interface{}{
			"url": "/api/equity/getFstTotal.json",
			"args": map[string]string{
				"beginDate":  "date",   //起始日期
				"endDate":    "date",   //截止日期
				"exchangeCD": "string", //交易市场,可选市场：XSHG，XSHE；XSHG表示上海证券交易所,XSHE表示深圳证券交易所。（可多值输入）
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"tradeDate":  "date",   //交易日期
				"exchangeCD": "string", //交易市场
				"currencyCD": "string", //交易货币代码
				"finVal":     "double", //本日融资余额
				"finBuyVal":  "double", //本日融资买入额
				"secVol":     "double", //本日融券余量
				"secVal":     "double", //本日融券余量金额
				"secSellVol": "double", //本日融券卖出量
				"tradeVal":   "double", //本日融资融券余额
				"finRatio":   "double", //本日融资占融资融券余额比
			},
			"indices": map[string]string{},
		},
		//说明获取股票的基本信息，包含股票交易代码及其简称、股票类型、上市状态、上市板块、上市日期等；上市状态为最新数据，不显示历史变动信息。
		"getEqu": map[string]interface{}{
			"url": "/api/equity/getEqu.json",
			"args": map[string]string{
				"secID":        "string", //证券内部编码，可通过交易代码和交易市场在getSecurityID获取到。（可多值输入）
				"ticker":       "string", //股票交易代码，如'000001'（可多值输入）
				"equTypeCD":    "string", //股票分类编码，输入A或B可查询获取到A股或B股
				"listStatusCD": "string", //上市状态，可选状态有L——上市，S——暂停，DE——已退市，UN——未上市，默认为L。（可多值输入）
				"field":        "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":              "string", //证券ID
				"ticker":             "string", //交易代码
				"exchangeCD":         "string", //交易市场
				"ListSectorCD":       "int32",  //上市板块编码
				"ListSector":         "string", //上市板块
				"transCurrCD":        "string", //交易货币
				"secShortName":       "string", //证券简称
				"secFullName":        "string", //证券全称
				"listStatusCD":       "string", //上市状态
				"listDate":           "date",   //上市日期
				"delistDate":         "date",   //摘牌日期
				"equTypeCD":          "string", //股票分类编码
				"equType":            "string", //股票类别
				"exCountryCD":        "string", //交易市场所属地区
				"partyID":            "int64",  //机构内部ID
				"totalShares":        "double", //总股本(最新)
				"nonrestFloatShares": "Double", //公司无限售流通股份合计(最新)
				"nonrestfloatA":      "double", //无限售流通股本(最新)。如果为A股，该列为最新无限售流通A股股本数量；如果为B股，该列为最新流通B股股本数量
				"officeAddr":         "String", //办公地址
				"primeOperating":     "text",   //主营业务范围
				"endDate":            "Date",   //财务报告日期
				"TShEquity":          "double", //所有者权益合计
			},
			"indices": map[string]string{},
		},
		//说明获取股票历次配股的基本信息，包含每次配股方案的内容、方案进度、历史配股预案公布次数以及最终是否配股成功。
		"getEquAllot": map[string]interface{}{
			"url": "/api/equity/getEquAllot.json",
			"args": map[string]string{
				"secID":       "string", //证券内部编码，可通过交易代码和交易市场在getSecurityID获取到。（可多值输入）
				"ticker":      "string", //股票交易代码，如'000001'（可多值输入）
				"beginDate":   "date",   //配股预案的起始公布日期
				"endDate":     "date",   //配股预案的结束公布日期
				"isAllotment": "int16",  //是否最终配股成功，0——不成功，1——成功，默认为1。
				"field":       "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":          "string", //证券内部ID
				"publishDate":    "date",   //公告日期
				"ticker":         "string", //交易代码
				"exchangeCD":     "string", //交易市场
				"secShortName":   "string", //证券简称
				"isAllotment":    "int16",  //是否配股
				"allotmentRatio": "double", //每股配股比例
				"allotmentPrice": "double", //配股价
				"currencyCD":     "string", //货币种类
				"allotFrPrice":   "double", //外币配股价
				"frCurrencyCD":   "string", //外币货币种类
				"recordDate":     "date",   //股权登记日
				"exRightsDate":   "date",   //除权日
				"payBeginDate":   "date",   //配股缴款起始日
				"payEndDate":     "date",   //配股缴款截止日
				"listDate":       "date",   //配股上市日
				"allotShares":    "double", //配股数量
				"sharesBfAllot":  "double", //配股前总股本
				"sharesAfAllot":  "double", //配股后总股本
			},
			"indices": map[string]string{},
		},
		//说明获取股票历次分红(派现、送股、转增股)的基本信息，包含历次分红预案的内容、实施进展情况以及历史宣告分红次数。
		"getEquDiv": map[string]interface{}{
			"url": "/api/equity/getEquDiv.json",
			"args": map[string]string{
				"secID":          "string", //证券内部编码，可通过交易代码和交易市场在getSecurityID获取到。（可多值输入）
				"ticker":         "string", //股票交易代码，如'000001'（可多值输入）
				"beginDate":      "date",   //起始日期，可空，输入格式“YYYYMMDD”。若起始日期和截止日期区间，包含本次分红所属的财政年度、财政半年度、财政季度的最后一个自然日，则返回该财政区间的所有分红。如输入起始日期‘20130101’，截止日期‘20131231’，可以查询到2013年期间所有分红；如输入起始日期‘20121231’，截止日期‘20131231’，可以查询到2012年和2013年期间所有分红。
				"endDate":        "date",   //截止日期，可空，输入格式“YYYYMMDD”。若起始日期和截止日期区间，包含本次分红所属的财政年度、财政半年度、财政季度的最后一个自然日，则返回该财政区间的所有分红。如输入起始日期‘20130101’，截止日期‘20131231’，可以查询到2013年期间所有分红；如输入起始日期‘20121231’，截止日期‘20131231’，可以查询到2012年和2013年期间所有分红。
				"eventProcessCD": "int64",  //分红方案的实施进程，可选：1-董事会预案，2-股东大会通过，3-股东大会否决，6-方案实施，7-停止实施。（可多值输入）
				"exDivDate":      "date",   //输入一个日期，格式“YYYYMMDD”，可以获取到这个输入日期当天的除权分红信息
				"field":          "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":              "string", //内部编码
				"endDate":            "date",   //分红年度截止日
				"ticker":             "string", //交易代码
				"exchangeCD":         "string", //交易市场
				"secShortName":       "string", //证券简称
				"publishDate":        "date",   //公告日期
				"eventProcessCD":     "int64",  //事件进程
				"perShareDivRatio":   "double", //每股送股比例
				"perShareTransRatio": "double", //每股转增股比例
				"perCashDiv":         "double", //每股派现（税前）
				"currencyCD":         "string", //货币种类
				"frPerCashDiv":       "double", //每股派现(税前)外币结算
				"frCurrencyCD":       "string", //外币货币种类
				"recordDate":         "date",   //股权登记日
				"exDivDate":          "date",   //除权除息日
				"bLastTradeDate":     "date",   //B股最后交易日
				"payCashDate":        "date",   //派现日
				"bonusShareListDate": "date",   //红股上市日
				"ftdAfExdiv":         "string", //除权除息后交易首日
				"sharesBfDiv":        "double", //分红前总股本
				"sharesAfDiv":        "double", //分红后总股本
			},
			"indices": map[string]string{},
		},
		//港股信息
		//说明获取香港交易所上市股票的基本信息，包含股票交易代码及其简称、股票类型、上市状态、上市板块、上市日期等；上市状态为最新状态。
		"getHKEqu": map[string]interface{}{
			"url": "/api/HKequity/getHKEqu.json",
			"args": map[string]string{
				"secID":        "string", //证券ID，证券统一数字编码，可通过交易代码在getSecID获取到。（可多值输入）
				"ticker":       "string", //证券交易代码，如'00001'（可多值输入）
				"listStatusCD": "string", //上市状态，可选状态有:L 上市，S 暂停，DE 已退市，UN 未上市，默认为L。
				"field":        "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":        "string", //证券ID
				"ticker":       "string", //证券交易代码
				"listStatusCD": "string", //上市状态
				"ListSectorCD": "int32",  //上市板块编码
				"ListSector":   "string", //上市板块
				"equTypeCD":    "string", //股票分类编码
				"equType":      "string", //股票类别
				"transCurrCD":  "string", //交易货币
				"tradingUnit":  "string", //买卖单位
				"secFullName":  "string", //证券全称
				"secShortName": "string", //证券简称
				"exCountryCD":  "string", //交易市场所属地区
				"listDate":     "date",   //上市日期
				"delistDate":   "date",   //摘牌日期
				"partyID":      "int64",  //上市公司ID
			},
			"indices": map[string]string{},
		},
		//说明获取香港交易所上市公司行为，包含有首发、现金增资、分红、拆细等。
		"getHKEquCA": map[string]interface{}{
			"url": "/api/HKequity/getHKEquCA.json",
			"args": map[string]string{
				"secID":       "string", //证券ID，证券统一数字编码，可通过交易代码在getSecID获取到。（可多值输入）
				"ticker":      "string", //证券交易代码，如'00001'（可多值输入）
				"eventTypeCD": "string", //公司行为类别，可选行为类别:1 首发，2 现金增资(不含供股)，3 现金增资(供股)，4 拆股，5 并股，6 回购，7 现金分红，8 股票分红；默认为7 现金分红。（可多值输入）
				"field":       "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":        "string", //证券ID
				"publishDate":  "date",   //公告发布日
				"ticker":       "string", //交易代码
				"exchangeCD":   "string", //交易市场
				"secShortName": "string", //证券简称
				"eventNum":     "int32",  //事件序号
				"eventTypeCD":  "string", //公司行为类别编码
				"eventType":    "string", //公司行为类别
				"statusCD":     "string", //事件状态编码
				"status":       "string", //事件状态
				"newShares":    "double", //新股数量
				"oldShares":    "double", //老股数量
				"currencyCD":   "string", //货币
				"cash":         "double", //现金数量
				"corShares":    "double", //关联股票数量
				"exePrice":     "double", //执行价格
				"recordDate":   "date",   //登记日
				"exDate":       "date",   //除权日
			},
			"indices": map[string]string{},
		},
		//说明沪港通额度信息，包括沪股通和港股通总额度、总额度余额、每日额度以及每日额度余额的信息。
		"getequSHHKQuota": map[string]interface{}{
			"url": "/api/HKequity/getequSHHKQuota.json",
			"args": map[string]string{
				"exchangeCD": "string", //交易类型，XSSC表示港股通，SHSC表示沪股通。
				"tradeDate":  "date",   //交易日期
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"exchangeCD":     "string", //交易类型
				"tradeDate":      "date",   //交易日期
				"aggregateQuota": "double", //总额度
				"closingBalance": "double", //总额度余额
				"dailyQuota":     "double", //每日额度
				"dailyBalance":   "double", //每日额度余额
				"currCD":         "string", //货币代码，CNY表示人民币，HKD表示港币。
			},
			"indices": map[string]string{},
		},
		//期权信息
		//说明获取期权品种名称、生效日期、履约方式、交割方式、申报单位等相关信息。
		"getOptVar": map[string]interface{}{
			"url": "/api/options/getOptVar.json",
			"args": map[string]string{
				"exchangeCD":   "string", //标的交易市场,如:XSHG 上海证券交易所
				"secID":        "string", //标的内部编码，与getOpt表VAR_UNI_CODE字段关联可获取该品种对应的期权合约
				"ticker":       "string", //标的交易代码，如:601318 中国平安
				"contractType": "string", //合约类型:CO 认购期权;PO 认沽期权;CP 认购、认沽期权
				"exerType":     "string", //期权履约方式:A 美式;E 欧式
				"field":        "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"ticker":         "string", //标的交易代码
				"effDate":        "date",   //生效日期
				"secID":          "string", //标的内部编码
				"secShortName":   "string", //标的简称
				"exchangeCD":     "string", //标的交易市场
				"contractType":   "string", //合约类型
				"exerType":       "string", //履约方式
				"contMultNum":    "int32",  //合约单位
				"expMonthDesc":   "string", //到期月份描述
				"tradeTimeDesc":  "string", //交易时间描述
				"exerTimeDesc":   "string", //行权日行权时间描述
				"lastTdateDesc":  "string", //最后交易日描述
				"expDateDesc":    "string", //到期日描述
				"exerDateDesc":   "string", //行权日描述
				"deliMethod":     "string", //交割方式
				"tradeCommiNum":  "double", //交易手续费(数值)
				"tradeCommiUnit": "string", //交易手续费(单位)
				"tickNum":        "double", //最小报价变动(数值)
				"tickUnit":       "string", //最小报价变动(单位)
				"contUnit":       "string", //申报单位
			},
			"indices": map[string]string{},
		},
		//说明获取期权合约每日涨幅上限价格、跌幅下限价格、单位保证金等相关信息。
		"getOptDpo": map[string]interface{}{
			"url": "/api/options/getOptDpo.json",
			"args": map[string]string{
				"optID":        "string", //合约编码,由交易所给出，标记唯一期权合约
				"secID":        "string", //合约内部编码,通联编制的证券唯一标识编码
				"ticker":       "string", //合约交易代码,如:510180P1506M03500
				"varSecID":     "string", //标的内部编码,如:601318.XSHG 中国平安
				"varTicker":    "string", //标的交易代码，如:601318 中国平安
				"varType":      "string", //品种类别:E 股票期权;F ETF期权
				"beginDate":    "date",   //交易日期，查询开始日期，输入格式“YYYYMMDD”
				"contractType": "string", //合约类型:CO 认购期权;PO 认沽期权;CP 认购、认沽期权
				"endDate":      "date",   //交易日期，查询截止日期，输入格式“YYYYMMDD”
				"exerType":     "string", //期权履约方式:A 美式;E 欧式
				"field":        "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"tradeDate":        "date",   //交易日期
				"optID":            "string", //合约编码
				"varTicker":        "string", //标的交易代码
				"secID":            "string", //合约内部编码
				"secShortName":     "string", //合约简称
				"ticker":           "string", //合约交易代码
				"exchangeCD":       "string", //交易所
				"varSecID":         "string", //标的内部编码
				"varName":          "string", //标的名称
				"varExchangeCD":    "string", //标的交易市场
				"varType":          "string", //品种类别
				"exerType":         "string", //履约方式
				"contractType":     "string", //合约类型
				"contMultNum":      "int32",  //合约单位
				"strikePrice":      "double", //行权价格
				"listDate":         "date",   //挂牌日期
				"lastTradeDate":    "date",   //最后交易日
				"exerDate":         "date",   //行权日
				"deliDate":         "date",   //交收日期
				"expDate":          "date",   //到期日
				"contVersion":      "string", //合约版本号
				"openInt":          "int32",  //持仓量
				"preClosePrice":    "double", //前收盘价
				"preSettPrice":     "double", //前结算价
				"varpreClosePrice": "double", //标的证券前收盘
				"isChgPctl":        "string", //涨跌幅限类型
				"limitUpNum":       "double", //涨幅上限价格
				"limitDownNum":     "double", //跌幅下限价格
				"marginUnit":       "double", //单位保证金
				"marginRatioP1":    "double", //保证金计算比例参数一
				"marginRatioP2":    "double", //保证金计算比例参数二
				"roundLot":         "int32",  //整手数
				"lmtOrdMin":        "int32",  //单笔限价申报下限
				"lmtOrdMax":        "int32",  //单笔限价申报上限
				"mktOrdMin":        "int32",  //单笔市价申报下限
				"mktOrdMax":        "int32",  //单笔市价申报上限
				"tickNum":          "double", //最小报价变动(数值)
				"contractStatus":   "string", //合约状态信息
			},
			"indices": map[string]string{},
		},
		//说明获取期权合约编码，交易代码，交易市场，标的等相关信息
		"getOpt": map[string]interface{}{
			"url": "/api/options/getOpt.json",
			"args": map[string]string{
				"optID":          "string", //合约编码，有交易所给出，标记唯一期权合约
				"secID":          "string", //证券ID，通联编制的证券唯一标识编码
				"ticker":         "string", //合约交易代码
				"varSecID":       "string", //标的证券ID
				"varticker":      "string", //标的交易代码
				"contractStatus": "string", //合约状态，UN-未上市、L-上市、S-暂停上市、DE-退市
				"field":          "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"optID":          "string", //合约编码
				"secID":          "string", //合约内部编码
				"secShortName":   "string", //合约简称
				"tickerSymbol":   "string", //合约交易代码
				"exchangeCD":     "string", //交易所
				"varSecID":       "string", //标的内部编码
				"varShortName":   "string", //标的简称
				"varTicker":      "string", //标的交易代码
				"varExchangeCD":  "string", //标的交易市场
				"varType":        "string", //标的品种类别
				"contractType":   "string", //合约类型
				"strikePrice":    "double", //行权价格
				"contMultNum":    "int32",  //合约单位
				"contractStatus": "string", //合约状态
				"listDate":       "date",   //挂牌日期
				"expYear":        "string", //到期年
				"expMonth":       "string", //到期月
				"expDate":        "date",   //到期日
				"lastTradeDate":  "date",   //最后交易日
				"exerDate":       "date",   //行权日
				"deliDate":       "date",   //交收日期
				"delistDate":     "date",   //摘牌日期
			},
			"indices": map[string]string{},
		},
		//债券信息
		//说明固定利率债券、浮动利率债券每个计息周期的票面利率，包括分段计息的具体利率。
		"getBondCoupon": map[string]interface{}{
			"url": "/api/bond/getBondCoupon.json",
			"args": map[string]string{
				"secID":  "string", //证券内部编码，可通过交易代码和交易市场在getSecurityID获取到。（可多值输入）
				"ticker": "string", //交易代码（可多值输入）
				"field":  "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":           "string", //证券ID
				"perValueDate":    "date",   //本段计息起始日
				"ticker":          "string", //交易代码
				"exchangeCD":      "string", //交易市场
				"secShortName":    "string", //债券简称
				"bondID":          "string", //债券ID
				"perValueEndDate": "date",   //本段计息截止日
				"refRatePer":      "double", //参考利率(%)
				"refRateDate":     "date",   //浮动利率计息基准日
				"stepMargin":      "double", //累进利差(%)
				"coupon":          "double", //票面年利率(%)
				"perAccrDate":     "date",   //本周期计息起始日
				"perAccrEndDate":  "date",   //本周期计息截止日
			},
			"indices": map[string]string{},
		},
		//说明债券在一级市场发行信息，记录的要素包括有发行方式、发行价格、当次发行规模等
		"getBondIssue": map[string]interface{}{
			"url": "/api/bond/getBondIssue.json",
			"args": map[string]string{
				"secID":       "string", //证券内部编码，可通过交易代码和交易市场在getSecurityID获取到。（可多值输入）
				"ticker":      "string", //交易代码（可多值输入）
				"raiseModeCD": "string", //募集方式，支持公开PUB和私募PRIV，默认为PUB
				"field":       "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":          "string", //证券ID
				"issueDate":      "date",   //发行日期
				"ticker":         "string", //交易代码
				"exchangeCD":     "string", //交易市场
				"bondID":         "string", //债券ID
				"secShortName":   "string", //债券简称
				"raiseModeCD":    "string", //募集方式
				"issueNum":       "int16",  //发行次数
				"issuePrice":     "double", //发行价格(元)
				"issueSize":      "double", //本次发行总额(亿元)
				"ytmAtIssue":     "double", //发行参考收益率(%)
				"auctionDate":    "date",   //招标日期
				"firstSettlDate": "date",   //划款日
				"recordDate":     "date",   //债权债务登记日
			},
			"indices": map[string]string{},
		},
		//说明记录债券在发行时约定在存续期内投资人或发行人可行使的选择权，如债券回售、债券赎回等。
		"getBondOption": map[string]interface{}{
			"url": "/api/bond/getBondOption.json",
			"args": map[string]string{
				"secID":  "string", //证券内部编码，可通过交易代码和交易市场在getSecurityID获取到。（可多值输入）
				"ticker": "string", //交易代码（可多值输入）
				"field":  "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":        "string", //证券ID
				"ticker":       "string", //交易代码
				"exchangeCD":   "string", //交易市场
				"bondID":       "string", //债券ID
				"secShortName": "string", //债券简称
				"optionCD":     "string", //选择权
			},
			"indices": map[string]string{},
		},
		//说明债券的长期评级、短期评级以及历史评级变动记录。
		"getBondRating": map[string]interface{}{
			"url": "/api/bond/getBondRating.json",
			"args": map[string]string{
				"secID":     "string", //证券内部编码，可通过交易代码和交易市场在getSecurityID获取到。（可多值输入）
				"ticker":    "string", //交易代码（可多值输入）
				"beginDate": "date",   //评级起始日期
				"endDate":   "date",   //评级结束日期
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":        "string", //证券ID
				"ratingDate":   "date",   //评级日期
				"ticker":       "string", //交易代码
				"exchangeCD":   "string", //交易市场
				"bondID":       "string", //债券ID
				"secShortName": "string", //债券简称
				"publishDate":  "date",   //公告日期
				"rating":       "string", //级别
				"ratComID":     "int16",  //评级公司ID
				"ratCom":       "string", //评级公司全称
				"ratTypeCD":    "string", //评级类型
			},
			"indices": map[string]string{},
		},
		//说明各评级机构公布的债券担保人的信用评级数据及历史变动记录。
		"getGuarRating": map[string]interface{}{
			"url": "/api/bond/getGuarRating.json",
			"args": map[string]string{
				"secID":     "string", //证券内部编码，可通过交易代码和交易市场在getSecurityID获取到。（可多值输入）
				"ticker":    "string", //交易代码（可多值输入）
				"beginDate": "date",   //评级起始日期
				"endDate":   "date",   //评级结束日期
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":         "string", //证券ID
				"ratingDate":    "date",   //评级日期
				"ticker":        "string", //交易代码
				"exchangeCD":    "string", //交易市场
				"secShortName":  "string", //债券简称
				"partyID":       "int64",  //担保人ID
				"partyFullName": "string", //担保人
				"rating":        "string", //级别
				"publishDate":   "date",   //评级公告发布日期
				"ratComID":      "int16",  //评级公司ID
				"ratCom":        "string", //评级公司
				"anticipateCD":  "int16",  //评级展望
			},
			"indices": map[string]string{},
		},
		//说明债券回购基本面信息，涵盖了回购交易代码、回购期限、申报价格最小变动单位(RMB)、申报参与价格单位(RMB)等。
		"getRepo": map[string]interface{}{
			"url": "/api/bond/getRepo.json",
			"args": map[string]string{
				"secID":  "string", //证券ID，证券统一编码，可通过交易代码在getSecID获取到。（可多值输入）
				"ticker": "string", //回购交易代码，如'204001'（可多值输入）
				"field":  "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"ticker":           "string", //回购代码
				"secID":            "string", //证券ID
				"exchangeCD":       "string", //交易市场
				"exCountryCD":      "string", //交易市场所属地区
				"secFullName":      "string", //回购全称
				"secShortName":     "string", //回购简称
				"maturity":         "int16",  //回购期限
				"repoTypeCD":       "string", //回购类型编码
				"repoType":         "string", //回购类型
				"interestBaseDays": "int16",  //全年计息基准天数
				"minQuoteUnit":     "double", //申报价格最小变动单位(RMB)
				"priceUnit":        "int32",  //申报参与价格单位(RMB)
				"listStatusCD":     "String", //上市状态：L-上市，S-暂停，DE-终止上市，UN-未上市。
			},
			"indices": map[string]string{},
		},
		//说明记录债券的存量变动情况。
		"getBondSizeChg": map[string]interface{}{
			"url": "/api/bond/getBondSizeChg.json",
			"args": map[string]string{
				"secID":  "string", //证券ID，通联数据编制，全商城唯一标识证券编码，可通过交易代码和交易市场在getSecID获取到。
				"ticker": "string", //债券交易代码，输入债券在交易所交易代码。
				"field":  "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"bondID":         "string", //债券ID
				"secID":          "string", //证券ID
				"ticker":         "string", //交易代码
				"secShortName":   "string", //债券简称
				"publishDate":    "string", //公告日期
				"changeDate":     "string", //截止日期
				"changeSize":     "double", //本次变动规模
				"remainSize":     "double", //存量
				"changeReasonCD": "string", //变动原因
			},
			"indices": map[string]string{},
		},
		//说明债券在每个计息周期付息兑付的相关数据。日期区间默认为过去一年。
		"getBondCf": map[string]interface{}{
			"url": "/api/bond/getBondCf.json",
			"args": map[string]string{
				"secID":      "string", //证券内部编码，可通过交易代码和交易市场在getSecurityID获取到。（可多值输入）
				"ticker":     "string", //交易代码（可多值输入）
				"beginDate":  "date",   //现金流起始日，默认为1年前
				"cashTypeCD": "string", //现金流类型，默认为INTPAY。（可多值输入）
				"endDate":    "date",   //现金流截止日
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":          "string", //证券ID
				"paymentDate":    "date",   //现金流日期
				"ticker":         "string", //交易代码
				"exchangeCD":     "string", //交易市场
				"bondID":         "string", //债券ID
				"secShortName":   "string", //债券简称
				"cashTypeCD":     "string", //现金流类型
				"perAccrDate":    "date",   //本周期计息起始日
				"perAccrEndDate": "date",   //本周期计息截止日
				"recordDate":     "date",   //债权登记日
				"exDivDate":      "date",   //除权除息日
				"interest":       "double", //每百元面额付息（元）
				"payment":        "double", //每百元面额兑付本金
				"totalSize":      "double", //债券总规模(亿元)
			},
			"indices": map[string]string{},
		},
		//说明债券在一个日期区间的每日应计利息。 日期区间默认是过去一年。
		"getBondAi": map[string]interface{}{
			"url": "/api/bond/getBondAi.json",
			"args": map[string]string{
				"secID":     "string", //证券内部编码，可通过交易代码和交易市场在getSecurityID获取到。（可多值输入）
				"ticker":    "string", //交易代码（可多值输入）
				"beginDate": "date",   //计息起始日,默认为1年前
				"endDate":   "date",   //计息截止日，默认为今天
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":        "string", //证券ID
				"dataDate":     "date",   //日期
				"ticker":       "string", //交易代码
				"exchangeCD":   "string", //交易市场
				"AI":           "double", //每百元应计利息(元)
				"secShortName": "string", //债券简称
			},
			"indices": map[string]string{},
		},
		//说明收录债券担保信息，如担保人、担保类型、担保方式、担保期限等。
		"getBondGuar": map[string]interface{}{
			"url": "/api/bond/getBondGuar.json",
			"args": map[string]string{
				"secID":      "string", //证券内部编码，可通过交易代码和交易市场在getSecurityID获取到。（可多值输入）
				"ticker":     "string", //交易代码（可多值输入）
				"guarModeCD": "string", //担保方式,默认为0101（可多值输入）
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":        "string", //证券ID
				"beginDate":    "date",   //担保起始日期
				"ticker":       "string", //交易代码
				"exchangeCD":   "string", //交易市场
				"bondID":       "string", //债券ID
				"secShortName": "string", //债券简称
				"guarID":       "int16",  //担保人ID
				"guar":         "string", //担保人全称
				"endDate":      "date",   //担保结束日期
				"guarModeCD":   "string", //担保方式
				"assureGuaCD":  "string", //保证担保方式
				"guarRange":    "string", //担保范围与对象
				"guarTypeCD":   "string", //担保类型
				"publishDate":  "date",   //公告日期
			},
			"indices": map[string]string{},
		},
		//说明各评级机构公布的债券发行人的信用评级数据及历史变动记录。
		"getIssuerRating": map[string]interface{}{
			"url": "/api/bond/getIssuerRating.json",
			"args": map[string]string{
				"secID":     "string", //证券内部编码，可通过交易代码和交易市场在getSecurityID获取到。（可多值输入）
				"ticker":    "string", //交易代码（可多值输入）
				"beginDate": "date",   //评级起始日期
				"endDate":   "date",   //评级结束日期
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":         "string", //证券ID
				"ratingDate":    "date",   //评级日期
				"ticker":        "string", //交易代码
				"exchangeCD":    "string", //交易市场
				"partyID":       "int64",  //发行人ID
				"secShortName":  "string", //债券简称
				"partyFullName": "string", //发行人
				"rating":        "string", //级别
				"publishDate":   "date",   //评级公告发布日期
				"ratComID":      "int16",  //评级公司ID
				"ratCom":        "string", //评级公司
				"anticipateCD":  "int16",  //评级展望
			},
			"indices": map[string]string{},
		},
		//说明1、存储可转债发行时约定的初始转股价格，以及历次转股价格变动记录； 2、历史数据追溯至1992年，每日更新。
		"getBondConvPriceChg": map[string]interface{}{
			"url": "/api/bond/getBondConvPriceChg.json",
			"args": map[string]string{
				"secID":  "string", //证券ID，通联数据编制，全商城唯一标识证券编码，可通过交易代码和交易市场在getSecID获取到。
				"ticker": "string", //债券交易代码，输入可转债在交易所交易代码。
				"field":  "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"bondID":        "string", //债券ID
				"secID":         "string", //证券ID
				"ticker":        "string", //交易代码
				"secShortName":  "string", //债券简称
				"equSecID":      "string", //正股证券ID
				"equTicker":     "string", //正股交易代码
				"equShortName":  "string", //正股简称
				"convCode":      "string", //申报转股代码
				"convShortName": "string", //申报转股简称
				"convPrice":     "double", //转股价
				"convDate":      "string", //转股价变动生效日期
			},
			"indices": map[string]string{},
		},
		//说明债券的基本面信息，涵盖了债券计息方式、付息频率、起息日、到期日、兑付形式等。
		"getBond": map[string]interface{}{
			"url": "/api/bond/getBond.json",
			"args": map[string]string{
				"secID":      "string", //证券内部编码，可通过交易代码和交易市场在getSecurityID获取到。（可多值输入）
				"ticker":     "string", //交易代码（可多值输入）
				"exchangeCD": "string", //输入交易市场，可多值输入，XIBE-中国银行间市场,XSHE-深圳证券交易所,XSHG-上海证券交易所,OTC-柜台,OTER-其他市场
				"partyID":    "Int64",  //输入发行人ID，可得到该发行人发行的债券列表
			},
			"rets": map[string]string{
				"secID":           "string", //证券ID
				"ticker":          "string", //交易代码
				"exchangeCD":      "string", //交易市场
				"bondID":          "string", //债券ID
				"secFullName":     "string", //债券全称
				"secShortName":    "string", //债券简称
				"partyID":         "int64",  //发行人ID
				"issuer":          "string", //发行人
				"totalSize":       "double", //累计发行总额(亿元)
				"currencyCD":      "string", //货币代码
				"couponTypeCD":    "string", //计息方式
				"cpnFreqCD":       "string", //付息频率
				"paymentCD":       "string", //兑付方式
				"coupon":          "double", //票面年利率(%)
				"par":             "double", //债券面值(元)
				"hybridCD":        "string", //混合方式
				"typeID":          "string", //分类ID
				"typeName":        "string", //分类名称
				"publishDate":     "date",   //信息发布日期
				"listDate":        "date",   //[废弃]上市日期
				"delistDate":      "date",   //[废弃]退市日期
				"actMaturityDate": "date",   //实际到期日
				"firstAccrDate":   "date",   //起息日
				"maturityDate":    "date",   //到期日
				"firstRedempDate": "date",   //首次兑付日
				"minCoupon":       "double", //保底利率(%)
				"frnRefRateCD":    "string", //浮息债参考利率指标
				"frnCurrBmkRate":  "double", //浮息债首次参考利率(%)
				"frnMargin":       "double", //浮息债利差(%)
				"privPlacemFlag":  "int16",  //非公开定向标志
				"issueInvNum":     "int16",  //初始定向投资人数量
				"privInvNum":      "int16",  //定向投资人数量
				"absIssuerID":     "int64",  //发起机构ID
				"absIssuer":       "string", //发起机构
				"absLevelCD":      "string", //分层级别
				"absLevelRatio":   "double", //规模分层占比(%)
				"absCouponCap":    "string", //利率上限说明
				"isOption":        "int16",  //是否含权
				"remainSize":      "double", //债券余额
			},
			"indices": map[string]string{},
		},
		//基金信息
		//说明获取基金的净值调整信息，包括基金分红和基金拆分两种调整情况
		"getFundDiv": map[string]interface{}{
			"url": "/api/fund/getFundDiv.json",
			"args": map[string]string{
				"secID":        "string", //证券内部编码，可通过交易代码和交易市场在getSecurityID获取到。（可多值输入）
				"ticker":       "string", //输入一个或多个基金代码，用","分隔，如"000001"、"000001,000003"（可多值输入）
				"adjustedType": "string", //基金净值调整类型，D为分红，S为拆分,默认为D。
				"beginDate":    "date",   //起始日期，默认为1年前当前日期
				"endDate":      "date",   //截止日期，默认为今天
				"field":        "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":         "string", //内部编码
				"effectDate":    "date",   //当调整类型为D即分红时，实施日期指除息日；当调整类型为S即拆分时，实施日期为拆分日
				"adjustedType":  "string", //调整类型
				"ticker":        "string", //基金代码
				"secShortName":  "string", //基金中文简称
				"dividendAfTax": "double", //税后每份分红（元）
				"dividendBfTax": "double", //税前每份分红（元）
				"reinvestDate":  "date",   //分红再投资日
				"splitRatio":    "double", //拆分折算比例
				"publishDate":   "date",   //公告日期
				"currencyCd":    "string", //币种
			},
			"indices": map[string]string{},
		},
		//说明获取基金定期披露的持仓明细，包含所持有的股票、债券、基金的持仓明细数据。收录了2005年以来的历史数据，数据更新频率为季度。获取方式支持：1）输入一个或多个secID/ticker，并输入beginDate和endDate，可以查询到指定基金，一段时间的基金持仓；2）输入reportDate,不输入其他参数，可以查询到输入日期的全部基金持仓数据。
		"getFundHoldings": map[string]interface{}{
			"url": "/api/fund/getFundHoldings.json",
			"args": map[string]string{
				"secID":      "string", //证券ID，可通过交易代码和交易市场在getSecID获取到。（可多值输入）
				"ticker":     "string", //输入一个或多个基金代码，用","分隔，如"000001"、"000001,000003"（可多值输入）
				"reportDate": "date",   //输入一个报告日期，如20141231，不输入其他请求参数，可获取这个报告期全部基金资产配置，输入格式“YYYYMMDD”
				"beginDate":  "date",   //起始日期，输入格式“YYYYMMDD”
				"endDate":    "date",   //截止日期，输入格式“YYYYMMDD”
				"secType":    "string", //报告期内基金投资组合中持有证券类型，可输入一种或多种证券类型，用","分隔，默认为e。（可多值输入）
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"holdingSecID":        "string", //持有证券内部编码
				"reportDate":          "date",   //报告日期
				"secID":               "string", //证券ID
				"ticker":              "string", //基金代码
				"SecShortName":        "string", //基金中文简称
				"holdingsecType":      "string", //持有证券类型
				"holdingTicker":       "string", //持有证券交易代码
				"holdingExchangeCd":   "string", //交易所代码
				"holdingsecShortName": "string", //持有证券中文简称
				"marketValue":         "double", //市值（元）
				"ratioInNa":           "double", //占净资产比例（%）
				"publishDate":         "date",   //发布日期
				"currencyCd":          "string", //币种
			},
			"indices": map[string]string{},
		},
		//说明获取场内基金的份额变动信息，包含基金名称、交易代码、交易市场、截止日期、流通份额等信息。收录了2005年以来的历史数据，数据更新频率为日。
		"getFundSharesChg": map[string]interface{}{
			"url": "/api/fund/getFundSharesChg.json",
			"args": map[string]string{
				"secID":     "string", //证券ID，证券统一编码，可通过交易代码在getSecID获取到。
				"ticker":    "string", //输入一个或多个基金代码，用","分隔，如"000001"、"000001,000003"
				"beginDate": "date",   //起始日期，输入格式“YYYYMMDD”
				"endDate":   "date",   //截止日期，输入格式“YYYYMMDD”
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"ticker":            "string", //交易代码
				"endDate":           "date",   //截止日期
				"secID":             "string", //基金内部编码
				"exchangeCD":        "string", //交易市场
				"secShortName":      "string", //基金简称
				"circulationShares": "int64",  //流通份额
			},
			"indices": map[string]string{},
		},
		//说明获取分级基金A份额每期约定收益及预期约定收益，约定收益表达式。
		"getFundLeverRate": map[string]interface{}{
			"url": "/api/fund/getFundLeverRate.json",
			"args": map[string]string{
				"secID":     "string", //证券内部编码，可通过交易代码和交易市场在getSecID获取到。
				"ticker":    "string", //输入一个或多个基金代码，用","分隔，如"000001"，"000001,000003"
				"beginDate": "date",   //开始日期
				"endDate":   "date",   //截止日期
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":       "int64",  //基金内部编码
				"ticker":      "string", //交易代码
				"exchangeCD":  "string", //交易所代码
				"beginDate":   "date",   //开始日期
				"endDate":     "date",   //截止日期
				"returnRateA": "double", //A约定收益率
				"note":        "string", //约定收益表达式
			},
			"indices": map[string]string{},
		},
		//说明获取基金的基本档案信息，涵盖基金名称、交易代码、分级情况、所属类别、保本情况、上市信息、相关机构、投资描述等信息
		"getFund": map[string]interface{}{
			"url": "/api/fund/getFund.json",
			"args": map[string]string{
				"secID":         "string", //证券内部编码，可通过交易代码和交易市场在getSecurityID获取到。（可多值输入）
				"ticker":        "string", //输入一个或多个基金代码，用","分隔，如"000001"，"000001,000003"（可多值输入）
				"etfLof":        "string", //Lof'会获取LOF类型基金
				"listStatusCd":  "string", //（可多值输入）
				"category":      "string", //按投资对象分基金类型，E为股票型，H为混合型，B为债券型，SB为短期理财债券，M为货币型，O为其他
				"operationMode": "string", //基金运作模式，O为开放式，C为封闭式
				"field":         "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":              "string", //内部编码
				"ticker":             "string", //基金代码
				"secShortName":       "string", //基金中文简称
				"category":           "string", //按投资对象分基金类型，E为股票型，H为混合型，B为债券型，SB为短期理财债券，M为货币型，O为其他
				"operationMode":      "string", //基金运作模式，O为开放式，C为封闭式
				"indexFund":          "string", //基金指数型属性，I为指数型，EI为指数增强型
				"etfLof":             "string", //ETF或LOF型基金
				"isQdii":             "int16",  //是否为QDII基金，1为是，0为否
				"isFof":              "int16",  //是否为FOF基金，1为是，0为否
				"isGuarFund":         "int16",  //是否为保本基金，1为是，0为否
				"guarPeriod":         "double", //保本周期（月）
				"guarRatio":          "double", //保本比例（%）
				"exchangeCd":         "string", //交易所代码
				"listStatusCd":       "string", //基金上市状态，L上市，DE终止上市，UN未上市
				"establishDate":      "date",   //基金成立日期
				"listDate":           "date",   //上市日期
				"delistDate":         "date",   //终止上市日期
				"managementCompany":  "int64",  //基金管理人编码
				"managementFullName": "string", //基金管理人中文全称
				"custodian":          "int64",  //基金托管人编码
				"custodianFullName":  "string", //基金托管人中文全称
				"investField":        "string", //投资领域
				"investTarget":       "string", //投资目标
				"perfBenchmark":      "string", //业绩比较基准
				"circulationShares":  "Double", //最新流通份额
				"isClass":            "int16",  //是否分级基金,1为母基金,2为子基金A,3为子基金B,0为否
				"tradeAbbrName":      "String", //交易简称
				"managerName":        "String", //基金经理
			},
			"indices": map[string]string{},
		},
		//说明获取每日基金(货币型、短期理财债券型除外)净值，涵盖单位净值及累计净值等信息
		"getFundNav": map[string]interface{}{
			"url": "/api/fund/getFundNav.json",
			"args": map[string]string{
				"secID":     "string", //证券内部编码，可通过交易代码和交易市场在getSecID获取到。（可多值输入）
				"ticker":    "string", //输入一个或多个基金代码，用","分隔，如"000001"、"000001,000003"（可多值输入）
				"dataDate":  "date",   //输入一个日期，不输入其他请求参数，可获取到一天全部基金净值数据，输入格式“YYYYMMDD”
				"beginDate": "date",   //起始净值日期，输入格式“YYYYMMDD”
				"endDate":   "date",   //截止净值日期，输入格式“YYYYMMDD”
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":        "string", //内部编码
				"endDate":      "date",   //净值日期
				"ticker":       "string", //基金代码
				"secShortName": "string", //基金中文简称
				"NAV":          "double", //单位净值（元）
				"publishDate":  "date",   //净值公布日期
				"currencyCd":   "string", //计量货币，"CNY"为人民币，"HKD"为港币，"USD"为美元
				"ACCUM_NAV":    "Double", //累计净值(元)
				"ADJUST_NAV":   "Double", //复权净值(元)
			},
			"indices": map[string]string{},
		},
		//说明获取ETF基金交易日的申赎清单基本信息，包括标的指数名称，上一交易日的现金差额、最小申赎单位净值、单位净值，交易日当日的预估现金差额、最小申赎单位、现金替代比例上限、是否允许申购赎回、是否公布IOPV等信息
		"getFundETFPRList": map[string]interface{}{
			"url": "/api/fund/getFundETFPRList.json",
			"args": map[string]string{
				"beginDate": "date",   //起始日期，默认值为今天
				"endDate":   "date",   //截止日期，默认值为今天
				"secID":     "string", //证券ID，证券统一编码，可通过交易代码在getSecID获取到。（可多值输入）
				"ticker":    "string", //输入一个或多个基金代码，用","分隔，如"000001"、"000001,000003"（可多值输入）
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":           "string", //基金ID
				"ticker":          "string", //交易代码
				"fundShortName":   "string", //基金简称
				"underLyingIndex": "string", //指数ID
				"idxShortName":    "string", //指数简称
				"preTradeDate":    "date",   //上一交易日期
				"cashComp":        "double", //现金差额(元)
				"NAVPreCu":        "double", //最小申赎单位资产净值(元)
				"NAV":             "double", //单位净值
				"tradeDate":       "date",   //交易日期
				"estCahComp":      "double", //预估现金差额(元)
				"maxCashRatio":    "double", //现金替代比例上限(%)
				"creationUnit":    "int32",  //最小申赎单位(份)
				"ifIOPV":          "int16",  //是否需要公布IOPV
				"ifPurchaseble":   "int16",  //是否允许申购
				"ifRedeemable":    "int16",  //是否允许赎回
			},
			"indices": map[string]string{},
		},
		//说明获取ETF基金每个交易日的跟踪的标的指数成分券清单，包括成分券的代码、简称、股票数量、现金替代溢价比、固定替代金额等信息
		"getFundETFCons": map[string]interface{}{
			"url": "/api/fund/getFundETFCons.json",
			"args": map[string]string{
				"beginDate": "date",   //起始日期，默认值为今天
				"endDate":   "date",   //交易日期，默认为今天
				"secID":     "string", //证券ID，证券统一编码，可通过交易代码在getSecID获取到。（可多值输入）
				"ticker":    "string", //输入一个或多个基金代码，用","分隔，如"000001"、"000001,000003"（可多值输入）
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":           "string", //基金ID
				"ticker":          "string", //交易代码
				"secShortName":    "string", //基金简称
				"exchangeCD":      "string", //交易市场
				"tradeDate":       "date",   //交易日期
				"consID":          "string", //成份股ID
				"consTicker":      "string", //成份股代码
				"consName":        "string", //成份股简称
				"consExchangeCD":  "string", //交易市场
				"quantity":        "int64",  //股票数量(股)
				"cashSubsSign":    "int16",  //现金替代标志：1为允许，2为必须，3为禁止，4为退补
				"CashRatio":       "double", //现金替代溢价比例(%)
				"fixedCahsAmount": "double", //固定替代金额(元)
			},
			"indices": map[string]string{},
		},
		//说明获取基金评级信息，包括最新与历史评级情况
		"getFundRating": map[string]interface{}{
			"url": "/api/fund/getFundRating.json",
			"args": map[string]string{
				"beginDate": "date",   //起始日期，默认为一年前
				"endDate":   "date",   //截止日期，默认为今天
				"secID":     "string", //证券ID，证券统一编码，可通过交易代码在getSecID获取到。（可多值输入）
				"ticker":    "string", //输入一个或多个基金代码，用","分隔，如"000001"、"000001,000003"（可多值输入）
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":         "string", //基金内部编码
				"ticker":        "string", //交易代码
				"endDate":       "date",   //截止日期
				"exchangeCD":    "string", //交易市场
				"secShortName":  "string", //基金简称
				"period":        "int16",  //评级年期，3为三年期，5为五年期
				"overallRating": "int16",  //上海证券基金综合评级，1为一星级，2为二星级，3为三星级，4为四星级，5为五星级
			},
			"indices": map[string]string{},
		},
		//说明获取分级基金的基本信息，包含母、子基金名称、交易代码、分拆比例、折算等信息。
		"getFundLeverageInfo": map[string]interface{}{
			"url": "/api/fund/getFundLeverageInfo.json",
			"args": map[string]string{
				"exchangeCDLeverage": "string", //子基金交易市场,支持多值输入
				"secID":              "string", //输入单个证券内部编码，可通过交易代码和交易市场在getSecurityID获取到。
				"ticker":             "string", //母基金交易代码
				"field":              "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"category":             "string", //基金类型
				"secID":                "string", //母基金ID
				"fundID":               "int64",  //基金主编码
				"ticker":               "string", //母基金交易代码
				"exchangeCD":           "string", //母基金交易市场
				"secShortName":         "string", //母基金简称
				"secIDLeverage":        "string", //子基金ID
				"shareType":            "string", //份额类别
				"tickerLeverage":       "string", //子基金交易代码
				"exchangeCDLeverage":   "string", //子基金交易市场
				"secShortNameLeverage": "string", //子基金简称
				"shareProp":            "double", //分拆配比
				"idxCn":                "string", //跟踪指数
				"estDate":              "date",   //成立日期
				"endDate":              "date",   //到期日期
				"regSplitDate":         "date",   //定折日
				"upThrshold":           "double", //上折阈值
				"downThrshold":         "double", //下折阈值
				"splitNote":            "string", //分拆说明
				"navLever":             "double", //最新净值杠杆
				"priceLever":           "double", //最新价格杠杆
				"maxApply":             "double", //母基金最高申购费率
				"maxRedeem":            "double", //母基金
			},
			"indices": map[string]string{},
		},
		//说明获取每日货币型基金及短期理财债券型基金的收益情况，涵盖了每万份基金单位当日收益，七日年化收益率等信息
		"getFundDivm": map[string]interface{}{
			"url": "/api/fund/getFundDivm.json",
			"args": map[string]string{
				"secID":     "string", //证券内部编码，可通过交易代码和交易市场在getSecID获取到。（可多值输入）
				"ticker":    "string", //输入一个或多个基金代码，用","分隔，如"000001"、"000001,000003"（可多值输入）
				"dataDate":  "date",   //输入一个日期，不输入其他请求参数，可获取到一天全部基金历史收益数据，输入格式“YYYYMMDD”
				"beginDate": "date",   //起始日期，输入格式“YYYYMMDD”
				"endDate":   "date",   //截止日期，输入格式“YYYYMMDD”
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":        "string", //内部编码
				"endDate":      "date",   //收益日期
				"ticker":       "string", //交易代码
				"secShortName": "string", //基金中文简称
				"dailyProfit":  "double", //每万份基金单位当日收益(元)
				"weeklyYield":  "double", //7日年化收益率（%）
				"publishDate":  "date",   //发布日期
				"currencyCd":   "string", //币种
			},
			"indices": map[string]string{},
		},
		//说明获取基金定期披露的资产配置情况，包含了资产总值、资产净值，以及资产总值中权益类、固定收益类、现金及其他四种资产的市值与占比情况。收录了2005年以来的历史数据，数据更新频率为季度。获取方式支持：1）输入一个或多个secID/ticker，并输入beginDate和endDate，可以查询到指定基金，一段时间的资产配置；2）输入reportDate,不输入其他参数，可以查询到输入日期的全部基金资产配置
		"getFundAssets": map[string]interface{}{
			"url": "/api/fund/getFundAssets.json",
			"args": map[string]string{
				"secID":      "string", //证券ID，可通过交易代码和交易市场在getSecID获取到。
				"ticker":     "string", //输入一个或多个基金代码，用","分隔，如"000001"、"000001,000003"（可多值输入）
				"reportDate": "date",   //输入一个报告日期，如20141231，不输入其他请求参数，可获取这个报告期全部基金资产配置，输入格式“YYYYMMDD”
				"updateTime": "date",   //数据更新日期
				"beginDate":  "date",   //起始日期，输入格式“YYYYMMDD”
				"endDate":    "date",   //截止日期，输入格式“YYYYMMDD”
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":             "string", //内部编码
				"reportDate":        "date",   //报告日期
				"ticker":            "string", //基金代码
				"secShortName":      "string", //基金中文简称
				"totalAsset":        "double", //资产总值（元）
				"netAsset":          "double", //资产净值（元）
				"equityMarketValue": "double", //股票市值（元）
				"bondMarketValue":   "double", //债券市值（元）
				"cashMarketValue":   "double", //现金市值（元）
				"otherMarketValue":  "double", //其他资产市值（元）
				"equityRatioInTa":   "double", //股票占资产总值的比例（%）
				"bondRatioInTa":     "double", //债券占资产总值的比例（%）
				"cashRatioInTa":     "double", //货币占资产总值的比例（%）
				"otherRatioInTa":    "double", //其它资产占资产总值的比例（%）
				"publishDate":       "date",   //发布日期
				"currencyCd":        "string", //币种
				"updateTime":        "date",   //数据更新时间
			},
			"indices": map[string]string{},
		},
		//期货信息
		//说明获取国债期货转换因子信息，包括合约可交割国债名称、可交割国债交易代码、转换因子等。
		"getFutuConvf": map[string]interface{}{
			"url": "/api/future/getFutuConvf.json",
			"args": map[string]string{
				"secID":  "string", //合约内部编码，可通过交易代码和证券类型在getSecID获取到。（可多值输入）
				"ticker": "string", //合约交易代码，如'cu1106'（可多值输入）
				"field":  "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":            "string", //合约内部编码
				"publishDate":      "date",   //发布日期
				"ticker":           "string", //合约代码
				"secShortName":     "string", //合约简称
				"exchangeCD":       "string", //交易所
				"bondFullName":     "string", //债券全称
				"bondTickerIB":     "string", //银行间交易代码
				"bondTickerSH":     "string", //上交所交易代码
				"bondTickerSZ":     "string", //深交所交易代码
				"bondCoupon":       "double", //债券票面利率(%)
				"bondMaturityDate": "date",   //债券到期日期
				"conversionFactor": "double", //转换因子
			},
			"indices": map[string]string{},
		},
		//说明获取国内四大期货交易所期货合约的基本要素信息，包括合约名称、合约代码、合约类型、合约标的、报价单位、最小变动价位、涨跌停板幅度、交易货币、合约乘数、交易保证金、上市日期、最后交易日、交割日期、交割方式、交易手续费、交割手续费、挂牌基准价、合约状态等。
		"getFutu": map[string]interface{}{
			"url": "/api/future/getFutu.json",
			"args": map[string]string{
				"secID":          "string", //合约内部编码，可通过交易代码和证券类型在getSecID获取到。（可多值输入）
				"ticker":         "string", //合约交易代码，如'cu1106'（可多值输入）
				"exchangeCD":     "string", //期货交易所,如:XZCE 郑州商品交易所, CCFX 中国金融期货交易所, XSGE 上海期货交易所, XDCE 大连期货交易所
				"contractStatus": "string", //合约状态
				"contractObject": "string", //期货合约标的,如:cu 铜
				"field":          "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":             "string", //合约内部编码
				"listDate":          "date",   //合约上市日
				"secFullName":       "string", //合约全称
				"secShortName":      "string", //合约简称
				"ticker":            "string", //合约代码
				"exchangeCD":        "string", //交易所
				"contractType":      "string", //合约类型
				"contractObject":    "string", //合约标的
				"priceUnit":         "string", //报价单位
				"minChgPriceNum":    "double", //最小变动价位(数值)
				"minChgPriceUnit":   "string", //最小变动价位(单位)
				"priceValidDecimal": "int16",  //行情有效小数位
				"limitUpNum":        "double", //涨停板幅度(数值)
				"limitUpUnit":       "string", //涨停板幅度(单位)
				"limitDownNum":      "double", //跌停板幅度(数值)
				"limitDownUnit":     "string", //跌停板幅度(单位)
				"transCurrCD":       "string", //交易货币
				"contMultNum":       "double", //合约乘数(数值)
				"contMultUnit":      "string", //合约乘数(单位)
				"tradeMarginRatio":  "double", //交易保证金(%)
				"deliYear":          "int16",  //交割年
				"deliMonth":         "int16",  //交割月
				"lastTradeDate":     "date",   //最后交易日
				"firstDeliDate":     "date",   //开始交割日
				"lastDeliDate":      "date",   //最后交割日
				"deliMethod":        "string", //交割方式
				"deliGrade":         "string", //交割品级
				"tradeCommiNum":     "double", //交易手续费(数值)
				"tradeCommiUnit":    "string", //交易手续费(单位)
				"deliCommiNum":      "double", //交割手续费(数值)
				"deliCommiUnit":     "string", //交割手续费(单位)
				"listBasisPrice":    "double", //挂牌基准价
				"settPriceMethod":   "string", //当日结算价计算方式
				"deliPriceMethod":   "string", //交割结算价计算方式
				"contractStatus":    "string", //合约状态
			},
			"indices": map[string]string{},
		},
		//指数信息
		//说明获取指数成分股权重，获取信息包括成分股名称、成分股代码、权重生效日、成分股权重等。
		"getIdxCloseWeight": map[string]interface{}{
			"url": "/api/idx/getIdxCloseWeight.json",
			"args": map[string]string{
				"secID":     "string", //合约内部编码，可通过交易代码和证券类型在getSecID获取到。
				"ticker":    "string", //指数代码，指数编制或发布机构提供的代码，如沪深300为“000300”。当前提供的指数是000001-上证综指，000002-A股指数，000003-B股指数，000004-工业指数，000005-商业指数，000006-地产指数，000007-公用指数，000008-综合指数，000009-上证380，000010-上证180，000011-上证基金，000016-上证50，000020-中型综指，000090-上证流通，000132-上证100，000133-上证150，000300-沪深300，000852-中证1000，000902-中证流通，000903-中证100，000904-中证200，000905-中证500，000906-中证800，000907-中证700，399002-成份A指，399004-深证100R，399005-中小板指，399006-创业板指，399007-深证300，399008-中小300，399009-深证200，399010-深证700，399011-深证1000，399012-创业300，399013-深市精选，399015-中小创新，399107-深证A指，399108-深证B指，399306-深证ETF，399330-深证100，399333-中小板R，399400-大中盘，399401-中小盘
				"beginDate": "date",   //权重生效起始日期，输入格式“YYYYMMDD”
				"endDate":   "date",   //权重生效截止日期，输入格式“YYYYMMDD”
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":            "string", //指数内部编码
				"effDate":          "date",   //生效日期
				"secShortName":     "string", //指数简称
				"ticker":           "string", //指数代码
				"consID":           "string", //成分股内部编码
				"consShortName":    "string", //成分股简称
				"consTickerSymbol": "string", //成分股代码
				"consExchangeCD":   "string", //成分股市场代码
				"weight":           "double", //权重(%)
			},
			"indices": map[string]string{},
		},
		//说明获取国内外指数的基本要素信息，包括指数名称、指数代码、发布机构、发布日期、基日、基点等。
		"getIdx": map[string]interface{}{
			"url": "/api/idx/getIdx.json",
			"args": map[string]string{
				"secID":  "string", //合约内部编码，可通过交易代码和证券类型在getSecID获取到。（可多值输入）
				"ticker": "string", //指数代码，可先通过getSecID获取到，如在getSecID，选择证券类型'IDX',输入SECURITY_ID'1',可获取到TICKER_SYMBOL'000001',在此输入'000001'（可多值输入）
				"field":  "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":        "string", //指数内部编码
				"publishDate":  "date",   //发布日期
				"secShortName": "string", //指数简称
				"ticker":       "string", //指数代码
				"indexTypeCD":  "string", //指数类型编码
				"indexType":    "string", //指数类型
				"pubOrgCD":     "int64",  //发布机构编码
				"porgFullName": "string", //发布机构全称
				"baseDate":     "date",   //基日
				"basePoint":    "double", //基点
				"endDate":      "date",   //停用日期
			},
			"indices": map[string]string{},
		},
		//说明获取国内外指数的成分构成情况，包括指数成份股名称、成份股代码、入选日期、剔除日期等。
		"getIdxCons": map[string]interface{}{
			"url": "/api/idx/getIdxCons.json",
			"args": map[string]string{
				"secID":    "string", //合约内部编码，可通过交易代码和证券类型在getSecID获取到。（可多值输入）
				"ticker":   "string", //指数代码，可先通过getSecID获取到，如在getSecID，选择证券类型'IDX',输入SECURITY_ID'1',可获取到TICKER_SYMBOL'000001',在此输入'000001'（可多值输入）
				"isNew":    "int16",  //是否最新，0为非最新，1为最新，默认为所有。（可多值输入）
				"intoDate": "date",   //输入日期，可以获取这一天证券的指数成分
				"field":    "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":            "string", //指数内部编码
				"intoDate":         "date",   //入选日期
				"secShortName":     "string", //指数简称
				"ticker":           "string", //指数代码
				"consID":           "string", //成分股内部编码
				"consShortName":    "string", //成分股简称
				"consTickerSymbol": "string", //成分股代码
				"consExchangeCD":   "string", //成分股市场代码
				"outDate":          "date",   //剔除日期
				"isNew":            "int16",  //是否最新
			},
			"indices": map[string]string{},
		},
		//证券概况
		//说明输入证券ID或证券交易代码，查询证券停牌期间，获取证券停牌起始时间、截止时间。
		"getSecHalt": map[string]interface{}{
			"url": "/api/master/getSecHalt.json",
			"args": map[string]string{
				"secID":     "string", //证券内部编码，可通过交易代码和交易市场在getSecID获取到。
				"ticker":    "string", //证券交易代码，如'000001'
				"beginDate": "date",   //起始日期，默认为当前日期，输入格式“YYYYMMDD”
				"endDate":   "date",   //截止日期，默认为当前日期，输入格式“YYYYMMDD”
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":         "string", //证券ID
				"haltBeginTime": "string", //停牌起始时间
				"haltEndTime":   "string", //停牌结束时间
				"ticker":        "string", //交易代码
				"secShortName":  "string", //证券简称
				"exchangeCD":    "string", //交易市场
			},
			"indices": map[string]string{},
		},
		//说明输入交易所，选取日期范围，可查询获取日历日期当天是否开市信息
		"getTradeCal": map[string]interface{}{
			"url": "/api/master/getTradeCal.json",
			"args": map[string]string{
				"exchangeCD": "string", //证券
				"beginDate":  "date",   //起始日期
				"endDate":    "date",   //截止日期
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"exchangeCD":    "string", //证券交易所
				"calendarDate":  "date",   //日期
				"isOpen":        "int16",  //日期当天是否开市。0表示否，1表示是
				"prevTradeDate": "date",   //所在日期前一交易日
				"isWeekEnd":     "int16",  //当前日期是否当周最后交易日。0表示否，1表示是
				"isMonthEnd":    "int16",  //当前日期是否当月最后交易日。0表示否，1表示是
				"isQuarterEnd":  "int16",  //当前日期是否当季最后交易日。0表示否，1表示是
				"isYearEnd":     "int16",  //当前日期是否当年最后交易日。0表示否，1表示是
			},
			"indices": map[string]string{},
		},
		//说明根据拼音或股票代码，匹配股票代码、名称。包含正在上市的全部沪深股票。
		"getEquInfo": map[string]interface{}{
			"url": "/api/master/getEquInfo.json",
			"args": map[string]string{
				"ticker":   "string", //输入交易代码，如'000001';或证券简称拼音，如'wx'
				"pagesize": "int",    //分页大小，最大为30，默认是10
				"pagenum":  "int",    //分页页数，默认是1
				"field":    "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"ticker":     "string", //交易代码
				"exchangeCD": "string", //交易市场
				"shortNM":    "string", //证券简称
			},
			"indices": map[string]string{},
		},
		//说明获取沪深股票地域分类，以注册地所在行政区域为标准。
		"getSecTypeRegionRel": map[string]interface{}{
			"url": "/api/master/getSecTypeRegionRel.json",
			"args": map[string]string{
				"secID":  "string", //输入证券ID，可查询该证券属于哪些分类的成分，支持输入多个证券ID。
				"ticker": "string", //输入证券交易代码，可查询该证券交易代码属于哪些分类的成分，支持输入多个证券交易代码。
				"typeID": "string", //输入某个证券分类的typeID值（分类的子节点），可获取该子节点分类下全部成分。分类可在api如getSecTypeXXX获取，如地域类getSecTypeRegion。
				"field":  "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"ticker":       "string", //交易代码
				"typeID":       "string", //板块ID
				"typeName":     "string", //板块名称
				"secID":        "string", //证券ID
				"exchangeCD":   "string", //交易市场
				"secShortName": "string", //证券简称
			},
			"indices": map[string]string{},
		},
		//说明各api接口有枚举值特性的输出列，如getSecID输出项exchangeCD值，编码分别代表的是什么市场，所有枚举值都可以在这个接口获取。
		"getSysCode": map[string]interface{}{
			"url": "/api/master/getSysCode.json",
			"args": map[string]string{
				"codeTypeID": "int",    //输入获知的参数分类ID，可以查询到这个参数分类下所有枚举值，非必须输入项
				"valueCD":    "string", //输入获知的参数下常量值，可以查询到这个常量值的说明，非必须输入项
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"codeTypeID": "int",    //参数分类ID
				"valueCD":    "string", //常量值
				"valueName":  "string", //常量值说明
			},
			"indices": map[string]string{},
		},
		//说明输入一个或多个证券交易代码，获取证券ID，证券在数据结构中的一个唯一识别的编码；同时可以获取输入证券的基本上市信息，如交易市场，上市状态，交易币种，ISIN编码等。
		"getSecID": map[string]interface{}{
			"url": "/api/master/getSecID.json",
			"args": map[string]string{
				"partyID":    "int64",  //机构ID，assetClass为股票和债券时，是发行人的机构ID；assetClass为基金时，是基金管理人的机构ID；assetClass为指数时，是指数发布机构ID。（可多值输入）
				"ticker":     "string", //证券在交易所的交易代码，可输入一个或多个，用","分隔，如"000001"、"000001,600001"。（可多值输入）
				"assetClass": "string", //证券类型，可供选择类型：E 股票,B 债券,F 基金,IDX 指数,FU 期货,OP 期权。
				"cnSpell":    "string", //通过输入证券简称拼音的方式，获取证券交易代码等信息，拼音以使用汉字拼音首位连接，如“平安银行”，即“PAYH”。
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":        "string", //证券ID
				"ticker":       "string", //交易代码
				"secShortName": "string", //证券简称
				"cnSpell":      "string", //证券简称拼音
				"exchangeCD":   "string", //交易市场
				"assetClass":   "string", //证券类型
				"listStatusCD": "string", //上市状态，L——上市，S——暂停，DE——已退市，UN——未上市
				"listDate":     "date",   //上市日期
				"transCurrCD":  "string", //交易货币
				"ISIN":         "string", //ISIN编码
				"partyID":      "int64",  //机构ID
			},
			"indices": map[string]string{},
		},
		//说明输入行业分类通联编码(如，010303表示申万行业分类2014版)或输入一个行业分类标准名称，获取行业分类标准下行业划分
		"getIndustry": map[string]interface{}{
			"url": "/api/master/getIndustry.json",
			"args": map[string]string{
				"industryVersion":   "string", //行业分类标准。SW- 申万行业分类
				"industryVersionCD": "string", //行业分类标准数字编码。010301-证监会行业V2012、010303-申万行业、010308-中证行业
				"industryLevel":     "int16",  //输入数字如1,2,3,4，可指定查询的第几级行业
				"isNew":             "int16",  //是否最新，1表示是，0表示否
				"field":             "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"industryVersionCD": "string", //行业分类标准数字编码
				"industryID":        "string", //行业划分数字编码，通联赋予
				"industryVersion":   "string", //行业分类标准(字母缩写)
				"industry":          "string", //行业分类标准名称
				"industrySymbol":    "string", //行业划分(编制机构)
				"industryName":      "string", //行业名称
				"industryLevel":     "int16",  //行业划分所属级别
				"isNew":             "int16",  //是否最新
				"indexSymbol":       "string", //对应行业指数代码(编制机构)
				"updateTime":        "date",   //最近的一次更新时间
			},
			"indices": map[string]string{},
		},
		//说明获取中国地域分类，以行政划分为标准。
		"getSecTypeRegion": map[string]interface{}{
			"url": "/api/master/getSecTypeRegion.json",
			"args": map[string]string{
				"typeID": "string",    //typeName
				"string": "分类名称",      //string
				"父节点":    "typeLevel", //行业级别
			},
			"rets":    map[string]string{},
			"indices": map[string]string{},
		},
		//说明证券分类列表，一级分类包含有沪深股票、港股、基金、债券、期货、期权等，每个分类又细分有不同类型；可一次获取全部分类。
		"getSecType": map[string]interface{}{
			"url": "/api/master/getSecType.json",
			"args": map[string]string{
				"field": "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"typeID":    "string", //分类ID
				"typeName":  "string", //分类名称
				"parentID":  "string", //父节点
				"typeLevel": "int16",  //分类所在级别
			},
			"indices": map[string]string{},
		},
		//说明记录证券每个分类的成分，证券分类可通过在getSecType获取。
		"getSecTypeRel": map[string]interface{}{
			"url": "/api/master/getSecTypeRel.json",
			"args": map[string]string{
				"secID":  "string", //输入证券ID，可查询该证券属于哪些分类的成分，支持输入多个证券ID。
				"ticker": "string", //输入证券交易代码，可查询该证券交易代码属于哪些分类的成分，支持输入多个证券交易代码。
				"typeID": "string", //输入某个证券分类的typeID值（分类的子节点），可获取该子节点分类下全部成分。分类可在api如getSecTypeXXX获取，如地域类getSecTypeRegion。
				"field":  "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"ticker":       "string", //交易代码
				"typeID":       "string", //板块ID
				"typeName":     "string", //板块名称
				"secID":        "string", //证券ID
				"exchangeCD":   "string", //交易市场
				"secShortName": "string", //证券简称
			},
			"indices": map[string]string{},
		},
		//市场行情数据
		//说明获取当日期货分钟线
		"getFutureBarRTIntraDay": map[string]interface{}{
			"url": "/api/market/getFutureBarRTIntraDay.json",
			"args": map[string]string{
				"instrumentID": "string", //期货代码，即期货合约交易代码，如‘CF605‘，该字段在其他API中表述为ticker
				"endTime":      "string", //结束时间，例如14:20，如果为空则是当前时间
				"startTime":    "string", //开始时间，如果为空则是前一天的20:00（前一天夜盘开始时间）
				"unit":         "int",    //Bar(s)的时间宽度，单位分钟,
			},
			"rets": map[string]string{
				"unit":         "int32",  //Bar(s)的时间宽度，单位分钟
				"instrumentID": "string", //期货代码
				"exchangeCD":   "string", //交易所代码
				"utcOffset":    "string", //UTC 时间偏移
				"currencyCD":   "string", //货币代码
				"barTime":      "string", //当前bar的起始分钟点，如 11:10
				"closePrice":   "double", //当前bar的收盘价
				"openPrice":    "double", //当前bar的开盘价
				"highPrice":    "double", //当前bar的最高价
				"lowPrice":     "double", //当前bar的最低价
				"totalVolume":  "int64",  //当前bar的总交易量
				"totalValue":   "double", //当前bar的总交易金额
			},
			"indices": map[string]string{},
		},
		//说明获取一只或多只证券最新Level1股票信息。 输入一只或多只证券代码，如000001.XSHG (上证指数） 或000001.XSHE（平安银行）， 还有所选字段， 得到证券的最新交易快照。 证券可以是股票，指数， 部分债券或 基金。
		"getTickRTSnapshot": map[string]interface{}{
			"url": "/api/market/getTickRTSnapshot.json",
			"args": map[string]string{
				"securityID": "string", //一只或多只证券代码，用,分隔，格式是“数字.交易所代码”，如000001.XSHG。如果为空，则为全部证券。
				"assetClass": "string", //assetClass，返回该类型证券，E是股票，F是基金，B是债券，IDX是指数
				"exchangeCD": "string", //交易所代码，当
				"field":      "string", //可选参数，用,分隔，如field=lastPrice,shortNM，默认为空，返回全部字段，不选即为默认值。返回字段见下方。
			},
			"rets": map[string]string{
				"timestamp":  "i64",    //交易所时间戳（Unix时间）
				"ticker":     "string", //证券6位代码， 如‘000001'
				"exchangeCD": "string", //交易所代码，如’XSHE'. 和Ticker一起组成了证券唯一的代码。
			},
			"indices": map[string]string{},
		},
		//说明高频数据，获取期权最新市场信息快照
		"getOptionTickRTSnapshot": map[string]interface{}{
			"url": "/api/market/getOptionTickRTSnapshot.json",
			"args": map[string]string{
				"optionId": "string", //期权代码，如果为空，返回当前所有期权
				"field":    "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"optionId":  "string", //期权代码
				"timestamp": "Int64",  //交易所时间戳
			},
			"indices": map[string]string{},
		},
		//说明获取股票月后复权行情，包含开高低收量价、涨跌幅、换手率等
		"getMktEqumAdjAf": map[string]interface{}{
			"url": "/api/market/getMktEqumAdjAf.json",
			"args": map[string]string{
				"monthEndDate": "date",   //交易日期，输入格式“YYYYMMDD”
				"secID":        "string", //证券内部编码，可通过交易代码在getSecID获取到。（可多值输入，最多输入50只）
				"ticker":       "string", //股票交易代码，如'000001'（可多值输入，最多输入50只）
				"beginDate":    "date",   //起始日期，输入格式“YYYYMMDD”
				"endDate":      "date",   //截止日期，输入格式“YYYYMMDD”
				"isOpen":       "int32",  //股票当月是否开盘标记位：0-未开盘，1-开盘
				"field":        "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"ticker":          "string", //证券交易代码
				"endDate":         "date",   //交易日
				"secID":           "string", //证券内部代码
				"secShortName":    "string", //证券简称
				"exchangeCD":      "string", //交易市场
				"tradeDays":       "int32",  //交易天数
				"preClosePrice":   "double", //上月收盘
				"openPrice":       "double", //本月开盘
				"highestPrice":    "double", //本月最高
				"lowestPrice":     "double", //本月最低
				"closePrice":      "double", //本月收盘
				"turnoverVol":     "double", //成交量
				"turnoverValue":   "double", //成交金额
				"chg":             "double", //涨跌额
				"chgPct":          "double", //涨跌幅
				"return":          "double", //月
				"turnoverRate":    "double", //月
				"avgTurnoverRate": "double", //月
				"varReturn24":     "double", //月
				"sdReturn24":      "double", //月
				"avgReturn24":     "double", //月
				"varReturn60":     "double", //月
				"sdReturn60":      "double", //月
				"avgReturn60":     "double", //月
			},
			"indices": map[string]string{},
		},
		//说明获取一只或多只期货的最新市场信息快照
		"getFutureTickRTSnapshot": map[string]interface{}{
			"url": "/api/market/getFutureTickRTSnapshot.json",
			"args": map[string]string{
				"instrumentID": "string", //一只或多只期货代码，若空，则表示获取所有上市期货的行情数据
				"field":        "string", //可选参数，用,分隔，如field=lastPrice,shortNM，默认为空，返回全部字段，不选即为默认值。返回字段见下方。
			},
			"rets": map[string]string{
				"instrumentID": "string", //期货合约代码，如'IF1412'
				"timestamp":    "int64",  //交易所时间戳(UNIX时间)
			},
			"indices": map[string]string{},
		},
		//说明获取一只证券当日的分钟线信息。 输入一只证券代码，如000001.XSHE（平安银行）， 得到此证券的当日的分钟线。 证券目前是股票，指数，基金和部分债券。分钟线规则是向后归结，即 [9:31, 9:32)归结为9:32，有效时间为09:30-11:30,13:01-15:00，共241根分钟线。
		"getBarRTIntraDay": map[string]interface{}{
			"url": "/api/market/getBarRTIntraDay.json",
			"args": map[string]string{
				"securityID": "string", //证券代码，格式是“数字.交易所代码”，如000001.XSHG.证券可以是股票，指数， 部分债券或 基金。
				"startTime":  "string", //可选参数，起始时间， 如09:40，就是从早上09:39开始， 默认开始时间早上开市时间，即09:30，不选即为默认值
				"endTime":    "string", //可选参数，终止时间， 如14:00, 就是到下午14点结束， 如终止时间是空， 则截止到最新数据或到关市为止，即15:00，不选即为默认值
				"unit":       "int",    //Bar(s)的时间宽度，单位分钟, 如 1（分钟）/5（分钟）/15（分钟）/30（分钟）/60（分钟）
			},
			"rets":    map[string]string{},
			"indices": map[string]string{},
		},
		//说明获取沪深交易所交易日大宗交易成交价，成交量等信息。
		"getMktBlockd": map[string]interface{}{
			"url": "/api/market/getMktBlockd.json",
			"args": map[string]string{
				"secID":      "string", //证券ID，证券统一数字编码，可通过交易代码在getSecID获取到。
				"ticker":     "string", //证券交易代码，如'000001'
				"assetClass": "string", //证券类型，可供选择类型：E 股票,B 债券,F 基金；默认为 E。
				"tradeDate":  "date",   //输入一个日期，不输入其他请求参数，可获取到一天全部沪深证券大宗交易数据，输入格式“YYYYMMDD”
				"beginDate":  "date",   //起始日期，输入格式“YYYYMMDD”
				"endDate":    "date",   //截止日期，输入格式“YYYYMMDD”
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"tradeDate":   "date",   //交易日期
				"secID":       "string", //ID
				"exchangeCD":  "string", //交易市场
				"ticker":      "string", //交易代码
				"assetClass":  "string", //证券类型
				"secFullName": "string", //证券简称
				"currencyCD":  "string", //货币种类
				"tradePrice":  "double", //成交价
				"tradeVal":    "double", //成交金额
				"tradeVol":    "double", //成交量
				"buyerBD":     "string", //买方营业部
				"sellerBD":    "string", //卖方营业部
			},
			"indices": map[string]string{},
		},
		//说明获取央行公开市场操作信息，包含正回购、逆回购、短期流动性调节工具交易等信息
		"getMktCBOMO": map[string]interface{}{
			"url": "/api/market/getMktCBOMO.json",
			"args": map[string]string{
				"omoType":   "int16",  //公开市场操作:1-正回购、2-逆回购、3-SLO投放、4-SLO回笼
				"beginDate": "date",   //以公开市场操作日期为基准，查询起始日期，输入格式：YYYYMMDD
				"endDate":   "date",   //以公开市场操作日期为基准，查询截止日期，输入格式：YYYYMMDD
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"operateDate":    "date",   //起始日
				"turnoverVol":    "double", //交易量
				"omoTerm":        "string", //回购期限
				"operateDueDate": "date",   //到期日
				"publishDate":    "date",   //发布日期
				"tenderRate":     "double", //中标利率
				"omoType":        "int16",  //交易方式
				"tenderType":     "int16",  //招标方式
				"remark":         "string", //备注
			},
			"indices": map[string]string{},
		},
		//说明获取股票周行情，包含周开高低收量价、涨跌幅、换手率等
		"getMktEquw": map[string]interface{}{
			"url": "/api/market/getMktEquw.json",
			"args": map[string]string{
				"secID":       "string", //证券内部编码，可通过交易代码在getSecID获取到。（可多值输入，最多输入50只）
				"ticker":      "string", //股票交易代码，如'000001'（可多值输入，最多输入50只）
				"weekEndDate": "date",   //交易日期，输入格式“YYYYMMDD”
				"beginDate":   "date",   //起始日期，输入格式“YYYYMMDD”
				"endDate":     "date",   //截止日期，输入格式“YYYYMMDD”
				"isOpen":      "Int32",  //股票当周是否开盘标记位：0-今日未开盘，1-今日有开盘
				"field":       "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"ticker":          "string", //证券交易代码
				"endDate":         "date",   //交易日
				"secID":           "string", //证券内部代码
				"secShortName":    "string", //证券简称
				"exchangeCD":      "string", //交易市场
				"tradeDays":       "int32",  //交易天数
				"preClosePrice":   "double", //上周收盘
				"openPrice":       "double", //本周开盘
				"highestPrice":    "double", //本周最高
				"lowestPrice":     "double", //本周最低
				"closePrice":      "double", //本周收盘
				"turnoverVol":     "double", //成交量
				"turnoverValue":   "double", //成交金额
				"chg":             "double", //涨跌额
				"chgPct":          "double", //涨跌幅
				"return":          "double", //周回报率
				"turnoverRate":    "double", //周累计换手率
				"avgTurnoverRate": "double", //周平均换手率
				"varReturn100":    "double", //周回报率方差
				"sdReturn100":     "double", //周回报率标准差
				"avgReturn100":    "double", //周平均回报率
			},
			"indices": map[string]string{},
		},
		//说明获取四大期货交易所期货合约行情信息。 默认日期区间是过去一年。日线数据第一次更新为交易结束后（如遇线路不稳定情况数据可能存在误差），第二次更新为18:00pm，其中主力合约是以成交量最大为基准计算的。
		"getMktFutdVol": map[string]interface{}{
			"url": "/api/market/getMktFutdVol.json",
			"args": map[string]string{
				"secID":     "string", //证券内部编码，可通过交易代码在getmktequd获取到。（可多值输入）
				"ticker":    "string", //合约交易代码，如'cu1106'（可多值输入）
				"beginDate": "date",   //起始日期，默认为一年前当前日期
				"endDate":   "date",   //截止日期，默认为当前日期
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":          "string", //合约内部编码
				"tradeDate":      "date",   //交易日
				"ticker":         "string", //合约代码
				"exchangeCD":     "string", //交易所代码
				"secShortName":   "string", //合约简称
				"contractObject": "string", //合约标的
				"contractMark":   "string", //连续标志
				"preSettlePrice": "double", //昨结算
				"preClosePrice":  "double", //昨收盘
				"openPrice":      "double", //今开盘
				"highestPrice":   "double", //最高价
				"lowestPrice":    "double", //最低价
				"closePrice":     "double", //今收盘
				"settlePrice":    "double", //结算价
				"turnoverVol":    "int64",  //成交数量，单位为手
				"turnoverValue":  "double", //成交金额
				"openInt":        "int64",  //持仓量，单位为手
				"CHG":            "double", //涨跌
				"CHG1":           "double", //涨跌1
				"CHGPct":         "double", //涨跌幅
				"mainCon":        "int16",  //是否主力
				"smainCon":       "int16",  //是否次主力
			},
			"indices": map[string]string{},
		},
		//说明获取股票周后复权行情，包含开高低收量价、涨跌幅、换手率等
		"getMktEquwAdjAf": map[string]interface{}{
			"url": "/api/market/getMktEquwAdjAf.json",
			"args": map[string]string{
				"secID":       "string", //证券内部编码，可通过交易代码在getSecID获取到。（可多值输入，最多输入50只）
				"ticker":      "string", //股票交易代码，如'000001'（可多值输入，最多输入50只）
				"weekEndDate": "date",   //交易日期，输入格式“YYYYMMDD”
				"beginDate":   "date",   //起始日期，输入格式“YYYYMMDD”
				"endDate":     "date",   //截止日期，输入格式“YYYYMMDD”
				"isOpen":      "Int32",  //股票当周是否开盘标记位：0-今日未开盘，1-今日有开盘
				"field":       "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"ticker":          "string", //证券交易代码
				"endDate":         "date",   //交易日
				"secID":           "string", //证券内部代码
				"secShortName":    "string", //证券简称
				"exchangeCD":      "string", //交易市场
				"tradeDays":       "int32",  //交易天数
				"preClosePrice":   "double", //上周收盘
				"openPrice":       "double", //本周开盘
				"highestPrice":    "double", //本周最高
				"lowestPrice":     "double", //本周最低
				"closePrice":      "double", //本周收盘
				"turnoverVol":     "double", //成交量
				"turnoverValue":   "double", //成交金额
				"chg":             "double", //涨跌额
				"chgPct":          "double", //涨跌幅
				"return":          "double", //周回报率
				"turnoverRate":    "double", //周累计换手率
				"avgTurnoverRate": "double", //周平均换手率
				"varReturn100":    "double", //周回报率方差
				"sdReturn100":     "double", //周回报率标准差
				"avgReturn100":    "double", //周平均回报率
			},
			"indices": map[string]string{},
		},
		//说明获取沪深AB股日行情信息，包含昨收价、开盘价、最高价、最低价、收盘价、成交量、成交金额等字段，每日16:00更新
		"getMktEqud": map[string]interface{}{
			"url": "/api/market/getMktEqud.json",
			"args": map[string]string{
				"secID":     "string", //证券内部编码，可通过交易代码在getmktequd获取到。
				"ticker":    "string", //股票交易代码，如'000001'
				"tradeDate": "date",   //输入一个日期，不输入其他请求参数，可获取到一天全部沪深股票日行情数据，输入格式“YYYYMMDD”
				"beginDate": "date",   //起始日期，输入格式“YYYYMMDD”
				"endDate":   "date",   //截止日期，输入格式“YYYYMMDD”
				"isOpen":    "Int",    //股票今日是否开盘标记位：0-今日未开盘，1-今日有开盘
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":            "string", //证券内部编码
				"tradeDate":        "date",   //交易日
				"ticker":           "string", //证券代码
				"secShortName":     "string", //证券简称
				"exchangeCD":       "string", //交易所代码
				"preClosePrice":    "double", //昨收盘
				"actPreClosePrice": "double", //实际昨收盘
				"openPrice":        "double", //今开盘
				"highestPrice":     "double", //最高价
				"lowestPrice":      "double", //最低价
				"closePrice":       "double", //今收盘
				"turnoverVol":      "double", //成交量
				"turnoverValue":    "double", //成交金额
				"dealAmount":       "int32",  //成交笔数
				"turnoverRate":     "double", //日换手率
				"accumAdjFactor":   "double", //累积复权因子，注：前复权是对历史行情进行调整，除权除息日的行情不受本次除权除息影响，不需进行调整。最新一次除权除息日至最新行情期间的价格不需要进行任何的调整，该期间前复权因子都是1。
				"negMarketValue":   "double", //流通市值
				"marketValue":      "double", //总市值
				"isOpen":           "Int",    //股票今日是否开盘标记位：0-今日未开盘，1-今日有开盘
				"PE":               "double", //滚动市盈率，将于2016-01-01停止维护，推荐使用/api/market/getStockFactorsOneDay或/api/market/getStockFactorsDateRange，返回字段PE
				"PE1":              "double", //（根据最新财报得到预测值），将于2016-01-01停止维护，推荐使用/api/market/getStockFactorsOneDay或/api/market/getStockFactorsDateRange
				"PB":               "double", //（根据最新财报得到预测值），
			},
			"indices": map[string]string{},
		},
		//说明获取债券交易开、收、高、低，成交等日行情信息。
		"getMktBondd": map[string]interface{}{
			"url": "/api/market/getMktBondd.json",
			"args": map[string]string{
				"secID":     "string", //证券ID，证券统一编码，可通过交易代码在getSecID获取到。
				"ticker":    "string", //债券交易代码，如'100501'
				"tradeDate": "date",   //输入一个日期，不输入其他请求参数，可获取到一天全部债券日行情数据，输入格式“YYYYMMDD”
				"beginDate": "date",   //起始日期，输入格式“YYYYMMDD”
				"endDate":   "date",   //截止日期，输入格式“YYYYMMDD”
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"tradeDate":     "date",   //交易日期
				"secID":         "string", //证券ID
				"exchangeCD":    "string", //交易市场
				"ticker":        "string", //债券交易代码
				"secShortName":  "string", //债券简称
				"preClosePrice": "double", //昨收盘
				"openPrice":     "double", //今开盘
				"highestPrice":  "double", //最高价
				"lowestPrice":   "double", //最低价
				"closePrice":    "double", //收盘价
				"turnoverVol":   "double", //成交量
				"turnoverValue": "double", //成交金额
				"dealAmount":    "int32",  //成交笔数
				"accrInterest":  "double", //应计利息
				"YTM":           "double", //到期收益率
			},
			"indices": map[string]string{},
		},
		//说明获取现货日行情，包含上海黄金交易所所有现货行情。
		"getMktSpotd": map[string]interface{}{
			"url": "/api/market/getMktSpotd.json",
			"args": map[string]string{
				"secID":     "string", //证券内部编码，一串流水号,可先通过getSecID获取到，如在getSecID，选择证券类型为'SPOT',输入'Au(T+D)'，可获取到ID'AU(T+D).SGEX'后，在此输入'AU(T+D).SGEX'
				"ticker":    "string", //证券在交易所的交易代码
				"beginDate": "date",   //起始日期，格式为yyyymmdd
				"endDate":   "date",   //截止日期，格式为yyyymmdd
				"tradeDate": "date",   //交易日期，格式为yyyymmdd
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"tradeDate":      "date",   //交易日
				"sec_ID":         "string", //证券内部代码
				"ticker":         "string", //证券交易代码
				"secShortName":   "string", //证券简称
				"exchangeCD":     "string", //交易市场
				"openPrice":      "double", //开盘价
				"highestPrice":   "double", //最高价
				"lowestPrice":    "double", //最低价
				"closePrice":     "double", //收盘价
				"chg":            "double", //涨跌
				"chgPct":         "double", //涨跌幅
				"wPrice":         "double", //加权平均价
				"turnoverVolume": "double", //成交量
				"turnoverValue":  "double", //成交金额
				"openInt":        "double", //持仓量
				"deliverySide":   "int16",  //交收方向
				"deliveryVol":    "double", //交收量
			},
			"indices": map[string]string{},
		},
		//说明获取一只证券当日内时间段的Level1信息。 证券可以是股票，指数， 部分债券或 基金。
		"getTickRTIntraDay": map[string]interface{}{
			"url": "/api/market/getTickRTIntraDay.json",
			"args": map[string]string{
				"securityID": "string", //一只证券代码，格式是“数字.交易所代码”，如000001.XSHG。证券可以是股票，指数， 部分债券或 基金。
				"startTime":  "string", //开始时间，必须大于等于09:15，例如10:05，如果为空则是09:15
				"endTime":    "string", //结束时间，必须小于等于15:05，例如14:20，如果为空则是15：05
				"field":      "string", //可选参数，用逗号分隔，如field=lastPrice,shortNM，默认为空，返回全部字段，不选即为默认值。返回字段见下方。
			},
			"rets": map[string]string{
				"dataDate":   "string", //交易日期（yyyy-MM-dd）
				"dataTime":   "string", //交易时间'HH:mm:ss'
				"ticker":     "string", //证券6位代码， 如‘000001'
				"exchangeCD": "string", //交易所代码，如’XSHE'. 和Ticker一起组成了证券唯一的代码。
			},
			"indices": map[string]string{},
		},
		//说明获取一只期货在本清算日内某时间段的行情信息
		"getFutureTickRTIntraDay": map[string]interface{}{
			"url": "/api/market/getFutureTickRTIntraDay.json",
			"args": map[string]string{
				"instrumentID": "string", //一个期货合约代码，如A1009（大小写均可），此处不可以是合约类型代码。
				"startTime":    "string", //开始时间，如10:10，若空，则表示没有上界
				"endTime":      "string", //结束时间，如10:15，若空，则表示没有下界
				"field":        "string", //可选参数，用,分隔，默认为空，返回全部字段，不选即为默认值。注意：返回所有字段数据量大，将会增加延迟。建议用户只选需要的字段。返回字段见下方。
			},
			"rets": map[string]string{
				"dataDate":     "string", //交易日期（年月日部分），格式是‘YYYY-MM-DD'.
				"dataTime":     "string", //交易时间（时分秒部分）， 格式是’HH:MM:SS'.
				"dataMillisec": "int32",  //交易时间（毫秒部分）
				"instrumentID": "string", //期货代码
			},
			"indices": map[string]string{},
		},
		//说明证券及股票期权投资者的资金余额及变动情况
		"getInvestorCapitalChg": map[string]interface{}{
			"url": "/api/market/getInvestorCapitalChg.json",
			"args": map[string]string{
				"beginDate": "date",   //统计起始日范围左侧，输入格式“YYYYMMDD”
				"endDate":   "date",   //统计起始日范围右侧，输入格式“YYYYMMDD”
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"statsBeginDate": "date",   //起始日期
				"statsEnddate":   "date",   //结束日期
				"capitalType":    "double", //资金类别，1-证券交易结算资金，2-融资融券担保资金，3-股票期权保证金；还可通过getSysCode获取，令CODE_TYPE_ID = 12001
				"endingBalance":  "double", //(亿元）
				"dailyAverage":   "double", //(亿元）
				"transIn":        "double", //(亿元）
				"transOut":       "double", //(亿元）
				"netTransIn":     "double", //(亿元）
				"remark":         "string", //备注
			},
			"indices": map[string]string{},
		},
		//说明获取指数日线行情信息，包含昨收价、开盘价、最高价、最低价、收盘价、成交量、成交金额等字段。(参数加上type=2，能够每日16:00更新；否则是17:00更新)
		"getMktIdxd": map[string]interface{}{
			"url": "/api/market/getMktIdxd.json",
			"args": map[string]string{
				"indexID":   "string", //指数交易代码，可通过交易代码在getmktidxd获取。
				"ticker":    "string", //指数交易代码
				"tradeDate": "date",   //输入一个日期，不输入其他请求参数，可获取到一天全部指数日行情数据，输入格式“YYYYMMDD”
				"beginDate": "date",   //起始日期，输入格式“YYYYMMDD”
				"endDate":   "date",   //截止日期，输入格式“YYYYMMDD”
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"indexID":       "string", //指数内部编码
				"tradeDate":     "date",   //交易日
				"ticker":        "string", //指数代码
				"porgFullName":  "string", //发布机构全称
				"secShortName":  "string", //证券简称
				"exchangeCD":    "string", //交易所代码
				"preCloseIndex": "double", //昨收盘指数
				"openIndex":     "double", //今开盘指数
				"lowestIndex":   "double", //最低价指数
				"highestIndex":  "double", //最高价指数
				"closeIndex":    "double", //今收盘指数
				"turnoverVol":   "double", //成交量
				"turnoverValue": "double", //成交金额
				"CHG":           "double", //涨跌
				"CHGPct":        "double", //涨跌幅
			},
			"indices": map[string]string{},
		},
		//说明获取基金买卖交易开、收、高、低，成交等日行情信息。
		"getMktFundd": map[string]interface{}{
			"url": "/api/market/getMktFundd.json",
			"args": map[string]string{
				"beginDate": "date",   //起始日期，输入格式“YYYYMMDD”
				"endDate":   "date",   //截止日期，输入格式“YYYYMMDD”
				"tradeDate": "date",   //输入一个日期，不输入其他请求参数，可获取到一天全部基金日行情数据，输入格式“YYYYMMDD”
				"secID":     "string", //证券ID，证券统一编码，可通过交易代码在getSecID获取到。
				"ticker":    "string", //基金交易代码，如'150001'
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":             "string", //证券ID
				"tradeDate":         "date",   //交易日期
				"ticker":            "string", //基金交易代码
				"exchangeCD":        "string", //交易市场编码
				"secShortName":      "string", //证券简称
				"preClosePrice":     "double", //昨收盘
				"openPrice":         "double", //今开盘
				"highestPrice":      "double", //最高价
				"lowestPrice":       "double", //最低价
				"closePrice":        "double", //收盘价
				"CHG":               "double", //涨跌
				"CHGPct":            "double", //涨跌幅
				"turnoverVol":       "double", //成交量
				"turnoverValue":     "double", //成交金额
				"discount":          "double", //贴水
				"discountRatio":     "double", //贴水率
				"circulationShares": "double", //流通份额
			},
			"indices": map[string]string{},
		},
		//说明获取获取沪深A股和B股用来调整行情的前复权因子数据，包含除权除息日、除权除息事项具体数据、本次复权因子、累积复权因子以及因子调整的截止日期。该因子用来调整历史行情，不作为预测使用，于除权除息日进行计算调整。
		"getMktAdjf": map[string]interface{}{
			"url": "/api/market/getMktAdjf.json",
			"args": map[string]string{
				"secID":     "string", //一只或多只证券代码，用,分隔，格式是“数字.交易所代码”，如000001.XSHE。如果为空，则为全部证券。
				"ticker":    "string", //一只或多只股票代码，用,分隔，如000001,000002。
				"exDivDate": "Date",   //除权除息日，输入格式：YYYYMMDD
				"beginDate": "Date",   //除权除息日查询为基准，查询开始日期,输入格式"YYYYMMDD"
				"endDate":   "Date",   //除权除息日查询为基准，查询截止日期,输入格式"YYYYMMDD"
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":              "string", //证券展示代码
				"exDivDate":          "date",   //除权除息日
				"ticker":             "string", //交易代码
				"exchangeCD":         "string", //交易市场
				"secShortName":       "string", //证券简称
				"secShortNameEn":     "string", //证券英文简称
				"perCashDiv":         "double", //每股派现
				"perShareDivRatio":   "double", //每股送股比例
				"perShareTransRatio": "double", //每股转增股比例
				"allotmentRatio":     "double", //每股配股比例
				"allotmentPrice":     "double", //配股价
				"adjFactor":          "double", //复权因子(前复权)
				"accumAdjFactor":     "double", //累积复权因子
				"endDate":            "date",   //累积复权因子截止日期
			},
			"indices": map[string]string{},
		},
		//说明获取四大期货交易所主力合约、上海黄金交易所黄金(T+D)、白银(T+D)以及国外主要期货连续合约行情信息。 历史追溯至2006年，每日16:00更新。
		"getMktMFutd": map[string]interface{}{
			"url": "/api/market/getMktMFutd.json",
			"args": map[string]string{
				"contractObject": "string", //合约标的，具体请参照期货合约类型代码和描述表
				"contractMark":   "string", //L0为当月连续，L1为下月连续，L3为当季连续，L6为下季连续，在此输入L0查询当月连续
				"mainCon":        "int16",  //1为主力合约，0为非主力合约，在此输入1查询主力合约
				"tradeDate":      "date",   //输入一个日期，不输入其他请求参数，可获取到一天全部期货合约日行情数据，输入格式“YYYYMMDD”
				"endDate":        "date",   //结束日期，输入格式为yyyymmdd
				"startDate":      "date",   //起始日期，输入格式为yyyymmdd
				"field":          "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":          "string", //合约内部编码
				"ticker":         "string", //合约代码
				"exchangeCD":     "string", //交易市场
				"secShortName":   "string", //合约简称
				"secShortNameEN": "string", //合约英文简称
				"tradeDate":      "date",   //交易日期
				"contractObject": "string", //交易品种
				"contractMark":   "string", //连续合约标记
				"preSettlePrice": "double", //前结算价
				"preClosePrice":  "double", //前收盘价
				"openPrice":      "double", //开盘价
				"highestPrice":   "double", //最高价
				"lowestPrice":    "double", //最低价
				"settlePrice":    "double", //结算价
				"closePrice":     "double", //收盘价
				"turnoverVol":    "double", //成交量
				"turnoverValue":  "double", //成交金额
				"openInt":        "double", //持仓量
				"chg":            "double", //涨跌
				"chg1":           "double", //涨跌1
				"chgPct":         "double", //涨跌幅
				"mainCon":        "int16",  //主力合约
				"smainCon":       "int16",  //次主力合约
			},
			"indices": map[string]string{},
		},
		//说明获取期货会员在各交易日期货合约的成交量、成交排名及成交量增减信息
		"getMktFutMTR": map[string]interface{}{
			"url": "/api/market/getMktFutMTR.json",
			"args": map[string]string{
				"beginDate": "date",   //起始日期，默认为一年前当前日期
				"endDate":   "date",   //截止日期，默认为当前日期
				"secID":     "string", //证券ID，证券统一编码，可通过交易代码在getSecID获取到。（可多值输入）
				"ticker":    "string", //期货合约代码，如'SR501'（可多值输入）
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":          "string", //证券ID
				"tradeDate":      "date",   //交易日
				"ticker":         "string", //期货合约代码
				"secShortName":   "string", //期货合约简称
				"partyShortName": "string", //期货会员简称
				"exchangeCD":     "string", //交易所编码
				"turnoverVol":    "int64",  //成交量
				"CHG":            "double", //成交量增减
				"rank":           "int32",  //排名
			},
			"indices": map[string]string{},
		},
		//说明获取人民币汇率中间价
		"getMktFxRefRate": map[string]interface{}{
			"url": "/api/market/getMktFxRefRate.json",
			"args": map[string]string{
				"currencyPair": "string", //货币对,例如USD/CNY
				"beginDate":    "date",   //起始日期，格式为yyyymmdd
				"endDate":      "date",   //截止日期，格式为yyyymmdd
				"field":        "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"tradeDate":    "date",   //交易日期，格式为yyyymmdd
				"currencyPair": "string", //货币对,例如USD/CNY
				"midRate":      "double", //中间价
			},
			"indices": map[string]string{},
		},
		//说明获取沪深股票涨跌幅排行
		"getEquRTRank": map[string]interface{}{
			"url": "/api/market/getEquRTRank.json",
			"args": map[string]string{
				"desc":       "int32",  //是否是跌幅排行；不输入返回涨幅排行；输入1返回跌幅排行
				"exchangeCD": "string", //交易所代码；不输入返回沪深，输入XSHG只返回沪市，输入XSHE只返回深市
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"changePct":      "double", //变动率(变动/前收)
				"dataTime":       "string", //交易时间"hh:mm:ss"
				"ticker":         "string", //证券6位代码
				"exchangeCD":     "string", //交易所代码
				"shortNM":        "string", //证券简称
				"prevClosePrice": "double", //昨收盘价格
				"openPrice":      "double", //开盘价
				"volume":         "int64",  //成交数量
				"value":          "double", //成交金额
				"deal":           "int32",  //成交笔数(上证股票，债券，基金输出为0)
				"highPrice":      "double", //最高价
				"lowPrice":       "double", //最低价
				"lastPrice":      "double", //最新价格
				"change":         "double", //变动(最新-前收)
			},
			"indices": map[string]string{},
		},
		//说明获取四大期货交易所期货合约、上海黄金交易所黄金(T+D)、白银(T+D)以及国外主要期货连续合约行情信息。 日线数据第一次更新为交易结束后（如遇线路不稳定情况数据可能存在误差），第二次更新为18:00pm，其中主力合约是以连续三个交易日持仓量最大为基准计算的。
		"getMktFutd": map[string]interface{}{
			"url": "/api/market/getMktFutd.json",
			"args": map[string]string{
				"secID":     "string", //证券内部编码，可通过交易代码在getmktequd获取到。
				"ticker":    "string", //合约交易代码，如'cu1106'
				"tradeDate": "date",   //输入一个日期，不输入其他请求参数，可获取到一天全部期货合约日行情数据，输入格式“YYYYMMDD”
				"beginDate": "date",   //起始日期，输入格式“YYYYMMDD”
				"endDate":   "date",   //截止日期，输入格式“YYYYMMDD”
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":          "string", //合约内部编码
				"tradeDate":      "date",   //交易日
				"ticker":         "string", //合约代码
				"exchangeCD":     "string", //交易所代码
				"secShortName":   "string", //合约简称
				"contractObject": "string", //合约标的
				"contractMark":   "string", //连续标志
				"preSettlePrice": "double", //昨结算
				"preClosePrice":  "double", //昨收盘
				"openPrice":      "double", //今开盘
				"highestPrice":   "double", //最高价
				"lowestPrice":    "double", //最低价
				"closePrice":     "double", //今收盘
				"settlePrice":    "double", //结算价
				"turnoverVol":    "int64",  //成交数量，国内期货成交量单位为手，国际期货成交量单位为张，黄金（T+D）与白银（T+D）成交量单位为千克
				"turnoverValue":  "int64",  //成交金额
				"openInt":        "int64",  //持仓量，国内期货持仓量单位为手，国际期货持仓量单位为张，黄金（T+D）与白银（T+D）持仓量单位为千克
				"CHG":            "double", //涨跌
				"CHG1":           "double", //涨跌1
				"CHGPct":         "double", //涨跌幅
				"mainCon":        "int16",  //是否主力
				"smainCon":       "int16",  //是否次主力
			},
			"indices": map[string]string{},
		},
		//说明获取股票月前复权行情，包含开高低收量价、涨跌幅、换手率等
		"getMktEqumAdj": map[string]interface{}{
			"url": "/api/market/getMktEqumAdj.json",
			"args": map[string]string{
				"monthEndDate": "date",   //交易日期，输入格式“YYYYMMDD”
				"secID":        "string", //证券内部编码，可通过交易代码在getSecID获取到。（可多值输入，最多输入50只）
				"ticker":       "string", //股票交易代码，如'000001'（可多值输入，最多输入50只）
				"beginDate":    "date",   //起始日期，输入格式“YYYYMMDD”
				"endDate":      "date",   //截止日期，输入格式“YYYYMMDD”
				"isOpen":       "int32",  //股票当月是否开盘标记位：0-未开盘，1-开盘
				"field":        "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"ticker":          "string", //证券交易代码
				"endDate":         "date",   //交易日
				"secID":           "string", //证券内部代码
				"secShortName":    "string", //证券简称
				"exchangeCD":      "string", //交易市场
				"tradeDays":       "int32",  //交易天数
				"preClosePrice":   "double", //上月收盘
				"openPrice":       "double", //本月开盘
				"highestPrice":    "double", //本月最高
				"lowestPrice":     "double", //本月最低
				"closePrice":      "double", //本月收盘
				"turnoverVol":     "double", //成交量
				"turnoverValue":   "double", //成交金额
				"chg":             "double", //涨跌额
				"chgPct":          "double", //涨跌幅
				"return":          "double", //月
				"turnoverRate":    "double", //月
				"avgTurnoverRate": "double", //月
				"varReturn24":     "double", //月
				"sdReturn24":      "double", //月
				"avgReturn24":     "double", //月
				"varReturn60":     "double", //月
				"sdReturn60":      "double", //月
				"avgReturn60":     "double", //月
			},
			"indices": map[string]string{},
		},
		//说明获取期货会员在各交易日期货合约的多头持仓、排名及多头持仓增减信息
		"getMktFutMLR": map[string]interface{}{
			"url": "/api/market/getMktFutMLR.json",
			"args": map[string]string{
				"beginDate": "date",   //起始日期，默认为一年前当前日期
				"endDate":   "date",   //截止日期，默认为当前日期
				"secID":     "string", //证券内部编码，一串流水号,可先通过getSecID获取到，如在getSecID，选择证券类型为'FU 期货',输入'cu1501'，可获取到ID'CU1501.XSGE'后，在此输入'CU1501.XSGE'（可多值输入）
				"ticker":    "string", //期货合约代码，如'cu1501'（可多值输入）
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":          "string", //证券ID
				"tradeDate":      "date",   //交易日
				"ticker":         "string", //期货合约代码
				"secShortName":   "string", //期货合约简称
				"partyShortName": "string", //期货会员简称
				"exchangeCD":     "string", //交易所编码
				"longVol":        "int64",  //多头持仓量
				"CHG":            "double", //多头增减
				"rank":           "int32",  //排名
			},
			"indices": map[string]string{},
		},
		//说明记录周期内，如一周内新增投资者数量、期末投资者数量、期末持仓投资者数量、期间参与交易的投资者数量的统计数据。
		"getInvestorsStats": map[string]interface{}{
			"url": "/api/market/getInvestorsStats.json",
			"args": map[string]string{
				"beginDate": "date",   //统计周期起始日期为基准，查询起始日，输入格式“YYYYMMDD”
				"endDate":   "date",   //统计周期起始日期为基准，查询截止日，输入格式“YYYYMMDD”
				"statCycle": "string", //统计周期，W-每周
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"statsBeginDate":           "date",   //统计周期起始日期
				"statsEnddate":             "date",   //统计周期结束日期
				"newlyIncreased":           "double", //新增投资者数量
				"newlyIncreasedNatural":    "double", //新增自然人
				"newlyIncreasedNonnatural": "double", //新增非自然人
				"endingNumber":             "double", //期末投资者数量
				"endingNatural":            "double", //期末自然人
				"endingNaturalA":           "double", //期末已开立A股账户自然人投资者
				"endingNaturalB":           "double", //期末已开立B股账户自然人投资者
				"endingNonnatural":         "double", //期末非自然人
				"endingNonnaturalA":        "double", //期末已开立A股账户非自然人投资者
				"endingNonnaturalB":        "double", //期末已开立B股账户非自然人投资者
				"endingHolding":            "double", //期末持仓投资者数量
				"endingHoldingA":           "double", //期末持有A股的投资者
				"endingHoldingB":           "double", //期末持有B股的投资者
				"endingTrading":            "double", //期间参与交易的投资者数量
				"endingTradingA":           "double", //期间交易A股的投资者
				"endingTradingB":           "double", //期间交易B股的投资者
				"statCycle":                "string", //统计周期
			},
			"indices": map[string]string{},
		},
		//说明获取银行间同业拆借利率，包括shibor，libor，hibor。
		"getMktIbor": map[string]interface{}{
			"url": "/api/market/getMktIbor.json",
			"args": map[string]string{
				"secID":     "string", //拆借品种ID
				"ticker":    "string", //拆借品种，如隔夜，Hibor1D
				"tradeDate": "date",   //拆借日期
				"beginDate": "date",   //查询起始日期，格式为yyyymmdd
				"currency":  "string", //拆借币种，如CNY，USD
				"endDate":   "date",   //查询截止日期，格式为yyyymmdd
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"tradeDate": "date",   //交易日期，格式为yyyymmdd
				"secID":     "string", //拆借品种ID
				"ticker":    "string", //拆借品种
				"currency":  "string", //拆借币种
				"rate":      "double", //利率
			},
			"indices": map[string]string{},
		},
		//说明获取沪深交易所每日股票交易公开信息，包括涨跌幅异常、换手率异常等各种公开信息
		"getMktRankListStocks": map[string]interface{}{
			"url": "/api/market/getMktRankListStocks.json",
			"args": map[string]string{
				"abnormalTypeCD": "int32",  //异动类型代码，可在getSysCode中获得，令codeTypeID=12006
				"secID":          "string", //证券内部编码，可通过交易代码在getSecID获取到。
				"ticker":         "string", //证券交易代码，如'600000'（可多值输入，最多输入50只）
				"tradeDate":      "date",   //输入一个日期，不输入其他请求参数，可获取到一天全部数据，输入格式“YYYYMMDD”
				"beginDate":      "date",   //查询起始日期，格式为YYYYMMDD
				"endDate":        "date",   //查询截至日期，格式为YYYYMMDD
				"field":          "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"abnormalBeginDate": "date",   //异动起始日期
				"abnormalEndDate":   "date",   //异动截至日期
				"abnormalType":      "string", //异动类型
				"abnormalTypeCD":    "int32",  //异动类型代码
				"deviation":         "double", //偏离值
				"exchangeCD":        "string", //交易市场
				"secID":             "string", //证券内部代码
				"secShortName":      "string", //证券简称
				"ticker":            "string", //证券交易代码
				"tradeDate":         "date",   //交易日
				"turnoverValue":     "double", //成交金额
				"turnoverVol":       "double", //成交量
			},
			"indices": map[string]string{},
		},
		//说明获取一个指数的成份股的最新Level1股票信息。 输入一个指数的证券代码，如000001.XSHG (上证指数） 或000300.XSHG（沪深300）， 还有所选字段， 得到指数成份股的最新交易快照。
		"getTickRTSnapshotIndex": map[string]interface{}{
			"url": "/api/market/getTickRTSnapshotIndex.json",
			"args": map[string]string{
				"securityID": "string", //000903.XSHG（中证100）
				"field":      "string", //可选参数，用,分隔，如field=lastPrice,shortNM，默认为空，返回全部字段，不选即为默认值。返回字段见下方。
			},
			"rets": map[string]string{
				"timestamp":  "i64",    //交易所时间戳（Unix时间）
				"ticker":     "string", //证券6位代码， 如‘000001'
				"exchangeCD": "string", //交易所代码，如’XSHE'. 和Ticker一起组成了证券唯一的代码。
			},
			"indices": map[string]string{},
		},
		//说明获取获取沪深A股和B股后复权日行情信息（以上市价格为基准），包含后复权昨收价、后复权开盘价、后复权最高价、后复权最低价、后复权收盘价，每日16:00更新。
		"getMktEqudAdjAf": map[string]interface{}{
			"url": "/api/market/getMktEqudAdjAf.json",
			"args": map[string]string{
				"secID":     "string", //输入证券代码，格式是“数字.交易所代码”，如000001.XSHE。
				"ticker":    "string", //输入股票代码，如000001。
				"tradeDate": "date",   //输入一个日期，不输入其他请求参数，可获取到一天全部沪深股票复权行情数据，输入格式“YYYYMMDD”
				"beginDate": "date",   //开始日期，输入格式“YYYYMMDD”
				"endDate":   "date",   //结束日期，输入格式“YYYYMMDD”
				"isOpen":    "Int",    //股票今日是否开盘标记位：0-今日未开盘，1-今日有开盘
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"ticker":           "string", //股票代码
				"tradeDate":        "date",   //交易日期
				"secID":            "string", //证券展示代码
				"exchangeCD":       "string", //交易市场
				"secShortName":     "string", //证券简称
				"secShortNameEn":   "string", //证券英文简称
				"preClosePrice":    "double", //昨收盘(后复权)
				"openPrice":        "double", //今开盘(后复权)
				"highestPrice":     "double", //最高价(后复权)
				"lowestPrice":      "double", //最低价(后复权)
				"closePrice":       "double", //今收盘(后复权)
				"actPreClosePrice": "Double", //实际昨收盘
				"turnoverVol":      "Double", //成交量
				"turnoverValue":    "Double", //成交金额
				"dealAmount":       "Int32",  //成交笔数
				"turnoverRate":     "Double", //日换手率
				"accumAdjFactor":   "Double", //累积前复权因子，后复权是对未来行情进行调整。
				"negMarketValue":   "Double", //流通市值
				"marketValue":      "Double", //总市值
				"isOpen":           "Int",    //股票今日是否开盘标记位：0-今日未开盘，1-今日有开盘
			},
			"indices": map[string]string{},
		},
		//说明基金前复权行情，包含前收盘、今开盘、最高价、最低价、收盘价等
		"getMktFunddAdjBf": map[string]interface{}{
			"url": "/api/market/getMktFunddAdjBf.json",
			"args": map[string]string{
				"secID":     "string", //证券内部编码
				"ticker":    "string", //证券代码
				"beginDate": "date",   //起始日期，输入格式“YYYYMMDD”
				"endDate":   "date",   //截止日期，输入格式“YYYYMMDD”
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"ticker":         "string", //证券代码
				"tradeDate":      "date",   //交易日
				"secID":          "string", //证券内部编码
				"secShortName":   "string", //证券简称
				"secShortNameEn": "string", //证券英文简称
				"exchangeCD":     "string", //交易市场
				"preClosePrice":  "double", //前收盘(前复权)
				"openPrice":      "double", //开盘价(前复权)
				"highestPrice":   "double", //最高价(前复权)
				"lowestPrice":    "double", //最低价(前复权)
				"closePrice":     "double", //收盘价(前复权)
				"turnoverVol":    "double", //成交量(前复权)
				"turnoverValue":  "double", //成交金额(前复权)
			},
			"indices": map[string]string{},
		},
		//说明获取股票日资金流向按照单类大小明细，包含小单资金流入，小单资金流出，小单资金净流入等
		"getMktEquFlowOrder": map[string]interface{}{
			"url": "/api/market/getMktEquFlowOrder.json",
			"args": map[string]string{
				"secID":     "string", //证券内部编码,可先通过getSecID获取到，如在getSecID，选择证券类型为'fu',输入'IF1505'，可获取到secID'IF1505.CCFX'后，在此输入'IF1505.CCFX'
				"ticker":    "string", //证券代码,证券在交易所交易的代码
				"beginDate": "date",   //起始日期，输入格式为yyyymmdd
				"endDate":   "date",   //截止日期，输入格式为yyyymmdd
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"ticker":       "string", //证券交易代码
				"tradeDate":    "date",   //交易日
				"secID":        "string", //证券内部编码
				"secShortName": "date",   //证券简称
				"inflowS":      "double", //资金流入(小单)
				"inflowM":      "double", //资金流入(中单)
				"inflowL":      "double", //资金流入(大单)
				"inflowXl":     "double", //资金流入(超大单)
				"outflowS":     "double", //资金流出(小单)
				"outflowM":     "double", //资金流出(中单)
				"outflowL":     "double", //资金流出(大单)
				"outflowXl":    "double", //资金流出(超大单)
				"netInflowS":   "double", //资金净流入(小单)
				"netInflowM":   "double", //资金净流入(中单)
				"netInflowL":   "double", //资金净流入(大单)
				"netInflowXl":  "double", //资金净流入(超大单)
			},
			"indices": map[string]string{},
		},
		//说明获取银行间现券交易日行情
		"getMktIBBondd": map[string]interface{}{
			"url": "/api/market/getMktIBBondd.json",
			"args": map[string]string{
				"secID":     "string", //证券内部编码，可通过交易代码在getSecID获取到。（可多值输入，最多输入50只）
				"ticker":    "string", //证券交易代码，如'111593425'（可多值输入，最多输入50只）
				"tradeDate": "date",   //输入一个日期，不输入其他请求参数，可获取到一天全部银行间现券日行情数据，输入格式“YYYYMMDD”
				"beginDate": "date",   //起始日期，输入格式“YYYYMMDD”
				"endDate":   "date",   //截止日期，输入格式“YYYYMMDD”
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"ticker":        "string", //证券交易代码
				"tradeDate":     "date",   //交易日
				"secID":         "string", //证券内部代码
				"secShortName":  "string", //证券简称
				"exchangeCD":    "string", //交易市场
				"preClosePrice": "double", //前收盘净价
				"preWAVGPrice":  "double", //前加权平均净价
				"openPrice":     "double", //开盘净价
				"highestPrice":  "double", //最高净价
				"lowestPrice":   "double", //最低净价
				"closePrice":    "double", //收盘净价
				"wAvgPrice":     "double", //加权平均净价
				"chgPct":        "double", //净价涨跌幅
				"turnoverVol":   "double", //成交量
				"preCloseYield": "double", //前收盘收益率
				"preWAVGYield":  "double", //前加权平均收益率
				"openYield":     "double", //开盘收益率
				"highestYield":  "double", //最高收益率
				"lowestYield":   "double", //最低收益率
				"closeYield":    "double", //收盘收益率
				"wAvgYield":     "double", //加权平均收益率
			},
			"indices": map[string]string{},
		},
		//说明获取行业（证监会行业标准）资金流向，内容包括小单，中单，大单，超大单的资金流入，流出，净流入等。
		"getIndustryTickRTSnapshot": map[string]interface{}{
			"url": "/api/market/getIndustryTickRTSnapshot.json",
			"args": map[string]string{
				"securityID": "string", //一只或多只行业代码，用“,”分隔。为空返回所有行业
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"type": "string", //行业分类编码
			},
			"indices": map[string]string{},
		},
		//说明获取债券回购交易开、收、高、低，成交等日行情信息。
		"getMktRepod": map[string]interface{}{
			"url": "/api/market/getMktRepod.json",
			"args": map[string]string{
				"secID":     "string", //证券ID，证券统一编码，可通过交易代码在getSecID获取到。
				"ticker":    "string", //回购交易代码，如'204001'
				"tradeDate": "date",   //输入一个日期，不输入其他请求参数，可获取到一天全部证券债券回购日行情数据，输入格式“YYYYMMDD”
				"beginDate": "date",   //起始日期，输入格式“YYYYMMDD”
				"endDate":   "date",   //截止日期，输入格式“YYYYMMDD”
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"tradeDate":     "date",   //交易日期
				"secID":         "string", //证券ID
				"exchangeCD":    "string", //交易市场
				"ticker":        "string", //回购代码
				"secShortName":  "string", //回购简称
				"preCloseRate":  "double", //前收盘利率
				"openRate":      "double", //开盘利率
				"highestRate":   "double", //最高利率
				"lowestRate":    "double", //最低利率
				"closeRate":     "double", //收盘利率
				"turnoverVol":   "double", //成交量
				"turnoverValue": "double", //成交额
				"dealAmount":    "double", //成交笔数
			},
			"indices": map[string]string{},
		},
		//说明主要记录上交所期权行情，包含昨结算、昨收盘、开盘价、最高价、最低价、收盘价、结算价、成交量、成交金额、持仓量等字段，每日16:00前更新。
		"getMktOptd": map[string]interface{}{
			"url": "/api/market/getMktOptd.json",
			"args": map[string]string{
				"optID":     "string", //交易所指定8位代码
				"secID":     "string", //证券内部编码，一串流水号,可先通过getSecID获取到，如在getSecID，选择证券类型为'op',输入'11000149'，可获取到ID'510180C1506M02950.XSHG2'后，在此输入'510180C1506M02950.XSHG2'
				"ticker":    "string", //合约交易代码
				"tradeDate": "date",   //输入一个日期，不输入其他请求参数，可获取到一天全部期权日行情数据，输入格式“YYYYMMDD”
				"beginDate": "date",   //起始日期，输入格式为yyyymmdd
				"endDate":   "date",   //结束日期，输入格式为yyyymmdd
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"tradeDate":     "date",   //交易日期
				"secID":         "string", //证券内部编码
				"optID":         "string", //合约编码
				"ticker":        "string", //合约交易代码
				"secShortName":  "string", //证券简称
				"exchangeCD":    "string", //交易市场
				"preClosePrice": "double", //前结算价
				"openPrice":     "double", //今开盘
				"highestPrice":  "double", //最高价
				"lowestPrice":   "double", //最低价
				"closePrice":    "double", //收盘价
				"settlPrice":    "double", //结算价
				"turnoverVol":   "double", //成交量
				"turnoverValue": "double", //成交金额
				"openInt":       "int32",  //持仓量
			},
			"indices": map[string]string{},
		},
		//说明获取每日个股资金流向数据，包含个股资金流入、资金流出和资金净流入等
		"getMktEquFlow": map[string]interface{}{
			"url": "/api/market/getMktEquFlow.json",
			"args": map[string]string{
				"secID":     "string", //证券内部编码，可通过交易代码在getSecID获取到。（可多值输入，最多输入50只）
				"ticker":    "string", //股票交易代码，如'000001'（可多值输入，最多输入50只）
				"tradeDate": "date",   //输入一个日期，不输入其他请求参数，可获取到一天全部沪深股票日资金流向数据，输入格式“YYYYMMDD”
				"beginDate": "date",   //起始日期，输入格式“YYYYMMDD”
				"endDate":   "date",   //截止日期，输入格式“YYYYMMDD”
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"tradeDate":      "date",   //交易日
				"ticker":         "string", //证券代码
				"secID":          "string", //证券内部编码
				"secShortName":   "string", //证券简称
				"secShortNameEn": "string", //证券英文简称
				"exchangeCD":     "string", //交易所代码
				"moneyInflow":    "double", //资金流入
				"moneyOutflow":   "double", //资金流出
				"netMoneyInflow": "double", //资金净流入
			},
			"indices": map[string]string{},
		},
		//说明获取当日所有股票某一分钟的分钟线。分钟线规则是向后归结，即 [9:31, 9:32)归结为9:32，有效时间为09:30-11:30,13:01-15:00，共241根分钟线。
		"getBarRTIntraDayOneMinute": map[string]interface{}{
			"url": "/api/market/getBarRTIntraDayOneMinute.json",
			"args": map[string]string{
				"time":       "string", //查询时间，如'09:33'
				"assetClass": "string", //证券类型，输入assetClass，返回该类型证券，E是股票，F是基金，B是债券，IDX是指数
				"exchangeCD": "string", //交易所代码，输入exchangeCD，返回该交易所证券，XSHG是上交所，XSHE是深交所
				"unit":       "int",    //可选参数，
			},
			"rets": map[string]string{
				"unit":        "i32",    //Bar(s)的时间宽度,目前只有一分钟
				"ticker":      "string", //证券代码
				"exchangeCD":  "string", //交易所代码
				"shortNM":     "string", //证劵简称
				"utcOffset":   "string", //UTC 时间偏移
				"currencyCD":  "string", //货币代码
				"barTime":     "string", //当前bar的起始分钟点，如 11:10
				"closePrice":  "double", //当前bar的收盘价
				"openPrice":   "double", //当前bar的开盘价
				"highPrice":   "double", //当前bar的最高价
				"lowPrice":    "double", //当前bar的最低价
				"totalVolume": "int64",  //(
				"totalValue":  "double", //当前bar的总交易金额
			},
			"indices": map[string]string{},
		},
		//说明主要记录盘前每日个股及基金涨跌停板价格，每日9:00更新
		"getMktLimit": map[string]interface{}{
			"url": "/api/market/getMktLimit.json",
			"args": map[string]string{
				"secID":     "string", //证券内部编码，可通过交易代码在getSecID获取到。
				"ticker":    "string", //证券交易代码
				"tradeDate": "date",   //交易日期，输入格式YYYYMMDD
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"ticker":                   "string", //证券交易代码
				"secID":                    "string", //证券内部编码
				"secShortName":             "string", //证券简称
				"secShortNameEn":           "string", //证券英文简称
				"exchangeCD":               "string", //交易所代码
				"tradeDate":                "date",   //交易日期
				"limitUpPrice":             "double", //涨停价
				"limitDownPrice":           "double", //跌停价
				"UP_LIMIT_REACHED_TIMES":   "string", //当日涨停次数
				"DOWN_LIMIT_REACHED_TIMES": "string", //当日跌停次数
			},
			"indices": map[string]string{},
		},
		//说明获取中国银行发布的实时外汇牌价，包含现钞现汇买入卖出价等
		"getMktFxQtRate": map[string]interface{}{
			"url": "/api/market/getMktFxQtRate.json",
			"args": map[string]string{
				"currencyPair": "string", //货币代码，HKD-港元,BRL-巴西雷亚尔,JPY-日元,GBP-英镑,SGD-新加坡元,MOP-澳门元,SEK-瑞典克朗,FRF-法国法郎,INR-印度卢比,AED-阿联酋迪拉姆,IDR-印度尼西亚卢比,FIM-芬兰马克,NZD-新西兰元,KRW-韩元,DEM-德国马克,NOK-挪威克朗,TWD-台湾元,AUD-澳元,BEF-比利时法郎,DKK-丹麦克朗,MYR-马来西亚林吉特,RUB-俄罗斯卢布,PHP-菲律宾比索,USD-美元,CAD-加元,NLG-荷兰盾,CHF-瑞士法郎,ESP-西班牙比塞塔,EUR-欧元,THB-泰铢,ITL-意大利里拉,ZAR-南非兰特
				"beginTime":    "string", //查询发布开始时间，格式为YYYYMMDDHHMMSS，如20160222093000
				"endTime":      "string", //查询发布截止时间，格式为YYYYMMDDHHMMSS，如20160222113000
				"field":        "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"publishTime":     "date",   //发布时间
				"partyFullName":   "string", //报价银行全称
				"partyID":         "int64",  //报价银行ID
				"currencyPair":    "string", //货币代码
				"currencyPairCn":  "string", //货币代码中文简称
				"currencyBid":     "double", //现汇买入价
				"cashBid":         "double", //现钞买入价
				"currencyAsk":     "double", //现汇卖出价
				"cashAsk":         "double", //现钞卖出价
				"conversionPrice": "double", //折算价
			},
			"indices": map[string]string{},
		},
		//说明获取沪深交易所每日营业部交易公开信息，包括涨跌幅异常、换手率异常等各营业部排名
		"getMktRankListSales": map[string]interface{}{
			"url": "/api/market/getMktRankListSales.json",
			"args": map[string]string{
				"rank":      "int32",  //排名，提供第1~5名的排名
				"secID":     "string", //证券内部编码，可通过交易代码在getSecID获取到。
				"side":      "string", //买卖方向，B-买入、S-卖出、N-买卖综合
				"ticker":    "string", //证券交易代码，如'600000'（可多值输入，最多输入50只）
				"tradeDate": "date",   //输入一个日期，不输入其他请求参数，可获取到一天全部数据，输入格式“YYYYMMDD”
				"beginDate": "date",   //查询起始日期，格式为YYYYMMDD
				"endDate":   "date",   //查询截至日期，格式为YYYYMMDD
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"buyValue":     "double", //买入金额
				"exchangeCD":   "string", //交易市场
				"rank":         "int32",  //排名
				"sales":        "string", //营业部名称
				"secID":        "string", //证券内部代码
				"secShortName": "string", //证券简称
				"sellValue":    "double", //卖出金额
				"side":         "string", //买卖方向
				"ticker":       "string", //证券交易代码
				"totalValue":   "double", //买卖金额
				"tradeDate":    "date",   //交易日
			},
			"indices": map[string]string{},
		},
		//说明获取期货会员在各交易日期货合约的空头持仓、排名及空头持仓增减信息
		"getMktFutMSR": map[string]interface{}{
			"url": "/api/market/getMktFutMSR.json",
			"args": map[string]string{
				"beginDate": "date",   //起始日期，默认为一年前当前日期
				"endDate":   "date",   //截止日期，默认为当前日期
				"secID":     "string", //证券ID，证券统一编码，可通过交易代码在getSecID获取到。（可多值输入）
				"ticker":    "string", //期货合约代码，如'SR501'（可多值输入）
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":          "string", //证券ID
				"tradeDate":      "date",   //交易日
				"ticker":         "string", //期货合约代码
				"secShortName":   "string", //期货合约简称
				"partyShortName": "string", //期货会员简称
				"exchangeCD":     "string", //交易所编码
				"shortVol":       "int64",  //空头持仓量
				"CHG":            "double", //空头增减
				"rank":           "int32",  //排名
			},
			"indices": map[string]string{},
		},
		//说明获取最活跃3个合约品种及持仓最大3个合约交易排名情况
		"getMktOptTdRank": map[string]interface{}{
			"url": "/api/market/getMktOptTdRank.json",
			"args": map[string]string{
				"secID":     "string", //证券内部编码，一串流水号,可先通过getSecID获取到，如在getSecID，选择证券类型为'F 基金',输入'510050'，可获取到ID'510050.XSHG'后，在此输入510050.XSHG'
				"ticker":    "string", //输入证券交易代码，如“510050”
				"beginDate": "date",   //起始日期，格式为yyyymmdd
				"endDate":   "date",   //截止日期，格式为yyyymmdd
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"tradeDate":      "date",   //交易日
				"ticker":         "string", //证券交易代码
				"secID":          "string", //证券内部编码
				"secShortName":   "string", //证券简称
				"exchangeCD":     "string", //交易市场
				"rank":           "int32",  //排名
				"volume":         "double", //数量
				"partyID":        "int64",  //经营机构ID
				"partyShortName": "string", //经营机构简称
				"statsType":      "string", //统计项目类型，可通过getSysCode获取，令CODE_TYPE_ID = 90002
			},
			"indices": map[string]string{},
		},
		//说明获取行业日资金流向按照单类大小明细，包含小单资金流入，小单资金流出，小单资金净流入等
		"getMktIndustryFlowOrder": map[string]interface{}{
			"url": "/api/market/getMktIndustryFlowOrder.json",
			"args": map[string]string{
				"tradeDate": "date",   //交易日期
				"beginDate": "date",   //起始日期，输入格式为yyyymmdd
				"endDate":   "date",   //截止日期，输入格式为yyyymmdd
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"industryName": "string", //行业名称
				"tradeDate":    "date",   //交易日
				"industryID":   "string", //行业ID
				"inflowS":      "double", //资金流入(小单)
				"inflowM":      "double", //资金流入(中单)
				"inflowL":      "double", //资金流入(大单)
				"inflowXl":     "double", //资金流入(超大单)
				"outflowS":     "double", //资金流出(小单)
				"outflowM":     "double", //资金流出(中单)
				"outflowL":     "double", //资金流出(大单)
				"outflowXl":    "double", //资金流出(超大单)
				"netInflowS":   "double", //资金净流入(小单)
				"netInflowM":   "double", //资金净流入(中单)
				"netInflowL":   "double", //资金净流入(大单)
				"netInflowXl":  "double", //资金净流入(超大单)
			},
			"indices": map[string]string{},
		},
		//说明获取期权交易统计
		"getMktOptStats": map[string]interface{}{
			"url": "/api/market/getMktOptStats.json",
			"args": map[string]string{
				"secID":         "string", //证券内部编码，一串流水号,可先通过getSecID获取到，如在getSecID，选择证券类型为'F 基金',输入'510050'，可获取到ID'510050.XSHG'后，在此输入510050.XSHG'
				"ticker":        "string", //证券交易代码
				"tradeDate":     "date",   //交易日期，格式为yyyymmdd
				"beginDate":     "date",   //起始日期，格式为yyyymmdd
				"endDate":       "date",   //截止日期，格式为yyyymmdd
				"statsInterval": "int16",  //数据统计区间，2=月；4=日
				"field":         "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"tradeDate":     "date",   //交易日
				"ticker":        "string", //证券交易代码
				"secID":         "string", //证券内部编码
				"secShortName":  "string", //证券简称
				"exchangeCD":    "string", //交易市场
				"statsInterval": "int16",  //数据统计区间，可通过getSysCode获取，令CODE_TYPE_ID = 50042
				"turnoverVol":   "int32",  //成交量
				"cVol":          "int32",  //认购成交量
				"pVol":          "int32",  //认沽成交量
				"pcRate":        "double", //认沽认购比
				"openInt":       "int32",  //未平仓合约总张数
				"cOpenInt":      "int32",  //未平仓认购合约张数
				"pOpenInt":      "int32",  //未平仓认沽合约张数
			},
			"indices": map[string]string{},
		},
		//说明获取沪深股票个股最近一次日行情，默认日期区间是过去1年，包含昨收价、开盘价、最高价、最低价、收盘价、成交量、成交金额等字段，每日15:30更新
		"getMktEqudLately": map[string]interface{}{
			"url": "/api/market/getMktEqudLately.json",
			"args": map[string]string{
				"field": "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":            "string", //证券内部编码
				"tradeDate":        "date",   //交易日
				"ticker":           "string", //证券代码
				"secShortName":     "string", //证券简称
				"exchangeCD":       "string", //交易所代码
				"preClosePrice":    "double", //昨收盘
				"actPreClosePrice": "double", //实际昨收盘
				"openPrice":        "double", //今开盘
				"highestPrice":     "double", //最高价
				"lowestPrice":      "double", //最低价
				"closePrice":       "double", //今收盘
				"turnoverVol":      "double", //成交量
				"turnoverValue":    "double", //成交金额
				"dealAmount":       "int32",  //成交笔数
				"turnoverRate":     "double", //日换手率
				"accumAdjFactor":   "double", //累积复权因子
				"exDivDate":        "Date",   //最近一次的除权除息日
				"negMarketValue":   "double", //流通市值
				"marketValue":      "double", //总市值
			},
			"indices": map[string]string{},
		},
		//说明获取当日期货分钟线
		"getOptionBarRTIntraDay": map[string]interface{}{
			"url": "/api/market/getOptionBarRTIntraDay.json",
			"args": map[string]string{
				"optionId":  "string", //期权代码
				"endTime":   "string", //结束时间
				"startTime": "string", //开始时间
				"unit":      "int",    //Bar(s)的时间宽度，单位分钟，
			},
			"rets": map[string]string{
				"unit":              "int32",  //Bar(s)的时间宽度，单位分钟
				"ticker":            "string", //期权代码
				"exchangeCD":        "string", //交易所代码
				"utcOffset":         "string", //UTC 时间偏移
				"currencyCD":        "string", //货币代码
				"barTime":           "string", //当前bar的起始分钟点，如 11:10
				"closePrice":        "double", //当前bar的收盘价
				"openPrice":         "double", //当前bar的开盘价
				"highPrice":         "double", //当前bar的最高价
				"lowPrice":          "double", //当前bar的最低价
				"totalVolume":       "int64",  //当前bar的总交易量
				"totalValue":        "double", //当前bar的总交易金额
				"totalLongPosition": "int",    //当前合约未平仓数量
			},
			"indices": map[string]string{},
		},
		//说明获取每日行业资金流向数据，包含行业资金流入、资金流出和资金净流入等，行业分类标准为证监会行业分类标准
		"getMktIndustryFlow": map[string]interface{}{
			"url": "/api/market/getMktIndustryFlow.json",
			"args": map[string]string{
				"tradeDate": "date",   //输入一个日期，不输入其他请求参数，可获取到一天全部行业日资金流向数据，输入格式“YYYYMMDD”
				"beginDate": "date",   //起始日期，输入格式“YYYYMMDD”
				"endDate":   "date",   //截止日期，输入格式“YYYYMMDD”
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"tradeDate":      "date",   //交易日
				"industryID":     "string", //行业ID
				"industryName":   "string", //行业名称
				"moneyInflow":    "double", //资金流入
				"moneyOutflow":   "double", //资金流出
				"netMoneyInflow": "double", //资金净流入
			},
			"indices": map[string]string{},
		},
		//说明获取获取沪深A股和B股前复权日行情信息，包含前复权昨收价、前复权开盘价、前复权最高价、前复权最低价、前复权收盘价。(参数加上type=2，能够每日16:00更新；否则是17:20更新)
		"getMktEqudAdj": map[string]interface{}{
			"url": "/api/market/getMktEqudAdj.json",
			"args": map[string]string{
				"secID":     "string", //可多值输入）
				"ticker":    "string", //可多值输入）
				"tradeDate": "date",   //输入一个日期，不输入其他请求参数，可获取到一天全部沪深股票复权行情数据，输入格式“YYYYMMDD”
				"beginDate": "date",   //开始日期，输入格式“YYYYMMDD”
				"endDate":   "date",   //结束日期，输入格式“YYYYMMDD”
				"isOpen":    "Int",    //股票今日是否开盘标记位：0-今日未开盘，1-今日有开盘
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"ticker":           "string", //交易代码
				"tradeDate":        "date",   //交易日期
				"secID":            "string", //证券展示代码
				"exchangeCD":       "string", //交易市场
				"secShortName":     "string", //证券简称
				"secShortNameEn":   "string", //证券英文简称
				"preClosePrice":    "double", //昨收盘(前复权)
				"openPrice":        "double", //今开盘(前复权)
				"highestPrice":     "double", //最高价(前复权)
				"lowestPrice":      "double", //最低价(前复权)
				"closePrice":       "double", //今收盘(前复权)
				"turnoverValue":    "Double", //成交金额(元)
				"turnoverVol":      "Double", //成交量(股)
				"actPreClosePrice": "Double", //实际昨收盘
				"dealAmount":       "Int32",  //成交笔数
				"turnoverRate":     "Double", //日换手率
				"accumAdjFactor":   "Double", //累积前复权因子，前复权是对历史行情进行调整，除权除息日的行情不受本次除权除息影响，不需进行调整。最新一次除权除息日至最新行情期间的价格不需要进行任何的调整，该期间前复权...
				"negMarketValue":   "Double", //流通市值
				"marketValue":      "Double", //总市值
				"isOpen":           "Int",    //股票今日是否开盘标记位：0-今日未开盘，1-今日有开盘
			},
			"indices": map[string]string{},
		},
		//说明高频数据，获取多只A股历史上某一天的因子数据
		"getStockFactorsOneDay": map[string]interface{}{
			"url": "/api/market/getStockFactorsOneDay.json",
			"args": map[string]string{
				"tradeDate": "int16",  //日期
				"secID":     "string", //多只股票ID，用逗号隔开
				"ticker":    "int16",  //多只股票交易代码，用逗号隔开
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"ticker":                 "string", //股票代码
				"secID":                  "string", //证券ID
				"tradeDate":              "string", //交易日期
				"AccountsPayablesTDays":  "double", //因子描述：应付账款周转天数（Accounts payable turnover days）
				"AccountsPayablesTRate":  "double", //因子描述：应付账款周转率（Accounts payable turnover rate）
				"AdminiExpenseRate":      "double", //因子描述：管理费用与营业总收入之比（Administrative expense rate）
				"ARTDays":                "double", //因子描述：应收账款周转天数（Accounts receivable turnover days）
				"ARTRate":                "double", //因子描述：应收账款周转率（Accounts receivable turnover rate）
				"ASSI":                   "double", //因子描述：对数总资产（Natural logarithm of total assets）
				"BLEV":                   "double", //因子描述：账面杠杆（Book leverage）
				"BondsPayableToAsset":    "double", //因子描述：应付债券与总资产之比（Bonds payable to total assets）
				"CashRateOfSales":        "double", //因子描述：经营活动产生的现金流量净额与营业收入之比（Cash rate of sales）
				"CashToCurrentLiability": "double", //因子描述：现金比率（Cash to current liability）
				"CMRA":               "double", //因子描述：24 月累计收益（Monthly cumulative return range over the past 24 months）
				"CTOP":               "double", //因子描述：现金流市值比（Cash flow to price）
				"CTP5":               "double", //因子描述：5 年平均现金流市值比（Five-year average cash flow to price）
				"CurrentAssetsRatio": "double", //因子描述：流动资产比率（Current assets ratio）
				"CurrentAssetsTRate": "double", //因子描述：流动资产周转率（Current assets turnover rate）
				"CurrentRatio":       "double", //因子描述：流动比率（Current ratio）
				"DAVOL10":            "double", //10日相对换手率
				"DAVOL20":            "double", //20日相对换手率
				"DAVOL5":             "double", //5日相对换手率
				"DDNBT":              "double", //因子描述：下跌贝塔（Downside beta），过往12 个月中，市场组合日收益为负时，个股日收益关于市场组合日收益的回归系数。
				"DDNCR":              "double", //因子描述：下跌相关系数（Downside correlation），过往12个月中，市场组合日收益为负时，个股日收益关于市场组合日收益的相关系数。
				"DDNSR":              "double", //因子描述：下跌波动（Downside standard deviations ratio），过往12个月中，市场组合日收益为负时，个股日收益标准差和市场组合日收益标准差之比。
				"DebtEquityRatio":    "double", //因子描述：产权比率（Debt equity ratio）
				"DebtsAssetRatio":    "double", //因子描述：债务总资产比（Debt to total assets）
				"DHILO":              "double", //波幅中位数
				"DilutedEPS":         "double", //因子描述：稀释每股收益（Diluted earnings per share trailing twelve months）
				"DVRAT":              "double", //因子描述：收益相对波动（Daily returns variance ratio-serial dependence in daily returns）。
				"EBITToTOR":          "double", //因子描述：息税前利润与营业总收入之比（Earnings before interest and tax to total operating revenues）
				"EGRO":               "double", //因子描述：5 年收益增长率（Five-year earnings growth）
				"EMA10":              "double", //10日指数移动均线
				"EMA120":             "double", //120日指数移动均线
				"EMA20":              "double", //20日指数移动均线
				"EMA5":               "double", //5日指数移动均线
				"EMA60":              "double", //60日指数移动均线
				"EPS":                "double", //因子描述：每股收益（Earnings per share trailing twelve months）
				"EquityFixedAssetRatio":        "double", //因子描述：股东权益与固定资产比率（Equity fixed assets ratio）
				"EquityToAsset":                "double", //因子描述：股东权益比率（Equity to total assets）
				"EquityTRate":                  "double", //因子描述：股东权益周转率（Equity turnover rate）
				"ETOP":                         "double", //因子描述：收益市值比（Earnings to price）
				"ETP5":                         "double", //因子描述：5 年平均收益市值比（Five-year average earnings to price）
				"FinancialExpenseRate":         "double", //因子描述：财务费用与营业总收入之比（Financial expense rate）
				"FinancingCashGrowRate":        "double", //因子描述：筹资活动产生的现金流量净额增长率（Growth rate of cash provided by (used for) financing
				"FixAssetRatio":                "double", //因子描述：固定资产比率（Fixed assets ratio）
				"FixedAssetsTRate":             "double", //因子描述：固定资产周转率（Fixed assets turnover rate）
				"GrossIncomeRatio":             "double", //因子描述：销售毛利率（Gross income ratio）
				"HBETA":                        "double", //因子描述：历史贝塔（Historical daily beta），过往12个月中，个股日收益关于市场组合日收益的三阶自回归，市场组合日收益的系数。
				"HSIGMA":                       "double", //因子描述：历史波动（Historical daily sigma），过往12个月中，个股日收益关于市场组合日收益的三阶自回归，市场组合日收益的残差标准差。
				"IntangibleAssetRatio":         "double", //因子描述：无形资产比率（Intangible assets ratio）
				"InventoryTDays":               "double", //因子描述：存货周转天数（Inventory turnover days）
				"InventoryTRate":               "double", //因子描述：存货周转率（Inventory turnover rate）
				"InvestCashGrowRate":           "double", //因子描述：投资活动产生的现金流量净额增长率（Growth rate of cash flows from investments）
				"LCAP":                         "double", //因子描述：对数市值（Natural logarithm of total market values）
				"LFLO":                         "double", //因子描述：对数流通市值（Natural logarithm of float market values）
				"LongDebtToAsset":              "double", //因子描述：长期借款与资产总计之比（Long term loan to total assets）
				"LongDebtToWorkingCapital":     "double", //因子描述：长期负债与营运资金比率（Long term debt to working capital）
				"LongTermDebtToAsset":          "double", //因子描述：长期负债与资产总计之比（Long term debt to total assets）
				"MA10":                         "double", //10日移动均线
				"MA120":                        "double", //120日移动均线
				"MA20":                         "double", //20日移动均线
				"MA5":                          "double", //5日移动均线
				"MA60":                         "double", //60日移动均线
				"MAWVAD":                       "double", //威廉变异离散量的移动平均
				"MFI":                          "double", //资金流量指标
				"MLEV":                         "double", //因子描述：市场杠杆（Market leverage）
				"NetAssetGrowRate":             "double", //因子描述：净资产增长率（Net assets growth rate）
				"NetProfitGrowRate":            "double", //因子描述：净利润增长率（Net profit growth rate）
				"NetProfitRatio":               "double", //因子描述：销售净利率（Net profit ratio）
				"NOCFToOperatingNI":            "double", //因子描述：经营活动产生的现金流量净额与经营活动净收益之比（Cash provided by operations to
				"NonCurrentAssetsRatio":        "double", //因子描述：非流动资产比率（Non-current assets ratio）
				"NPParentCompanyGrowRate":      "double", //因子描述：归属母公司股东的净利润增长率（Growth rate of net income attributable to shareholders of parent company）
				"NPToTOR":                      "double", //因子描述：净利润与营业总收入之比（Net profit to total operating revenues）
				"OperatingExpenseRate":         "double", //因子描述：营业费用与营业总收入之比（Operating expense rate）
				"OperatingProfitGrowRate":      "double", //因子描述：营业利润增长率（Operating profit growth rate）
				"OperatingProfitRatio":         "double", //因子描述：营业利润率（Operating profit ratio）
				"OperatingProfitToTOR":         "double", //因子描述：营业利润与营业总收入之比（Operating profit to total operating revenues）
				"OperatingRevenueGrowRate":     "double", //因子描述：营业收入增长率（Operating revenue growth rate）
				"OperCashGrowRate":             "double", //因子描述：经营活动产生的现金流量净额增长率（Growth rate of cash provided by operations）
				"OperCashInToCurrentLiability": "double", //因子描述：现金流动负债比（Cash provided by operations to current liability）
				"PB":                  "double", //因子描述：市净率（Price-to-book ratio）
				"PCF":                 "double", //因子描述：市现率（Price-to-cash-flow ratio）
				"PE":                  "double", //因子描述：市盈率（Price-earnings ratio）
				"PS":                  "double", //因子描述：市销率（Price-to-sales ratio）
				"PSY":                 "double", //心理线指标
				"QuickRatio":          "double", //因子描述：速动比率（Quick ratio）
				"REVS10":              "double", //10日股价收益
				"REVS20":              "double", //20日股价收益
				"REVS5":               "double", //5日股价收益
				"ROA":                 "double", //因子描述：资产回报率（Return on assets）
				"ROA5":                "double", //因子描述：5 年资产回报率（Five-year average return on assets）
				"ROE":                 "double", //因子描述：权益回报率（Return on equity）
				"ROE5":                "double", //因子描述：5 年权益回报率（Five-year average return on equity）
				"RSI":                 "double", //相对强弱指标
				"RSTR12":              "double", //因子描述：12月相对强势（Relative strength for the last 12 months）
				"RSTR24":              "double", //因子描述：24月相对强势（Relative strength for the last 24 months）
				"SalesCostRatio":      "double", //因子描述：销售成本率（Sales cost ratio）
				"SaleServiceCashToOR": "double", //因子描述：销售商品提供劳务收到的现金与营业收入之比（Sale service cash to operating revenues）
				"SUE":                  "double", //因子描述：未预期盈余（Standardized unexpected earnings）
				"TaxRatio":             "double", //因子描述：销售税金率（Tax ratio）
				"TOBT":                 "double", //超额流动
				"TotalAssetGrowRate":   "double", //因子描述：总资产增长率（Total assets growth rate）
				"TotalAssetsTRate":     "double", //因子描述：总资产周转率（Total assets turnover rate）
				"TotalProfitCostRatio": "double", //因子描述：成本费用利润率（Total profit cost ratio）
				"TotalProfitGrowRate":  "double", //因子描述：利润总额增长率（Total profit growth rate）
				"VOL10":                "double", //10日换手率
				"VOL120":               "double", //120日换手率
				"VOL20":                "double", //20日换手率
				"VOL240":               "double", //240日换手率
				"VOL5":                 "double", //5日换手率
				"VOL60":                "double", //60日换手率
				"WVAD":                 "double", //威廉变异离散量
				"REC":                  "double", //因子描述：分析师推荐评级（Recommended rating score by analyst）
				"DAREC":                "double", //因子描述：分析师推荐评级变化（Changes of recommended rating score by analyst），相比于60 个交易日前。
				"GREC":                 "double", //因子描述：分析师推荐评级变化趋势（Change tendency of recommended rating score by analyst），过去60 个交易日内的  符号加和。
				"FY12P":                "double", //因子描述：分析师盈利预测（Forecast earnings by analyst to market values）
				"DAREV":                "double", //因子描述：分析师盈利预测变化（Changes of forecast earnings by analyst），相比于60个交易日前。
				"GREV":                 "double", //因子描述：分析师盈利预测变化趋势（Change tendency of forecast earnings by analyst），过去60 个
				"SFY12P":               "double", //因子描述：分析师营收预测（Forecast sales by analyst to market values）
				"DASREV":               "double", //因子描述：分析师盈收预测变化（Changes of forecast sales by analyst），相比于60 个交易日前。
				"GSREV":                "double", //因子描述：分析师盈收预测变化趋势（Change tendency of forecast sales by analyst），过去60 个交易日内的DASREV 符号加和。
				"FEARNG":               "double", //因子描述：未来预期盈利增长（Growth of forecast earnings）
				"FSALESG":              "double", //因子描述：未来预期盈收增长（Growth of forecast sales）
				"TA2EV":                "double", //因子描述：资产总计与企业价值之比（Assets to enterprise value）
				"CFO2EV":               "double", //因子描述：经营活动产生的现金流量净额与企业价值之比（Cash provided by operations to enterprise
				"ACCA":                 "double", //因子描述：现金流资产比和资产回报率之差（Cash flow assets ratio minus return on assets）
				"DEGM":                 "double", //因子描述：毛利率增长（Growth rate of gross income ratio），去年同期相比。
				"SUOI":                 "double", //因子描述：未预期毛利（Standardized unexpected gross income）
				"EARNMOM":              "double", //因子描述：八季度净利润变化趋势（Change tendency of net profit in the past eight quarters），前8个季度的净利润，如果同比（去年同期）增长记为+1，同比下滑记为-1，再将8 个值相加。
				"FiftyTwoWeekHigh":     "double", //因子描述：当前价格处于过去1 年股价的位置（Price level during the pasted 52 weeks）。
				"Volatility":           "double", //N
				"Skewness":             "double", //
				"ILLIQUIDITY":          "double", //因子描述：收益相对金额比（Daily return to turnover value during the last  days），过去20 个交易日收益相对金额的比例。
				"BackwardADJ":          "double", //因子描述：股价向后复权因子（Backward adjustment factor for restoration of rights）。复权是对股价和成交量进行权息修复，按照股票的实际涨跌绘制股价走势图，并把成交量调整为相同的股本口径。
				"MACD":                 "double", //指数平滑异同平均线
			},
			"indices": map[string]string{},
		},
		//说明获取股票月行情，包含月开高低收量价、涨跌幅、换手率等
		"getMktEqum": map[string]interface{}{
			"url": "/api/market/getMktEqum.json",
			"args": map[string]string{
				"monthEndDate": "date",   //交易日期，输入格式“YYYYMMDD”
				"secID":        "string", //证券内部编码，可通过交易代码在getSecID获取到。（可多值输入，最多输入50只）
				"ticker":       "string", //股票交易代码，如'000001'（可多值输入，最多输入50只）
				"beginDate":    "date",   //起始日期，输入格式“YYYYMMDD”
				"endDate":      "date",   //截止日期，输入格式“YYYYMMDD”
				"isOpen":       "int32",  //股票当月是否开盘标记位：0-未开盘，1-开盘
				"field":        "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"ticker":          "string", //证券交易代码
				"endDate":         "date",   //交易日
				"secID":           "string", //证券内部代码
				"secShortName":    "string", //证券简称
				"exchangeCD":      "string", //交易市场
				"tradeDays":       "int32",  //交易天数
				"preClosePrice":   "double", //上月收盘
				"openPrice":       "double", //本月开盘
				"highestPrice":    "double", //本月最高
				"lowestPrice":     "double", //本月最低
				"closePrice":      "double", //本月收盘
				"turnoverVol":     "double", //成交量
				"turnoverValue":   "double", //成交金额
				"chg":             "double", //涨跌额
				"chgPct":          "double", //涨跌幅
				"return":          "double", //月
				"turnoverRate":    "double", //月累计换手率
				"avgTurnoverRate": "double", //月
				"varReturn24":     "double", //月
				"sdReturn24":      "double", //月
				"avgReturn24":     "double", //月
				"varReturn60":     "double", //月
				"sdReturn60":      "double", //月
				"avgReturn60":     "double", //月
			},
			"indices": map[string]string{},
		},
		//说明主要记录基金每日后复权行情，包括开高低收、成交量、成交价格等
		"getMktFunddAdj": map[string]interface{}{
			"url": "/api/market/getMktFunddAdj.json",
			"args": map[string]string{
				"secID":     "string", //证券内部编码
				"ticker":    "string", //证券交易代码
				"beginDate": "date",   //起始日期
				"endDate":   "date",   //结束日期
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"ticker":         "string", //证券交易代码
				"tradeDate":      "date",   //交易日期
				"secID":          "string", //证券内部编码
				"secShortName":   "string", //证券中文简称
				"secShortNameEn": "string", //证券英文简称
				"exchangeCD":     "string", //交易所代码
				"preClosePrice":  "double", //昨收盘（后复权）
				"openPrice":      "double", //今开盘（后复权）
				"highestPrice":   "double", //最高价（后复权）
				"lowestPrice":    "double", //最低价（后复权）
				"closePrice":     "double", //收盘价（后复权）
				"turnoverVol":    "double", //成交量
				"turnoverValue":  "double", //成交金额
			},
			"indices": map[string]string{},
		},
		//说明获取沪深A股和B股调整行情的后复权因子数据（以上市价格为基准），包含除权除息日、除权除息事项具体数据、本次复权因子、累积复权因子以及因子调整的截止日期。该因子用来调整历史行情，不作为预测使用，于除权除息日进行计算调整。
		"getMktAdjfAf": map[string]interface{}{
			"url": "/api/market/getMktAdjfAf.json",
			"args": map[string]string{
				"exchangeCD": "string", //股票交易市场，可选择：XSHG上海证券交易所，XSHE深圳证券交易所。
				"exDivDate":  "date",   //除权除息日，输入格式：YYYYMMDD
				"secID":      "string", //一只或多只证券代码，用,分隔，格式是“数字.交易所代码”，如000001.XSHE。如果为空，则为全部证券。
				"ticker":     "string", //一只或多只股票代码，用,分隔，如000001,000002。
				"beginDate":  "date",   //除权除息日查询为基准，查询开始日期，输入格式："YYYYMMDD"
				"endDate":    "date",   //除权除息日查询为基准，查询截止日期，输入格式："YYYYMMDD"
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":              "string", //证券展示代码
				"exDivDate":          "date",   //除权除息日
				"ticker":             "string", //交易代码
				"exchangeCD":         "string", //交易市场
				"secShortName":       "string", //证券简称
				"secShortNameEn":     "string", //证券英文简称
				"perCashDiv":         "double", //每股派现
				"perShareDivRatio":   "double", //每股送股比例
				"perShareTransRatio": "double", //每股转增股比例
				"allotmentRatio":     "double", //每股配股比例
				"allotmentPrice":     "double", //配股价
				"splitsRatio":        "double", //拆股比例
				"adjFactor":          "double", //复权因子(后复权)
				"accumAdjFactor":     "double", //累积复权因子(后复权)
				"endDate":            "date",   //累积复权因子截止日期
			},
			"indices": map[string]string{},
		},
		//说明获取股票周前复权行情，包含开高低收量价、涨跌幅、换手率等
		"getMktEquwAdj": map[string]interface{}{
			"url": "/api/market/getMktEquwAdj.json",
			"args": map[string]string{
				"secID":       "string", //证券内部编码，可通过交易代码在getSecID获取到。（可多值输入，最多输入50只）
				"ticker":      "string", //股票交易代码，如'000001'（可多值输入，最多输入50只）
				"weekEndDate": "date",   //交易日期，输入格式“YYYYMMDD”
				"beginDate":   "date",   //起始日期，输入格式“YYYYMMDD”
				"endDate":     "date",   //截止日期，输入格式“YYYYMMDD”
				"isOpen":      "Int32",  //股票当周是否开盘标记位：0-今日未开盘，1-今日有开盘
				"field":       "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"ticker":          "string", //证券交易代码
				"endDate":         "date",   //交易日
				"secID":           "string", //证券内部代码
				"secShortName":    "string", //证券简称
				"exchangeCD":      "string", //交易市场
				"tradeDays":       "int32",  //交易天数
				"preClosePrice":   "double", //上周收盘
				"openPrice":       "double", //本周开盘
				"highestPrice":    "double", //本周最高
				"lowestPrice":     "double", //本周最低
				"closePrice":      "double", //本周收盘
				"turnoverVol":     "double", //成交量
				"turnoverValue":   "double", //成交金额
				"chg":             "double", //涨跌额
				"chgPct":          "double", //涨跌幅
				"return":          "double", //周回报率
				"turnoverRate":    "double", //周累计换手率
				"avgTurnoverRate": "double", //周平均换手率
				"varReturn100":    "double", //周回报率方差
				"sdReturn100":     "double", //周回报率标准差
				"avgReturn100":    "double", //周平均回报率
			},
			"indices": map[string]string{},
		},
		//说明获取香港交易所股票开、收、高、低，成交等日行情信息。
		"getMktHKEqud": map[string]interface{}{
			"url": "/api/market/getMktHKEqud.json",
			"args": map[string]string{
				"secID":     "string", //证券ID，证券统一编码，可通过交易代码在getSecID获取到。
				"ticker":    "string", //交易代码，如'00001'
				"tradeDate": "date",   //输入一个日期，不输入其他请求参数，可获取到一天全部港股日行情数据，输入格式“YYYYMMDD”
				"beginDate": "date",   //起始日期，输入格式“YYYYMMDD”
				"endDate":   "date",   //截止日期，输入格式“YYYYMMDD”
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"tradeDate":        "date",   //交易日
				"secID":            "string", //证券ID
				"ticker":           "string", //交易代码
				"exchangeCD":       "string", //交易市场
				"secShortName":     "string", //证券简称
				"preClosePrice":    "double", //前收价
				"actPreClosePrice": "double", //实际前收
				"openPrice":        "double", //开市价
				"highestPrice":     "double", //最高价
				"lowestPrice":      "double", //最低价
				"closePrice":       "double", //今收盘
				"turnoverVol":      "double", //成交股数
				"turnoverValue":    "double", //成交金额
				"SMA10":            "double", //10日均价
				"SMA20":            "double", //20日均价
				"SMA50":            "double", //50日均价
				"SMA250":           "double", //250日均价
			},
			"indices": map[string]string{},
		},
		//说明获取期货仓单日报情况，包括品种代码、交易市场、仓库、上期仓单量、本期仓单量等字段，每日16:00更新
		"getMktFutWRd": map[string]interface{}{
			"url": "/api/market/getMktFutWRd.json",
			"args": map[string]string{
				"beginDate":      "date",   //起始日期，默认为一周前，输入格式“YYYYMMDD”
				"contractObject": "string", //品种代码
				"endDate":        "date",   //截止日期，默认为当前日期，输入格式“YYYYMMDD”
				"exchangeCD":     "string", //交易市场
				"field":          "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"tradeDate":      "date",   //交易日期
				"exchangeCD":     "string", //交易市场
				"contractObject": "string", //品种代码
				"unit":           "string", //单位
				"warehouse":      "string", //仓库
				"preWrVOL":       "int32",  //上期仓单量
				"wrVOL":          "int32",  //本期仓单量
				"chg":            "int32",  //增减
			},
			"indices": map[string]string{},
		},
		//说明高频数据，获取一只A股历史上某一时间段的因子数据
		"getStockFactorsDateRange": map[string]interface{}{
			"url": "/api/market/getStockFactorsDateRange.json",
			"args": map[string]string{
				"secID":     "string", //一只股票ID
				"ticker":    "string", //一只股票交易代码
				"beginDate": "int16",  //开始日期
				"endDate":   "int16",  //结束日期
				"field":     "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"ticker":                 "string", //股票代码
				"secID":                  "string", //证券ID
				"tradeDate":              "string", //交易日期
				"AccountsPayablesTDays":  "double", //因子描述：应付账款周转天数（Accounts payable turnover days）
				"AccountsPayablesTRate":  "double", //因子描述：应付账款周转率（Accounts payable turnover rate）
				"AdminiExpenseRate":      "double", //因子描述：管理费用与营业总收入之比（Administrative expense rate）
				"ARTDays":                "double", //因子描述：应收账款周转天数（Accounts receivable turnover days）
				"ARTRate":                "double", //因子描述：应收账款周转率（Accounts receivable turnover rate）
				"ASSI":                   "double", //因子描述：对数总资产（Natural logarithm of total assets）
				"BLEV":                   "double", //因子描述：账面杠杆（Book leverage）
				"BondsPayableToAsset":    "double", //因子描述：应付债券与总资产之比（Bonds payable to total assets）
				"CashRateOfSales":        "double", //因子描述：经营活动产生的现金流量净额与营业收入之比（Cash rate of sales）
				"CashToCurrentLiability": "double", //因子描述：现金比率（Cash to current liability）
				"CMRA":               "double", //因子描述：24 月累计收益（Monthly cumulative return range over the past 24 months）
				"CTOP":               "double", //因子描述：现金流市值比（Cash flow to price）
				"CTP5":               "double", //因子描述：5 年平均现金流市值比（Five-year average cash flow to price）
				"CurrentAssetsRatio": "double", //因子描述：流动资产比率（Current assets ratio）
				"CurrentAssetsTRate": "double", //因子描述：流动资产周转率（Current assets turnover rate）
				"CurrentRatio":       "double", //因子描述：流动比率（Current ratio）
				"DAVOL10":            "double", //10日相对换手率
				"DAVOL20":            "double", //20日相对换手率
				"DAVOL5":             "double", //5日相对换手率
				"DDNBT":              "double", //因子描述：下跌贝塔（Downside beta），过往12 个月中，市场组合日收益为负时，个股日收益关于市场组合日收益的回归系数。
				"DDNCR":              "double", //因子描述：下跌相关系数（Downside correlation），过往12个月中，市场组合日收益为负时，个股日收益关于市场组合日收益的相关系数。
				"DDNSR":              "double", //因子描述：下跌波动（Downside standard deviations ratio），过往12个月中，市场组合日收益为负时，个股日收益标准差和市场组合日收益标准差之比。
				"DebtEquityRatio":    "double", //因子描述：产权比率（Debt equity ratio）
				"DebtsAssetRatio":    "double", //因子描述：债务总资产比（Debt to total assets）
				"DHILO":              "double", //波幅中位数
				"DilutedEPS":         "double", //因子描述：稀释每股收益（Diluted earnings per share trailing twelve months）
				"DVRAT":              "double", //因子描述：收益相对波动（Daily returns variance ratio-serial dependence in daily returns）。
				"EBITToTOR":          "double", //因子描述：息税前利润与营业总收入之比（Earnings before interest and tax to total operating revenues）
				"EGRO":               "double", //因子描述：5 年收益增长率（Five-year earnings growth）
				"EMA10":              "double", //10日指数移动均线
				"EMA120":             "double", //120日指数移动均线
				"EMA20":              "double", //20日指数移动均线
				"EMA5":               "double", //5日指数移动均线
				"EMA60":              "double", //60日指数移动均线
				"EPS":                "double", //因子描述：每股收益（Earnings per share trailing twelve months）
				"EquityFixedAssetRatio":        "double", //因子描述：股东权益与固定资产比率（Equity fixed assets ratio）
				"EquityToAsset":                "double", //因子描述：股东权益比率（Equity to total assets）
				"EquityTRate":                  "double", //因子描述：股东权益周转率（Equity turnover rate）
				"ETOP":                         "double", //因子描述：收益市值比（Earnings to price）
				"ETP5":                         "double", //因子描述：5 年平均收益市值比（Five-year average earnings to price）
				"FinancialExpenseRate":         "double", //因子描述：财务费用与营业总收入之比（Financial expense rate）
				"FinancingCashGrowRate":        "double", //因子描述：筹资活动产生的现金流量净额增长率（Growth rate of cash provided by (used for) financing
				"FixAssetRatio":                "double", //因子描述：固定资产比率（Fixed assets ratio）
				"FixedAssetsTRate":             "double", //因子描述：固定资产周转率（Fixed assets turnover rate）
				"GrossIncomeRatio":             "double", //因子描述：销售毛利率（Gross income ratio）
				"HBETA":                        "double", //因子描述：历史贝塔（Historical daily beta），过往12个月中，个股日收益关于市场组合日收益的三阶自回归，市场组合日收益的系数。
				"HSIGMA":                       "double", //因子描述：历史波动（Historical daily sigma），过往12个月中，个股日收益关于市场组合日收益的三阶自回归，市场组合日收益的残差标准差。
				"IntangibleAssetRatio":         "double", //因子描述：无形资产比率（Intangible assets ratio）
				"InventoryTDays":               "double", //因子描述：存货周转天数（Inventory turnover days）
				"InventoryTRate":               "double", //因子描述：存货周转率（Inventory turnover rate）
				"InvestCashGrowRate":           "double", //因子描述：投资活动产生的现金流量净额增长率（Growth rate of cash flows from investments）
				"LCAP":                         "double", //因子描述：对数市值（Natural logarithm of total market values）
				"LFLO":                         "double", //因子描述：对数流通市值（Natural logarithm of float market values）
				"LongDebtToAsset":              "double", //因子描述：长期借款与资产总计之比（Long term loan to total assets）
				"LongDebtToWorkingCapital":     "double", //因子描述：长期负债与营运资金比率（Long term debt to working capital）
				"LongTermDebtToAsset":          "double", //因子描述：长期负债与资产总计之比（Long term debt to total assets）
				"MA10":                         "double", //10日移动均线
				"MA120":                        "double", //120日移动均线
				"MA20":                         "double", //20日移动均线
				"MA5":                          "double", //5日移动均线
				"MA60":                         "double", //60日移动均线
				"MAWVAD":                       "double", //威廉变异离散量的移动平均
				"MFI":                          "double", //资金流量指标
				"MLEV":                         "double", //因子描述：市场杠杆（Market leverage）
				"NetAssetGrowRate":             "double", //因子描述：净资产增长率（Net assets growth rate）
				"NetProfitGrowRate":            "double", //因子描述：净利润增长率（Net profit growth rate）
				"NetProfitRatio":               "double", //因子描述：销售净利率（Net profit ratio）
				"NOCFToOperatingNI":            "double", //因子描述：经营活动产生的现金流量净额与经营活动净收益之比（Cash provided by operations to
				"NonCurrentAssetsRatio":        "double", //因子描述：非流动资产比率（Non-current assets ratio）
				"NPParentCompanyGrowRate":      "double", //因子描述：归属母公司股东的净利润增长率（Growth rate of net income attributable to shareholders of parent company）
				"NPToTOR":                      "double", //因子描述：净利润与营业总收入之比（Net profit to total operating revenues）
				"OperatingExpenseRate":         "double", //因子描述：营业费用与营业总收入之比（Operating expense rate）
				"OperatingProfitGrowRate":      "double", //因子描述：营业利润增长率（Operating profit growth rate）
				"OperatingProfitRatio":         "double", //因子描述：营业利润率（Operating profit ratio）
				"OperatingProfitToTOR":         "double", //因子描述：营业利润与营业总收入之比（Operating profit to total operating revenues）
				"OperatingRevenueGrowRate":     "double", //因子描述：营业收入增长率（Operating revenue growth rate）
				"OperCashGrowRate":             "double", //因子描述：经营活动产生的现金流量净额增长率（Growth rate of cash provided by operations）
				"OperCashInToCurrentLiability": "double", //因子描述：现金流动负债比（Cash provided by operations to current liability）
				"PB":                  "double", //因子描述：市净率（Price-to-book ratio）
				"PCF":                 "double", //因子描述：市现率（Price-to-cash-flow ratio）
				"PE":                  "double", //因子描述：市盈率（Price-earnings ratio）
				"PS":                  "double", //因子描述：市销率（Price-to-sales ratio）
				"PSY":                 "double", //心理线指标
				"QuickRatio":          "double", //因子描述：速动比率（Quick ratio）
				"REVS10":              "double", //10日股价收益
				"REVS20":              "double", //20日股价收益
				"REVS5":               "double", //5日股价收益
				"ROA":                 "double", //因子描述：资产回报率（Return on assets）
				"ROA5":                "double", //因子描述：5 年资产回报率（Five-year average return on assets）
				"ROE":                 "double", //因子描述：权益回报率（Return on equity）
				"ROE5":                "double", //因子描述：5 年权益回报率（Five-year average return on equity）
				"RSI":                 "double", //相对强弱指标
				"RSTR12":              "double", //因子描述：12月相对强势（Relative strength for the last 12 months）
				"RSTR24":              "double", //因子描述：24月相对强势（Relative strength for the last 24 months）
				"SalesCostRatio":      "double", //因子描述：销售成本率（Sales cost ratio）
				"SaleServiceCashToOR": "double", //因子描述：销售商品提供劳务收到的现金与营业收入之比（Sale service cash to operating revenues）
				"SUE":                  "double", //因子描述：未预期盈余（Standardized unexpected earnings）
				"TaxRatio":             "double", //因子描述：销售税金率（Tax ratio）
				"TOBT":                 "double", //超额流动
				"TotalAssetGrowRate":   "double", //因子描述：总资产增长率（Total assets growth rate）
				"TotalAssetsTRate":     "double", //因子描述：总资产周转率（Total assets turnover rate）
				"TotalProfitCostRatio": "double", //因子描述：成本费用利润率（Total profit cost ratio）
				"TotalProfitGrowRate":  "double", //因子描述：利润总额增长率（Total profit growth rate）
				"VOL10":                "double", //10日换手率
				"VOL120":               "double", //120日换手率
				"VOL20":                "double", //20日换手率
				"VOL240":               "double", //240日换手率
				"VOL5":                 "double", //5日换手率
				"VOL60":                "double", //60日换手率
				"WVAD":                 "double", //威廉变异离散量
				"REC":                  "double", //因子描述：分析师推荐评级（Recommended rating score by analyst）
				"DAREC":                "double", //因子描述：分析师推荐评级变化（Changes of recommended rating score by analyst），相比于60 个交易日前。
				"GREC":                 "double", //因子描述：分析师推荐评级变化趋势（Change tendency of recommended rating score by analyst），过去60 个交易日内的  符号加和。
				"FY12P":                "double", //因子描述：分析师盈利预测（Forecast earnings by analyst to market values）
				"DAREV":                "double", //因子描述：分析师盈利预测变化（Changes of forecast earnings by analyst），相比于60个交易日前。
				"GREV":                 "double", //因子描述：分析师盈利预测变化趋势（Change tendency of forecast earnings by analyst），过去60 个
				"SFY12P":               "double", //因子描述：分析师营收预测（Forecast sales by analyst to market values）
				"DASREV":               "double", //因子描述：分析师盈收预测变化（Changes of forecast sales by analyst），相比于60 个交易日前。
				"GSREV":                "double", //因子描述：分析师盈收预测变化趋势（Change tendency of forecast sales by analyst），过去60 个交易日内的DASREV 符号加和。
				"FEARNG":               "double", //因子描述：未来预期盈利增长（Growth of forecast earnings）
				"FSALESG":              "double", //因子描述：未来预期盈收增长（Growth of forecast sales）
				"TA2EV":                "double", //因子描述：资产总计与企业价值之比（Assets to enterprise value）
				"CFO2EV":               "double", //因子描述：经营活动产生的现金流量净额与企业价值之比（Cash provided by operations to enterprise
				"ACCA":                 "double", //因子描述：现金流资产比和资产回报率之差（Cash flow assets ratio minus return on assets）
				"DEGM":                 "double", //因子描述：毛利率增长（Growth rate of gross income ratio），去年同期相比。
				"SUOI":                 "double", //因子描述：未预期毛利（Standardized unexpected gross income）
				"EARNMOM":              "double", //因子描述：八季度净利润变化趋势（Change tendency of net profit in the past eight quarters），前8个季度的净利润，如果同比（去年同期）增长记为+1，同比下滑记为-1，再将8 个值相加。
				"FiftyTwoWeekHigh":     "double", //因子描述：当前价格处于过去1 年股价的位置（Price level during the pasted 52 weeks）。
				"Volatility":           "double", //N
				"Skewness":             "double", //
				"ILLIQUIDITY":          "double", //因子描述：收益相对金额比（Daily return to turnover value during the last  days），过去20 个交易日收益相对金额的比例。
				"BackwardADJ":          "double", //因子描述：股价向后复权因子（Backward adjustment factor for restoration of rights）。复权是对股价和成交量进行权息修复，按照股票的实际涨跌绘制股价走势图，并把成交量调整为相同的股本口径。
				"MACD":                 "double", //指数平滑异同平均线
			},
			"indices": map[string]string{},
		},
		//说明上海证券交易所、深圳证券交易所今日停复牌股票列表。数据更新频率：日。
		"getSecTips": map[string]interface{}{
			"url": "/api/market/getSecTips.json",
			"args": map[string]string{
				"tipsTypeCD": "string", //交易提示类型：H表示停牌，R表示复牌；支持多值输入
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"ticker":       "string", //交易代码
				"secID":        "string", //证券ID
				"exchangeCD":   "string", //交易市场
				"secShortName": "string", //证券简称
				"tipsDesc":     "string", //交易提示描述
			},
			"indices": map[string]string{},
		},
		//说明获取一只股票当日资金流向
		"getTickRTIntraDayMoneyFlowOrder": map[string]interface{}{
			"url": "/api/market/getTickRTIntraDayMoneyFlowOrder.json",
			"args": map[string]string{
				"securityID": "string", //一只证券代码，格式是“数字.交易所代码”，如000001.XSHG。证券可以是股票，指数， 部分债券或 基金。
				"endTime":    "string", //结束时间，必须小于等于15:05，例如14:20，如果为空则是15：05
				"startTime":  "string", //开始时间，必须大于等于09:25，例如10:05，如果为空则是09:25
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"dataDate":   "string", //交易日期(yyyy-MM-dd)
				"dataTime":   "string", //交易时间(HH:mm:ss)
				"ticker":     "string", //证券交易代码
				"exchangeCD": "string", //交易市场编码
			},
			"indices": map[string]string{},
		},
		//基本面数据
		//说明1、根据2007年新会计准则制定的合并利润表模板，收集了2007年以来沪深上市公司定期报告中各个会计期间的利润表数据； 2、仅收集合并报表数据，包括本期和上期数据； 3、如果上市公司对外财务报表进行更正，调整，均有采集并对外展示； 4、本表中单位为人民币元； 5、每季更新。
		"getFdmtIS": map[string]interface{}{
			"url": "/api/fundamental/getFdmtIS.json",
			"args": map[string]string{
				"secID":            "string", //证券内部编码,可通过交易代码和交易市场在getSecurityID获取到,如'000002.XSHE'
				"ticker":           "string", //股票代码，如'000002'
				"beginDate":        "date",   //会计期间截止日期,起始时间,如‘20121231’
				"endDate":          "date",   //会计期间截止日期,结束时间,如‘20131231’
				"fiscalPeriod":     "String", //会计期间长度,可多值输入，如：3，6，9，12，3表示季度值
				"publishDateBegin": "date",   //证券交易所披露的信息发布日期,起始时间,默认为1年前当前日期,如‘20130407’
				"publishDateEnd":   "date",   //证券交易所披露的信息发布日期,结束时间,默认为当前日期,如‘20140407’
				"beginDateRep":     "date",   //报表会计期末，起始时间，输入格式“YYYYMMDD”
				"endDateRep":       "date",   //报表会计期末，结束时间，输入格式“YYYYMMDD”
				"beginYear":        "string", //会计期间所属年份，起始时间，输入格式"YYYY"
				"endYear":          "string", //会计期间所属年份，结束时间，输入格式"YYYY"
				"reportType":       "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，Q3-第三季报，CQ3-三季报（累计1-9月），A-年报。
				"field":            "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"reportType":         "string", //报告类型
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"tRevenue":           "double", //营业总收入
				"revenue":            "double", //营业收入
				"intIncome":          "double", //利息收入
				"intExp":             "double", //利息支出
				"premEarned":         "double", //已赚保费
				"commisIncome":       "double", //手续费及佣金收入
				"commisExp":          "double", //手续费及佣金支出
				"TCogs":              "double", //营业总成本
				"COGS":               "double", //营业成本
				"premRefund":         "double", //退保金
				"NCompensPayout":     "double", //赔付支出净额
				"reserInsurContr":    "double", //提取保险合同准备金净额
				"policyDivPayt":      "double", //保单红利支出
				"reinsurExp":         "double", //分保费用
				"bizTaxSurchg":       "double", //营业税金及附加
				"sellExp":            "double", //销售费用
				"adminExp":           "double", //管理费用
				"finanExp":           "double", //财务费用
				"assetsImpairLoss":   "double", //资产减值损失
				"fValueChgGain":      "double", //公允价值变动收益
				"investIncome":       "double", //投资收益
				"AJInvestIncome":     "double", //其中:对联营企业和合营企业的投资收益
				"forexGain":          "double", //汇兑收益
				"operateProfit":      "double", //营业利润
				"NoperateIncome":     "double", //营业外收入
				"NoperateExp":        "double", //营业外支出
				"NCADisploss":        "double", //非流动资产处置损失
				"TProfit":            "double", //利润总额
				"incomeTax":          "double", //所得税费用
				"NIncome":            "double", //净利润
				"NIncomeAttrP":       "double", //归属于母公司所有者的净利润
				"minorityGain":       "double", //少数股东损益
				"basicEPS":           "double", //基本每股收益
				"dilutedEPS":         "double", //稀释每股收益
				"othComprIncome":     "double", //其他综合收益
				"TComprIncome":       "double", //综合收益总额
				"comprIncAttrP":      "double", //归属于母公司所有者的综合收益总额
				"comprIncAttrMS":     "double", //归属于少数股东的综合收益总额
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的银行业利润表模板，收集了2007年以来沪深上市公司定期报告中所有以此模板披露的利润表数据；（主要是银行业上市公司） 2、仅收集合并报表数据，包括本期和上期数据； 3、如果上市公司对外财务报表进行更正，调整，均有采集并对外展示； 4、本表中单位为人民币元； 5、每季更新。
		"getFdmtISBank": map[string]interface{}{
			"url": "/api/fundamental/getFdmtISBank.json",
			"args": map[string]string{
				"secID":            "string", //证券内部编码,可通过交易代码和交易市场在getSecurityID获取到,如'000001.XSHE'
				"ticker":           "string", //股票代码，如'000001'
				"beginDate":        "date",   //会计期间截止日期,起始时间,如‘20121231’
				"endDate":          "date",   //会计期间截止日期,结束时间,如‘20131231’
				"fiscalPeriod":     "String", //会计期间长度,可多值输入，如：3，6，9，12，3表示季度值
				"publishDateBegin": "date",   //证券交易所披露的信息发布日期,起始时间,默认为1年前当前日期,如‘20130407’
				"publishDateEnd":   "date",   //证券交易所披露的信息发布日期,结束时间,默认为当前日期,如‘20140407’
				"beginDateRep":     "date",   //报表会计期末，起始时间，输入格式“YYYYMMDD”
				"endDateRep":       "date",   //报表会计期末，结束时间，输入格式“YYYYMMDD”
				"beginYear":        "string", //会计期间所属年份，起始时间，输入格式"YYYY"
				"endYear":          "string", //会计期间所属年份，结束时间，输入格式"YYYY"
				"reportType":       "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，Q3-第三季报，CQ3-三季报（累计1-9月），A-年报。
				"field":            "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"reportType":         "string", //报告类型
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"revenue":            "double", //营业收入
				"NIntIncome":         "double", //利息净收入
				"intIncome":          "double", //利息收入
				"intExp":             "double", //利息支出
				"NCommisIncome":      "double", //手续费及佣金净收入
				"commisIncome":       "double", //手续费及佣金收入
				"commisExp":          "double", //手续费及佣金支出
				"othOperRev":         "double", //其他业务收入
				"COGS":               "double", //营业支出
				"bizTaxSurchg":       "double", //营业税金及附加
				"genlAdminExp":       "double", //业务及管理费
				"assetsImpairLoss":   "double", //资产减值损失
				"othOperCosts":       "double", //其他业务成本
				"fValueChgGain":      "double", //公允价值变动收益
				"investIncome":       "double", //投资收益
				"AJInvestIncome":     "double", //其中:对联营企业和合营企业的投资收益
				"forexGain":          "double", //汇兑收益
				"operateProfit":      "double", //营业利润
				"NoperateIncome":     "double", //营业外收入
				"NoperateExp":        "double", //营业外支出
				"TProfit":            "double", //利润总额
				"incomeTax":          "double", //所得税费用
				"NIncome":            "double", //净利润
				"NIncomeAttrP":       "double", //归属于母公司所有者的净利润
				"minorityGain":       "double", //少数股东损益
				"basicEPS":           "double", //基本每股收益
				"dilutedEPS":         "double", //稀释每股收益
				"othComprIncome":     "double", //其他综合收益
				"TComprIncome":       "double", //综合收益总额
				"comprIncAttrP":      "double", //归属于母公司所有者的综合收益总额
				"comprIncAttrMS":     "double", //归属于少数股东的综合收益总额
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的证券业利润表模板，收集了2007年以来沪深上市公司定期报告中所有以此模板披露的利润表数据；（主要是证券业上市公司） 2、仅收集合并报表数据，包括本期和上期数据； 3、如果上市公司对外财务报表进行更正，调整，均有采集并对外展示； 4、本表中单位为人民币元； 5、每季更新。
		"getFdmtISSecu": map[string]interface{}{
			"url": "/api/fundamental/getFdmtISSecu.json",
			"args": map[string]string{
				"secID":            "string", //证券内部编码,可通过交易代码和交易市场在getSecurityID获取到,如'600369.XSHG'
				"ticker":           "string", //股票代码，如'600369'
				"beginDate":        "date",   //会计期间截止日期,起始时间,如‘20121231’
				"endDate":          "date",   //会计期间截止日期,结束时间,如‘20131231’
				"fiscalPeriod":     "String", //会计期间长度,可多值输入，如：3，6，9，12，3表示季度值
				"publishDateBegin": "date",   //证券交易所披露的信息发布日期,起始时间,默认为1年前当前日期,如‘20130407’
				"publishDateEnd":   "date",   //证券交易所披露的信息发布日期,结束时间,默认为当前日期,如‘20140407’
				"beginDateRep":     "date",   //报表会计期末，起始时间，输入格式“YYYYMMDD”
				"endDateRep":       "date",   //报表会计期末，结束时间，输入格式“YYYYMMDD”
				"beginYear":        "string", //会计期间所属年份，起始时间，输入格式"YYYY"
				"endYear":          "string", //会计期间所属年份，结束时间，输入格式"YYYY"
				"reportType":       "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，Q3-第三季报，CQ3-三季报（累计1-9月），A-年报。
				"field":            "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"reportType":         "string", //报告类型
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"revenue":            "double", //营业收入
				"NIntIncome":         "double", //利息净收入
				"NCommisIncome":      "double", //手续费及佣金净收入
				"NSecTaIncome":       "double", //代理买卖证券业务净收入
				"NUndwrtSecIncome":   "double", //证券承销业务净收入
				"NTrustIncome":       "double", //委托客户资产管理业务净收入
				"othOperRev":         "double", //其他业务收入
				"COGS":               "double", //营业支出
				"bizTaxSurchg":       "double", //营业税金及附加
				"genlAdminExp":       "double", //业务及管理费
				"assetsImpairLoss":   "double", //资产减值损失
				"othOperCosts":       "double", //其他业务成本
				"fValueChgGain":      "double", //公允价值变动收益
				"investIncome":       "double", //投资收益
				"AJInvestIncome":     "double", //其中:对联营企业和合营企业的投资收益
				"forexGain":          "double", //汇兑收益
				"operateProfit":      "double", //营业利润
				"NoperateIncome":     "double", //营业外收入
				"NoperateExp":        "double", //营业外支出
				"TProfit":            "double", //利润总额
				"incomeTax":          "double", //所得税费用
				"NIncome":            "double", //净利润
				"NIncomeAttrP":       "double", //归属于母公司所有者的净利润
				"minorityGain":       "double", //少数股东损益
				"basicEPS":           "double", //基本每股收益
				"dilutedEPS":         "double", //稀释每股收益
				"othComprIncome":     "double", //其他综合收益
				"TComprIncome":       "double", //综合收益总额
				"comprIncAttrP":      "double", //归属于母公司所有者的综合收益总额
				"comprIncAttrMS":     "double", //归属于少数股东的综合收益总额
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的一般工商业利润表模板，收集了2007年以来沪深上市公司定期报告中所有以此模板披露的利润表数据；（主要是一般工商业上市公司） 2、仅收集合并报表数据，包括本期和上期数据； 3、如果上市公司对外财务报表进行更正，调整，均有采集并对外展示； 4、本表中单位为人民币元； 5、每季更新。
		"getFdmtISIndu": map[string]interface{}{
			"url": "/api/fundamental/getFdmtISIndu.json",
			"args": map[string]string{
				"secID":            "string", //证券内部编码,可通过交易代码和交易市场在getSecurityID获取到,如'000002.XSHE'
				"ticker":           "string", //股票代码，如'000002'
				"beginDate":        "date",   //会计期间截止日期,起始时间,如‘20121231’
				"endDate":          "date",   //会计期间截止日期,结束时间,如‘20131231’
				"fiscalPeriod":     "String", //会计期间长度,可多值输入，如：3，6，9，12，3表示季度值
				"publishDateBegin": "date",   //证券交易所披露的信息发布日期,起始时间,默认为1年前当前日期,如‘20130407’
				"publishDateEnd":   "date",   //证券交易所披露的信息发布日期,结束时间,默认为当前日期,如‘20140407’
				"beginDateRep":     "date",   //报表会计期末，起始时间，输入格式“YYYYMMDD”
				"endDateRep":       "date",   //报表会计期末，结束时间，输入格式“YYYYMMDD”
				"beginYear":        "string", //会计期间所属年份，起始时间，输入格式"YYYY"
				"endYear":          "string", //会计期间所属年份，结束时间，输入格式"YYYY"
				"reportType":       "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，Q3-第三季报，CQ3-三季报（累计1-9月），A-年报。
				"field":            "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"reportType":         "string", //报告类型
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"tRevenue":           "double", //营业总收入
				"revenue":            "double", //营业收入
				"intIncome":          "double", //利息收入
				"intExp":             "double", //利息支出
				"premEarned":         "double", //已赚保费
				"commisIncome":       "double", //手续费及佣金收入
				"commisExp":          "double", //手续费及佣金支出
				"TCogs":              "double", //营业总成本
				"COGS":               "double", //营业成本
				"premRefund":         "double", //退保金
				"NCompensPayout":     "double", //赔付支出净额
				"reserInsurContr":    "double", //提取保险合同准备金净额
				"policyDivPayt":      "double", //保单红利支出
				"reinsurExp":         "double", //分保费用
				"bizTaxSurchg":       "double", //营业税金及附加
				"sellExp":            "double", //销售费用
				"adminExp":           "double", //管理费用
				"finanExp":           "double", //财务费用
				"assetsImpairLoss":   "double", //资产减值损失
				"fValueChgGain":      "double", //公允价值变动收益
				"investIncome":       "double", //投资收益
				"AJInvestIncome":     "double", //其中:对联营企业和合营企业的投资收益
				"forexGain":          "double", //汇兑收益
				"operateProfit":      "double", //营业利润
				"NoperateIncome":     "double", //营业外收入
				"NoperateExp":        "double", //营业外支出
				"NCADisploss":        "double", //非流动资产处置损失
				"TProfit":            "double", //利润总额
				"incomeTax":          "double", //所得税费用
				"NIncome":            "double", //净利润
				"NIncomeAttrP":       "double", //归属于母公司所有者的净利润
				"minorityGain":       "double", //少数股东损益
				"basicEPS":           "double", //基本每股收益
				"dilutedEPS":         "double", //稀释每股收益
				"othComprIncome":     "double", //其他综合收益
				"TComprIncome":       "double", //综合收益总额
				"comprIncAttrP":      "double", //归属于母公司所有者的综合收益总额
				"comprIncAttrMS":     "double", //归属于少数股东的综合收益总额
			},
			"indices": map[string]string{},
		},
		//说明获取2007年及以后年度上市公司披露的业绩快报中的主要财务指标等其他数据，包括本期，去年同期，及本期与期初数值同比数据。每季证券交易所披露相关公告时更新数据，公司ipo时发布相关信息也会同时更新。每日9：00前完成证券交易所披露的数据更新，中午发布公告每日12：45前完成更新。
		"getFdmtEe": map[string]interface{}{
			"url": "/api/fundamental/getFdmtEe.json",
			"args": map[string]string{
				"secID":            "string", //证券内部编码,可通过交易代码和交易市场在getSecurityID获取到,如'600000.XSHG'
				"ticker":           "string", //股票交易代码，如'600000'，默认为'600000’
				"beginDate":        "date",   //会计期间截止日期,起始时间,如‘20121231’
				"endDate":          "date",   //会计期间截止日期,结束时间,如‘20131231’
				"publishDateBegin": "date",   //证券交易所披露的信息发布日期，起始时间，如‘20130812’
				"publishDateEnd":   "date",   //证券交易所披露的信息发布日期，结束时间，如‘20140812’
				"reportType":       "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，Q3-第三季报，CQ3-三季报（累计1-9月），A-年报。
				"field":            "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":              "string", //证券内部编码
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"partyID":            "int64",  //公司代码
				"ticker":             "string", //交易代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易所代码
				"actPubtime":         "string", //实际发布时间
				"mergedFlag":         "string", //合并标志
				"reportType":         "string", //报告类型
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币
				"revenue":            "double", //营业收入
				"primeOperRev":       "double", //主营业务收入
				"grossProfit":        "double", //主营业务利润
				"operateProfit":      "double", //营业利润
				"TProfit":            "double", //利润总额
				"NIncomeAttrP":       "double", //归属于母公司所有者的净利润(若公司未披露该科目，则为净利润数据)
				"NIncomeCut":         "double", //扣除非经常性损益后净利润
				"NCfOperA":           "double", //经营活动现金流量净额
				"basicEPS":           "double", //基本每股收益(元/股)
				"EPSW":               "double", //加权平均每股收益(元/股)
				"EPSCut":             "double", //扣除非经常损益后的基本每股收益(元/股)
				"EPSCutW":            "double", //扣除非经常损益后的加权每股收益(元/股)
				"ROE":                "double", //全面摊薄净资产收益率(%)
				"ROEW":               "double", //加权平均净资产收益率(%)
				"ROECut":             "double", //扣除非经常性损益的全面摊薄净资产收益率(%)
				"ROECutW":            "double", //扣除非经常性损益后的加权平均净资产收益率(%)
				"NCfOperAPs":         "double", //每股经营活动现金流量净额(元/股)
				"TAssets":            "double", //总资产
				"TEquityAttrP":       "double", //归属于母公司所有者权益合计
				"paidInCapital":      "double", //股本
				"NAssetPS":           "double", //每股净资产(元)
				"revenueLY":          "double", //上年同期营业收入
				"primeOperRevLY":     "double", //上年同期主营业务收入
				"grossProfitLY":      "double", //上年同期主营业务利润
				"operProfitLY":       "double", //上年同期营业利润
				"TProfitLY":          "double", //上年同期利润总额
				"NIncomeAttrPLY":     "double", //上年同期归属于母公司所有者的净利润(若公司未披露该科目，则为净利润数据)
				"NIncomeCutLY":       "double", //上年同期扣除非经常性损益后净利润（若公司快报中仅披露扣除非经常性损益后归属于上市公司股东的净利润，则该科目为扣除非经常性损益后归属于上市公司股东的净利润）
				"NCfOperALY":         "double", //上年同期经营活动现金流量净额
				"basicEPSLY":         "double", //上年同期基本每股收益(元/股)
				"EPSWLY":             "double", //上年同期加权平均每股收益(元/股)
				"EPSCutLY":           "double", //上年同期扣除非经常损益后的基本每股收益(元/股)
				"EPSCutWLY":          "double", //上年同期扣除非经常损益后的加权每股收益(元/股)
				"ROELY":              "double", //上年同期期末净资产收益率(%)
				"ROEWLY":             "double", //上年同期加权平均净资产收益率(%)
				"ROECutLY":           "double", //上年同期扣除非经常性损益的期末净资产收益率(%)
				"ROECutWLY":          "double", //上年同期扣除非经常性损益后的加权平均净资产收益率(%)
				"NCfOperAPsLY":       "double", //上年同期每股经营活动现金流量净额(元/股)
				"TAssetsLY":          "double", //期初总资产
				"TEquityAttrPLY":     "double", //期初归属于母公司所有者权益合计
				"NAssetPsLY":         "double", //期初每股净资产(元)
				"revenueYOY":         "double", //营业收入同比(%)
				"primeOperRevYOY":    "double", //主营业务收入同比(%)
				"grossProfitYOY":     "double", //主营业务利润同比(%)
				"operProfitYOY":      "double", //营业利润同比(%)
				"TProfitYOY":         "double", //利润总额同比(%)
				"NIncomeAttrPYOY":    "double", //归属于母公司所有者的净利润同比(%)(若公司未披露该科目，则为净利润数据)
				"NIncomeCutYOY":      "double", //扣除非经常性损益后净利润同比(%)
				"NCFOperAYOY":        "double", //经营活动现金流量净额同比(%)
				"basicEPSYOY":        "double", //基本每股收益同比(%)
				"EPSWYOY":            "double", //加权平均每股收益同比(%)
				"EPSCutYOY":          "double", //扣除非经常损益后的基本每股收益同比(%)
				"EPSCutWYOY":         "double", //扣除非经常损益后的加权每股收益同比(%)
				"ROEYOY":             "double", //期末净资产收益率同比(%)
				"ROEWYOY":            "double", //加权平均净资产收益率同比(%)
				"ROECutYOY":          "double", //扣除非经常性损益的期末净资产收益率同比(%)
				"ROECutWYOY":         "double", //扣除非经常性损益后的加权平均净资产收益率同比(%)
				"NCfOperAPsYOY":      "double", //每股经营活动现金流量净额同比(%)
				"TAssetsYOY":         "double", //总资产同比(%)
				"TEquityAttrPYOY":    "double", //归属于母公司所有者权益同比(%)
				"NAssetPsYOY":        "double", //每股净资产同比(%)
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的银行业资产负债表模板，收集了2007年以来沪深上市公司定期报告中所有以此模板披露的资产负债表数据；（主要是银行业上市公司） 2、仅收集合并报表数据，包括期末和期初数据； 3、如果上市公司对外财务报表进行更正，调整，展示上市公司最新披露一期财务数据； 4、本表中单位为人民币元； 5、每季更新。
		"getFdmtBSBankAllLatest": map[string]interface{}{
			"url": "/api/fundamental/getFdmtBSBankAllLatest.json",
			"args": map[string]string{
				"reportType": "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，Q3-第三季报，A-年报。
				"secID":      "string", //证券内部编码,可通过交易代码和交易市场在getSecID获取到,如'000001.XSHE'
				"ticker":     "string", //股票代码，如'000001'
				"beginDate":  "date",   //会计期间截止日期，起始时间，输入格式“YYYYMMDD”
				"endDate":    "date",   //会计期间截止日期，结束时间，输入格式“YYYYMMDD”
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志:1-合并
				"reportType":         "string", //报告类型
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"loanToOthBankFi":    "double", //拆出资金
				"tradingFA":          "double", //交易性金融资产
				"intReceiv":          "double", //应收利息
				"purResaleFa":        "double", //买入返售金融资产
				"disburLA":           "double", //发放委托贷款及垫款
				"availForSaleFA":     "double", //可供出售金融资产
				"htmInvest":          "double", //持有至到期投资
				"LTEquityInvest":     "double", //长期股权投资
				"investRealEstate":   "double", //投资性房地产
				"fixedAssets":        "double", //固定资产
				"CIP":                "double", //在建工程
				"intanAssets":        "double", //无形资产
				"goodwill":           "double", //商誉
				"deferTaxAssets":     "double", //递延所得税资产
				"CReserCB":           "double", //现金及存放中央银行款项
				"deposInOthBfi":      "double", //存放同业款项
				"preciMetals":        "double", //贵金属
				"derivAssets":        "double", //衍生金融资产
				"finanLeaseReceiv":   "double", //应收融资租赁款
				"investAsReceiv":     "double", //应收款项类投资
				"othAssets":          "double", //其他资产
				"TAssets":            "double", //资产总计
				"CBBorr":             "double", //向中央银行借款
				"depos":              "double", //吸收存款
				"loanFrOthBankFi":    "double", //拆入资金
				"tradingFL":          "double", //交易性金融负债
				"soldForRepurFa":     "double", //卖出回购金融资产款
				"payrollPayable":     "double", //应付职工薪酬
				"taxesPayable":       "double", //应交税费
				"intPayable":         "double", //应付利息
				"bondPayable":        "double", //应付债券
				"estimatedLiab":      "double", //预计负债
				"deferTaxLiab":       "double", //递延所得税负债
				"deposFrOthBfi":      "double", //同业及其他金融机构存放款项
				"derivLiab":          "double", //衍生金融负债
				"othLiab":            "double", //其他负债
				"TLiab":              "double", //负债合计
				"paidInCapital":      "double", //实收资本(或股本)
				"capitalReser":       "double", //资本公积
				"treasuryShare":      "double", //减:库存股
				"surplusReser":       "double", //盈余公积
				"ordinRiskReser":     "double", //一般风险准备
				"retainedEarnings":   "double", //未分配利润
				"forexDiffer":        "double", //外币报表折算差额
				"TEquityAttrP":       "double", //归属于母公司所有者权益合计
				"minorityInt":        "double", //少数股东权益
				"TShEquity":          "double", //所有者权益合计
				"TLiabEquity":        "double", //负债和所有者权益总计
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的证券业现金流量表模板，收集了2007年以来沪深上市公司定期报告中所有以此模板披露的现金流量表数据（主要是证券业上市公司）；2、仅收集合并报表数据，包括本期和上期数据；3、如果上市公司对外财务报表进行更正，调整，展示上市公司最新披露一期财务数据；4、本表中单位为人民币元；5、每季更新。
		"getFdmtCFSecuAllLatest": map[string]interface{}{
			"url": "/api/fundamental/getFdmtCFSecuAllLatest.json",
			"args": map[string]string{
				"reportType": "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，CQ3-三季报（累计1-9月），A-年报。
				"secID":      "string", //证券内部编码,可通过交易代码和交易市场在getSecID获取到,如'600369.XSHG '
				"ticker":     "string", //股票代码，如'600369'
				"beginDate":  "date",   //会计期间截止日期，起始时间，输入格式“YYYYMMDD”
				"endDate":    "date",   //会计期间截止日期，结束时间，输入格式“YYYYMMDD”
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志
				"reportType":         "string", //报告类型
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"NIncBorrOthFI":      "double", //向其他金融机构拆入资金净增加额
				"NIncDispTradFA":     "double", //处置交易性金融资产净增加额
				"NIncDispFaFS":       "double", //处置可供出售金融资产净增加额
				"IFCCashIncr":        "double", //收取利息、手续费及佣金的现金
				"NIncFrBorr":         "double", //拆入资金净增加额
				"NCApIncrRepur":      "double", //回购业务资金净增加额
				"refundOfTax":        "double", //收到的税费返还
				"CFrOthOperateA":     "double", //收到其他与经营活动有关的现金
				"CInfFrOperateA":     "double", //经营活动现金流入小计
				"CPaidIFC":           "double", //支付利息、手续费及佣金的现金
				"CPaidToForEmpl":     "double", //支付给职工以及为职工支付的现金
				"CPaidForTaxes":      "double", //支付的各项税费
				"CPaidForOthOpA":     "double", //支付其他与经营活动有关的现金
				"COutfOperateA":      "double", //经营活动现金流出小计
				"NCFOperateA":        "double", //经营活动产生的现金流量净额
				"procSellInvest":     "double", //收回投资收到的现金
				"gainInvest":         "double", //取得投资收益收到的现金
				"dispFixAssetsOth":   "double", //处置固定资产、无形资产和其他长期资产收回的现金净额
				"NDispSubsOthBizC":   "double", //处置子公司及其他营业单位收到的现金净额
				"CFrOthInvestA":      "double", //收到其他与投资活动有关的现金
				"CInfFrInvestA":      "double", //投资活动现金流入小计
				"purFixAssetsOth":    "double", //购建固定资产、无形资产和其他长期资产支付的现金
				"CPaidInvest":        "double", //投资支付的现金
				"NCPaidAcquis":       "double", //取得子公司及其他营业单位支付的现金净额
				"CPaidOthInvestA":    "double", //支付其他与投资活动有关的现金
				"COutfFrInvestA":     "double", //投资活动现金流出小计
				"NCFFrInvestA":       "double", //投资活动产生的现金流量净额
				"CFrCapContr":        "double", //吸收投资收到的现金
				"CFrMinoSSubs":       "double", //其中:子公司吸收少数股东投资收到的现金
				"CFrBorr":            "double", //取得借款收到的现金
				"CFrIssueBond":       "double", //发行债券收到的现金
				"CFrOthFinanA":       "double", //收到其他与筹资活动有关的现金
				"CInfFrFinanA":       "double", //筹资活动现金流入小计
				"CPaidForDebts":      "double", //偿还债务支付的现金
				"CPaidDivProfInt":    "double", //分配股利、利润或偿付利息支付的现金
				"divProfSubsMinoS":   "double", //其中:子公司支付给少数股东的股利、利润
				"CPaidOthFinanA":     "double", //支付其他与筹资活动有关的现金
				"COutfFrFinanA":      "double", //筹资活动现金流出小计
				"NCFFrFinanA":        "double", //筹资活动产生的现金流量净额
				"forexEffects":       "double", //汇率变动对现金及现金等价物的影响
				"NChangeInCash":      "double", //现金及现金等价物净增加额
				"NCEBegBal":          "double", //加:期初现金及现金等价物余额
				"NCEEndBal":          "double", //期末现金及现金等价物余额
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的合并现金流量表模板，收集了2007年以来沪深上市公司定期报告中各个会计期间的现金流量表数据；2、仅收集合并报表数据，包括本期和上期数据； 3、如果上市公司对外财务报表进行更正，调整，展示上市公司最新披露一期财务数据；4、本表中单位为人民币元；5、每季更新。
		"getFdmtCFAllLatest": map[string]interface{}{
			"url": "/api/fundamental/getFdmtCFAllLatest.json",
			"args": map[string]string{
				"reportType": "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，CQ3-三季报（累计1-9月），A-年报。
				"secID":      "string", //证券内部编码,可通过交易代码和交易市场在getSecID获取到,如'000002.XSHE'
				"ticker":     "string", //股票代码，如'000002'
				"beginDate":  "date",   //会计期间截止日期，起始时间，输入格式“YYYYMMDD”
				"endDate":    "date",   //会计期间截止日期，结束时间，输入格式“YYYYMMDD”
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志
				"reportType":         "string", //报告类型
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"CFrSaleGS":          "double", //销售商品、提供劳务收到的现金
				"NDeposIncrCFI":      "double", //客户存款和同业存放款项净增加额
				"NIncrBorrFrCB":      "double", //向中央银行借款净增加额
				"NIncBorrOthFI":      "double", //向其他金融机构拆入资金净增加额
				"premFrOrigContr":    "double", //收到原保险合同保费取得的现金
				"NReinsurPrem":       "double", //收到再保险业务现金净额
				"NIncPhDeposInv":     "double", //保户储金及投资款净增加额
				"NIncDispTradFA":     "double", //处置交易性金融资产净增加额
				"IFCCashIncr":        "double", //收取利息、手续费及佣金的现金
				"NIncFrBorr":         "double", //拆入资金净增加额
				"NCApIncrRepur":      "double", //回购业务资金净增加额
				"refundOfTax":        "double", //收到的税费返还
				"CFrOthOperateA":     "double", //收到其他与经营活动有关的现金
				"CInfFrOperateA":     "double", //经营活动现金流入小计
				"CPaidGS":            "double", //购买商品、接受劳务支付的现金
				"NIncDisburOfLA":     "double", //客户贷款及垫款净增加额
				"NIncrDeposInFI":     "double", //存放中央银行和同业款项净增加额
				"origContrCIndem":    "double", //支付原保险合同赔付款项的现金
				"CPaidIFC":           "double", //支付利息、手续费及佣金的现金
				"CPaidPolDiv":        "double", //支付保单红利的现金
				"CPaidToForEmpl":     "double", //支付给职工以及为职工支付的现金
				"CPaidForTaxes":      "double", //支付的各项税费
				"CPaidForOthOpA":     "double", //支付其他与经营活动有关的现金
				"COutfOperateA":      "double", //经营活动现金流出小计
				"NCFOperateA":        "double", //经营活动产生的现金流量净额
				"procSellInvest":     "double", //收回投资收到的现金
				"gainInvest":         "double", //取得投资收益收到的现金
				"dispFixAssetsOth":   "double", //处置固定资产、无形资产和其他长期资产收回的现金净额
				"NDispSubsOthBizC":   "double", //处置子公司及其他营业单位收到的现金净额
				"CFrOthInvestA":      "double", //收到其他与投资活动有关的现金
				"CInfFrInvestA":      "double", //投资活动现金流入小计
				"purFixAssetsOth":    "double", //购建固定资产、无形资产和其他长期资产支付的现金
				"CPaidInvest":        "double", //投资支付的现金
				"NIncrPledgeLoan":    "double", //质押贷款净增加额
				"NCPaidAcquis":       "double", //取得子公司及其他营业单位支付的现金净额
				"CPaidOthInvestA":    "double", //支付其他与投资活动有关的现金
				"COutfFrInvestA":     "double", //投资活动现金流出小计
				"NCFFrInvestA":       "double", //投资活动产生的现金流量净额
				"CFrCapContr":        "double", //吸收投资收到的现金
				"CFrMinoSSubs":       "double", //其中:子公司吸收少数股东投资收到的现金
				"CFrBorr":            "double", //取得借款收到的现金
				"CFrIssueBond":       "double", //发行债券收到的现金
				"CFrOthFinanA":       "double", //收到其他与筹资活动有关的现金
				"CInfFrFinanA":       "double", //筹资活动现金流入小计
				"CPaidForDebts":      "double", //偿还债务支付的现金
				"CPaidDivProfInt":    "double", //分配股利、利润或偿付利息支付的现金
				"divProfSubsMinoS":   "double", //其中:子公司支付给少数股东的股利、利润
				"CPaidOthFinanA":     "double", //支付其他与筹资活动有关的现金
				"COutfFrFinanA":      "double", //筹资活动现金流出小计
				"NCFFrFinanA":        "double", //筹资活动产生的现金流量净额
				"forexEffects":       "double", //汇率变动对现金及现金等价物的影响
				"NChangeInCash":      "double", //现金及现金等价物净增加额
				"NCEBegBal":          "double", //加:期初现金及现金等价物余额
				"NCEEndBal":          "double", //期末现金及现金等价物余额
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的银行业利润表模板，收集了2007年以来沪深上市公司定期报告中所有以此模板披露的利润表数据（主要是银行业上市公司）；2、仅收集合并报表数据，包括本期和上期数据；3、如果上市公司对外财务报表进行更正，调整，展示上市公司最新披露一期财务数据；4、本表中单位为人民币元；5、每季更新。
		"getFdmtISBankAllLatest": map[string]interface{}{
			"url": "/api/fundamental/getFdmtISBankAllLatest.json",
			"args": map[string]string{
				"reportType": "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，Q3-第三季报，CQ3-三季报（累计1-9月），A-年报。
				"secID":      "string", //证券内部编码,可通过交易代码和交易市场在getSecID获取到,如'000001.XSHE'
				"ticker":     "string", //股票代码，如'000001'
				"beginDate":  "date",   //会计期间截止日期，起始时间，输入格式“YYYYMMDD”
				"endDate":    "date",   //会计期间截止日期，结束时间，输入格式“YYYYMMDD”
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志
				"reportType":         "string", //报告类型
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"revenue":            "double", //营业收入
				"NIntIncome":         "double", //利息净收入
				"intIncome":          "double", //利息收入
				"intExp":             "double", //利息支出
				"NCommisIncome":      "double", //手续费及佣金净收入
				"commisIncome":       "double", //手续费及佣金收入
				"commisExp":          "double", //手续费及佣金支出
				"othOperRev":         "double", //其他业务收入
				"COGS":               "double", //营业支出
				"bizTaxSurchg":       "double", //营业税金及附加
				"genlAdminExp":       "double", //业务及管理费
				"assetsImpairLoss":   "double", //资产减值损失
				"othOperCosts":       "double", //其他业务成本
				"fValueChgGain":      "double", //公允价值变动收益
				"investIncome":       "double", //投资收益
				"AJInvestIncome":     "double", //其中:对联营企业和合营企业的投资收益
				"forexGain":          "double", //汇兑收益
				"operateProfit":      "double", //营业利润
				"NoperateIncome":     "double", //营业外收入
				"NoperateExp":        "double", //营业外支出
				"TProfit":            "double", //利润总额
				"incomeTax":          "double", //所得税费用
				"NIncome":            "double", //净利润
				"NIncomeAttrP":       "double", //归属于母公司所有者的净利润
				"minorityGain":       "double", //少数股东损益
				"basicEPS":           "double", //基本每股收益
				"dilutedEPS":         "double", //稀释每股收益
				"othComprIncome":     "double", //其他综合收益
				"TComprIncome":       "double", //综合收益总额
				"comprIncAttrP":      "double", //归属于母公司所有者的综合收益总额
				"comprIncAttrMS":     "double", //归属于少数股东的综合收益总额
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的一般工商业利润表模板，收集了2007年以来沪深上市公司定期报告中所有以此模板披露的利润表数据（主要是一般工商业上市公司）；2、仅收集合并报表数据，包括本期和上期数据；3、如果上市公司对外财务报表进行更正，调整，展示上市公司最新披露一期财务数据；4、本表中单位为人民币元；5、每季更新。
		"getFdmtISInduAllLatest": map[string]interface{}{
			"url": "/api/fundamental/getFdmtISInduAllLatest.json",
			"args": map[string]string{
				"reportType": "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，Q3-第三季报，CQ3-三季报（累计1-9月），A-年报。
				"secID":      "string", //证券内部编码,可通过交易代码和交易市场在getSecID获取到,如'000002.XSHE'
				"ticker":     "string", //股票代码，如'000002'
				"beginDate":  "date",   //会计期间截止日期，起始时间，输入格式“YYYYMMDD”
				"endDate":    "date",   //会计期间截止日期，结束时间，输入格式“YYYYMMDD”
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志
				"reportType":         "string", //报告类型
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"tRevenue":           "double", //营业总收入
				"revenue":            "double", //营业收入
				"intIncome":          "double", //利息收入
				"intExp":             "double", //利息支出
				"premEarned":         "double", //已赚保费
				"commisIncome":       "double", //手续费及佣金收入
				"commisExp":          "double", //手续费及佣金支出
				"TCogs":              "double", //营业总成本
				"COGS":               "double", //营业成本
				"premRefund":         "double", //退保金
				"NCompensPayout":     "double", //赔付支出净额
				"reserInsurContr":    "double", //提取保险合同准备金净额
				"policyDivPayt":      "double", //保单红利支出
				"reinsurExp":         "double", //分保费用
				"bizTaxSurchg":       "double", //营业税金及附加
				"sellExp":            "double", //销售费用
				"adminExp":           "double", //管理费用
				"finanExp":           "double", //财务费用
				"assetsImpairLoss":   "double", //资产减值损失
				"fValueChgGain":      "double", //公允价值变动收益
				"investIncome":       "double", //投资收益
				"AJInvestIncome":     "double", //其中:对联营企业和合营企业的投资收益
				"forexGain":          "double", //汇兑收益
				"operateProfit":      "double", //营业利润
				"NoperateIncome":     "double", //营业外收入
				"NoperateExp":        "double", //营业外支出
				"NCADisploss":        "double", //非流动资产处置损失
				"TProfit":            "double", //利润总额
				"incomeTax":          "double", //所得税费用
				"NIncome":            "double", //净利润
				"NIncomeAttrP":       "double", //归属于母公司所有者的净利润
				"minorityGain":       "double", //少数股东损益
				"basicEPS":           "double", //基本每股收益
				"dilutedEPS":         "double", //稀释每股收益
				"othComprIncome":     "double", //其他综合收益
				"TComprIncome":       "double", //综合收益总额
				"comprIncAttrP":      "double", //归属于母公司所有者的综合收益总额
				"comprIncAttrMS":     "double", //归属于少数股东的综合收益总额
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的合并现金流量表模板，收集了2007年以来沪深上市公司定期报告中各个会计期间的现金流量表数据； 2、仅收集合并报表数据，包括本期和上期数据； 3、如果上市公司对外财务报表进行更正，调整，均有采集并对外展示； 4、本表中单位为人民币元； 5、每季更新。
		"getFdmtCF": map[string]interface{}{
			"url": "/api/fundamental/getFdmtCF.json",
			"args": map[string]string{
				"secID":            "string", //证券内部编码,可通过交易代码和交易市场在getSecurityID获取到,如'000002.XSHE'
				"ticker":           "string", //股票代码，如'000002'
				"beginDate":        "date",   //会计期间截止日期,起始时间,如‘20121231’
				"endDate":          "date",   //会计期间截止日期,结束时间,如‘20131231’
				"fiscalPeriod":     "String", //，如：3，6，9，12，3表示季度值
				"publishDateBegin": "date",   //证券交易所披露的信息发布日期,起始时间,默认为1年前当前日期,如‘20130407’
				"publishDateEnd":   "date",   //证券交易所披露的信息发布日期,结束时间,默认为当前日期,如‘20140407’
				"beginDateRep":     "date",   //报表会计期末，起始时间，输入格式“YYYYMMDD”
				"endDateRep":       "date",   //报表会计期末，结束时间，输入格式“YYYYMMDD”
				"beginYear":        "string", //会计期间所属年份，起始时间，输入格式"YYYY"
				"endYear":          "string", //会计期间所属年份，结束时间，输入格式"YYYY"
				"reportType":       "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，CQ3-三季报（累计1-9月），A-年报。
				"field":            "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"reportType":         "string", //报告类型
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"CFrSaleGS":          "double", //销售商品、提供劳务收到的现金
				"NDeposIncrCFI":      "double", //客户存款和同业存放款项净增加额
				"NIncrBorrFrCB":      "double", //向中央银行借款净增加额
				"NIncBorrOthFI":      "double", //向其他金融机构拆入资金净增加额
				"premFrOrigContr":    "double", //收到原保险合同保费取得的现金
				"NReinsurPrem":       "double", //收到再保险业务现金净额
				"NIncPhDeposInv":     "double", //保户储金及投资款净增加额
				"NIncDispTradFA":     "double", //处置交易性金融资产净增加额
				"IFCCashIncr":        "double", //收取利息、手续费及佣金的现金
				"NIncFrBorr":         "double", //拆入资金净增加额
				"NCApIncrRepur":      "double", //回购业务资金净增加额
				"refundOfTax":        "double", //收到的税费返还
				"CFrOthOperateA":     "double", //收到其他与经营活动有关的现金
				"CInfFrOperateA":     "double", //经营活动现金流入小计
				"CPaidGS":            "double", //购买商品、接受劳务支付的现金
				"NIncDisburOfLA":     "double", //客户贷款及垫款净增加额
				"NIncrDeposInFI":     "double", //存放中央银行和同业款项净增加额
				"origContrCIndem":    "double", //支付原保险合同赔付款项的现金
				"CPaidIFC":           "double", //支付利息、手续费及佣金的现金
				"CPaidPolDiv":        "double", //支付保单红利的现金
				"CPaidToForEmpl":     "double", //支付给职工以及为职工支付的现金
				"CPaidForTaxes":      "double", //支付的各项税费
				"CPaidForOthOpA":     "double", //支付其他与经营活动有关的现金
				"COutfOperateA":      "double", //经营活动现金流出小计
				"NCFOperateA":        "double", //经营活动产生的现金流量净额
				"procSellInvest":     "double", //收回投资收到的现金
				"gainInvest":         "double", //取得投资收益收到的现金
				"dispFixAssetsOth":   "double", //处置固定资产、无形资产和其他长期资产收回的现金净额
				"NDispSubsOthBizC":   "double", //处置子公司及其他营业单位收到的现金净额
				"CFrOthInvestA":      "double", //收到其他与投资活动有关的现金
				"CInfFrInvestA":      "double", //投资活动现金流入小计
				"purFixAssetsOth":    "double", //购建固定资产、无形资产和其他长期资产支付的现金
				"CPaidInvest":        "double", //投资支付的现金
				"NIncrPledgeLoan":    "double", //质押贷款净增加额
				"NCPaidAcquis":       "double", //取得子公司及其他营业单位支付的现金净额
				"CPaidOthInvestA":    "double", //支付其他与投资活动有关的现金
				"COutfFrInvestA":     "double", //投资活动现金流出小计
				"NCFFrInvestA":       "double", //投资活动产生的现金流量净额
				"CFrCapContr":        "double", //吸收投资收到的现金
				"CFrMinoSSubs":       "double", //其中:子公司吸收少数股东投资收到的现金
				"CFrBorr":            "double", //取得借款收到的现金
				"CFrIssueBond":       "double", //发行债券收到的现金
				"CFrOthFinanA":       "double", //收到其他与筹资活动有关的现金
				"CInfFrFinanA":       "double", //筹资活动现金流入小计
				"CPaidForDebts":      "double", //偿还债务支付的现金
				"CPaidDivProfInt":    "double", //分配股利、利润或偿付利息支付的现金
				"divProfSubsMinoS":   "double", //其中:子公司支付给少数股东的股利、利润
				"CPaidOthFinanA":     "double", //支付其他与筹资活动有关的现金
				"COutfFrFinanA":      "double", //筹资活动现金流出小计
				"NCFFrFinanA":        "double", //筹资活动产生的现金流量净额
				"forexEffects":       "double", //汇率变动对现金及现金等价物的影响
				"NChangeInCash":      "double", //现金及现金等价物净增加额
				"NCEBegBal":          "double", //加:期初现金及现金等价物余额
				"NCEEndBal":          "double", //期末现金及现金等价物余额
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的银行业现金流量表模板，收集了2007年以来沪深上市公司定期报告中所有以此模板披露的现金流量表数据；（主要是银行业上市公司） 2、仅收集合并报表数据，包括本期和上期数据； 3、如果上市公司对外财务报表进行更正，调整，均有采集并对外展示； 4、本表中单位为人民币元；5、每季更新。
		"getFdmtCFBank": map[string]interface{}{
			"url": "/api/fundamental/getFdmtCFBank.json",
			"args": map[string]string{
				"secID":            "string", //证券内部编码,可通过交易代码和交易市场在getSecurityID获取到,如'000001.XSHE'
				"ticker":           "string", //股票代码，如'000001'
				"beginDate":        "date",   //会计期间截止日期,起始时间,如‘20121231’
				"endDate":          "date",   //会计期间截止日期,结束时间,如‘20131231’
				"fiscalPeriod":     "String", //会计期间长度,可多值输入，如：3，6，9，12，3表示季度值
				"publishDateBegin": "date",   //证券交易所披露的信息发布日期,起始时间,默认为1年前当前日期,如‘20130407’
				"publishDateEnd":   "date",   //证券交易所披露的信息发布日期,结束时间,默认为当前日期,如‘20140407’
				"beginDateRep":     "date",   //报表会计期末，起始时间，输入格式“YYYYMMDD”
				"endDateRep":       "date",   //报表会计期末，结束时间，输入格式“YYYYMMDD”
				"beginYear":        "string", //会计期间所属年份，起始时间，输入格式"YYYY"
				"endYear":          "string", //会计期间所属年份，结束时间，输入格式"YYYY"
				"reportType":       "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，CQ3-三季报（累计1-9月），A-年报。
				"field":            "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"reportType":         "string", //报告类型
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"NDeposIncrCFI":      "double", //客户存款和同业存放款项净增加额
				"NIncrBorrFrCB":      "double", //向中央银行借款净增加额
				"NIncBorrOthFI":      "double", //向其他金融机构拆入资金净增加额
				"NDecrInDisburOfLa":  "double", //发放贷款及垫款净减少额
				"NDecrInDeposInFI":   "double", //存放中央银行和同业款项净减少额
				"NDecrLoanToOthFI":   "double", //向其他金融机构拆出资金净减少额
				"IFCCashIncr":        "double", //收取利息、手续费及佣金的现金
				"CFrOthOperateA":     "double", //收到其他与经营活动有关的现金
				"CInfFrOperateA":     "double", //经营活动现金流入小计
				"NDeposDecrFrFI":     "double", //客户存款和同业存放款项净减少额
				"NDecrBorrFrCB":      "double", //向中央银行借款净减少额
				"NDecrBorrFrOthFI":   "double", //向其他金融机构拆入资金净减少额
				"NIncDisburOfLA":     "double", //客户贷款及垫款净增加额
				"NIncrDeposInFI":     "double", //存放中央银行和同业款项净增加额
				"NIncrLoansToOthFi":  "double", //向其他金融机构拆出资金净增加额
				"CPaidIFC":           "double", //支付利息、手续费及佣金的现金
				"CPaidToForEmpl":     "double", //支付给职工以及为职工支付的现金
				"CPaidForTaxes":      "double", //支付的各项税费
				"CPaidForOthOpA":     "double", //支付其他与经营活动有关的现金
				"COutfOperateA":      "double", //经营活动现金流出小计
				"NCFOperateA":        "double", //经营活动产生的现金流量净额
				"procSellInvest":     "double", //收回投资收到的现金
				"gainInvest":         "double", //取得投资收益收到的现金
				"dispFixAssetsOth":   "double", //处置固定资产、无形资产和其他长期资产收回的现金净额
				"NDispSubsOthBizC":   "double", //处置子公司及其他营业单位收到的现金净额
				"CFrOthInvestA":      "double", //收到其他与投资活动有关的现金
				"CInfFrInvestA":      "double", //投资活动现金流入小计
				"purFixAssetsOth":    "double", //购建固定资产、无形资产和其他长期资产支付的现金
				"CPaidInvest":        "double", //投资支付的现金
				"NCPaidAcquis":       "double", //取得子公司及其他营业单位支付的现金净额
				"CPaidOthInvestA":    "double", //支付其他与投资活动有关的现金
				"COutfFrInvestA":     "double", //投资活动现金流出小计
				"NCFFrInvestA":       "double", //投资活动产生的现金流量净额
				"CFrCapContr":        "double", //吸收投资收到的现金
				"CFrMinoSSubs":       "double", //其中:子公司吸收少数股东投资收到的现金
				"CFrIssueBond":       "double", //发行债券收到的现金
				"CFrOthFinanA":       "double", //收到其他与筹资活动有关的现金
				"CInfFrFinanA":       "double", //筹资活动现金流入小计
				"CPaidForDebts":      "double", //偿还债务支付的现金
				"CPaidDivProfInt":    "double", //分配股利、利润或偿付利息支付的现金
				"divProfSubsMinoS":   "double", //其中:子公司支付给少数股东的股利、利润
				"CPaidOthFinanA":     "double", //支付其他与筹资活动有关的现金
				"COutfFrFinanA":      "double", //筹资活动现金流出小计
				"NCFFrFinanA":        "double", //筹资活动产生的现金流量净额
				"forexEffects":       "double", //汇率变动对现金及现金等价物的影响
				"NChangeInCash":      "double", //现金及现金等价物净增加额
				"NCEBegBal":          "double", //加:期初现金及现金等价物余额
				"NCEEndBal":          "double", //期末现金及现金等价物余额
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的证券业现金流量表模板，收集了2007年以来沪深上市公司定期报告中所有以此模板披露的现金流量表数据；（主要是证券业上市公司） 2、仅收集合并报表数据，包括本期和上期数据； 3、如果上市公司对外财务报表进行更正，调整，均有采集并对外展示； 4、本表中单位为人民币元；5、每季更新。
		"getFdmtCFSecu": map[string]interface{}{
			"url": "/api/fundamental/getFdmtCFSecu.json",
			"args": map[string]string{
				"secID":            "string", //证券内部编码,可通过交易代码和交易市场在getSecurityID获取到,如'600369.XSHG '
				"ticker":           "string", //股票代码，如'600369'
				"beginDate":        "date",   //会计期间截止日期,起始时间,如‘20121231’
				"endDate":          "date",   //会计期间截止日期,结束时间,如‘20131231’
				"fiscalPeriod":     "String", //会计期间长度,可多值输入，如：3，6，9，12，3表示季度值
				"publishDateBegin": "date",   //证券交易所披露的信息发布日期,起始时间,默认为1年前当前日期,如‘20130407’
				"publishDateEnd":   "date",   //证券交易所披露的信息发布日期,结束时间,默认为当前日期,如‘20140407’
				"beginDateRep":     "date",   //报表会计期末，起始时间，输入格式“YYYYMMDD”
				"endDateRep":       "date",   //报表会计期末，结束时间，输入格式“YYYYMMDD”
				"beginYear":        "string", //会计期间所属年份，起始时间，输入格式"YYYY"
				"endYear":          "string", //会计期间所属年份，结束时间，输入格式"YYYY"
				"reportType":       "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，CQ3-三季报（累计1-9月），A-年报。
				"field":            "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"reportType":         "string", //报告类型
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"NIncBorrOthFI":      "double", //向其他金融机构拆入资金净增加额
				"NIncDispTradFA":     "double", //处置交易性金融资产净增加额
				"NIncDispFaFS":       "double", //处置可供出售金融资产净增加额
				"IFCCashIncr":        "double", //收取利息、手续费及佣金的现金
				"NIncFrBorr":         "double", //拆入资金净增加额
				"NCApIncrRepur":      "double", //回购业务资金净增加额
				"refundOfTax":        "double", //收到的税费返还
				"CFrOthOperateA":     "double", //收到其他与经营活动有关的现金
				"CInfFrOperateA":     "double", //经营活动现金流入小计
				"CPaidIFC":           "double", //支付利息、手续费及佣金的现金
				"CPaidToForEmpl":     "double", //支付给职工以及为职工支付的现金
				"CPaidForTaxes":      "double", //支付的各项税费
				"CPaidForOthOpA":     "double", //支付其他与经营活动有关的现金
				"COutfOperateA":      "double", //经营活动现金流出小计
				"NCFOperateA":        "double", //经营活动产生的现金流量净额
				"procSellInvest":     "double", //收回投资收到的现金
				"gainInvest":         "double", //取得投资收益收到的现金
				"dispFixAssetsOth":   "double", //处置固定资产、无形资产和其他长期资产收回的现金净额
				"NDispSubsOthBizC":   "double", //处置子公司及其他营业单位收到的现金净额
				"CFrOthInvestA":      "double", //收到其他与投资活动有关的现金
				"CInfFrInvestA":      "double", //投资活动现金流入小计
				"purFixAssetsOth":    "double", //购建固定资产、无形资产和其他长期资产支付的现金
				"CPaidInvest":        "double", //投资支付的现金
				"NCPaidAcquis":       "double", //取得子公司及其他营业单位支付的现金净额
				"CPaidOthInvestA":    "double", //支付其他与投资活动有关的现金
				"COutfFrInvestA":     "double", //投资活动现金流出小计
				"NCFFrInvestA":       "double", //投资活动产生的现金流量净额
				"CFrCapContr":        "double", //吸收投资收到的现金
				"CFrMinoSSubs":       "double", //其中:子公司吸收少数股东投资收到的现金
				"CFrBorr":            "double", //取得借款收到的现金
				"CFrIssueBond":       "double", //发行债券收到的现金
				"CFrOthFinanA":       "double", //收到其他与筹资活动有关的现金
				"CInfFrFinanA":       "double", //筹资活动现金流入小计
				"CPaidForDebts":      "double", //偿还债务支付的现金
				"CPaidDivProfInt":    "double", //分配股利、利润或偿付利息支付的现金
				"divProfSubsMinoS":   "double", //其中:子公司支付给少数股东的股利、利润
				"CPaidOthFinanA":     "double", //支付其他与筹资活动有关的现金
				"COutfFrFinanA":      "double", //筹资活动现金流出小计
				"NCFFrFinanA":        "double", //筹资活动产生的现金流量净额
				"forexEffects":       "double", //汇率变动对现金及现金等价物的影响
				"NChangeInCash":      "double", //现金及现金等价物净增加额
				"NCEBegBal":          "double", //加:期初现金及现金等价物余额
				"NCEEndBal":          "double", //期末现金及现金等价物余额
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的保险业现金流量表模板，收集了2007年以来沪深上市公司定期报告中所有以此模板披露的现金流量表数据；（主要是保险业上市公司） 2、仅收集合并报表数据，包括本期和上期数据； 3、如果上市公司对外财务报表进行更正，调整，均有采集并对外展示； 4、本表中单位为人民币元； 5、每季更新。
		"getFdmtCFInsu": map[string]interface{}{
			"url": "/api/fundamental/getFdmtCFInsu.json",
			"args": map[string]string{
				"secID":            "string", //证券内部编码,可通过交易代码和交易市场在getSecurityID获取到,如'601318.XSHG'
				"ticker":           "string", //股票代码，如'601318'
				"beginDate":        "date",   //会计期间截止日期,起始时间,如‘20121231’
				"endDate":          "date",   //会计期间截止日期,结束时间,如‘20131231’
				"fiscalPeriod":     "String", //会计期间长度,可多值输入，如：3，6，9，12，3表示季度值
				"publishDateBegin": "date",   //证券交易所披露的信息发布日期,起始时间,默认为1年前当前日期,如‘20130407’
				"publishDateEnd":   "date",   //证券交易所披露的信息发布日期,结束时间,默认为当前日期,如‘20140407’
				"beginDateRep":     "date",   //报表会计期末，起始时间，输入格式“YYYYMMDD”
				"endDateRep":       "date",   //报表会计期末，结束时间，输入格式“YYYYMMDD”
				"beginYear":        "string", //会计期间所属年份，起始时间，输入格式"YYYY"
				"endYear":          "string", //会计期间所属年份，结束时间，输入格式"YYYY"
				"reportType":       "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，CQ3-三季报（累计1-9月），A-年报。
				"field":            "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"reportType":         "string", //报告类型
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"NDeposIncrCFI":      "double", //客户存款和同业存放款项净增加额
				"NIncrBorrFrCB":      "double", //向中央银行借款净增加额
				"premFrOrigContr":    "double", //收到原保险合同保费取得的现金
				"NReinsurPrem":       "double", //收到再保险业务现金净额
				"NIncPhDeposInv":     "double", //保户储金及投资款净增加额
				"IFCCashIncr":        "double", //收取利息、手续费及佣金的现金
				"refundOfTax":        "double", //收到的税费返还
				"CFrOthOperateA":     "double", //收到其他与经营活动有关的现金
				"CInfFrOperateA":     "double", //经营活动现金流入小计
				"NIncDisburOfLA":     "double", //客户贷款及垫款净增加额
				"NIncrDeposInFI":     "double", //存放中央银行和同业款项净增加额
				"origContrCIndem":    "double", //支付原保险合同赔付款项的现金
				"CPaidIFC":           "double", //支付利息、手续费及佣金的现金
				"CPaidPolDiv":        "double", //支付保单红利的现金
				"CPaidToForEmpl":     "double", //支付给职工以及为职工支付的现金
				"CPaidForTaxes":      "double", //支付的各项税费
				"CPaidForOthOpA":     "double", //支付其他与经营活动有关的现金
				"COutfOperateA":      "double", //经营活动现金流出小计
				"NCFOperateA":        "double", //经营活动产生的现金流量净额
				"procSellInvest":     "double", //收回投资收到的现金
				"gainInvest":         "double", //取得投资收益收到的现金
				"dispFixAssetsOth":   "double", //处置固定资产、无形资产和其他长期资产收回的现金净额
				"NDispSubsOthBizC":   "double", //处置子公司及其他营业单位收到的现金净额
				"CFrOthInvestA":      "double", //收到其他与投资活动有关的现金
				"CInfFrInvestA":      "double", //投资活动现金流入小计
				"PurFixAssetsOth":    "double", //购建固定资产、无形资产和其他长期资产支付的现金
				"CPaidInvest":        "double", //投资支付的现金
				"NIncrPledgeLoan":    "double", //质押贷款净增加额
				"NCPaidAcquis":       "double", //取得子公司及其他营业单位支付的现金净额
				"CPaidOthInvestA":    "double", //支付其他与投资活动有关的现金
				"COutfFrInvestA":     "double", //投资活动现金流出小计
				"NCFFrInvestA":       "double", //投资活动产生的现金流量净额
				"CFrCapContr":        "double", //吸收投资收到的现金
				"CFrMinoSSubs":       "double", //其中:子公司吸收少数股东投资收到的现金
				"CFrBorr":            "double", //取得借款收到的现金
				"CFrIssueBond":       "double", //发行债券收到的现金
				"CFrOthFinanA":       "double", //收到其他与筹资活动有关的现金
				"CInfFrFinanA":       "double", //筹资活动现金流入小计
				"CPaidForDebts":      "double", //偿还债务支付的现金
				"CPaidDivProfInt":    "double", //分配股利、利润或偿付利息支付的现金
				"divProfSubsMinoS":   "double", //其中:子公司支付给少数股东的股利、利润
				"CPaidOthFinanA":     "double", //支付其他与筹资活动有关的现金
				"COutfFrFinanA":      "double", //筹资活动现金流出小计
				"NCFFrFinanA":        "double", //筹资活动产生的现金流量净额
				"forexEffects":       "double", //汇率变动对现金及现金等价物的影响
				"NChangeInCash":      "double", //现金及现金等价物净增加额
				"NCEBegBal":          "double", //加:期初现金及现金等价物余额
				"NCEEndBal":          "double", //期末现金及现金等价物余额
			},
			"indices": map[string]string{},
		},
		//说明1、获取2007年及以后年度上市公司披露的业绩快报中的主要财务指标等其他数据，包括本期，去年同期，及本期与期初数值同比数据。 2、如果上市公司对外财务报表进行更正，调整，展示上市公司最新披露一期财务数据。
		"getFdmtEeAllLatest": map[string]interface{}{
			"url": "/api/fundamental/getFdmtEeAllLatest.json",
			"args": map[string]string{
				"reportType": "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，Q3-第三季报，CQ3-三季报（累计1-9月），A-年报。。
				"secID":      "string", //证券内部编码,可通过交易代码和交易市场在getSecID获取到,如'600000.XSHG'
				"ticker":     "string", //股票交易代码，如'600000'
				"beginDate":  "date",   //会计期间截止日期，起始时间，输入格式“YYYYMMDD”
				"endDate":    "date",   //会计期间截止日期，结束时间，输入格式“YYYYMMDD”
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":              "string", //证券内部编码
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"partyID":            "int64",  //公司代码
				"ticker":             "string", //交易代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易所代码
				"actPubtime":         "string", //实际发布时间
				"mergedFlag":         "string", //合并标志:1-合并
				"reportType":         "string", //报告类型
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币
				"revenue":            "double", //营业收入
				"primeOperRev":       "double", //主营业务收入
				"grossProfit":        "double", //主营业务利润
				"operateProfit":      "double", //营业利润
				"TProfit":            "double", //利润总额
				"NIncomeAttrP":       "double", //归属于母公司所有者的净利润
				"NIncomeCut":         "double", //扣除非经常性损益后净利润
				"NCfOperA":           "double", //经营活动现金流量净额
				"basicEPS":           "double", //基本每股收益(元/股)
				"EPSW":               "double", //加权平均每股收益(元/股)
				"EPSCut":             "double", //扣除非经常损益后的基本每股收益(元/股)
				"EPSCutW":            "double", //扣除非经常损益后的加权每股收益(元/股)
				"ROE":                "double", //全面摊薄净资产收益率(%)
				"ROEW":               "double", //加权平均净资产收益率(%)
				"ROECut":             "double", //扣除非经常性损益的全面摊薄净资产收益率(%)
				"ROECutW":            "double", //扣除非经常性损益后的加权平均净资产收益率(%)
				"NCfOperAPs":         "double", //每股经营活动现金流量净额(元/股)
				"TAssets":            "double", //总资产
				"TEquityAttrP":       "double", //归属于母公司所有者权益合计
				"paidInCapital":      "double", //股本
				"NAssetPS":           "double", //每股净资产(元)
				"revenueLY":          "double", //上年同期营业收入
				"primeOperRevLY":     "double", //上年同期主营业务收入
				"grossProfitLY":      "double", //上年同期主营业务利润
				"operProfitLY":       "double", //上年同期营业利润
				"TProfitLY":          "double", //上年同期利润总额
				"NIncomeAttrPLY":     "double", //上年同期归属于母公司所有者的净利润
				"NIncomeCutLY":       "double", //上年同期扣除非经常性损益后净利润
				"NCfOperALY":         "double", //上年同期经营活动现金流量净额
				"basicEPSLY":         "double", //上年同期基本每股收益(元/股)
				"EPSWLY":             "double", //上年同期加权平均每股收益(元/股)
				"EPSCutLY":           "double", //上年同期扣除非经常损益后的基本每股收益(元/股)
				"EPSCutWLY":          "double", //上年同期扣除非经常损益后的加权每股收益(元/股)
				"ROELY":              "double", //上年同期期末净资产收益率(%)
				"ROEWLY":             "double", //上年同期加权平均净资产收益率(%)
				"ROECutLY":           "double", //上年同期扣除非经常性损益的期末净资产收益率(%)
				"ROECutWLY":          "double", //上年同期扣除非经常性损益后的加权平均净资产收益率(%)
				"NCfOperAPsLY":       "double", //上年同期每股经营活动现金流量净额(元/股)
				"TAssetsLY":          "double", //期初总资产
				"TEquityAttrPLY":     "double", //期初归属于母公司所有者权益合计
				"NAssetPsLY":         "double", //期初每股净资产(元)
				"revenueYOY":         "double", //营业收入同比(%)
				"primeOperRevYOY":    "double", //主营业务收入同比(%)
				"grossProfitYOY":     "double", //主营业务利润同比(%)
				"operProfitYOY":      "double", //营业利润同比(%)
				"TProfitYOY":         "double", //利润总额同比(%)
				"NIncomeAttrPYOY":    "double", //归属于母公司所有者的净利润同比(%)
				"NIncomeCutYOY":      "double", //扣除非经常性损益后净利润同比(%)
				"NCFOperAYOY":        "double", //经营活动现金流量净额同比(%)
				"basicEPSYOY":        "double", //基本每股收益同比(%)
				"EPSWYOY":            "double", //加权平均每股收益同比(%)
				"EPSCutYOY":          "double", //扣除非经常损益后的基本每股收益同比(%)
				"EPSCutWYOY":         "double", //扣除非经常损益后的加权每股收益同比(%)
				"ROEYOY":             "double", //期末净资产收益率同比(%)
				"ROEWYOY":            "double", //加权平均净资产收益率同比(%)
				"ROECutYOY":          "double", //扣除非经常性损益的期末净资产收益率同比(%)
				"ROECutWYOY":         "double", //扣除非经常性损益后的加权平均净资产收益率同比(%)
				"NCfOperAPsYOY":      "double", //每股经营活动现金流量净额同比(%)
				"TAssetsYOY":         "double", //总资产同比(%)
				"TEquityAttrPYOY":    "double", //归属于母公司所有者权益同比(%)
				"NAssetPsYOY":        "double", //每股净资产同比(%)
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的合并利润表模板，收集了2007年以来沪深上市公司定期报告中各个会计期间的利润表数据；2、仅收集合并报表数据，包括本期和上期数据； 3、如果上市公司对外财务报表进行更正，调整，展示上市公司最新披露一期财务数据；4、本表中单位为人民币元；5、每季更新。
		"getFdmtISAllLatest": map[string]interface{}{
			"url": "/api/fundamental/getFdmtISAllLatest.json",
			"args": map[string]string{
				"reportType": "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，Q3-第三季报，CQ3-三季报（累计1-9月），A-年报。
				"secID":      "string", //证券内部编码,可通过交易代码和交易市场在getSecID获取到,如'000002.XSHE'
				"ticker":     "string", //股票代码，如'000002'
				"beginDate":  "date",   //会计期间截止日期，起始时间，输入格式“YYYYMMDD”
				"endDate":    "date",   //会计期间截止日期，结束时间，输入格式“YYYYMMDD”
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志
				"reportType":         "string", //报告类型
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"tRevenue":           "double", //营业总收入
				"revenue":            "double", //营业收入
				"intIncome":          "double", //利息收入
				"intExp":             "double", //利息支出
				"premEarned":         "double", //已赚保费
				"commisIncome":       "double", //手续费及佣金收入
				"commisExp":          "double", //手续费及佣金支出
				"TCogs":              "double", //营业总成本
				"COGS":               "double", //营业成本
				"premRefund":         "double", //退保金
				"NCompensPayout":     "double", //赔付支出净额
				"reserInsurContr":    "double", //提取保险合同准备金净额
				"policyDivPayt":      "double", //保单红利支出
				"reinsurExp":         "double", //分保费用
				"bizTaxSurchg":       "double", //营业税金及附加
				"sellExp":            "double", //销售费用
				"adminExp":           "double", //管理费用
				"finanExp":           "double", //财务费用
				"assetsImpairLoss":   "double", //资产减值损失
				"fValueChgGain":      "double", //公允价值变动收益
				"investIncome":       "double", //投资收益
				"AJInvestIncome":     "double", //其中:对联营企业和合营企业的投资收益
				"forexGain":          "double", //汇兑收益
				"operateProfit":      "double", //营业利润
				"NoperateIncome":     "double", //营业外收入
				"NoperateExp":        "double", //营业外支出
				"NCADisploss":        "double", //非流动资产处置损失
				"TProfit":            "double", //利润总额
				"incomeTax":          "double", //所得税费用
				"NIncome":            "double", //净利润
				"NIncomeAttrP":       "double", //归属于母公司所有者的净利润
				"minorityGain":       "double", //少数股东损益
				"basicEPS":           "double", //基本每股收益
				"dilutedEPS":         "double", //稀释每股收益
				"othComprIncome":     "double", //其他综合收益
				"TComprIncome":       "double", //综合收益总额
				"comprIncAttrP":      "double", //归属于母公司所有者的综合收益总额
				"comprIncAttrMS":     "double", //归属于少数股东的综合收益总额
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的证券业利润表模板，收集了2007年以来沪深上市公司定期报告中所有以此模板披露的利润表数据（主要是证券业上市公司）；2、仅收集合并报表数据，包括本期和上期数据；3、如果上市公司对外财务报表进行更正，调整，展示上市公司最新披露一期财务数据；4、本表中单位为人民币元；5、每季更新。
		"getFdmtISSecuAllLatest": map[string]interface{}{
			"url": "/api/fundamental/getFdmtISSecuAllLatest.json",
			"args": map[string]string{
				"reportType": "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，Q3-第三季报，CQ3-三季报（累计1-9月），A-年报。
				"secID":      "string", //证券内部编码,可通过交易代码和交易市场在getSecID获取到,如'600369.XSHG'
				"ticker":     "string", //股票代码，如'600369'
				"beginDate":  "date",   //会计期间截止日期，起始时间，输入格式“YYYYMMDD”
				"endDate":    "date",   //会计期间截止日期，结束时间，输入格式“YYYYMMDD”
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志
				"reportType":         "string", //报告类型
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"revenue":            "double", //营业收入
				"NIntIncome":         "double", //利息净收入
				"NCommisIncome":      "double", //手续费及佣金净收入
				"NSecTaIncome":       "double", //代理买卖证券业务净收入
				"NUndwrtSecIncome":   "double", //证券承销业务净收入
				"NTrustIncome":       "double", //委托客户资产管理业务净收入
				"othOperRev":         "double", //其他业务收入
				"COGS":               "double", //营业支出
				"bizTaxSurchg":       "double", //营业税金及附加
				"genlAdminExp":       "double", //业务及管理费
				"assetsImpairLoss":   "double", //资产减值损失
				"othOperCosts":       "double", //其他业务成本
				"fValueChgGain":      "double", //公允价值变动收益
				"investIncome":       "double", //投资收益
				"AJInvestIncome":     "double", //其中:对联营企业和合营企业的投资收益
				"forexGain":          "double", //汇兑收益
				"operateProfit":      "double", //营业利润
				"NoperateIncome":     "double", //营业外收入
				"NoperateExp":        "double", //营业外支出
				"TProfit":            "double", //利润总额
				"incomeTax":          "double", //所得税费用
				"NIncome":            "double", //净利润
				"NIncomeAttrP":       "double", //归属于母公司所有者的净利润
				"minorityGain":       "double", //少数股东损益
				"basicEPS":           "double", //基本每股收益
				"dilutedEPS":         "double", //稀释每股收益
				"othComprIncome":     "double", //其他综合收益
				"TComprIncome":       "double", //综合收益总额
				"comprIncAttrP":      "double", //归属于母公司所有者的综合收益总额
				"comprIncAttrMS":     "double", //归属于少数股东的综合收益总额
			},
			"indices": map[string]string{},
		},
		//说明1、获取2007年及以后年度上市公司披露的公告中的预期下一报告期收入、净利润、归属于母公司净利润、基本每股收益及其幅度变化数据。 2、上市公司对经营成果科目的预计情况数据一般为其上限与下限，上限取值为公告中披露该科目中绝对值较大值，下限取值为公告中披露该科目中绝对值较小值。 3、数值为"正"代表该公司预计盈利，数值为"负"代表该公司预计亏损。若上下限"正"、"负"符号不同，代表该公司盈利亏损情况尚不确定。 4、业绩预期类型以公告中文字披露预期类型为准，若公告中未有文字披露预期类型，则根据数据情况判断预期类型。 5、每季证券交易所披露相关公告时更新数据，公司ipo时发布相关信息也会同时更新。每日9：00前完成证券交易所披露的数据更新，中午发布公告每日12：45前完成更新。
		"getFdmtEf": map[string]interface{}{
			"url": "/api/fundamental/getFdmtEf.json",
			"args": map[string]string{
				"secID":            "string", //证券内部编码,可通过交易代码和交易市场在getSecurityID获取到,如'000005.XSHE'
				"ticker":           "string", //股票交易代码，如'000005'，默认为'000005’
				"beginDate":        "date",   //会计期间截止日期,起始时间,如‘20121231’
				"endDate":          "date",   //会计期间截止日期,结束时间,如‘20131231’
				"forecastType":     "int32",  //业绩预期类型，可供选择类型：11 预计亏损，21 预计盈利，22盈利增加，12 亏损减少，24 盈利下降，23 盈利减缓，13 下降减缓，31 业绩不确定，41 基本持平，51 经营预期及其他
				"publishDateBegin": "date",   //证券交易所披露的信息发布日期，起始时间，如‘20130812’
				"publishDateEnd":   "date",   //证券交易所披露的信息发布日期，结束时间，如‘20140812’
				"reportType":       "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，Q3-第三季报，CQ3-三季报（累计1-9月），A-年报。
				"field":            "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":              "string", //证券内部编码
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"partyID":            "int64",  //公司代码
				"ticker":             "string", //交易代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易所代码
				"actPubtime":         "string", //实际发布时间
				"mergedFlag":         "string", //合并标志
				"reportType":         "string", //报告类型
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币
				"forecastType":       "double", //业绩预期类型
				"revChgrLL":          "double", //收入增减幅度下限(%)
				"revChgrUPL":         "double", //收入增减幅度上限(%)
				"expRevLL":           "double", //预期收入下限
				"expRevUPL":          "double", //预期收入上限
				"NIncomeChgrLL":      "double", //净利润增减幅度下限(%)
				"NIncomeChgrUPL":     "double", //净利润增减幅度上限(%)
				"expnIncomeLL":       "double", //预期净利润下限
				"expnIncomeUPL":      "double", //预期净利润上限
				"NIncAPChgrLL":       "double", //归属于母公司所有者的净利润增减幅度下限(%)
				"NIncAPChgrUPL":      "double", //归属于母公司所有者的净利润增减幅度上限(%)
				"expnIncAPLL":        "double", //预期归属于母公司所有者的净利润下限
				"expnIncAPUPL":       "double", //预期归属于母公司所有者的净利润上限
				"expEPSLL":           "double", //预期每股收益下限
				"expEPSUPL":          "double", //预期每股收益上限
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的合并资产负债表模板，收集了2007年以来沪深上市公司定期报告中各个会计期间的资产负债表数据； 2、仅收集合并报表数据，包括期末和期初数据； 3、如果上市公司对外财务报表进行更正，调整，展示上市公司最新披露一期财务数据； 4、本表中单位为人民币元； 5、每季更新。
		"getFdmtBSAllLatest": map[string]interface{}{
			"url": "/api/fundamental/getFdmtBSAllLatest.json",
			"args": map[string]string{
				"reportType": "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，Q3-第三季报，A-年报。
				"secID":      "string", //证券内部编码,可通过交易代码和交易市场在getSecID获取到,如'000002.XSHE'
				"ticker":     "string", //股票代码，如'000002'
				"beginDate":  "date",   //会计期间截止日期，起始时间，输入格式“YYYYMMDD”
				"endDate":    "date",   //会计期间截止日期，结束时间，输入格式“YYYYMMDD”
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志:1-合并
				"reportType":         "string", //报告类型
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"cashCEquiv":         "double", //货币资金
				"settProv":           "double", //结算备付金
				"loanToOthBankFi":    "double", //拆出资金
				"tradingFA":          "double", //交易性金融资产
				"NotesReceiv":        "double", //应收票据
				"AR":                 "double", //应收账款
				"prepayment":         "double", //预付款项
				"premiumReceiv":      "double", //应收保费
				"reinsurReceiv":      "double", //应收分保账款
				"reinsurReserReceiv": "double", //应收分保合同准备金
				"intReceiv":          "double", //应收利息
				"divReceiv":          "double", //应收股利
				"othReceiv":          "double", //其他应收款
				"purResaleFa":        "double", //买入返售金融资产
				"inventories":        "double", //存货
				"NCAWithin1Y":        "double", //一年内到期的非流动资产
				"othCA":              "double", //其他流动资产
				"TCA":                "double", //流动资产合计
				"disburLA":           "double", //发放委托贷款及垫款
				"availForSaleFa":     "double", //可供出售金融资产
				"htmInvest":          "double", //持有至到期投资
				"LTReceive":          "double", //长期应收款
				"LTEquityInvest":     "double", //长期股权投资
				"investRealEstate":   "double", //投资性房地产
				"fixedAssets":        "double", //固定资产
				"CIP":                "double", //在建工程
				"constMaterials":     "double", //工程物资
				"fixedAssetsDisp":    "double", //固定资产清理
				"producBiolAssets":   "double", //生产性生物资产
				"oilAndGasAssets":    "double", //油气资产
				"intanAssets":        "double", //无形资产
				"RD":                 "double", //开发支出
				"goodwill":           "double", //商誉
				"LTAmorExp":          "double", //长期待摊费用
				"deferTaxAssets":     "double", //递延所得税资产
				"othNCA":             "double", //其他非流动资产
				"TNCA":               "double", //非流动资产合计
				"TAssets":            "double", //资产总计
				"STBorr":             "double", //短期借款
				"CBBorr":             "double", //向中央银行借款
				"depos":              "double", //吸收存款及同业存放
				"loanFrOthBankFi":    "double", //拆入资金
				"tradingFL":          "double", //交易性金融负债
				"NotesPayable":       "double", //应付票据
				"AP":                 "double", //应付账款
				"advanceReceipts":    "double", //预收款项
				"soldForRepurFa":     "double", //卖出回购金融资产款
				"commisPayable":      "double", //应付手续费及佣金
				"payrollPayable":     "double", //应付职工薪酬
				"taxesPayable":       "double", //应交税费
				"intPayable":         "double", //应付利息
				"divPayable":         "double", //应付股利
				"othPayable":         "double", //其他应付款
				"reinsurPayable":     "double", //应付分保账款
				"insurReser":         "double", //保险合同准备金
				"fundsSecTradAgen":   "double", //代理买卖证券款
				"fundsSecUndwAgen":   "double", //代理承销证券款
				"NCLWithin1Y":        "double", //一年内到期的非流动负债
				"othCL":              "double", //其他流动负债
				"TCL":                "double", //流动负债合计
				"LTBorr":             "double", //长期借款
				"bondPayable":        "double", //应付债券
				"LTPayable":          "double", //长期应付款
				"specificPayables":   "double", //专项应付款
				"estimatedLiab":      "double", //预计负债
				"deferTaxLiab":       "double", //递延所得税负债
				"othNCL":             "double", //其他非流动负债
				"TNCL":               "double", //非流动负债合计
				"TLiab":              "double", //负债合计
				"paidInCapital":      "double", //实收资本(或股本)
				"capitalReser":       "double", //资本公积
				"treasuryShare":      "double", //减:库存股
				"specialReser":       "double", //专项储备
				"surplusReser":       "double", //盈余公积
				"ordinRiskReser":     "double", //一般风险准备
				"retainedEarnings":   "double", //未分配利润
				"forexDiffer":        "double", //外币报表折算差额
				"TEquityAttrP":       "double", //归属于母公司所有者权益合计
				"minorityInt":        "double", //少数股东权益
				"TShEquity":          "double", //所有者权益合计
				"TLiabEquity":        "double", //负债和所有者权益总计
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的证券业资产负债表模板，收集了2007年以来沪深上市公司定期报告中所有以此模板披露的资产负债表数据；（主要是证券业上市公司） 2、仅收集合并报表数据，包括期末和期初数据； 3、如果上市公司对外财务报表进行更正，调整，展示上市公司最新披露一期财务数据； 4、本表中单位为人民币元； 5、每季更新
		"getFdmtBSSecuAllLatest": map[string]interface{}{
			"url": "/api/fundamental/getFdmtBSSecuAllLatest.json",
			"args": map[string]string{
				"reportType": "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，Q3-第三季报，A-年报。
				"secID":      "string", //证券内部编码,可通过交易代码和交易市场在getSecID获取到,如'600369.XSHG'
				"ticker":     "string", //股票代码，如'600369'
				"beginDate":  "date",   //会计期间截止日期，起始时间，输入格式“YYYYMMDD”
				"endDate":    "date",   //会计期间截止日期，结束时间，输入格式“YYYYMMDD”
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志:1-合并
				"reportType":         "string", //报告类型
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"cashCEquiv":         "double", //货币资金
				"clientDepos":        "double", //其中:客户资金存款
				"settProv":           "double", //结算备付金
				"clientProv":         "double", //其中:客户备付金
				"loanToOthBankFi":    "double", //拆出资金
				"tradingFA":          "double", //交易性金融资产
				"intReceiv":          "double", //应收利息
				"purResaleFa":        "double", //买入返售金融资产
				"availForSaleFa":     "double", //可供出售金融资产
				"htmInvest":          "double", //持有至到期投资
				"LTEquityInvest":     "double", //长期股权投资
				"investRealEstate":   "double", //投资性房地产
				"fixedAssets":        "double", //固定资产
				"intanAssets":        "double", //无形资产
				"transacSeatFee":     "double", //其中:交易席位费
				"deferTaxAssets":     "double", //递延所得税资产
				"derivAssets":        "double", //衍生金融资产
				"refundDepos":        "double", //存出保证金
				"othAssets":          "double", //其他资产
				"TAssets":            "double", //资产总计
				"STBorr":             "double", //短期借款
				"pledgeBorr":         "double", //其中:质押借款
				"loanFrOthBankFi":    "double", //拆入资金
				"tradingFL":          "double", //交易性金融负债
				"soldForRepurFa":     "double", //卖出回购金融资产款
				"payrollPayable":     "double", //应付职工薪酬
				"taxesPayable":       "double", //应交税费
				"intPayable":         "double", //应付利息
				"fundsSecTradAgen":   "double", //代理买卖证券款
				"fundsSecUndwAgen":   "double", //代理承销证券款
				"LTBorr":             "double", //长期借款
				"bondPayable":        "double", //应付债券
				"estimatedLiab":      "double", //预计负债
				"deferTaxLiab":       "double", //递延所得税负债
				"derivLiab":          "double", //衍生金融负债
				"othLiab":            "double", //其他负债
				"TLiab":              "double", //负债合计
				"paidInCapital":      "double", //实收资本(或股本)
				"capitalReser":       "double", //资本公积
				"treasuryShare":      "double", //减:库存股
				"surplusReser":       "double", //盈余公积
				"ordinRiskReser":     "double", //一般风险准备
				"transacRiskReser":   "double", //交易风险准备
				"retainedEarnings":   "double", //未分配利润
				"forexDiffer":        "double", //外币报表折算差额
				"TEquityAttrP":       "double", //归属于母公司所有者权益合计
				"minorityInt":        "double", //少数股东权益
				"TShEquity":          "double", //所有者权益合计
				"TLiabEquity":        "double", //负债和所有者权益总计
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的一般工商业资产负债表模板，收集了2007年以来沪深上市公司定期报告中所有以此模板披露的资产负债表数据；（主要是一般工商业上市公司） 2、仅收集合并报表数据，包括期末和期初数据； 3、如果上市公司对外财务报表进行更正，调整，展示上市公司最新披露一期财务数据； 4、本表中单位为人民币元； 5、每季更新。
		"getFdmtBSInduAllLatest": map[string]interface{}{
			"url": "/api/fundamental/getFdmtBSInduAllLatest.json",
			"args": map[string]string{
				"reportType": "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，Q3-第三季报，A-年报。
				"secID":      "string", //证券内部编码,可通过交易代码和交易市场在getSecID获取到,如'000002.XSHE'
				"ticker":     "string", //股票代码，如'000002'
				"beginDate":  "date",   //会计期间截止日期，起始时间，输入格式“YYYYMMDD”
				"endDate":    "date",   //会计期间截止日期，结束时间，输入格式“YYYYMMDD”
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志:1-合并
				"reportType":         "string", //报告类型
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"cashCEquiv":         "double", //货币资金
				"settProv":           "double", //结算备付金
				"loanToOthBankFi":    "double", //拆出资金
				"tradingFA":          "double", //交易性金融资产
				"NotesReceiv":        "double", //应收票据
				"AR":                 "double", //应收账款
				"prepayment":         "double", //预付款项
				"premiumReceiv":      "double", //应收保费
				"reinsurReceiv":      "double", //应收分保账款
				"reinsurReserReceiv": "double", //应收分保合同准备金
				"intReceiv":          "double", //应收利息
				"divReceiv":          "double", //应收股利
				"othReceiv":          "double", //其他应收款
				"purResaleFa":        "double", //买入返售金融资产
				"inventories":        "double", //存货
				"NCAWithin1Y":        "double", //一年内到期的非流动资产
				"othCA":              "double", //其他流动资产
				"TCA":                "double", //流动资产合计
				"disburLA":           "double", //发放委托贷款及垫款
				"availForSaleFa":     "double", //可供出售金融资产
				"htmInvest":          "double", //持有至到期投资
				"LTReceive":          "double", //长期应收款
				"LTEquityInvest":     "double", //长期股权投资
				"investRealEstate":   "double", //投资性房地产
				"fixedAssets":        "double", //固定资产
				"CIP":                "double", //在建工程
				"constMaterials":     "double", //工程物资
				"fixedAssetsDisp":    "double", //固定资产清理
				"producBiolAssets":   "double", //生产性生物资产
				"oilAndGasAssets":    "double", //油气资产
				"intanAssets":        "double", //无形资产
				"RD":                 "double", //开发支出
				"goodwill":           "double", //商誉
				"LTAmorExp":          "double", //长期待摊费用
				"deferTaxAssets":     "double", //递延所得税资产
				"othNCA":             "double", //其他非流动资产
				"TNCA":               "double", //非流动资产合计
				"TAssets":            "double", //资产总计
				"STBorr":             "double", //短期借款
				"CBBorr":             "double", //向中央银行借款
				"depos":              "double", //吸收存款及同业存放
				"loanFrOthBankFi":    "double", //拆入资金
				"tradingFL":          "double", //交易性金融负债
				"NotesPayable":       "double", //应付票据
				"AP":                 "double", //应付账款
				"advanceReceipts":    "double", //预收款项
				"soldForRepurFa":     "double", //卖出回购金融资产款
				"commisPayable":      "double", //应付手续费及佣金
				"payrollPayable":     "double", //应付职工薪酬
				"taxesPayable":       "double", //应交税费
				"intPayable":         "double", //应付利息
				"divPayable":         "double", //应付股利
				"othPayable":         "double", //其他应付款
				"reinsurPayable":     "double", //应付分保账款
				"insurReser":         "double", //保险合同准备金
				"fundsSecTradAgen":   "double", //代理买卖证券款
				"fundsSecUndwAgen":   "double", //代理承销证券款
				"NCLWithin1Y":        "double", //一年内到期的非流动负债
				"othCL":              "double", //其他流动负债
				"TCL":                "double", //流动负债合计
				"LTBorr":             "double", //长期借款
				"bondPayable":        "double", //应付债券
				"LTPayable":          "double", //长期应付款
				"specificPayables":   "double", //专项应付款
				"estimatedLiab":      "double", //预计负债
				"deferTaxLiab":       "double", //递延所得税负债
				"othNCL":             "double", //其他非流动负债
				"TNCL":               "double", //非流动负债合计
				"TLiab":              "double", //负债合计
				"paidInCapital":      "double", //实收资本(或股本)
				"capitalReser":       "double", //资本公积
				"treasuryShare":      "double", //减:库存股
				"specialReser":       "double", //专项储备
				"surplusReser":       "double", //盈余公积
				"ordinRiskReser":     "double", //一般风险准备
				"retainedEarnings":   "double", //未分配利润
				"forexDiffer":        "double", //外币报表折算差额
				"TEquityAttrP":       "double", //归属于母公司所有者权益合计
				"minorityInt":        "double", //少数股东权益
				"TShEquity":          "double", //所有者权益合计
				"TLiabEquity":        "double", //负债和所有者权益总计
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的保险业现金流量表模板，收集了2007年以来沪深上市公司定期报告中所有以此模板披露的现金流量表数据（主要是保险业上市公司）；2、仅收集合并报表数据，包括本期和上期数据；3、如果上市公司对外财务报表进行更正，调整，展示上市公司最新披露一期财务数据；4、本表中单位为人民币元；5、每季更新。
		"getFdmtCFInsuAllLatest": map[string]interface{}{
			"url": "/api/fundamental/getFdmtCFInsuAllLatest.json",
			"args": map[string]string{
				"reportType": "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，CQ3-三季报（累计1-9月），A-年报。
				"secID":      "string", //证券内部编码,可通过交易代码和交易市场在getSecID获取到,如'601318.XSHG'
				"ticker":     "string", //股票代码，如'601318'
				"beginDate":  "date",   //会计期间截止日期，起始时间，输入格式“YYYYMMDD”
				"endDate":    "date",   //会计期间截止日期，结束时间，输入格式“YYYYMMDD”
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志
				"reportType":         "string", //报告类型
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"NDeposIncrCFI":      "double", //客户存款和同业存放款项净增加额
				"NIncrBorrFrCB":      "double", //向中央银行借款净增加额
				"premFrOrigContr":    "double", //收到原保险合同保费取得的现金
				"NReinsurPrem":       "double", //收到再保险业务现金净额
				"NIncPhDeposInv":     "double", //保户储金及投资款净增加额
				"IFCCashIncr":        "double", //收取利息、手续费及佣金的现金
				"refundOfTax":        "double", //收到的税费返还
				"CFrOthOperateA":     "double", //收到其他与经营活动有关的现金
				"CInfFrOperateA":     "double", //经营活动现金流入小计
				"NIncDisburOfLA":     "double", //客户贷款及垫款净增加额
				"NIncrDeposInFI":     "double", //存放中央银行和同业款项净增加额
				"origContrCIndem":    "double", //支付原保险合同赔付款项的现金
				"CPaidIFC":           "double", //支付利息、手续费及佣金的现金
				"CPaidPolDiv":        "double", //支付保单红利的现金
				"CPaidToForEmpl":     "double", //支付给职工以及为职工支付的现金
				"CPaidForTaxes":      "double", //支付的各项税费
				"CPaidForOthOpA":     "double", //支付其他与经营活动有关的现金
				"COutfOperateA":      "double", //经营活动现金流出小计
				"NCFOperateA":        "double", //经营活动产生的现金流量净额
				"procSellInvest":     "double", //收回投资收到的现金
				"gainInvest":         "double", //取得投资收益收到的现金
				"dispFixAssetsOth":   "double", //处置固定资产、无形资产和其他长期资产收回的现金净额
				"NDispSubsOthBizC":   "double", //处置子公司及其他营业单位收到的现金净额
				"CFrOthInvestA":      "double", //收到其他与投资活动有关的现金
				"CInfFrInvestA":      "double", //投资活动现金流入小计
				"PurFixAssetsOth":    "double", //购建固定资产、无形资产和其他长期资产支付的现金
				"CPaidInvest":        "double", //投资支付的现金
				"NIncrPledgeLoan":    "double", //质押贷款净增加额
				"NCPaidAcquis":       "double", //取得子公司及其他营业单位支付的现金净额
				"CPaidOthInvestA":    "double", //支付其他与投资活动有关的现金
				"COutfFrInvestA":     "double", //投资活动现金流出小计
				"NCFFrInvestA":       "double", //投资活动产生的现金流量净额
				"CFrCapContr":        "double", //吸收投资收到的现金
				"CFrMinoSSubs":       "double", //其中:子公司吸收少数股东投资收到的现金
				"CFrBorr":            "double", //取得借款收到的现金
				"CFrIssueBond":       "double", //发行债券收到的现金
				"CFrOthFinanA":       "double", //收到其他与筹资活动有关的现金
				"CInfFrFinanA":       "double", //筹资活动现金流入小计
				"CPaidForDebts":      "double", //偿还债务支付的现金
				"CPaidDivProfInt":    "double", //分配股利、利润或偿付利息支付的现金
				"divProfSubsMinoS":   "double", //其中:子公司支付给少数股东的股利、利润
				"CPaidOthFinanA":     "double", //支付其他与筹资活动有关的现金
				"COutfFrFinanA":      "double", //筹资活动现金流出小计
				"NCFFrFinanA":        "double", //筹资活动产生的现金流量净额
				"forexEffects":       "double", //汇率变动对现金及现金等价物的影响
				"NChangeInCash":      "double", //现金及现金等价物净增加额
				"NCEBegBal":          "double", //加:期初现金及现金等价物余额
				"NCEEndBal":          "double", //期末现金及现金等价物余额
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的合并资产负债表模板，收集了2007年以来沪深上市公司定期报告中各个会计期间的资产负债表数据； 2、仅收集合并报表数据，包括期末和期初数据； 3、如果上市公司对外财务报表进行更正，调整，均有采集并对外展示； 4、本表中单位为人民币元； 5、每季更新。
		"getFdmtBS": map[string]interface{}{
			"url": "/api/fundamental/getFdmtBS.json",
			"args": map[string]string{
				"secID":            "string", //证券内部编码,可通过交易代码和交易市场在getSecurityID获取到,如'000002.XSHE'
				"ticker":           "string", //股票代码，如'000002'
				"beginDate":        "date",   //会计期间截止日期,起始时间,如‘20121231’
				"endDate":          "date",   //会计期间截止日期,结束时间,如‘20131231’
				"publishDateBegin": "date",   //证券交易所披露的信息发布日期,起始时间,默认为1年前当前日期,如‘20130407’
				"publishDateEnd":   "date",   //证券交易所披露的信息发布日期,结束时间,默认为当前日期,如‘20140407’
				"beginDateRep":     "date",   //报表会计期末，起始时间，输入格式“YYYYMMDD”
				"endDateRep":       "date",   //报表会计期末，结束时间，输入格式“YYYYMMDD”
				"beginYear":        "string", //会计期间所属年份
				"endYear":          "string", //会计期间所属年份，结束时间，输入格式"YYYY"
				"reportType":       "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，Q3-第三季报，A-年报。
				"field":            "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志
				"reportType":         "string", //报告类型
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"cashCEquiv":         "double", //货币资金
				"settProv":           "double", //结算备付金
				"loanToOthBankFi":    "double", //拆出资金
				"tradingFA":          "double", //交易性金融资产
				"NotesReceiv":        "double", //应收票据
				"AR":                 "double", //应收账款
				"prepayment":         "double", //预付款项
				"premiumReceiv":      "double", //应收保费
				"reinsurReceiv":      "double", //应收分保账款
				"reinsurReserReceiv": "double", //应收分保合同准备金
				"intReceiv":          "double", //应收利息
				"divReceiv":          "double", //应收股利
				"othReceiv":          "double", //其他应收款
				"purResaleFa":        "double", //买入返售金融资产
				"inventories":        "double", //存货
				"NCAWithin1Y":        "double", //一年内到期的非流动资产
				"othCA":              "double", //其他流动资产
				"TCA":                "double", //流动资产合计
				"disburLA":           "double", //发放委托贷款及垫款
				"availForSaleFa":     "double", //可供出售金融资产
				"htmInvest":          "double", //持有至到期投资
				"LTReceive":          "double", //长期应收款
				"LTEquityInvest":     "double", //长期股权投资
				"investRealEstate":   "double", //投资性房地产
				"fixedAssets":        "double", //固定资产
				"CIP":                "double", //在建工程
				"constMaterials":     "double", //工程物资
				"fixedAssetsDisp":    "double", //固定资产清理
				"producBiolAssets":   "double", //生产性生物资产
				"oilAndGasAssets":    "double", //油气资产
				"intanAssets":        "double", //无形资产
				"RD":                 "double", //开发支出
				"goodwill":           "double", //商誉
				"LTAmorExp":          "double", //长期待摊费用
				"deferTaxAssets":     "double", //递延所得税资产
				"othNCA":             "double", //其他非流动资产
				"TNCA":               "double", //非流动资产合计
				"TAssets":            "double", //资产总计
				"STBorr":             "double", //短期借款
				"CBBorr":             "double", //向中央银行借款
				"depos":              "double", //吸收存款及同业存放
				"loanFrOthBankFi":    "double", //拆入资金
				"tradingFL":          "double", //交易性金融负债
				"NotesPayable":       "double", //应付票据
				"AP":                 "double", //应付账款
				"advanceReceipts":    "double", //预收款项
				"soldForRepurFa":     "double", //卖出回购金融资产款
				"commisPayable":      "double", //应付手续费及佣金
				"payrollPayable":     "double", //应付职工薪酬
				"taxesPayable":       "double", //应交税费
				"intPayable":         "double", //应付利息
				"divPayable":         "double", //应付股利
				"othPayable":         "double", //其他应付款
				"reinsurPayable":     "double", //应付分保账款
				"insurReser":         "double", //保险合同准备金
				"fundsSecTradAgen":   "double", //代理买卖证券款
				"fundsSecUndwAgen":   "double", //代理承销证券款
				"NCLWithin1Y":        "double", //一年内到期的非流动负债
				"othCL":              "double", //其他流动负债
				"TCL":                "double", //流动负债合计
				"LTBorr":             "double", //长期借款
				"bondPayable":        "double", //应付债券
				"LTPayable":          "double", //长期应付款
				"specificPayables":   "double", //专项应付款
				"estimatedLiab":      "double", //预计负债
				"deferTaxLiab":       "double", //递延所得税负债
				"othNCL":             "double", //其他非流动负债
				"TNCL":               "double", //非流动负债合计
				"TLiab":              "double", //负债合计
				"paidInCapital":      "double", //实收资本(或股本)
				"capitalReser":       "double", //资本公积
				"treasuryShare":      "double", //减:库存股
				"specialReser":       "double", //专项储备
				"surplusReser":       "double", //盈余公积
				"ordinRiskReser":     "double", //一般风险准备
				"retainedEarnings":   "double", //未分配利润
				"forexDiffer":        "double", //外币报表折算差额
				"TEquityAttrP":       "double", //归属于母公司所有者权益合计
				"minorityInt":        "double", //少数股东权益
				"TShEquity":          "double", //所有者权益合计
				"TLiabEquity":        "double", //负债和所有者权益总计
			},
			"indices": map[string]string{},
		},
		//说明1、可获取上市公司最近一次数据，根据2007年新会计准则制定的合并利润表模板，仅收集合并报表数据；2、如果上市公司对外财务报表进行更正，调整，展示上市公司最新披露一期财务数据；3、本表中单位为人民币元；4、每季更新。
		"getFdmtISLately": map[string]interface{}{
			"url": "/api/fundamental/getFdmtISLately.json",
			"args": map[string]string{
				"field": "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志
				"reportType":         "string", //报告类型
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"tRevenue":           "double", //营业总收入
				"revenue":            "double", //营业收入
				"intIncome":          "double", //利息收入
				"intExp":             "double", //利息支出
				"premEarned":         "double", //已赚保费
				"commisIncome":       "double", //手续费及佣金收入
				"commisExp":          "double", //手续费及佣金支出
				"TCogs":              "double", //营业总成本
				"COGS":               "double", //营业成本
				"premRefund":         "double", //退保金
				"NCompensPayout":     "double", //赔付支出净额
				"reserInsurContr":    "double", //提取保险合同准备金净额
				"policyDivPayt":      "double", //保单红利支出
				"reinsurExp":         "double", //分保费用
				"bizTaxSurchg":       "double", //营业税金及附加
				"sellExp":            "double", //销售费用
				"adminExp":           "double", //管理费用
				"finanExp":           "double", //财务费用
				"assetsImpairLoss":   "double", //资产减值损失
				"fValueChgGain":      "double", //公允价值变动收益
				"investIncome":       "double", //投资收益
				"AJInvestIncome":     "double", //其中:对联营企业和合营企业的投资收益
				"forexGain":          "double", //汇兑收益
				"operateProfit":      "double", //营业利润
				"NoperateIncome":     "double", //营业外收入
				"NoperateExp":        "double", //营业外支出
				"NCADisploss":        "double", //非流动资产处置损失
				"TProfit":            "double", //利润总额
				"incomeTax":          "double", //所得税费用
				"NIncome":            "double", //净利润
				"NIncomeAttrP":       "double", //归属于母公司所有者的净利润
				"minorityGain":       "double", //少数股东损益
				"basicEPS":           "double", //基本每股收益
				"dilutedEPS":         "double", //稀释每股收益
				"othComprIncome":     "double", //其他综合收益
				"TComprIncome":       "double", //综合收益总额
				"comprIncAttrP":      "double", //归属于母公司所有者的综合收益总额
				"comprIncAttrMS":     "double", //归属于少数股东的综合收益总额
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的保险业资产负债表模板，收集了2007年以来沪深上市公司定期报告中所有以此模板披露的资产负债表数据；（主要是保险业上市公司） 2、仅收集合并报表数据，包括期末和期初数据； 3、如果上市公司对外财务报表进行更正，调整，展示上市公司最新披露一期财务数据； 4、本表中单位为人民币元； 5、每季更新。
		"getFdmtBSInsuAllLatest": map[string]interface{}{
			"url": "/api/fundamental/getFdmtBSInsuAllLatest.json",
			"args": map[string]string{
				"reportType": "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，Q3-第三季报，A-年报。
				"secID":      "string", //证券内部编码,可通过交易代码和交易市场在getSecID获取到,如'601318.XSHG'
				"ticker":     "string", //股票代码，如'601318'
				"beginDate":  "date",   //会计期间截止日期，起始时间，输入格式“YYYYMMDD”
				"endDate":    "date",   //会计期间截止日期，结束时间，输入格式“YYYYMMDD”
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志:1-合并
				"reportType":         "string", //报告类型
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"cashCEquiv":         "double", //货币资金
				"loanToOthBankFi":    "double", //拆出资金
				"tradingFA":          "double", //交易性金融资产
				"premiumReceiv":      "double", //应收保费
				"reinsurReceiv":      "double", //应收分保账款
				"intReceiv":          "double", //应收利息
				"purResaleFa":        "double", //买入返售金融资产
				"availForSaleFa":     "double", //可供出售金融资产
				"htmInvest":          "double", //持有至到期投资
				"LTEquityInvest":     "double", //长期股权投资
				"investRealEstate":   "double", //投资性房地产
				"fixedAssets":        "double", //固定资产
				"intanAssets":        "double", //无形资产
				"deferTaxAssets":     "double", //递延所得税资产
				"derivAssets":        "double", //衍生金融资产
				"subrogRecoReceiv":   "double", //应收代位追偿款
				"RRReinsUnePrem":     "double", //应收分保未到期责任准备金
				"RRReinsOutstdCla":   "double", //应收分保未决赔款准备金
				"RRReinsLinsLiab":    "double", //应收分保寿险责任准备金
				"RRReinsLThinsLiab":  "double", //应收分保长期健康险责任准备金
				"PHPledgeLoans":      "double", //保户质押贷款
				"fixedTermDepos":     "double", //定期存款
				"refundCapDepos":     "double", //存出资本保证金
				"indepAccAssets":     "double", //独立账户资产
				"othAssets":          "double", //其他资产
				"TAssets":            "double", //资产总计
				"STBorr":             "double", //短期借款
				"loanFrOthBankFi":    "double", //拆入资金
				"tradingFL":          "double", //交易性金融负债
				"soldForRepurFa":     "double", //卖出回购金融资产款
				"commisPayable":      "double", //应付手续费及佣金
				"payrollPayable":     "double", //应付职工薪酬
				"taxesPayable":       "double", //应交税费
				"reinsurPayable":     "double", //应付分保账款
				"LTBorr":             "double", //长期借款
				"bondPayable":        "double", //应付债券
				"deferTaxLiab":       "double", //递延所得税负债
				"derivLiab":          "double", //衍生金融负债
				"premReceivAdva":     "double", //预收保费
				"indemAccPayable":    "double", //应付赔付款
				"policyDivPayable":   "double", //应付保单红利
				"PHInvest":           "double", //保户储金及投资款
				"reserUnePrem":       "double", //未到期责任准备金
				"reserOutstdClaims":  "double", //未决赔款准备金
				"reserLinsLiab":      "double", //寿险责任准备金
				"reserLthinsLiab":    "double", //长期健康险责任准备金
				"indeptAccLiab":      "double", //独立账户负债
				"othLiab":            "double", //其他负债
				"TLiab":              "double", //负债合计
				"paidInCapital":      "double", //实收资本(或股本)
				"capitalReser":       "double", //资本公积
				"treasuryShare":      "double", //减:库存股
				"surplusReser":       "double", //盈余公积
				"ordinRiskReser":     "double", //一般风险准备
				"retainedEarnings":   "double", //未分配利润
				"forexDiffer":        "double", //外币报表折算差额
				"TEquityAttrP":       "double", //归属于母公司所有者权益合计
				"minorityInt":        "double", //少数股东权益
				"TShEquity":          "double", //所有者权益合计
				"TLiabEquity":        "double", //负债和所有者权益总计
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的一般工商业现金流量表模板，收集了2007年以来沪深上市公司定期报告中所有以此模板披露的现金流量表数据（主要是一般工商业上市公司）；2、仅收集合并报表数据，包括本期和上期数据；3、如果上市公司对外财务报表进行更正，调整，展示上市公司最新披露一期财务数据；4、本表中单位为人民币元；5、每季更新。
		"getFdmtCFInduAllLatest": map[string]interface{}{
			"url": "/api/fundamental/getFdmtCFInduAllLatest.json",
			"args": map[string]string{
				"reportType": "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，CQ3-三季报（累计1-9月），A-年报。
				"secID":      "string", //证券内部编码,可通过交易代码和交易市场在getSecID获取到,如'000002.XSHE'
				"ticker":     "string", //股票代码，如'000002'
				"beginDate":  "date",   //会计期间截止日期，起始时间，输入格式“YYYYMMDD”
				"endDate":    "date",   //会计期间截止日期，结束时间，输入格式“YYYYMMDD”
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志
				"reportType":         "string", //报告类型
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"CFrSaleGS":          "double", //销售商品、提供劳务收到的现金
				"NDeposIncrCFI":      "double", //客户存款和同业存放款项净增加额
				"NIncrBorrFrCB":      "double", //向中央银行借款净增加额
				"NIncBorrOthFI":      "double", //向其他金融机构拆入资金净增加额
				"premFrOrigContr":    "double", //收到原保险合同保费取得的现金
				"NReinsurPrem":       "double", //收到再保险业务现金净额
				"NIncPhDeposInv":     "double", //保户储金及投资款净增加额
				"NIncDispTradFA":     "double", //处置交易性金融资产净增加额
				"IFCCashIncr":        "double", //收取利息、手续费及佣金的现金
				"NIncFrBorr":         "double", //拆入资金净增加额
				"NCApIncrRepur":      "double", //回购业务资金净增加额
				"refundOfTax":        "double", //收到的税费返还
				"CFrOthOperateA":     "double", //收到其他与经营活动有关的现金
				"CInfFrOperateA":     "double", //经营活动现金流入小计
				"CPaidGS":            "double", //购买商品、接受劳务支付的现金
				"NIncDisburOfLA":     "double", //客户贷款及垫款净增加额
				"NIncrDeposInFI":     "double", //存放中央银行和同业款项净增加额
				"origContrCIndem":    "double", //支付原保险合同赔付款项的现金
				"CPaidIFC":           "double", //支付利息、手续费及佣金的现金
				"CPaidPolDiv":        "double", //支付保单红利的现金
				"CPaidToForEmpl":     "double", //支付给职工以及为职工支付的现金
				"CPaidForTaxes":      "double", //支付的各项税费
				"CPaidForOthOpA":     "double", //支付其他与经营活动有关的现金
				"COutfOperateA":      "double", //经营活动现金流出小计
				"NCFOperateA":        "double", //经营活动产生的现金流量净额
				"procSellInvest":     "double", //收回投资收到的现金
				"gainInvest":         "double", //取得投资收益收到的现金
				"dispFixAssetsOth":   "double", //处置固定资产、无形资产和其他长期资产收回的现金净额
				"NDispSubsOthBizC":   "double", //处置子公司及其他营业单位收到的现金净额
				"CFrOthInvestA":      "double", //收到其他与投资活动有关的现金
				"CInfFrInvestA":      "double", //投资活动现金流入小计
				"purFixAssetsOth":    "double", //购建固定资产、无形资产和其他长期资产支付的现金
				"CPaidInvest":        "double", //投资支付的现金
				"NIncrPledgeLoan":    "double", //质押贷款净增加额
				"NCPaidAcquis":       "double", //取得子公司及其他营业单位支付的现金净额
				"CPaidOthInvestA":    "double", //支付其他与投资活动有关的现金
				"COutfFrInvestA":     "double", //投资活动现金流出小计
				"NCFFrInvestA":       "double", //投资活动产生的现金流量净额
				"CFrCapContr":        "double", //吸收投资收到的现金
				"CFrMinoSSubs":       "double", //其中:子公司吸收少数股东投资收到的现金
				"CFrBorr":            "double", //取得借款收到的现金
				"CFrIssueBond":       "double", //发行债券收到的现金
				"CFrOthFinanA":       "double", //收到其他与筹资活动有关的现金
				"CInfFrFinanA":       "double", //筹资活动现金流入小计
				"CPaidForDebts":      "double", //偿还债务支付的现金
				"CPaidDivProfInt":    "double", //分配股利、利润或偿付利息支付的现金
				"divProfSubsMinoS":   "double", //其中:子公司支付给少数股东的股利、利润
				"CPaidOthFinanA":     "double", //支付其他与筹资活动有关的现金
				"COutfFrFinanA":      "double", //筹资活动现金流出小计
				"NCFFrFinanA":        "double", //筹资活动产生的现金流量净额
				"forexEffects":       "double", //汇率变动对现金及现金等价物的影响
				"NChangeInCash":      "double", //现金及现金等价物净增加额
				"NCEBegBal":          "double", //加:期初现金及现金等价物余额
				"NCEEndBal":          "double", //期末现金及现金等价物余额
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的保险业利润表模板，收集了2007年以来沪深上市公司定期报告中所有以此模板披露的利润表数据；（主要是保险业上市公司） 2、仅收集合并报表数据，包括本期和上期数据； 3、如果上市公司对外财务报表进行更正，调整，均有采集并对外展示； 4、本表中单位为人民币元； 5、每季更新。
		"getFdmtISInsu": map[string]interface{}{
			"url": "/api/fundamental/getFdmtISInsu.json",
			"args": map[string]string{
				"secID":            "string", //证券内部编码,可通过交易代码和交易市场在getSecurityID获取到,如'601318.XSHG'
				"ticker":           "string", //股票代码，如'601318'
				"beginDate":        "date",   //会计期间截止日期,起始时间,如‘20121231’
				"endDate":          "date",   //会计期间截止日期,结束时间,如‘20131231’
				"fiscalPeriod":     "String", //会计期间长度,可多值输入，如：3，6，9，12，3表示季度值
				"publishDateBegin": "date",   //证券交易所披露的信息发布日期,起始时间,默认为1年前当前日期,如‘20130407’
				"publishDateEnd":   "date",   //证券交易所披露的信息发布日期,结束时间,默认为当前日期,如‘20140407’
				"beginDateRep":     "date",   //报表会计期末，起始时间，输入格式“YYYYMMDD”
				"endDateRep":       "date",   //报表会计期末，结束时间，输入格式“YYYYMMDD”
				"beginYear":        "string", //会计期间所属年份，起始时间，输入格式"YYYY"
				"endYear":          "string", //会计期间所属年份，结束时间，输入格式"YYYY"
				"reportType":       "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，Q3-第三季报，CQ3-三季报（累计1-9月），A-年报。
				"field":            "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"reportType":         "string", //报告类型
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"revenue":            "double", //营业收入
				"premEarned":         "double", //已赚保费
				"grossPremWrit":      "double", //保险业务收入
				"reinsIncome":        "double", //其中:分保费收入
				"reinsur":            "double", //减:分出保费
				"unePremReser":       "double", //提取未到期责任准备金
				"commisExp":          "double", //手续费及佣金支出
				"othOperRev":         "double", //其他业务收入
				"COGS":               "double", //营业支出
				"premRefund":         "double", //退保金
				"compensPayout":      "double", //赔付支出
				"compensPayoutRefu":  "double", //减:摊回赔付支出
				"reserInsurLiab":     "double", //提取保险责任准备金
				"insurLiabReserRefu": "double", //减:摊回保险责任准备金
				"policyDivPayt":      "double", //保单红利支出
				"reinsurExp":         "double", //分保费用
				"bizTaxSurchg":       "double", //营业税金及附加
				"genlAdminExp":       "double", //业务及管理费
				"reinsCostRefund":    "double", //减:摊回分保费用
				"assetsImpairLoss":   "double", //资产减值损失
				"othOperCosts":       "double", //其他业务成本
				"fValueChgGain":      "double", //公允价值变动收益
				"investIncome":       "double", //投资收益
				"AJInvestIncome":     "double", //其中:对联营企业和合营企业的投资收益
				"forexGain":          "double", //汇兑收益
				"operateProfit":      "double", //营业利润
				"NoperateIncome":     "double", //营业外收入
				"NoperateExp":        "double", //营业外支出
				"TProfit":            "double", //利润总额
				"incomeTax":          "double", //所得税费用
				"NIncome":            "double", //净利润
				"NIncomeAttrP":       "double", //归属于母公司所有者的净利润
				"minorityGain":       "double", //少数股东损益
				"basicEPS":           "double", //基本每股收益
				"dilutedEPS":         "double", //稀释每股收益
				"othComprIncome":     "double", //其他综合收益
				"TComprIncome":       "double", //综合收益总额
				"comprIncAttrP":      "double", //归属于母公司所有者的综合收益总额
				"comprIncAttrMS":     "double", //归属于少数股东的综合收益总额
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的保险业利润表模板，收集了2007年以来沪深上市公司定期报告中所有以此模板披露的利润表数据（主要是保险业上市公司）；2、仅收集合并报表数据，包括本期和上期数据；3、如果上市公司对外财务报表进行更正，调整，展示上市公司最新披露一期财务数据；4、本表中单位为人民币元；5、每季更新。
		"getFdmtISInsuAllLatest": map[string]interface{}{
			"url": "/api/fundamental/getFdmtISInsuAllLatest.json",
			"args": map[string]string{
				"reportType": "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，Q3-第三季报，CQ3-三季报（累计1-9月），A-年报。
				"secID":      "string", //证券内部编码,可通过交易代码和交易市场在getSecID获取到,如'601318.XSHG'
				"ticker":     "string", //股票代码，如'601318'
				"beginDate":  "date",   //会计期间截止日期，起始时间，输入格式“YYYYMMDD”
				"endDate":    "date",   //会计期间截止日期，结束时间，输入格式“YYYYMMDD”
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志
				"reportType":         "string", //报告类型
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"revenue":            "double", //营业收入
				"premEarned":         "double", //已赚保费
				"grossPremWrit":      "double", //保险业务收入
				"reinsIncome":        "double", //其中:分保费收入
				"reinsur":            "double", //减:分出保费
				"unePremReser":       "double", //提取未到期责任准备金
				"commisExp":          "double", //手续费及佣金支出
				"othOperRev":         "double", //其他业务收入
				"COGS":               "double", //营业支出
				"premRefund":         "double", //退保金
				"compensPayout":      "double", //赔付支出
				"compensPayoutRefu":  "double", //减:摊回赔付支出
				"reserInsurLiab":     "double", //提取保险责任准备金
				"insurLiabReserRefu": "double", //减:摊回保险责任准备金
				"policyDivPayt":      "double", //保单红利支出
				"reinsurExp":         "double", //分保费用
				"bizTaxSurchg":       "double", //营业税金及附加
				"genlAdminExp":       "double", //业务及管理费
				"reinsCostRefund":    "double", //减:摊回分保费用
				"assetsImpairLoss":   "double", //资产减值损失
				"othOperCosts":       "double", //其他业务成本
				"fValueChgGain":      "double", //公允价值变动收益
				"investIncome":       "double", //投资收益
				"AJInvestIncome":     "double", //其中:对联营企业和合营企业的投资收益
				"forexGain":          "double", //汇兑收益
				"operateProfit":      "double", //营业利润
				"NoperateIncome":     "double", //营业外收入
				"NoperateExp":        "double", //营业外支出
				"TProfit":            "double", //利润总额
				"incomeTax":          "double", //所得税费用
				"NIncome":            "double", //净利润
				"NIncomeAttrP":       "double", //归属于母公司所有者的净利润
				"minorityGain":       "double", //少数股东损益
				"basicEPS":           "double", //基本每股收益
				"dilutedEPS":         "double", //稀释每股收益
				"othComprIncome":     "double", //其他综合收益
				"TComprIncome":       "double", //综合收益总额
				"comprIncAttrP":      "double", //归属于母公司所有者的综合收益总额
				"comprIncAttrMS":     "double", //归属于少数股东的综合收益总额
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的银行业资产负债表模板，收集了2007年以来沪深上市公司定期报告中所有以此模板披露的资产负债表数据；（主要是银行业上市公司） 2、仅收集合并报表数据，包括期末和期初数据； 3、如果上市公司对外财务报表进行更正，调整，均有采集并对外展示； 4、本表中单位为人民币元； 5、每季更新。
		"getFdmtBSBank": map[string]interface{}{
			"url": "/api/fundamental/getFdmtBSBank.json",
			"args": map[string]string{
				"secID":            "string", //证券内部编码,可通过交易代码和交易市场在getSecurityID获取到,如'000001.XSHE'
				"ticker":           "string", //股票代码，如'000001'
				"beginDate":        "date",   //会计期间截止日期,起始时间,如‘20121231’
				"endDate":          "date",   //会计期间截止日期,结束时间,如‘20131231’
				"publishDateBegin": "date",   //证券交易所披露的信息发布日期，起始时间，默认为1年前当前日期，如‘20130812’
				"publishDateEnd":   "date",   //证券交易所披露的信息发布日期，结束时间，默认为当前日期，如‘20140812’
				"beginDateRep":     "date",   //报表会计期末，起始时间，输入格式“YYYYMMDD”
				"endDateRep":       "date",   //报表会计期末，结束时间，输入格式“YYYYMMDD”
				"beginYear":        "string", //会计期间所属年份，起始时间，输入格式"YYYY"
				"endYear":          "string", //会计期间所属年份，结束时间，输入格式"YYYY"
				"reportType":       "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，Q3-第三季报，A-年报。
				"field":            "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"reportType":         "string", //报告类型
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"loanToOthBankFi":    "double", //拆出资金
				"tradingFA":          "double", //交易性金融资产
				"intReceiv":          "double", //应收利息
				"purResaleFa":        "double", //买入返售金融资产
				"disburLA":           "double", //发放委托贷款及垫款
				"availForSaleFA":     "double", //可供出售金融资产
				"htmInvest":          "double", //持有至到期投资
				"LTEquityInvest":     "double", //长期股权投资
				"investRealEstate":   "double", //投资性房地产
				"fixedAssets":        "double", //固定资产
				"CIP":                "double", //在建工程
				"intanAssets":        "double", //无形资产
				"goodwill":           "double", //商誉
				"deferTaxAssets":     "double", //递延所得税资产
				"CReserCB":           "double", //现金及存放中央银行款项
				"deposInOthBfi":      "double", //存放同业款项
				"preciMetals":        "double", //贵金属
				"derivAssets":        "double", //衍生金融资产
				"finanLeaseReceiv":   "double", //应收融资租赁款
				"investAsReceiv":     "double", //应收款项类投资
				"othAssets":          "double", //其他资产
				"TAssets":            "double", //资产总计
				"CBBorr":             "double", //向中央银行借款
				"depos":              "double", //吸收存款
				"loanFrOthBankFi":    "double", //拆入资金
				"tradingFL":          "double", //交易性金融负债
				"soldForRepurFa":     "double", //卖出回购金融资产款
				"payrollPayable":     "double", //应付职工薪酬
				"taxesPayable":       "double", //应交税费
				"intPayable":         "double", //应付利息
				"bondPayable":        "double", //应付债券
				"estimatedLiab":      "double", //预计负债
				"deferTaxLiab":       "double", //递延所得税负债
				"deposFrOthBfi":      "double", //同业及其他金融机构存放款项
				"derivLiab":          "double", //衍生金融负债
				"othLiab":            "double", //其他负债
				"TLiab":              "double", //负债合计
				"paidInCapital":      "double", //实收资本(或股本)
				"capitalReser":       "double", //资本公积
				"treasuryShare":      "double", //减:库存股
				"surplusReser":       "double", //盈余公积
				"ordinRiskReser":     "double", //一般风险准备
				"retainedEarnings":   "double", //未分配利润
				"forexDiffer":        "double", //外币报表折算差额
				"TEquityAttrP":       "double", //归属于母公司所有者权益合计
				"minorityInt":        "double", //少数股东权益
				"TShEquity":          "double", //所有者权益合计
				"TLiabEquity":        "double", //负债和所有者权益总计
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的保险业资产负债表模板，收集了2007年以来沪深上市公司定期报告中所有以此模板披露的资产负债表数据；（主要是保险业上市公司） 2、仅收集合并报表数据，包括期末和期初数据； 3、如果上市公司对外财务报表进行更正，调整，均有采集并对外展示； 4、本表中单位为人民币元。 5、每季更新。
		"getFdmtBSInsu": map[string]interface{}{
			"url": "/api/fundamental/getFdmtBSInsu.json",
			"args": map[string]string{
				"secID":            "string", //证券内部编码,可通过交易代码和交易市场在getSecurityID获取到,如'601318.XSHG'
				"ticker":           "string", //股票代码，如'601318'
				"beginDate":        "date",   //会计期间截止日期,起始时间,如‘20121231’
				"endDate":          "date",   //会计期间截止日期,结束时间,如‘20131231’
				"publishDateBegin": "date",   //证券交易所披露的信息发布日期，起始时间，默认为1年前当前日期，如‘20130812’
				"publishDateEnd":   "date",   //证券交易所披露的信息发布日期，结束时间，默认为当前日期，如‘20140812’
				"beginDateRep":     "date",   //报表会计期末，起始时间，输入格式“YYYYMMDD”
				"endDateRep":       "date",   //报表会计期末，结束时间，输入格式“YYYYMMDD”
				"beginYear":        "string", //会计期间所属年份，起始时间，输入格式"YYYY"
				"endYear":          "string", //会计期间所属年份，结束时间，输入格式"YYYY"
				"reportType":       "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，Q3-第三季报，A-年报。
				"field":            "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"reportType":         "string", //报告类型
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"cashCEquiv":         "double", //货币资金
				"loanToOthBankFi":    "double", //拆出资金
				"tradingFA":          "double", //交易性金融资产
				"premiumReceiv":      "double", //应收保费
				"reinsurReceiv":      "double", //应收分保账款
				"intReceiv":          "double", //应收利息
				"purResaleFa":        "double", //买入返售金融资产
				"availForSaleFa":     "double", //可供出售金融资产
				"htmInvest":          "double", //持有至到期投资
				"LTEquityInvest":     "double", //长期股权投资
				"investRealEstate":   "double", //投资性房地产
				"fixedAssets":        "double", //固定资产
				"intanAssets":        "double", //无形资产
				"deferTaxAssets":     "double", //递延所得税资产
				"derivAssets":        "double", //衍生金融资产
				"subrogRecoReceiv":   "double", //应收代位追偿款
				"RRReinsUnePrem":     "double", //应收分保未到期责任准备金
				"RRReinsOutstdCla":   "double", //应收分保未决赔款准备金
				"RRReinsLinsLiab":    "double", //应收分保寿险责任准备金
				"RRReinsLThinsLiab":  "double", //应收分保长期健康险责任准备金
				"PHPledgeLoans":      "double", //保户质押贷款
				"fixedTermDepos":     "double", //定期存款
				"refundCapDepos":     "double", //存出资本保证金
				"indepAccAssets":     "double", //独立账户资产
				"othAssets":          "double", //其他资产
				"TAssets":            "double", //资产总计
				"STBorr":             "double", //短期借款
				"loanFrOthBankFi":    "double", //拆入资金
				"tradingFL":          "double", //交易性金融负债
				"soldForRepurFa":     "double", //卖出回购金融资产款
				"commisPayable":      "double", //应付手续费及佣金
				"payrollPayable":     "double", //应付职工薪酬
				"taxesPayable":       "double", //应交税费
				"reinsurPayable":     "double", //应付分保账款
				"LTBorr":             "double", //长期借款
				"bondPayable":        "double", //应付债券
				"deferTaxLiab":       "double", //递延所得税负债
				"derivLiab":          "double", //衍生金融负债
				"premReceivAdva":     "double", //预收保费
				"indemAccPayable":    "double", //应付赔付款
				"policyDivPayable":   "double", //应付保单红利
				"PHInvest":           "double", //保户储金及投资款
				"reserUnePrem":       "double", //未到期责任准备金
				"reserOutstdClaims":  "double", //未决赔款准备金
				"reserLinsLiab":      "double", //寿险责任准备金
				"reserLthinsLiab":    "double", //长期健康险责任准备金
				"indeptAccLiab":      "double", //独立账户负债
				"othLiab":            "double", //其他负债
				"TLiab":              "double", //负债合计
				"paidInCapital":      "double", //实收资本(或股本)
				"capitalReser":       "double", //资本公积
				"treasuryShare":      "double", //减:库存股
				"surplusReser":       "double", //盈余公积
				"ordinRiskReser":     "double", //一般风险准备
				"retainedEarnings":   "double", //未分配利润
				"forexDiffer":        "double", //外币报表折算差额
				"TEquityAttrP":       "double", //归属于母公司所有者权益合计
				"minorityInt":        "double", //少数股东权益
				"TShEquity":          "double", //所有者权益合计
				"TLiabEquity":        "double", //负债和所有者权益总计
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的一般工商业现金流量表模板，收集了2007年以来沪深上市公司定期报告中所有以此模板披露的现金流量表数据；（主要是一般工商业上市公司） 2、仅收集合并报表数据，包括本期和上期数据； 3、如果上市公司对外财务报表进行更正，调整，均有采集并对外展示； 4、本表中单位为人民币元； 5、每季更新。
		"getFdmtCFIndu": map[string]interface{}{
			"url": "/api/fundamental/getFdmtCFIndu.json",
			"args": map[string]string{
				"secID":            "string", //证券内部编码,可通过交易代码和交易市场在getSecurityID获取到,如'000002.XSHE'
				"ticker":           "string", //股票代码，如'000002'
				"beginDate":        "date",   //会计期间截止日期,起始时间,如‘20121231’
				"endDate":          "date",   //会计期间截止日期,结束时间,如‘20131231’
				"fiscalPeriod":     "String", //会计期间长度,可多值输入，如：3，6，9，12，3表示季度值
				"publishDateBegin": "date",   //证券交易所披露的信息发布日期,起始时间,默认为1年前当前日期,如‘20130407’
				"publishDateEnd":   "date",   //证券交易所披露的信息发布日期,结束时间,默认为当前日期,如‘20140407’
				"beginDateRep":     "date",   //报表会计期末，起始时间，输入格式“YYYYMMDD”
				"endDateRep":       "date",   //报表会计期末，结束时间，输入格式“YYYYMMDD”
				"beginYear":        "string", //会计期间所属年份，起始时间，输入格式"YYYY"
				"endYear":          "string", //会计期间所属年份，结束时间，输入格式"YYYY"
				"reportType":       "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，CQ3-三季报（累计1-9月），A-年报。
				"field":            "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"reportType":         "string", //报告类型
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"CFrSaleGS":          "double", //销售商品、提供劳务收到的现金
				"NDeposIncrCFI":      "double", //客户存款和同业存放款项净增加额
				"NIncrBorrFrCB":      "double", //向中央银行借款净增加额
				"NIncBorrOthFI":      "double", //向其他金融机构拆入资金净增加额
				"premFrOrigContr":    "double", //收到原保险合同保费取得的现金
				"NReinsurPrem":       "double", //收到再保险业务现金净额
				"NIncPhDeposInv":     "double", //保户储金及投资款净增加额
				"NIncDispTradFA":     "double", //处置交易性金融资产净增加额
				"IFCCashIncr":        "double", //收取利息、手续费及佣金的现金
				"NIncFrBorr":         "double", //拆入资金净增加额
				"NCApIncrRepur":      "double", //回购业务资金净增加额
				"refundOfTax":        "double", //收到的税费返还
				"CFrOthOperateA":     "double", //收到其他与经营活动有关的现金
				"CInfFrOperateA":     "double", //经营活动现金流入小计
				"CPaidGS":            "double", //购买商品、接受劳务支付的现金
				"NIncDisburOfLA":     "double", //客户贷款及垫款净增加额
				"NIncrDeposInFI":     "double", //存放中央银行和同业款项净增加额
				"origContrCIndem":    "double", //支付原保险合同赔付款项的现金
				"CPaidIFC":           "double", //支付利息、手续费及佣金的现金
				"CPaidPolDiv":        "double", //支付保单红利的现金
				"CPaidToForEmpl":     "double", //支付给职工以及为职工支付的现金
				"CPaidForTaxes":      "double", //支付的各项税费
				"CPaidForOthOpA":     "double", //支付其他与经营活动有关的现金
				"COutfOperateA":      "double", //经营活动现金流出小计
				"NCFOperateA":        "double", //经营活动产生的现金流量净额
				"procSellInvest":     "double", //收回投资收到的现金
				"gainInvest":         "double", //取得投资收益收到的现金
				"dispFixAssetsOth":   "double", //处置固定资产、无形资产和其他长期资产收回的现金净额
				"NDispSubsOthBizC":   "double", //处置子公司及其他营业单位收到的现金净额
				"CFrOthInvestA":      "double", //收到其他与投资活动有关的现金
				"CInfFrInvestA":      "double", //投资活动现金流入小计
				"purFixAssetsOth":    "double", //购建固定资产、无形资产和其他长期资产支付的现金
				"CPaidInvest":        "double", //投资支付的现金
				"NIncrPledgeLoan":    "double", //质押贷款净增加额
				"NCPaidAcquis":       "double", //取得子公司及其他营业单位支付的现金净额
				"CPaidOthInvestA":    "double", //支付其他与投资活动有关的现金
				"COutfFrInvestA":     "double", //投资活动现金流出小计
				"NCFFrInvestA":       "double", //投资活动产生的现金流量净额
				"CFrCapContr":        "double", //吸收投资收到的现金
				"CFrMinoSSubs":       "double", //其中:子公司吸收少数股东投资收到的现金
				"CFrBorr":            "double", //取得借款收到的现金
				"CFrIssueBond":       "double", //发行债券收到的现金
				"CFrOthFinanA":       "double", //收到其他与筹资活动有关的现金
				"CInfFrFinanA":       "double", //筹资活动现金流入小计
				"CPaidForDebts":      "double", //偿还债务支付的现金
				"CPaidDivProfInt":    "double", //分配股利、利润或偿付利息支付的现金
				"divProfSubsMinoS":   "double", //其中:子公司支付给少数股东的股利、利润
				"CPaidOthFinanA":     "double", //支付其他与筹资活动有关的现金
				"COutfFrFinanA":      "double", //筹资活动现金流出小计
				"NCFFrFinanA":        "double", //筹资活动产生的现金流量净额
				"forexEffects":       "double", //汇率变动对现金及现金等价物的影响
				"NChangeInCash":      "double", //现金及现金等价物净增加额
				"NCEBegBal":          "double", //加:期初现金及现金等价物余额
				"NCEEndBal":          "double", //期末现金及现金等价物余额
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的银行业现金流量表模板，收集了2007年以来沪深上市公司定期报告中所有以此模板披露的现金流量表数据（主要是银行业上市公司）；2、仅收集合并报表数据，包括本期和上期数据；3、如果上市公司对外财务报表进行更正，调整，展示上市公司最新披露一期财务数据；4、本表中单位为人民币元；5、每季更新。
		"getFdmtCFBankAllLatest": map[string]interface{}{
			"url": "/api/fundamental/getFdmtCFBankAllLatest.json",
			"args": map[string]string{
				"reportType": "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，CQ3-三季报（累计1-9月），A-年报。
				"secID":      "string", //证券内部编码,可通过交易代码和交易市场在getSecID获取到,如'000001.XSHE'
				"ticker":     "string", //股票代码，如'000001'
				"beginDate":  "date",   //会计期间截止日期，起始时间，输入格式“YYYYMMDD”
				"endDate":    "date",   //会计期间截止日期，结束时间，输入格式“YYYYMMDD”
				"field":      "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志
				"reportType":         "string", //报告类型
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"NDeposIncrCFI":      "double", //客户存款和同业存放款项净增加额
				"NIncrBorrFrCB":      "double", //向中央银行借款净增加额
				"NIncBorrOthFI":      "double", //向其他金融机构拆入资金净增加额
				"NDecrInDisburOfLa":  "double", //发放贷款及垫款净减少额
				"NDecrInDeposInFI":   "double", //存放中央银行和同业款项净减少额
				"NDecrLoanToOthFI":   "double", //向其他金融机构拆出资金净减少额
				"IFCCashIncr":        "double", //收取利息、手续费及佣金的现金
				"CFrOthOperateA":     "double", //收到其他与经营活动有关的现金
				"CInfFrOperateA":     "double", //经营活动现金流入小计
				"NDeposDecrFrFI":     "double", //客户存款和同业存放款项净减少额
				"NDecrBorrFrCB":      "double", //向中央银行借款净减少额
				"NDecrBorrFrOthFI":   "double", //向其他金融机构拆入资金净减少额
				"NIncDisburOfLA":     "double", //客户贷款及垫款净增加额
				"NIncrDeposInFI":     "double", //存放中央银行和同业款项净增加额
				"NIncrLoansToOthFi":  "double", //向其他金融机构拆出资金净增加额
				"CPaidIFC":           "double", //支付利息、手续费及佣金的现金
				"CPaidToForEmpl":     "double", //支付给职工以及为职工支付的现金
				"CPaidForTaxes":      "double", //支付的各项税费
				"CPaidForOthOpA":     "double", //支付其他与经营活动有关的现金
				"COutfOperateA":      "double", //经营活动现金流出小计
				"NCFOperateA":        "double", //经营活动产生的现金流量净额
				"procSellInvest":     "double", //收回投资收到的现金
				"gainInvest":         "double", //取得投资收益收到的现金
				"dispFixAssetsOth":   "double", //处置固定资产、无形资产和其他长期资产收回的现金净额
				"NDispSubsOthBizC":   "double", //处置子公司及其他营业单位收到的现金净额
				"CFrOthInvestA":      "double", //收到其他与投资活动有关的现金
				"CInfFrInvestA":      "double", //投资活动现金流入小计
				"purFixAssetsOth":    "double", //购建固定资产、无形资产和其他长期资产支付的现金
				"CPaidInvest":        "double", //投资支付的现金
				"NCPaidAcquis":       "double", //取得子公司及其他营业单位支付的现金净额
				"CPaidOthInvestA":    "double", //支付其他与投资活动有关的现金
				"COutfFrInvestA":     "double", //投资活动现金流出小计
				"NCFFrInvestA":       "double", //投资活动产生的现金流量净额
				"CFrCapContr":        "double", //吸收投资收到的现金
				"CFrMinoSSubs":       "double", //其中:子公司吸收少数股东投资收到的现金
				"CFrIssueBond":       "double", //发行债券收到的现金
				"CFrOthFinanA":       "double", //收到其他与筹资活动有关的现金
				"CInfFrFinanA":       "double", //筹资活动现金流入小计
				"CPaidForDebts":      "double", //偿还债务支付的现金
				"CPaidDivProfInt":    "double", //分配股利、利润或偿付利息支付的现金
				"divProfSubsMinoS":   "double", //其中:子公司支付给少数股东的股利、利润
				"CPaidOthFinanA":     "double", //支付其他与筹资活动有关的现金
				"COutfFrFinanA":      "double", //筹资活动现金流出小计
				"NCFFrFinanA":        "double", //筹资活动产生的现金流量净额
				"forexEffects":       "double", //汇率变动对现金及现金等价物的影响
				"NChangeInCash":      "double", //现金及现金等价物净增加额
				"NCEBegBal":          "double", //加:期初现金及现金等价物余额
				"NCEEndBal":          "double", //期末现金及现金等价物余额
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的证券业资产负债表模板，收集了2007年以来沪深上市公司定期报告中所有以此模板披露的资产负债表数据；（主要是证券业上市公司） 2、仅收集合并报表数据，包括期末和期初数据； 3、如果上市公司对外财务报表进行更正，调整，均有采集并对外展示； 4、本表中单位为人民币元； 5、每季更新。
		"getFdmtBSSecu": map[string]interface{}{
			"url": "/api/fundamental/getFdmtBSSecu.json",
			"args": map[string]string{
				"secID":            "string", //证券内部编码,可通过交易代码和交易市场在getSecurityID获取到,如'600369.XSHG'
				"ticker":           "string", //股票代码，如'600369'
				"beginDate":        "date",   //会计期间截止日期,起始时间,如‘20121231’
				"endDate":          "date",   //会计期间截止日期,结束时间,如‘20131231’
				"publishDateBegin": "date",   //证券交易所披露的信息发布日期，起始时间，默认为1年前当前日期，如‘20130812’
				"publishDateEnd":   "date",   //证券交易所披露的信息发布日期，结束时间，默认为当前日期，如‘20140812’
				"beginDateRep":     "date",   //报表会计期末，起始时间，输入格式“YYYYMMDD”
				"endDateRep":       "date",   //报表会计期末，结束时间，输入格式“YYYYMMDD”
				"beginYear":        "string", //会计期间所属年份，起始时间，输入格式"YYYY"
				"endYear":          "string", //会计期间所属年份，结束时间，输入格式"YYYY"
				"reportType":       "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，Q3-第三季报，A-年报。
				"field":            "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"reportType":         "string", //报告类型
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"cashCEquiv":         "double", //货币资金
				"clientDepos":        "double", //其中:客户资金存款
				"settProv":           "double", //结算备付金
				"clientProv":         "double", //其中:客户备付金
				"loanToOthBankFi":    "double", //拆出资金
				"tradingFA":          "double", //交易性金融资产
				"intReceiv":          "double", //应收利息
				"purResaleFa":        "double", //买入返售金融资产
				"availForSaleFa":     "double", //可供出售金融资产
				"htmInvest":          "double", //持有至到期投资
				"LTEquityInvest":     "double", //长期股权投资
				"investRealEstate":   "double", //投资性房地产
				"fixedAssets":        "double", //固定资产
				"intanAssets":        "double", //无形资产
				"transacSeatFee":     "double", //其中:交易席位费
				"deferTaxAssets":     "double", //递延所得税资产
				"derivAssets":        "double", //衍生金融资产
				"refundDepos":        "double", //存出保证金
				"othAssets":          "double", //其他资产
				"TAssets":            "double", //资产总计
				"STBorr":             "double", //短期借款
				"pledgeBorr":         "double", //其中:质押借款
				"loanFrOthBankFi":    "double", //拆入资金
				"tradingFL":          "double", //交易性金融负债
				"soldForRepurFa":     "double", //卖出回购金融资产款
				"payrollPayable":     "double", //应付职工薪酬
				"taxesPayable":       "double", //应交税费
				"intPayable":         "double", //应付利息
				"fundsSecTradAgen":   "double", //代理买卖证券款
				"fundsSecUndwAgen":   "double", //代理承销证券款
				"LTBorr":             "double", //长期借款
				"bondPayable":        "double", //应付债券
				"estimatedLiab":      "double", //预计负债
				"deferTaxLiab":       "double", //递延所得税负债
				"derivLiab":          "double", //衍生金融负债
				"othLiab":            "double", //其他负债
				"TLiab":              "double", //负债合计
				"paidInCapital":      "double", //实收资本(或股本)
				"capitalReser":       "double", //资本公积
				"treasuryShare":      "double", //减:库存股
				"surplusReser":       "double", //盈余公积
				"ordinRiskReser":     "double", //一般风险准备
				"transacRiskReser":   "double", //交易风险准备
				"retainedEarnings":   "double", //未分配利润
				"forexDiffer":        "double", //外币报表折算差额
				"TEquityAttrP":       "double", //归属于母公司所有者权益合计
				"minorityInt":        "double", //少数股东权益
				"TShEquity":          "double", //所有者权益合计
				"TLiabEquity":        "double", //负债和所有者权益总计
			},
			"indices": map[string]string{},
		},
		//说明1、根据2007年新会计准则制定的一般工商业资产负债表模板，收集了2007年以来沪深上市公司定期报告中所有以此模板披露的资产负债表数据；（主要是一般工商业上市公司） 2、仅收集合并报表数据，包括期末和期初数据； 3、如果上市公司对外财务报表进行更正，调整，均有采集并对外展示； 4、本表中单位为人民币元； 5、每季更新。
		"getFdmtBSIndu": map[string]interface{}{
			"url": "/api/fundamental/getFdmtBSIndu.json",
			"args": map[string]string{
				"secID":            "string", //证券内部编码,可通过交易代码和交易市场在getSecurityID获取到,如'000002.XSHE'
				"ticker":           "string", //股票代码，如'000002'
				"beginDate":        "date",   //会计期间截止日期,起始时间,如‘20121231’
				"endDate":          "date",   //会计期间截止日期,结束时间,如‘20131231’
				"publishDateBegin": "date",   //证券交易所披露的信息发布日期，起始时间，默认为1年前当前日期，如‘20130812’
				"publishDateEnd":   "date",   //证券交易所披露的信息发布日期，结束时间，默认为当前日期，如‘20140812’
				"beginDateRep":     "date",   //报表会计期末，起始时间，输入格式“YYYYMMDD”
				"endDateRep":       "date",   //报表会计期末，结束时间，输入格式“YYYYMMDD”
				"beginYear":        "string", //会计期间所属年份，起始时间，输入格式"YYYY"
				"endYear":          "string", //会计期间所属年份，结束时间，输入格式"YYYY"
				"reportType":       "string", //财务报告期间类型，可供选择类型：Q1-第一季报，S1-半年报，Q3-第三季报，A-年报。
				"field":            "string", //可选参数，用,分隔,默认为空，返回全部字段，不选即为默认值。返回字段见下方
			},
			"rets": map[string]string{
				"reportType":         "string", //报告类型
				"secID":              "string", //证券内部ID
				"endDate":            "date",   //截止日期
				"publishDate":        "date",   //发布日期
				"endDateRep":         "date",   //报表截止日期
				"partyID":            "int32",  //机构内部ID
				"ticker":             "string", //股票代码
				"secShortName":       "string", //证券简称
				"exchangeCD":         "string", //交易市场代码
				"actPubtime":         "string", //实际披露时间
				"mergedFlag":         "string", //合并标志
				"fiscalPeriod":       "string", //会计期间
				"accoutingStandards": "string", //会计准则
				"currencyCD":         "string", //货币代码
				"cashCEquiv":         "double", //货币资金
				"settProv":           "double", //结算备付金
				"loanToOthBankFi":    "double", //拆出资金
				"tradingFA":          "double", //交易性金融资产
				"NotesReceiv":        "double", //应收票据
				"AR":                 "double", //应收账款
				"prepayment":         "double", //预付款项
				"premiumReceiv":      "double", //应收保费
				"reinsurReceiv":      "double", //应收分保账款
				"reinsurReserReceiv": "double", //应收分保合同准备金
				"intReceiv":          "double", //应收利息
				"divReceiv":          "double", //应收股利
				"othReceiv":          "double", //其他应收款
				"purResaleFa":        "double", //买入返售金融资产
				"inventories":        "double", //存货
				"NCAWithin1Y":        "double", //一年内到期的非流动资产
				"othCA":              "double", //其他流动资产
				"TCA":                "double", //流动资产合计
				"disburLA":           "double", //发放委托贷款及垫款
				"availForSaleFa":     "double", //可供出售金融资产
				"htmInvest":          "double", //持有至到期投资
				"LTReceive":          "double", //长期应收款
				"LTEquityInvest":     "double", //长期股权投资
				"investRealEstate":   "double", //投资性房地产
				"fixedAssets":        "double", //固定资产
				"CIP":                "double", //在建工程
				"constMaterials":     "double", //工程物资
				"fixedAssetsDisp":    "double", //固定资产清理
				"producBiolAssets":   "double", //生产性生物资产
				"oilAndGasAssets":    "double", //油气资产
				"intanAssets":        "double", //无形资产
				"RD":                 "double", //开发支出
				"goodwill":           "double", //商誉
				"LTAmorExp":          "double", //长期待摊费用
				"deferTaxAssets":     "double", //递延所得税资产
				"othNCA":             "double", //其他非流动资产
				"TNCA":               "double", //非流动资产合计
				"TAssets":            "double", //资产总计
				"STBorr":             "double", //短期借款
				"CBBorr":             "double", //向中央银行借款
				"depos":              "double", //吸收存款及同业存放
				"loanFrOthBankFi":    "double", //拆入资金
				"tradingFL":          "double", //交易性金融负债
				"NotesPayable":       "double", //应付票据
				"AP":                 "double", //应付账款
				"advanceReceipts":    "double", //预收款项
				"soldForRepurFa":     "double", //卖出回购金融资产款
				"commisPayable":      "double", //应付手续费及佣金
				"payrollPayable":     "double", //应付职工薪酬
				"taxesPayable":       "double", //应交税费
				"intPayable":         "double", //应付利息
				"divPayable":         "double", //应付股利
				"othPayable":         "double", //其他应付款
				"reinsurPayable":     "double", //应付分保账款
				"insurReser":         "double", //保险合同准备金
				"fundsSecTradAgen":   "double", //代理买卖证券款
				"fundsSecUndwAgen":   "double", //代理承销证券款
				"NCLWithin1Y":        "double", //一年内到期的非流动负债
				"othCL":              "double", //其他流动负债
				"TCL":                "double", //流动负债合计
				"LTBorr":             "double", //长期借款
				"bondPayable":        "double", //应付债券
				"LTPayable":          "double", //长期应付款
				"specificPayables":   "double", //专项应付款
				"estimatedLiab":      "double", //预计负债
				"deferTaxLiab":       "double", //递延所得税负债
				"othNCL":             "double", //其他非流动负债
				"TNCL":               "double", //非流动负债合计
				"TLiab":              "double", //负债合计
				"paidInCapital":      "double", //实收资本(或股本)
				"capitalReser":       "double", //资本公积
				"treasuryShare":      "double", //减:库存股
				"specialReser":       "double", //专项储备
				"surplusReser":       "double", //盈余公积
				"ordinRiskReser":     "double", //一般风险准备
				"retainedEarnings":   "double", //未分配利润
				"forexDiffer":        "double", //外币报表折算差额
				"TEquityAttrP":       "double", //归属于母公司所有者权益合计
				"minorityInt":        "double", //少数股东权益
				"TShEquity":          "double", //所有者权益合计
				"TLiabEquity":        "double", //负债和所有者权益总计
			},
			"indices": map[string]string{},
		},
	}
}
