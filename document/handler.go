package document

import (
	"io"
	"log"
	"net/http"
	"path/filepath"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		return
	}

	// Parse multipart form with a max memory of 100MB
	err := r.ParseMultipartForm(100 << 20)
	if err != nil {
		log.Println("Error parsing form:", err)
		http.Error(w, "Unable to parse form.", http.StatusBadRequest)
		return
	}

	// Retrieve the file from form data
	file, handler, err := r.FormFile("docxFile")
	if err != nil {
		log.Println("Error retrieving the file:", err)
		http.Error(w, "Error retrieving the file.", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Check file extension
	ext := filepath.Ext(handler.Filename)
	if ext != ".docx" {
		log.Println("Invalid file type:", ext)
		http.Error(w, "Only .docx files are allowed.", http.StatusBadRequest)
		return
	}

	// Read file contents
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Println("Error reading file:", err)
		http.Error(w, "Error reading file.", http.StatusInternalServerError)
		return
	}

	// Parse file
	article, err := extractTextContentFromDocx(fileBytes)
	if err != nil {
		log.Println("Error extracting text from .docx file:", err)
		http.Error(w, "Error extracting text from file.", http.StatusInternalServerError)
		return
	}

	// Generate basic HTML
	html := `<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>` + article.Title + `</title>
	</head>
	<body>
		<h1>` + article.Title + `</h1>`
	for _, content := range article.Content {
		html += "\n" + content.Html()
	}
	html += `
	</body>
</html>`

	// Respond to client
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}
