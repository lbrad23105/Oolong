package main

import (
	"os"
	"os/exec"
)

type app struct {
	name    string
	version string
}

func directoryStructure(appName string, mode os.FileMode) {
	src := "./" + appName + "/src"
	style := "./" + appName + "/style"
	images := "./" + appName + "/images"
	os.Mkdir(appName, mode)
	os.Mkdir(src, mode)
	os.Mkdir(style, mode)
	os.Mkdir(images, mode)
}

func generateHTML(rootDir string) {
	// This defines the html document.
	html := "<!DOCTYPE html>\n"
	html += "<html lang = 'en'>\n"
	html += "<head></head>\n"
	html += "<body>\n"
	html += "<h1 align = 'center'>Oolong</h1>\n"
	html += "</body>\n"
	html += "</html>\n"

	index, err := os.Create("./index.html")
	if err != nil {
		panic(err)
	}

	index.WriteString(html)
	index.Sync()
	index.Close()
}

func generateWebServer() {
	var serverString string
	serverString = "package main\n\n"
	serverString += "import \"net/http\"\n"
	serverString += "import \"fmt\"\n\n"
	serverString += "func rootHandler(w http.ResponseWriter, r *http.Request) {\n"
	serverString += "   fmt.Fprintf(w, \"Hello World from Oolong\")\n"
	serverString += "}\n\n"
	serverString += "func main() {\n"
	serverString += "   mux := http.NewServeMux()\n"
	serverString += "   mux.HandleFunc(\"/\", rootHandler)\n"
	serverString += "   http.ListenAndServe(\":5000\", mux)\n"
	serverString += "}"

	server, err := os.Create("./server.go")
	if err != nil {
		panic(err)
	}

	server.WriteString(serverString)
	server.Sync()
	server.Close()
}

func compileWebServer() {
	server := exec.Command("go", "build", "-o", "server", "./server.go")
	execServer := exec.Command("./server")
	server.Run()
	execServer.Run()
}

func main() {
	var appName string
	var version string

	appName = os.Args[1]
	version = os.Args[2]

	app := new(app)
	app.name = appName
	app.version = version

	directoryStructure(app.name, os.FileMode(0755))
	generateHTML(app.name)
	generateWebServer()
	compileWebServer()
}
