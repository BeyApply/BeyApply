package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

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

const latexTemplate = `\documentclass[10pt, letterpaper]{article}

% Packages:
\usepackage[
    ignoreheadfoot,
    top=2 cm,
    bottom=2 cm,
    left=2 cm,
    right=2 cm,
    footskip=1.0 cm,
]{geometry}
\usepackage{titlesec}
\usepackage{tabularx}
\usepackage{array}
\usepackage[dvipsnames]{xcolor}
\definecolor{primaryColor}{RGB}{0, 0, 0}
\usepackage{enumitem}
\usepackage{fontawesome5}
\usepackage{amsmath}
\usepackage[
    pdftitle={ {{.Name}} 's CV},
    pdfauthor={ {{.Name}} },
    pdfcreator={LaTeX with RenderCV},
    colorlinks=true,
    urlcolor=primaryColor
]{hyperref}
\usepackage[pscoord]{eso-pic}
\usepackage{calc}
\usepackage{bookmark}
\usepackage{lastpage}
\usepackage{changepage}
\usepackage{paracol}
\usepackage{ifthen}
\usepackage{needspace}
\usepackage{iftex}

\ifPDFTeX
    \input{glyphtounicode}
    \pdfgentounicode=1
    \usepackage[T1]{fontenc}
    \usepackage[utf8]{inputenc}
    \usepackage{lmodern}
\fi

\usepackage{mathptmx}

\raggedright
\AtBeginEnvironment{adjustwidth}{\partopsep0pt}
\pagestyle{empty}
\setcounter{secnumdepth}{0}
\setlength{\parindent}{0pt}
\setlength{\topskip}{0pt}
\setlength{\columnsep}{0.15cm}
\pagenumbering{gobble}

\titleformat{\section}{\needspace{4\baselineskip}\bfseries\large}{}{0pt}{}[\vspace{1pt}\titlerule]
\titlespacing{\section}{-1pt}{0.3 cm}{0.2 cm}


\renewcommand\labelitemi{$\vcenter{\hbox{\small$\bullet$}}$}
\newenvironment{highlights}{
    \begin{itemize}[
        topsep=0.10 cm,
        parsep=0.10 cm,
        partopsep=0pt,
        itemsep=0pt,
        leftmargin=0 cm + 10pt
    ]
}{
    \end{itemize}
}

\newenvironment{onecolentry}{
    \begin{adjustwidth}{0 cm + 0.00001 cm}{0 cm + 0.00001 cm}
}{
    \end{adjustwidth}
}

\newenvironment{twocolentry}[2][]{
    \onecolentry
    \def\secondColumn{#2}
    \setcolumnwidth{\fill, 4.5 cm}
    \begin{paracol}{2}
}{
    \switchcolumn \raggedleft \secondColumn
    \end{paracol}
    \endonecolentry
}

\begin{document}
    \newcommand{\AND}{\quad | \quad}
    \newsavebox\ANDbox
    \sbox\ANDbox{$|$}

    \begin{center}
    \fontsize{25 pt}{25 pt}\selectfont {{.Name}}

    \vspace{5 pt}
    \normalsize
    \begin{tabular}{c}  % This ensures links are centered and on the same line
        {{ range $index, $link := .Links }}
            {{if $index}} \quad | \quad {{end}}
            \href{ {{$link}} }{ {{$link}} }
        {{end}}
    \end{tabular}
\end{center}

    \section{Education}
    {{range $degree, $details := .Education}}
        \begin{twocolentry}{ {{$details.Date}} }
            \textbf{ {{$details.Institution}} }, {{$degree}}
        \end{twocolentry}

        \vspace{0.10 cm}
        \begin{onecolentry}
            \begin{highlights}
                \item GPA: {{$details.GPA}}
            \end{highlights}
        \end{onecolentry}
    {{end}}

    \section{Technologies}
    \begin{onecolentry}
        \begin{highlights}
            \item Languages/Frameworks: {{join .TechnicalSkills.LanguagesFrameworks ", "}}
        \end{highlights}
    \end{onecolentry}
    \vspace{0.2 cm}
    \begin{onecolentry}
        \begin{highlights}
            \item Technologies: {{join .TechnicalSkills.Technologies ", "}}
        \end{highlights}
    \end{onecolentry}

    \section{Experience}
    {{range $company, $job := .Experience}}
        \begin{twocolentry}{ {{$job.Date}} }
            {{$job.Role}}, {{$company}}
        \end{twocolentry}

        \vspace{0.10 cm}
        \begin{onecolentry}
            \begin{highlights}
                {{range $job.Responsibilities}}
                    \item {{.}}
                {{end}}
            \end{highlights}
        \end{onecolentry}
        \vspace{0.4 cm}
    {{end}}

    \section{Projects}
    {{range $project, $details := .Projects}}
        \begin{twocolentry}{}
            { {{$project}} }
        \end{twocolentry}

        \vspace{0.10 cm}
        \begin{onecolentry}
            \begin{highlights}
                {{range $details}}
                    \item {{.}}
                {{end}}
            \end{highlights}
        \end{onecolentry}
        \vspace{0.4 cm}
    {{end}}

    \section{Organizations}
    {{range $org, $details := .Clubs}}
        \begin{twocolentry}{}
            \textbf{ {{$org}} }
        \end{twocolentry}
        \vspace{0.1 cm}
        \begin{onecolentry}
            \begin{highlights}
                {{range $details}}
                    \item {{.}}
                {{end}}
            \end{highlights}
        \end{onecolentry}
        \vspace{0.4 cm}
    {{end}}
\end{document}`

func generateResume(jsonString string) error {
	// Parse JSON string into struct
	var resume ResumeInfo
	if err := json.Unmarshal([]byte(jsonString), &resume); err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}

	// Create template with custom function
	tmpl := template.New("resume")
	tmpl.Funcs(template.FuncMap{
		"join": strings.Join,
	})

	// Parse template
	tmpl, err := tmpl.Parse(latexTemplate)
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	// Create output file
	texFile := "output.tex"
	f, err := os.Create(texFile)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer f.Close()

	// Execute template
	if err := tmpl.Execute(f, resume); err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}

	// Generate PDF
	cmd := exec.Command("pdflatex", texFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running pdflatex: %v", err)
	}

	return nil
}

func main() {
	jsonString := `{
		"Name": "Carolina Campos",
		"Links": [
			"github.com/Carol0427",
			"Carolmeister@gmail.com",
			"772-701-2041"
		],
		"Education": {
			"Bachelor of Science in Computer Science": {
				"Institution": "Florida Atlantic University",
				"GPA": "3.4",
				"Date": "December 2024"
			}
		},
		"Technical Skills": {
			"Languages/Frameworks": [
				"Java",
				"Python",
				"JavaScript",
				"C++",
				"PHP",
				"Spring",
				"Node",
				"React",
				"MySQL",
				"HTML5"
			],
			"Technologies": [
				"Git",
				"Linux",
				"AWS",
				"SaaS microservices",
				"Docker/Kubernetes"
			]
		},
		"Experience": {
			"Entrust": {
				"Date": "Sep 2023 – Sep 2024",
				"Role": "Software Engineering Intern",
				"Responsibilities": [
					"Integrated client and company products for seamless cryptographic solutions, boosting deal closure rates by 23 percent and facilitating sales worth up to 100,000",
					"Developed tools in Python to automate integration, reducing manual overhead and cutting integration errors by 12 percent",
					"Authored technical guides, saving client engineers hundreds of hours by documenting common errors, which accelerated client's project timelines"
				]
			},
			"Microsoft": {
				"Date": "Oct 2019 – Oct 2022",
				"Role": "Software Engineering Intern",
				"Responsibilities": [
					"Designed a UI for the VS open file switcher (Ctrl-Tab) and extended it to tool windows",
					"Created a service to provide gradient across VS and VS add-ins, optimizing its performance via caching"
				]
			}
		},
		"Projects": {
			"Multi-User Drawing Tool": [
				"Developed full-stack IoT application with client specification for real-time environmental monitoring using a network of sensors for data ingestion and visualization",
				"Technologies used: React, Grafana, MapBox API for front-end, and backend with Java, Spring, AWS IoT Core, Kafka, Telegraf for depth mathematics, InfluxDB for database, deployed on AWS EC2 instances"
			],
			"Synchronized Desktop Calendar": [
				"Developed a desktop calendar with globally shared and synchronized calendars, allowing users to schedule meetings with other users"
			]
		},
		"Clubs/Organizations": {
			"Girls Who Code - President": [
				"Delegated and oversaw tasks to a cross-functional 8-member Executive Board",
				"Increased membership from 15 to 52, a 247% growth in less than a year",
				"Successfully raised funds through car washes and donors which increased our budget by 55%"
			],
			"Medium- Technical Writing Blog": [
				"Launched a Medium blog focused on explaining complex software engineering concepts to beginners, covering topics like deploying full-stack Next.js applications on EC2 instances."
			]
		}
	}`

	if err := generateResume(jsonString); err != nil {
		fmt.Printf("Error generating resume: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully generated resume PDF!")
}
