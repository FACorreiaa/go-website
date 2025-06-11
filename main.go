package main

import (
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/FACorreiaa/go-website/ui/pages"
	"github.com/a-h/templ"
	"github.com/go-pdf/fpdf"
)

func main() {
	// Create dist directory
	distPath := "dist"
	os.RemoveAll(distPath) // Clean existing dist
	os.MkdirAll(distPath, 0755)

	// Generate static pages
	generateStaticPage(distPath, "index.html", pages.Landing())
	generateStaticPage(distPath, "about/index.html", pages.About())
	generateStaticPage(distPath, "projects/index.html", pages.Projects())

	// Copy assets to dist
	copyAssets(distPath)

	// Generate CV file
	generateCVFile(distPath)
}

func generateStaticPage(distPath, pagePath string, component templ.Component) {
	// Create directory if needed
	dir := filepath.Dir(filepath.Join(distPath, pagePath))
	os.MkdirAll(dir, 0755)

	// Create HTML file
	file, err := os.Create(filepath.Join(distPath, pagePath))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Render template to file
	err = component.Render(context.Background(), file)
	if err != nil {
		panic(err)
	}
}


func copyAssets(distPath string) {
	assetsDir := filepath.Join(distPath, "assets")
	os.MkdirAll(assetsDir, 0755)

	// Copy CSS files
	copyFile("assets/css/output.css", filepath.Join(assetsDir, "css", "output.css"))
}

func copyFile(src, dst string) {
	// Create directory if needed
	dir := filepath.Dir(dst)
	os.MkdirAll(dir, 0755)

	// Open source file
	srcFile, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()

	// Create destination file
	dstFile, err := os.Create(dst)
	if err != nil {
		panic(err)
	}
	defer dstFile.Close()

	// Copy content
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		panic(err)
	}
}

func generateCVFile(distPath string) {
	cvPath := filepath.Join(distPath, "cv.pdf")
	file, err := os.Create(cvPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Generate PDF (reusing the existing CV generation logic)
	generateCV(file)
}

func generateCV(w io.Writer) {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Set font
	pdf.SetFont("Arial", "B", 16)

	// Header
	pdf.Cell(0, 10, "Fernando António Torre Correia")
	pdf.Ln(8)
	pdf.SetFont("Arial", "", 14)
	pdf.Cell(0, 10, "Golang Developer")
	pdf.Ln(10)

	// Contact Info
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(0, 6, "Phone: +351 937042328")
	pdf.Ln(5)
	pdf.Cell(0, 6, "Email: fernandocorreia316@gmail.com")
	pdf.Ln(5)
	pdf.Cell(0, 6, "LinkedIn: https://www.linkedin.com/in/fernando-correia-ab018079/")
	pdf.Ln(5)
	pdf.Cell(0, 6, "GitHub: https://github.com/FACorreiaa")
	pdf.Ln(12)

	// Professional Summary
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 8, "Professional Summary")
	pdf.Ln(8)
	pdf.SetFont("Arial", "", 10)
	summary := "Software Engineer with 7+ years of experience architecting scalable web services using " +
		"Golang and cloud technologies. Proven track record in diverse sectors like finance and " +
		"telecommunications, delivering high-performance solutions in multinational environments. " +
		"Expertise in backend development, distributed systems, and API design (REST, gRPC), with a " +
		"strong focus on observability and optimization."
	pdf.MultiCell(0, 5, summary, "", "", false)
	pdf.Ln(8)

	// Professional Experience
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 8, "Professional Experience")
	pdf.Ln(8)

	// iDWELL
	pdf.SetFont("Arial", "B", 10)
	pdf.Cell(0, 6, "Golang Developer - iDWELL (Aug 2024 - Present)")
	pdf.Ln(6)
	pdf.SetFont("Arial", "", 9)
	pdf.Cell(0, 5, "Vienna, Austria (Remote)")
	pdf.Ln(5)
	experiences := []string{
		"• Developing web services for the finance sector with Go, Postgres, MSSQL, and GraphQL",
		"• Building scalable, robust APIs for financial data processing",
		"• Collaborating with cross-functional teams to enhance system architecture",
	}
	for _, exp := range experiences {
		pdf.Cell(0, 4, exp)
		pdf.Ln(4)
	}
	pdf.Ln(4)

	// GSMK
	pdf.SetFont("Arial", "B", 10)
	pdf.Cell(0, 6, "Software Developer - GSMK (Jan 2023 - Aug 2024)")
	pdf.Ln(6)
	pdf.SetFont("Arial", "", 9)
	pdf.Cell(0, 5, "Berlin, Germany")
	pdf.Ln(5)
	gsmkExperiences := []string{
		"• Introduced Golang for developing CLI tools across multi-technology projects",
		"• Built web version of mobile network tracking application using React and Rust",
		"• Developed gRPC-based microservices with OTLP and Prometheus monitoring",
	}
	for _, exp := range gsmkExperiences {
		pdf.Cell(0, 4, exp)
		pdf.Ln(4)
	}
	pdf.Ln(4)

	// Education
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 8, "Education")
	pdf.Ln(8)
	pdf.SetFont("Arial", "B", 10)
	pdf.Cell(0, 5, "Master's in Computer Science (2018-2020)")
	pdf.Ln(5)
	pdf.SetFont("Arial", "", 9)
	pdf.Cell(0, 4, "Instituto Politécnico do Cávado e Ave - Barcelos, Portugal")
	pdf.Ln(6)

	pdf.SetFont("Arial", "B", 10)
	pdf.Cell(0, 5, "BSc in Computer Science (2014-2017)")
	pdf.Ln(5)
	pdf.SetFont("Arial", "", 9)
	pdf.Cell(0, 4, "Instituto Politécnico do Cávado e Ave - Barcelos, Portugal")
	pdf.Ln(8)

	// Skills
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 8, "Skills")
	pdf.Ln(8)
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(0, 5, "Languages: Golang, C#, TypeScript, SQL, JavaScript")
	pdf.Ln(5)
	pdf.Cell(0, 5, "Technologies: Postgres, MSSQL, MongoDB, Redis, Kafka, GraphQL, Docker")
	pdf.Ln(5)
	pdf.Cell(0, 5, "Frameworks: React, Node.js, React Native, PostgREST, Redux")
	pdf.Ln(5)
	pdf.Cell(0, 5, "Tools: Git, Jira, Docker")

	// Output PDF to writer
	err := pdf.Output(w)
	if err != nil {
		panic(err)
	}
}
