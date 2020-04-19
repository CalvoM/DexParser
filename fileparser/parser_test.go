package fileparser_test

import (
	"os"
	"testing"

	"github.com/CalvoM/DexParser/fileparser"
)

func TestGetFirstDataElement(t *testing.T) {
	got := fileparser.GetFirstDataElement("Hello#World#Welcome#Back", "#")
	expected := "Hello"
	if got != expected {
		t.Errorf("Expected:%q , Got:%q", expected, got)
	}
}

func TestIsHeaderPresent(t *testing.T) {
	got := fileparser.IsHeaderPresent("Ozark", []string{"Ozark", "Sona", "Us"})
	expected := true
	if got != expected {
		t.Errorf("Expected:%t , Got:%t", expected, got)
	}
}

func TestStreamLine(t *testing.T) {
	file, _ := os.Open("./mock.txt")
	gotLine, gotCount := fileparser.StreamLine(*file, 0)
	expectedLine := "Mary had a little lamb"
	expectedCount := len(expectedLine) + 2
	if gotLine != expectedLine {
		t.Errorf("Expected:%s,Got:%s", expectedLine, gotLine)
	}
	if gotCount != expectedCount {
		t.Errorf("Expected:%d , Got:%d", expectedCount, gotCount)
	}
}
