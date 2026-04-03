package utils

import (
	"fmt"
	"os"
	"strings"
)

func AsciiArt(s string) string {
	standardAscii, err := os.ReadFile("utils/standard.ascii")
	if err != nil {
		fmt.Printf("Error reading files: %v\n",err)
		return ""
	}
	content := string(standardAscii)
	bannerLine := strings.Split(content,"\n")
	inputPart := strings.Split(s,"\\n")

	var AsciiArt strings.Builder
	for _, word := range inputPart{
		if word == ""{
			AsciiArt.WriteString("\n")
			continue
		}
		for i := range 8 {
			for _, char := range word {
				line := getLines(char,bannerLine)
				AsciiArt.WriteString(line[i])
			}
			AsciiArt.WriteString("\n")
		}
	}
	return AsciiArt.String()
}

func getLines(char rune, lines []string) []string {
	startLine := ((char - 32) * 9) +1
	return lines[startLine:startLine+8]
}