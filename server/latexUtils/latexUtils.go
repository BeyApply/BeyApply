package latexUtils

type ResumeInfo struct {
	Name            string              `json:"name"`
	Links           []string            `json:"links"`
	Education       map[string]Degree   `json:"Education"`
	TechnicalSkills TechnicalSkills     `json:"Technical Skills"`
	Experience      map[string]Job      `json:"Experience"`
	Projects        map[string][]string `json:"Projects"`
	Clubs           map[string][]string `json:"Clubs/Organizations"`
}

type Degree struct {
	Institution string `json:"Institution"`
	GPA         string `json:"GPA"`
	Date        string `json:"Date"`
}

type TechnicalSkills struct {
	LanguagesFrameworks []string `json:"Languages/Frameworks"`
	Technologies        []string `json:"Technologies"`
}

type Job struct {
	Date             string   `json:"Date"`
	Role             string   `json:"Role"`
	Responsibilities []string `json:"Responsibilities"`
}

func ConvertStructToLatex(resumeinfo ResumeInfo) {

}

func ConvertLatexToPDF(latexstring string) {

}
