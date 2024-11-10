package table

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"strings"
)

func addSpaces(s string) string {
	if !strings.HasPrefix(s, " ") {
		s = " " + s
	}
	if !strings.HasSuffix(s, " ") {
		s += " "
	}
	return s
}
func addSpacesToSlice(data []string) []string {
	for i, s := range data {
		data[i] = addSpaces(s)
	}
	return data
}
func addSpacesToMatrix(data [][]string) [][]string {
	for i, row := range data {
		data[i] = addSpacesToSlice(row)
	}
	return data
}

func Gen(headers []string, rows [][]string) {
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		Headers(addSpacesToSlice(headers)...).
		Rows(addSpacesToMatrix(rows)...)
	fmt.Println(t.Render())
}
