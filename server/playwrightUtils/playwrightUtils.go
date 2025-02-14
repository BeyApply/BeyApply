// package playwrightutils
package main

import (
	"fmt"
	"log"

	"os"

	// "github.com/sashabaranov/go-openai"

	"context"

	"github.com/joho/godotenv"
	"github.com/playwright-community/playwright-go"
	"github.com/sashabaranov/go-openai"
)

func main() {
	image := getJobDetails("https://wd1.myworkdaysite.com/en-US/recruiting/snapchat/snap/job/Los-Angeles-California/Software-Engineer--Android--1--Years-of-Experience_Q125SWEA1-1?source=LinkedIn")
	requirements := getTechRequirements(image)
	ogResume := `{
  "Name": "Carolina Campos",
  "Links": [
    "github.com/Carol0427"
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
      "Date": "Sept 2023 – Sept 2024",
      "Role": "Software Engineering Intern",
      "Responsibilities": [
        "Integrated client and company products for seamless cryptographic solutions, boosting deal closure rates by 23% and facilitating sales worth up to $100,000",
        "Developed tools in Python to automate integration, reducing manual overhead and cutting integration errors by 12%",
        "Authored technical guides, saving client engineers hundreds of hours by documenting common errors, which accelerated client’s project timelines"
      ]
    }
  },
  "Projects": {
    "FloodSense": [
      "Developed full-stack IoT application with client specification for real-time environmental monitoring using a network of sensors for data ingestion and visualization",
      "Technologies used: React, Grafana, MapBox API for front-end, and backend with Java, Spring, AWS IoT Core, Kafka, Telegraf for depth mathematics, InfluxDB for database, deployed on AWS EC2 instances"
    ],
    "AnonyVent (Waffle Hacks 2nd Place Winner)": [
      "Allows users to anonymously record and view venting sessions with an AI-generated transcription",
      "Built full-stack web app with React, Node, JavaScript, MongoDB, AWS S3, ChatGPT API, Assembly AI API, Netlify, and Heroku"
    ],
    "LeafSafe": [
      "Detects poisonous/inedible plants and provides customized advice on actions to take",
      "Full Stack web app developed with Next.js, utilizing the ChatGPT API for AI-driven plant identification and advice, Deployed on EC2"
    ]
  },
  "Clubs/Organizations": {
    "Girls Who Code - President": [
      "Delegated and oversaw tasks to a cross-functional 8-member Executive Board",
      "Hosted various events and technical workshops in collaboration with local companies",
      "Increased membership from 15 to 52, a 247 percent growth in less than a year",
      "Successfully raised funds through car washes and donors which increased our budget by 55 percent"     
    ],
    "Society of Hispanic Professional Engineers – Internal Affairs Committee": [
      "Oversee member engagement and retention",
      "Assist in resolving internal conflicts"
    ]
  }
}`
	updatedResume := updateJSON(ogResume, requirements)
	fmt.Println(updatedResume)
}

func getJobDetails(link string) string {
	// Launch Playwright and create a browser instance
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start Playwright: %v", err)
	}
	defer pw.Stop()

	// Launch a Chromium browser (you can also use firefox or webkit)
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	defer browser.Close()

	// Create a new page
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	// Navigate to a URL
	_, err = page.Goto(link, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle, // Wait until network is idle
	})
	if err != nil {
		log.Fatalf("could not navigate: %v", err)
	}

	// Retrieve the full HTML content of the page
	// html, err := page.Content()
	// if err != nil {
	// 	log.Fatalf("could not get page content: %v", err)
	// }
	// Extract only the text from the body tag
	bodyText, err := page.Locator("body").InnerText()

	if err != nil {
		log.Fatalf("could not extract text from body: %v", err)
	}

	// Print the text content of the body
	// fmt.Println("Body text:", bodyText)
	// Print the HTML
	// fmt.Println(html)
	// doc := soup.HTMLParse(html)
	// Extract the text content, excluding the image tags
	// text := doc.Find("body").Text()

	// fmt.Println(text) // Output: "Hello, world! Some text."
	return bodyText
	// // Navigate to the URL and wait until the network is idle
	// if _, err := page.Goto(url, playwright.PageGotoOptions{
	// 	WaitUntil: playwright.WaitUntilStateNetworkidle, // Wait until the network is idle
	// }); err != nil {
	// 	log.Fatalf("❌ Could not navigate to the URL: %v", err)
	// }

	// // Save the page as a PDF
	// // pdfPath := "local-report.pdf"
	// // if _, err := page.PDF(playwright.PagePdfOptions{
	// // 	Path:   &pdfPath,
	// // 	Format: playwright.String("A4"), // Standard format
	// // }); err != nil {
	// // 	log.Fatalf("❌ Could not save PDF: %v", err)
	// // }

	// // Take a screenshot and save it to a file
	// screenshotPath := "screenshot.png"
	// if _, err := page.Screenshot(playwright.PageScreenshotOptions{
	// 	Path:     playwright.String(screenshotPath), // Path to save the screenshot
	// 	FullPage: playwright.Bool(true),             // Capture the entire page
	// }); err != nil {
	// 	log.Fatalf("❌ Could not take screenshot: %v", err)
	// }

	// // fmt.Printf("✅ PDF saved successfully as %s\n", pdfPath)
	// imageBytes, err := os.ReadFile("screenshot.png")
	// if err != nil {
	// 	return ""
	// }
	// return base64.StdEncoding.EncodeToString(imageBytes)
}
func getTechRequirements(image string) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	apiKey := os.Getenv("OPENAI_KEY")
	// fmt.Println(apiKey)
	// Create a new OpenAI client
	client := openai.NewClient(apiKey)
	// Define the chat completion request
	// Call OpenAI API (Example: GPT-4 Vision)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: "gpt-4", // or another multimodal model that supports image input
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleUser,
					Content: fmt.Sprintf(`Persona: You are a text-based job requirements processor.

Context: You are processing the HTML content of a job description page to extract and organize technical requirements into a structured JSON object with two main sections: Languages/Frameworks and Technologies You Need to Know. These sections will contain arrays listing the specific languages, frameworks, and technologies mentioned in the job description.

Task: Analyze the HTML content and extract the following:

Languages/Frameworks: Identify and list all programming languages and frameworks explicitly mentioned in the job description.
Technologies You Need to Know: Identify and list relevant technologies, libraries, tools, or platforms required for the role.
Exemplar Output:
{
  "Languages/Frameworks": [
    "Python",
    "Java",
    "JavaScript",
    "C++"
  ],
  "Technologies": [
    "AWS",
    "Docker",
    "Kubernetes",
    "React",
    "Node.js",
    "PostgreSQL",
    "Git"
  ]
}
Format:

Input: Raw HTML content of a job description page.
Output: A structured JSON object with the specified sections.
Tone:

Maintain accuracy in extracting only the technical requirements.
Ensure that each programming language/framework or technology appears only once.
Do not include unrelated job qualifications or soft skills.
Prompt: Given the HTML content of a job description page, extract and organize the technical requirements into the
 specified JSON object. Ensure the output includes only the programming languages/frameworks and technologies mentioned in the
  job listing. Here is the HTML content: %s`, image),
				},
			},
		},
	)

	// Handle errors
	if err != nil {
		fmt.Printf("Error creating chat completion: %v\n", err)
		return ""
	}

	// Print the response
	fmt.Println(resp.Choices[0].Message.Content)
	return resp.Choices[0].Message.Content
}

func updateJSON(ogResume string, requirements string) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	apiKey := os.Getenv("OPENAI_KEY")
	// fmt.Println(apiKey)
	// Create a new OpenAI client
	client := openai.NewClient(apiKey)
	// Define the chat completion request
	// Call OpenAI API (Example: GPT-4 Vision)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: "gpt-4", // or another multimodal model that supports image input
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleUser,
					Content: fmt.Sprintf(`Task:
Take in two JSON objects:

Job Requirements JSON - Contains the required Programming Languages/Frameworks and Technologies for the job.
Candidate Resume JSON - Contains details about the candidate's experience, skills, education, projects, and organizations.
Your goal is to update the resume to better align with the job requirements by:

Ensuring that matching skills (programming languages/frameworks and technologies) from the job description are emphasized in the resume.
If a skill exists in the candidate’s experience but is not explicitly listed under technical skills, add it to the Technical Skills section.
If a skill is missing but is reasonably inferred from the candidate’s work (e.g., if the candidate has experience with AWS EC2 but "AWS" is not in their Technical Skills, add "AWS"), add it.
If the candidate does not have experience with a required skill, you can fabricate information and subtly adjust descriptions to highlight any related or transferable experience.
Input JSON Structure:
Job Requirements JSON
{
  "Languages/Frameworks": [
    "Python",
    "Java",
    "JavaScript",
    "C++"
  ],
  "Technologies": [
    "AWS",
    "Docker",
    "Kubernetes",
    "React",
    "Node.js",
    "PostgreSQL",
    "Git"
  ]
}
Here is the Job Requirements JSON: %s
Candidate Resume JSON
{
  "Name": "John Doe",
  "Links": [
    "https://linkedin.com/in/johndoe",
    "https://github.com/johndoe"
  ],
  "Education": {
    "B.S. in Computer Science": {
      "Institution": "University of XYZ",
      "GPA": "3.8",
      "Date": "Aug 2018 - May 2022"
    }
  },
  "Technical Skills": {
    "Languages/Frameworks": [
      "Python",
      "JavaScript",
      "React"
    ],
    "Technologies": [
      "AWS",
      "Docker",
      "Git",
      "VS Code"
    ]
  },
  "Experience": {
    "ABC Corp": {
      "Date": "Jun 2021 - Aug 2021",
      "Role": "Software Engineer Intern",
      "Responsibilities": [
        "Developed a REST API for internal tools",
        "Collaborated with cross-functional teams"
      ]
    }
  },
  "Projects": {
    "AI Chatbot": [
      "Built a chatbot using Python and NLP",
      "Technologies used: Python, NLTK, Flask",
      "Deployed on AWS EC2"
    ]
  },
  "Clubs/Organizations": {
    "Tech Club - President": [
      "Organized hackathons and workshops",
      "Increased membership by 50 percent",
      "Raised funds for club activities"
    ]
  }
}
Here is the original resume: %s
Expected Output:
An updated resume JSON that emphasizes the required programming languages/frameworks and technologies wherever applicable and replaces every percent sign with the english word percent:
{
  "Name": "John Doe",
  "Links": [
    "https://linkedin.com/in/johndoe",
    "https://github.com/johndoe"
  ],
  "Education": {
    "B.S. in Computer Science": {
      "Institution": "University of XYZ",
      "GPA": "3.8",
      "Date": "Aug 2018 - May 2022"
    }
  },
  "Technical Skills": {
    "Languages/Frameworks": [
      "Python",
      "Java",
      "JavaScript",
      "C++",
	  "React"
    ],
    "Technologies": [
      "AWS",
      "Docker",
      "Kubernetes",
      "Node.js",
      "PostgreSQL",
      "Git"
    ]
  },
  "Experience": {
    "ABC Corp": {
      "Date": "Jun 2021 - Aug 2021",
      "Role": "Software Engineer Intern",
      "Responsibilities": [
        "Developed a REST API for internal tools",
        "Collaborated with cross-functional teams"
      ]
    }
  },
  "Projects": {
    "AI Chatbot": [
      "Built a chatbot using Python and NLP",
      "Technologies used: Python, NLTK, Flask",
      "Deployed on AWS EC2"
    ]
  },
  "Clubs/Organizations": {
    "Tech Club - President": [
      "Organized hackathons and workshops",
      "Increased membership by 50 percent",
      "Raised funds for club activities"
    ]
  }
`, requirements, ogResume),
				},
			},
		},
	)

	// Handle errors
	if err != nil {
		fmt.Printf("Error creating chat completion: %v\n", err)
		return ""
	}

	// Print the response
	fmt.Println(resp.Choices[0].Message.Content)
	return ""
}
