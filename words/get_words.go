package words

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func GenerateValidWordsFile() error {
	file, err := os.Open("words.txt")
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	words := strings.Split(string(content), "\n")

	var validWords []string
	for _, word := range words {
		if word != "" {
			validWords = append(validWords, word)
		}
	}

	goSource := generateGoSource(validWords)

	err = os.WriteFile("valid_words.go", []byte(goSource), 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	fmt.Println("valid_words.go has been generated successfully.")
	return nil
}

func generateGoSource(words []string) string {
	var sb strings.Builder
	sb.WriteString("package words\n\n")
	sb.WriteString("var ValidWordList = []string{\n")
	for _, word := range words {
		sb.WriteString("\t\"" + word + "\",\n")
	}
	sb.WriteString("}\n")
	return sb.String()
}
