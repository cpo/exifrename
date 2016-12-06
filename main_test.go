package main

import (
	"testing"
	"os"
)

func Test_Main(m *testing.T) {
	UnittestMode = true

	os.Chdir("testfiles")
	main()

	testCases := map[string]string{
		"./img_1771.jpg": "2003/12/img_1771.jpg",
	}

	for r := range(testCases) {
		if Renames[r] != testCases[r] {
			m.Error("Error:", r, "was", Renames[r], "but expected", testCases[r])
		}
	}
}
