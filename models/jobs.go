package models

import "log"

type BossJobs struct {
	Model

	Jid                   string `json:"jid"`
	JobType               string `json:"job_type"`
	Title                 string `json:"title"`
	SalaryRange           string `json:"salary_range"`
	WorkYears             string `json:"work_years"`
	Education             string `json:"education"`
	CompanyName           string `json:"company_name"`
	CompanyAddress        string `json:"company_address"`
	CompanyLabel          string `json:"company_label"`
	CompanyEmployeesCount string `json:"company_employees_count"`
	FinancingStage        string `json:"financing_stage"`
}

func CreateOne(job BossJobs) BossJobs {
	if DB.Set("gorm:insert_modifier", "IGNORE").Create(&job).Error != nil {
		log.Fatalf("Should ignore duplicate user insert by insert modifier:IGNORE ")
	}
	return job
}
