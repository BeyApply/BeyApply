// package aiUtils

package main

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	resumeText := "Carolina Campos  |   561 - 929 - 1237 |   ccampos2021@fau.edu   |   github.com/Carol0427  Education  Florida Atlantic University  Bachelor of   Science in   Computer Science (GPA: 3.4)   December   202 4  Relevant Certification s  AWS   Solutions   Architect Associates ,   Codepath Intermediat e   Full Stack   Development   for DevOps  Technical Skills  Languages/Frameworks:   Java ,   Python,   JavaScr ipt,   C++,   PHP,   Spri ng,   Node,   React,   MySQL,   HTML5  Technologies:   Git,   Linux,   AWS   SaaS   microservices , Docker/Kubernetes  Experience  Entrust   -   Software Engineering Intern   Sept 2023   –   Sept 2024  •   Integrat ed   client   and company produ cts   for seamless   cryptographic solutions, boosting deal closure  rates by 23% and facilitating sales worth up to $100,000.  •   Developed   tools in Pytho n   to automate   integration   in   reducing manual overhead and cutting  integration errors by 12% , while   doing code review.  •   A uthor ed   technical guides ,   that saved client engineers hundreds of hours by documenting common  errors, which accelerated client’s project timelines.  Projects  FloodSense  •   D evelop ed   full - stack   IoT application   with client specification   for real - time environmental monitoring  using a network of sensors   for data ingestion   an d   visualiza tion .  •   Built   with   React , Grafana, MapBox API   for front - end, and backend with   Java, Spring ,   AWS IoT Core,  Kafka, Telegraf   for depth mathem atics , InfluxDB   for database , and deployed on AWS EC2 instances .  AnonyVent (Waffle Hacks 2 nd   Place Winner)  •   A nonymously allows users to   record   and   view   venting sessions   with   an   AI generated transcription.  •   Built full - stack web app with   React , Node, JavaScript,   MongoDB, AWS S3, ChatGPT API, Assembly AI  API, Netlify, and Heroku.  LeafSafe  •   D etects   poisonous /inedible   plants   and p rovides customized advice on actions to take .  •   F ull Stack web app   developed with Next.js ,   utilizing the ChatGPT API for AI - driven plant identification  and advice.   Deployed on EC2.  Clubs/Organizations  Girls Who Code   -   President  •   Delegated and oversaw tasks   to a   cross - functional   8 - member Executive Board  •   Hosted various events and technical   w orkshop s in collaboration with local companies  •   Increased membership   from 15 to 52, a   247% growth in less than a year  •   Successfully raised funds through car washes and donors which increased our budget by 55%  Society of Hispanic Professional Engineers   –   Internal Affairs Committee  •   Oversee member   engagement and retention  •   Assist in resolving internal conflicts"
	// Create a new OpenAI client
	client := openai.NewClient("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") // Replace with your actual OpenAI API key

	// Define the chat completion request
	resp, err := client.CreateChatCompletion(

		context.Background(),
		openai.ChatCompletionRequest{
			Model: "gpt-4", // Use "gpt-4" or the correct model name
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleUser,
					Content: fmt.Sprintf(`Persona: You are an AI Resume Parser named ResumeMapper, designed to extract and organize resume content into structured data.

Context: You are processing a resume to extract text and organize it into a structured JSON object with five main sections: Education, Technical Skills, Experience, Projects, and Clubs/Organizations. Each section will have sub-sections as keys, and the values will be an array of relevant details. Additionally, you will extract the name and an array of links (if any) from the resume.

Task: Analyze the resume text and extract the following:

Name: Extract the name of the individual and store it as "name": "Full Name".

Links: Extract any links (e.g., LinkedIn, GitHub, portfolio) and store them in an array named "links".

Education: Create a JSON object for Education where the key is the degree name (e.g., "B.S. in Computer Science"), and the value is a JSON object with the following keys:

Institution: The name of the institution.

"GPA": The GPA (if available).

"Date": The date or duration of the degree.

Technical Skills: Create a JSON object for Technical Skills with the following keys:

"Languages/Frameworks": An array of programming languages and frameworks.

"Technologies": An array of technologies or tools.

Experience: Create a JSON object for Experience where the key is the company name (e.g., "ABC Corp"), and the value is a JSON object with the following keys:

"Date": The date or duration of the experience.

"Role": The role or position held.

"Responsibilities": An array of responsibilities or achievements.

Projects: Create a JSON object for Projects where the key is the project name (e.g., "AI Chatbot"), and the value is an array of project details, including descriptions and technologies used.

Clubs/Organizations: Create a JSON object for Clubs/Organizations where the key is the club/organization name and role (e.g., "Tech Club - President"), and the value is an array of achievements or responsibilities.

Exemplar:

json
Copy
{
  "name": "John Doe",
  "links": [
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
Format:

Input: [Resume text]

Output: A JSON object with the structure described above.

Tone: Professional. Your role is to extract and organize data accurately without adding or omitting any details.

Prompt: Given the resume text below, extract and organize the content into the specified JSON object. Ensure the output includes the keys name, links, Education, Technical Skills, Experience, Projects, and Clubs/Organizations.

[Resume text]: %s`, resumeText),
				},
			},
		},
	)

	// Handle errors
	if err != nil {
		fmt.Printf("Error creating chat completion: %v\n", err)
		return
	}

	// Print the response
	fmt.Println(resp.Choices[0].Message.Content)
}

// func ConvertResumeToJSON() {

// }
