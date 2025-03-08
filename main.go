package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func toLower(val *string) *string {
	lower_val := strings.ToLower(*val)
	return &lower_val
}

func main() {
	framework := flag.String("framework", "", "Specificy framework: django, fastapi, flask")
	db := flag.String("db", "", "Specifiy Database: postgres, mysql")
	redis := flag.String("redis", "", "Specifiy Redis: redis")
	celery := flag.String("celery", "", "Specifiy Celery: celery")
	port := flag.Int("port", 8000, "Server exposed port")

	flag.Parse()

	if *framework == "" {
		fmt.Println("Error: Framework is required. Usage: go run main.go --framework=<django|fastapi|flask>")
	}

	framework = toLower(framework)
	generateProject(framework, db, redis, celery, port)
}

func generateProject(framework *string, db *string, redis *string, celery *string, port *int) {
	switch *framework {
	case "django":
		generateDjangoProject()
	case "fastapi":
		generateFastAPIProject()
	case "flask":
		generateFlaskProject()
	}

	createDockerfile(*port)
}

func generateDjangoProject() {
	fmt.Println("Django")
}

func generateFastAPIProject() {
	fmt.Println("FastAPI")
}

func generateFlaskProject() {
	fmt.Println("Flask")
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
