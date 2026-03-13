package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Repository struct {
	// В JSON это поле называется "name"
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Description string `json:"description"` // Описание репозитория
	Stars       int    `json:"stargazers_count"`
	Forks       int    `json:"forks_count"`
	Language    string `json:"language"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	HTMLURL     string `json:"html_url"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Ошибка")
	}
	repos := os.Args[1]
	parts := strings.Split(repos, "/")

	owner := parts[0]
	name := parts[1]

	url := fmt.Sprintf("http://api.github.com/repos/%s/%s", owner, name)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Ошибка при запросе:", err)
		return
	} else if resp.StatusCode == 404 {
		fmt.Printf("Ошибка:%s\n", resp.Status)
		return
	} else if resp.StatusCode == 200 {
		fmt.Printf("Запрос успешно отправлен:%s\n", resp.Status)
	} else if resp.StatusCode == 500 {
		fmt.Printf("Что то пошло не так:%s\n", resp.Status)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Ошибка при чтении")
		return
	}

	var repo Repository

	err = json.Unmarshal(body, &repo)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("\nИНФОРМАЦИЯ О РЕПОЗИТОРИИ:")
	fmt.Println("================================")
	fmt.Printf("Название:    %s\n", repo.Name)
	fmt.Printf("Полное имя: %s\n", repo.FullName)
	fmt.Printf("Описание:    %s\n", repo.Description)
	fmt.Printf("Звезд:       %d\n", repo.Stars)
	fmt.Printf("Форков:      %d\n", repo.Forks)
	fmt.Printf("Язык:        %s\n", repo.Language)
	fmt.Printf("URL:         %s\n", repo.HTMLURL)

	if repo.CreatedAt != "" {
		fmt.Printf("📅 Создан:%s\n", repo.CreatedAt[:10])
	}

}
