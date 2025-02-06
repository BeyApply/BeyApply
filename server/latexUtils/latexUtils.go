package latexUtils

import "fmt"

type ResumeInfo struct {
	resumeName       string
	name             string
	links            []string
	technicalSkills  []string
	workExperience   []string
	projects         []string
	extracurriculars []string
}

func PrintHi() {
	fmt.Println("hi")
}

func ConvertJSONToLatex(resumeinfo ResumeInfo) {

}

func ConvertLatexToPDF(latexstring string) {

}
