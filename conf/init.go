package conf

const (
	// HTTP请求头，简单起见，从自己浏览器copy过来
	USER_AGENT = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36"
	COOKIE     = "__c=1559283718; __g=-; __l=l=%2Fwww.zhipin.com%2Fgongsi%2F4e86008c2bd050e00HN82t64.html&r=https%3A%2F%2Fwww.google.com%2F; lastCity=101270100; __a=54573481.1556194697.1556257996.1559283718.262.3.190.192; Hm_lvt_194df3105ad7148dcf2b98a91b5e727a=1561438396,1561446563,1561459913,1561541594; Hm_lpvt_194df3105ad7148dcf2b98a91b5e727a=1561541594"

	// Mysql配置
	DB_USERNAME     string = "root"
	DB_PASSWORD     string = "123456"
	DB_NAME         string = "spiders"
	DB_TABLE_PREFIX string = "spiders_"
)
