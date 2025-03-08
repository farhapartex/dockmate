package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
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
	venv := flag.String("venv", "venv", "Virtual environment name")
	pyve := flag.String("pyve", "python3.10", "Python version")

	flag.Parse()

	if *framework == "" {
		fmt.Println("Error: Framework is required. Usage: go run main.go --framework=<django|fastapi|flask>")
	}

	framework = toLower(framework)
	generateProject(framework, db, redis, celery, port, venv, pyve)
}

func generateProject(framework *string, db *string, redis *string, celery *string, port *int, venv *string, pyve *string) {
	switch *framework {
	case "django":
		generateDjangoProject(pyve, venv)
	case "fastapi":
		generateFastAPIProject()
	case "flask":
		generateFlaskProject()
	}

	createDockerfile(*port)
}

func deleteFolderIfExists(dirName *string) {
	if _, err := os.Stat(*dirName); err == nil {
		fmt.Println("Project directory already exists. Removing...")
		if err := os.RemoveAll(*dirName); err != nil {
			fmt.Println("Error removing existing project directory:", err)
			return
		}
	}

}

func generateDjangoProject(pyve *string, venv *string) {
	projectDir := "dockmake_django"
	deleteFolderIfExists(&projectDir)

	err := os.Mkdir(projectDir, 0755)
	if err != nil {
		fmt.Println("Error to create django project root dir: ", err)
		return
	}

	cmd := exec.Command(*pyve, "-m", "venv", *venv)
	cmd.Dir = projectDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error creating virtual environment:", err)
		return
	}

	installCmd := exec.Command("bash", "-c", "source venv/bin/activate && pip install django djangorestframework")
	installCmd.Dir = projectDir
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	if err := installCmd.Run(); err != nil {
		fmt.Println("Error installing Django and DRF:", err)
		return
	}
	fmt.Println("Django and DRF installed successfully!")

	djangoCmd := exec.Command("bash", "-c", "source venv/bin/activate && django-admin startproject core .")
	djangoCmd.Dir = projectDir
	djangoCmd.Stdout = os.Stdout
	djangoCmd.Stderr = os.Stderr
	if err := djangoCmd.Run(); err != nil {
		fmt.Println("Error creating Django project:", err)
		return
	}
	fmt.Println("Django core project created successfully!")

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
