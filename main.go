package main

import (
	"flag"
	"fmt"
	"os"
)

var allowedFrameworks = map[string]bool{
	"Django":  true,
	"FastAPI": true,
	"Nodejs":  true,
	"Gin":     true,
}

func createDockerfile(port int) {
	fileName := "Dockerfile"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error to create Dockerfile")
		return
	}

	defer file.Close()

	content := fmt.Sprintf(`
	FROM python:3.10-slim

	WORKDIR /app

	COPY . /app/

	RUN pip install --upgrade pip
	RUN pip install --no-cache-dir -r requirements.txt

	EXPOSE 8000

	CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "%d"]
	`, port)
	_, err = file.WriteString(content)

	if err != nil {
		fmt.Println("Error to write in Dockerfile")
		return
	}

}

func main() {
	fmt.Println("Hello Dockmate!")
	framework := flag.String("framework", "", "Framework Name")

	if *framework == "" {
		fmt.Println("Error: --framework flag is missing")
		os.Exit(1)
	}

	if !allowedFrameworks[*framework] {
		fmt.Printf("Error: %s is not a valid framework", *framework)
		os.Exit(1)
	}
	port := flag.Int("port", 8000, "Server exposed port") // return an pointer
	flag.Parse()

	createDockerfile(*port)
}
