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
	getTechRequirements(image)
	// fmt.Print(image)
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
func getTechRequirements(image string) {
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

Context: You are processing the HTML content of a job description page to extract and organize technical requirements into a structured JSON object with two main sections: Programming Languages and Technologies You Need to Know. These sections will contain arrays listing the specific languages, frameworks, and technologies mentioned in the job description.

Task: Analyze the HTML content and extract the following:

Programming Languages: Identify and list all programming languages explicitly mentioned in the job description.
Technologies You Need to Know: Identify and list relevant technologies, frameworks, libraries, tools, or platforms required for the role.
Exemplar Output:

json
Copy
Edit
{
  "Programming Languages": [
    "Python",
    "Java",
    "JavaScript",
    "C++"
  ],
  "Technologies You Need to Know": [
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
Ensure that each programming language or technology appears only once.
Do not include unrelated job qualifications or soft skills.
Prompt: Given the HTML content of a job description page, extract and organize the technical requirements into the
 specified JSON object. Ensure the output includes only the programming languages and technologies mentioned in the
  job listing. Here is the HTML content: %s`, image),
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
