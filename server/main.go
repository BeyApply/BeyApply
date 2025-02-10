package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Example LaTeX content
	latexContent := `
	\documentclass[10pt, letterpaper]{article}

% Packages:
\usepackage[
    ignoreheadfoot, % set margins without considering header and footer
    top=2 cm, % separation between body and page edge from the top
    bottom=2 cm, % separation between body and page edge from the bottom
    left=2 cm, % separation between body and page edge from the left
    right=2 cm, % separation between body and page edge from the right
    footskip=1.0 cm, % separation between body and footer
    % showframe % for debugging 
]{geometry} % for adjusting page geometry
\usepackage{titlesec} % for customizing section titles
\usepackage{tabularx} % for making tables with fixed width columns
\usepackage{array} % tabularx requires this
\usepackage[dvipsnames]{xcolor} % for coloring text
\definecolor{primaryColor}{RGB}{0, 0, 0} % define primary color
\usepackage{enumitem} % for customizing lists
\usepackage{fontawesome5} % for using icons
\usepackage{amsmath} % for math
\usepackage[
    pdftitle={John Doe's CV},
    pdfauthor={John Doe},
    pdfcreator={LaTeX with RenderCV},
    colorlinks=true,
    urlcolor=primaryColor
]{hyperref} % for links, metadata and bookmarks
\usepackage[pscoord]{eso-pic} % for floating text on the page
\usepackage{calc} % for calculating lengths
\usepackage{bookmark} % for bookmarks
\usepackage{lastpage} % for getting the total number of pages
\usepackage{changepage} % for one column entries (adjustwidth environment)
\usepackage{paracol} % for two and three column entries
\usepackage{ifthen} % for conditional statements
\usepackage{needspace} % for avoiding page break right after the section title
\usepackage{iftex} % check if engine is pdflatex, xetex or luatex

% Ensure that generated pdf is machine readable/ATS parsable:
\ifPDFTeX
    \input{glyphtounicode}
    \pdfgentounicode=1
    \usepackage[T1]{fontenc}
    \usepackage[utf8]{inputenc}
    \usepackage{lmodern}
\fi

\usepackage{charter}

% Some settings:
\raggedright
\AtBeginEnvironment{adjustwidth}{\partopsep0pt} % remove space before adjustwidth environment
\pagestyle{empty} % no header or footer
\setcounter{secnumdepth}{0} % no section numbering
\setlength{\parindent}{0pt} % no indentation
\setlength{\topskip}{0pt} % no top skip
\setlength{\columnsep}{0.15cm} % set column separation
\pagenumbering{gobble} % no page numbering

\titleformat{\section}{\needspace{4\baselineskip}\bfseries\large}{}{0pt}{}[\vspace{1pt}\titlerule]

\titlespacing{\section}{
    % left space:
    -1pt
}{
    % top space:
    0.3 cm
}{
    % bottom space:
    0.2 cm
} % section title spacing

\renewcommand\labelitemi{$\vcenter{\hbox{\small$\bullet$}}$} % custom bullet points
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
} % new environment for highlights

\newenvironment{highlightsforbulletentries}{
    \begin{itemize}[
        topsep=0.10 cm,
        parsep=0.10 cm,
        partopsep=0pt,
        itemsep=0pt,
        leftmargin=10pt
    ]
}{
    \end{itemize}
} % new environment for highlights for bullet entries

\newenvironment{onecolentry}{
    \begin{adjustwidth}{
        0 cm + 0.00001 cm
    }{
        0 cm + 0.00001 cm
    }
}{
    \end{adjustwidth}
} % new environment for one column entries

\newenvironment{twocolentry}[2][]{
    \onecolentry
    \def\secondColumn{#2}
    \setcolumnwidth{\fill, 4.5 cm}
    \begin{paracol}{2}
}{
    \switchcolumn \raggedleft \secondColumn
    \end{paracol}
    \endonecolentry
} % new environment for two column entries

\newenvironment{threecolentry}[3][]{
    \onecolentry
    \def\thirdColumn{#3}
    \setcolumnwidth{, \fill, 4.5 cm}
    \begin{paracol}{3}
    {\raggedright #2} \switchcolumn
}{
    \switchcolumn \raggedleft \thirdColumn
    \end{paracol}
    \endonecolentry
} % new environment for three column entries

\newenvironment{header}{
    \setlength{\topsep}{0pt}\par\kern\topsep\centering\linespread{1.5}
}{
    \par\kern\topsep
} % new environment for the header

\newcommand{\placelastupdatedtext}{% \placetextbox{<horizontal pos>}{<vertical pos>}{<stuff>}
  \AddToShipoutPictureFG*{% Add <stuff> to current page foreground
    \put(
        \LenToUnit{\paperwidth-2 cm-0 cm+0.05cm},
        \LenToUnit{\paperheight-1.0 cm}
    ){\vtop{{\null}\makebox[0pt][c]{
        \small\color{gray}\textit{Last updated in September 2024}\hspace{\widthof{Last updated in September 2024}}
    }}}%
  }%
}%

% save the original href command in a new command:
\let\hrefWithoutArrow\href

% new command for external links:

\begin{document}
    \newcommand{\AND}{\unskip
        \cleaders\copy\ANDbox\hskip\wd\ANDbox
        \ignorespaces
    }
    \newsavebox\ANDbox
    \sbox\ANDbox{$|$}

    \begin{header}
        \fontsize{25 pt}{25 pt}\selectfont John Doe

        \vspace{5 pt}

        \normalsize
        
        \mbox{\hrefWithoutArrow{mailto:youremail@yourdomain.com}{youremail@yourdomain.com}}%
        \kern 5.0 pt%
        \AND%
        \kern 5.0 pt%
        \mbox{\hrefWithoutArrow{tel:+90-541-999-99-99}{0541 999 99 99}}%
        \kern 5.0 pt%
        \AND%
        \kern 5.0 pt%
        \mbox{\hrefWithoutArrow{https://yourwebsite.com/}{yourwebsite.com}}%
        \kern 5.0 pt%
        \AND%
        \kern 5.0 pt%
        \mbox{\hrefWithoutArrow{https://linkedin.com/in/yourusername}{linkedin.com/in/yourusername}}%
        \kern 5.0 pt%
    \end{header}

    \section{Education}        
    \begin{twocolentry}{
        Sept 2000 – May 2005
    }
        \textbf{University of Pennsylvania}, BS in Computer Science
    \end{twocolentry}

    \vspace{0.10 cm}
    \begin{onecolentry}
        \begin{highlights}
            \item GPA: 3.9/4.0 (\href{https://example.com}{a link to somewhere})
            \item \textbf{Coursework:} Computer Architecture, Comparison of Learning Algorithms, Computational Theory
        \end{highlights}
    \end{onecolentry}

    \section{Technologies}
    \begin{onecolentry}
        \textbf{Languages:} C++, C, Java, Objective-C, C\#, SQL, JavaScript
    \end{onecolentry}

    \vspace{0.2 cm}

    \begin{onecolentry}
        \textbf{Technologies:} .NET, Microsoft SQL Server, XCode, Interface Builder
    \end{onecolentry}

    \section{Experience}
    \begin{twocolentry}{
        June 2005 – Aug 2007
    }
        \textbf{Software Engineer}, Apple -- Cupertino, CA
    \end{twocolentry}

    \vspace{0.10 cm}
    \begin{onecolentry}
        \begin{highlights}
            \item Reduced time to render user buddy lists by 75\% by implementing a prediction algorithm
            \item Integrated iChat with Spotlight Search by creating a tool to extract metadata from saved chat transcripts and provide metadata to a system-wide search database
            \item Redesigned chat file format and implemented backward compatibility for search
        \end{highlights}
    \end{onecolentry}

    \vspace{0.2 cm}

    \begin{twocolentry}{
        June 2003 – Aug 2003
    }
        \textbf{Software Engineer Intern}, Microsoft -- Redmond, WA
    \end{twocolentry}

    \vspace{0.10 cm}
    \begin{onecolentry}
        \begin{highlights}
            \item Designed a UI for the VS open file switcher (Ctrl-Tab) and extended it to tool windows
            \item Created a service to provide gradient across VS and VS add-ins, optimizing its performance via caching
            \item Built an app to compute the similarity of all methods in a codebase, reducing the time from $\mathcal{O}(n^2)$ to $\mathcal{O}(n \log n)$
            \item Created a test case generation tool that creates random XML docs from XML Schema
            \item Automated the extraction and processing of large datasets from legacy systems using SQL and Perl scripts
        \end{highlights}
    \end{onecolentry}

    \section{Projects}
    \begin{twocolentry}{
        \href{https://github.com/sinaatalay/rendercv}{github.com/name/repo}
    }
        \textbf{Multi-User Drawing Tool}
    \end{twocolentry}

    \vspace{0.10 cm}
    \begin{onecolentry}
        \begin{highlights}
            \item Developed an electronic classroom where multiple users can simultaneously view and draw on a "chalkboard" with each person's edits synchronized
            \item Tools Used: C++, MFC
        \end{highlights}
    \end{onecolentry}

    \vspace{0.2 cm}

    \begin{twocolentry}{
        \href{https://github.com/sinaatalay/rendercv}{github.com/name/repo}
    }
        \textbf{Synchronized Desktop Calendar}
    \end{twocolentry}

    \vspace{0.10 cm}
    \begin{onecolentry}
        \begin{highlights}
            \item Developed a desktop calendar with globally shared and synchronized calendars, allowing users to schedule meetings with other users
            \item Tools Used: C\#, .NET, SQL, XML
        \end{highlights}
    \end{onecolentry}

    \vspace{0.2 cm}

    \begin{twocolentry}{
        2002
    }
        \textbf{Custom Operating System}
    \end{twocolentry}

    \vspace{0.10 cm}
    \begin{onecolentry}
        \begin{highlights}
            \item Built a UNIX-style OS with a scheduler, file system, text editor, and calculator
            \item Tools Used: C
        \end{highlights}
    \end{onecolentry}

    \section{Organizations}
    \begin{twocolentry}{
        \href{https://github.com/sinaatalay/rendercv}{github.com/name/repo}
    }
        \textbf{Multi-User Drawing Tool}
    \end{twocolentry}

    \vspace{0.10 cm}
    \begin{onecolentry}
        \begin{highlights}
            \item Developed an electronic classroom where multiple users can simultaneously view and draw on a "chalkboard" with each person's edits synchronized
            \item Tools Used: C++, MFC
        \end{highlights}
    \end{onecolentry}
\end{document}
	`

	// Define file paths
	texFile := "output.tex"
	pdfFile := "output.pdf"

	// Write LaTeX content to a file
	err := os.WriteFile(texFile, []byte(latexContent), 0644)
	if err != nil {
		fmt.Printf("Error writing .tex file: %v\n", err)
		os.Exit(1)
	}

	// Run pdflatex to generate PDF
	cmd := exec.Command("pdflatex", texFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error running pdflatex: %v\n", err)
		os.Exit(1)
	}

	// Check if the PDF was created
	if _, err := os.Stat(pdfFile); os.IsNotExist(err) {
		fmt.Println("PDF generation failed.")
		os.Exit(1)
	}

	fmt.Printf("Successfully generated PDF: %s\n", pdfFile)
}
