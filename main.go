package main

import (
	"fmt"
	"github.com/gushasha/boss-jobs/models"
	"github.com/gushasha/boss-jobs/parse"
	"strconv"
)

const BASE_URL = "https://www.zhipin.com/c101270100/?"

var jobTypes = []string{
	"java",
	"go",
	//"php",
	//"c++",
	//"c",
	//"c#",
	//"android",
	//"ios",
	//"web前端",
}

func main() {
	for _, languageType := range jobTypes {
		for i := 1; ; i++ {
			url := BASE_URL + "query=" + languageType + "&page=" + strconv.Itoa(i)
			// 通过url获取一个页面信息，解析出job列表和下一页
			jobList, hasNext := parse.GetJobs(url)
			// 数据入库
			saveJob(jobList)
			if !hasNext {
				break
			}
		}
	}
}

func saveJob(jobs []models.BossJobs) {
	for index, job := range jobs {
		createdJob := models.CreateOne(job)
		fmt.Printf("create %d:  %v", index, createdJob)
	}
}
