package parse

import (
	"fmt"
	"github.com/gushasha/boss-jobs/conf"

	"github.com/PuerkitoBio/goquery"
	"github.com/gpmgo/gopm/modules/log"

	"github.com/gushasha/boss-jobs/models"

	"net/http"
	netUrl "net/url"

	"strings"
	"time"
)



type Page struct {
	Page int
	Url  string
}

func GetJobs(url string) (jobs []models.BossJobs, hasNext bool) {
	fmt.Println("爬取地址：", url)

	u, err := netUrl.Parse(url)
	languageType := u.Query().Get("query")

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", conf.USER_AGENT)
	request.Header.Add("cookie", conf.COOKIE)
	resp, err := client.Do(request)

	if err != nil {
		log.Fatal("抓取页面错误[%s]: %v", url, err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("http状态码异常[%d %s]", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("NewDocumentFromReader Error: %v", err)
	}

	// 判断页面是否出错
	checkPageError(doc)

	// 解析页面内容
	jobs = parseJobs(doc, languageType)

	// 是否含有下一页
	hasNext = hasNextPage(doc)

	// 防止请求过于频繁，IP被禁
	time.Sleep(time.Second * 5)

	return
}

func parseJobs(doc *goquery.Document, JobType string) (jobs []models.BossJobs) {
	doc.Find("#main > .job-box > .job-list > ul > li").Each(func(i int, s *goquery.Selection) {
		jobId, ok := s.Find(".info-primary > .name > a").Eq(0).Attr("data-jid")
		if !ok {
			log.Fatal("data-jid Error")
		}

		title := s.Find(".info-primary > .name > a > .job-title").Eq(0).Text()
		salaryRange := s.Find(".info-primary > .name > a > .red").Eq(0).Text()

		primaryExt, _ := s.Find(".info-primary > p").Eq(0).Html()
		primaryExtSplits := strings.Split(primaryExt, `<em class="vline"></em>`)
		companyAddress := primaryExtSplits[0]
		workYears := primaryExtSplits[1]
		education := primaryExtSplits[2]

		companyName := s.Find(".info-company > .company-text a").Eq(0).Text()
		companyExt, _ := s.Find(".info-company > .company-text p").Eq(0).Html()
		companyExtSplits := strings.Split(companyExt, `<em class="vline"></em>`)
		var companyLabel, financingStage, companyEmployeesCount string
		if len(companyExtSplits) == 3 {
			companyLabel = companyExtSplits[0]
			financingStage = companyExtSplits[1]
			companyEmployeesCount = companyExtSplits[2]
		} else if len(companyExtSplits) == 2 {
			companyLabel = companyExtSplits[0]
			companyEmployeesCount = companyExtSplits[1]
		}

		job := models.BossJobs{
			Jid:                   jobId,
			JobType:               JobType,
			Title:                 title,
			SalaryRange:           salaryRange,
			CompanyAddress:        companyAddress,
			WorkYears:             workYears,
			Education:             education,
			CompanyName:           companyName,
			CompanyLabel:          companyLabel,
			FinancingStage:        financingStage,
			CompanyEmployeesCount: companyEmployeesCount,
		}
		jobs = append(jobs, job)
	})
	return
}

func hasNextPage(doc *goquery.Document) bool {
	pageList := doc.Find("#main > .job-box > .job-list > .page > a")
	nextPage := pageList.Last()
	// 审查元素，在最后一页中，"下一页图标" class为 "disabled"
	return !nextPage.HasClass("disabled")
}

func checkPageError(doc *goquery.Document) {
	errorContent := doc.Find("#main .error-content").Text()
	if errorContent != "" {
		log.Fatal(errorContent)
	}
}
