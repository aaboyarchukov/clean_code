package lesson6

import "time"

type Report struct {
	ID   int
	Name string
}

// old name: PrintReportDocument
// new name: Print
func (report *Report) Print(computerHost string) {}

// old name: SendReportDocumentByMail
// new name: SendByMail
func (report *Report) SendByMail(destEmail string) {}

// old name: SignReportDocument
// new name: Sign
func (report *Report) Sign() {}

type Employee struct {
	EmployeeID int
	Name       string
	Post       string
}

// old name: CountSalaryPerQuarterAndFormingReportInExcel
func (employee *Employee) CountSalaryPerQuarterAndFormingReportInExcel(salaries []float64) Report {}

// new name: CountSalary, FormingReport, ConvertToExcel, PrepareSalaryReportPerPeriod
func (employee *Employee) CountSalary(salaries []float64) map[string]float64   {}
func (employee *Employee) FormingReport(salaryInfo map[string]float64) Report  {}
func (employee *Employee) PrepareSalaryReportByPeriod(period time.Time) Report {}
func (employee *Employee) ConvertToExcel(document Report) file.Excel           {}

type OS struct {
	OsID    int
	Name    string
	Version string
}

// old name: Version
// new name: GetVersion
func (os *OS) GetVersion() {}

// old name: Info
// new name: GetOSParametrs
func (os *OS) GetOSParametrs() {}

// old name: ActiveThreads
// new name: GetActiveThreads
func (os *OS) GetActiveThreads() {}

// old name: Program
// new name: ExecuteProgram
func (os *OS) ExecuteProgram(programID int) {}

// old name: Starting
// new name: Start
func (os *OS) Start() {}
