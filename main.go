package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
)

// Κώδικας των quad functions ως Go κώδικας
var quadSources = map[string]string{
	"quadA": `package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: quadA width height")
		return
	}
	width, _ := strconv.Atoi(os.Args[1])
	height, _ := strconv.Atoi(os.Args[2])
	if width <= 0 || height <= 0 {
		return
	}
	var result strings.Builder
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == 0 && j == 0 || i == 0 && j == width-1 || i == height-1 && j == 0 || i == height-1 && j == width-1 {
				result.WriteByte('o')
			} else if i == 0 || i == height-1 {
				result.WriteByte('-')
			} else if j == 0 || j == width-1 {
				result.WriteByte('|')
			} else {
				result.WriteByte(' ')
			}
		}
		result.WriteByte('\n')
	}
	fmt.Print(result.String())
}
`,
	"quadB": `package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: quadB width height")
		return
	}
	width, _ := strconv.Atoi(os.Args[1])
	height, _ := strconv.Atoi(os.Args[2])
	if width <= 0 || height <= 0 {
		return
	}
	var result strings.Builder
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == 0 && j == 0 {
				result.WriteByte('/')
			} else if i == 0 && j == width-1 {
				result.WriteByte('\\')
			} else if i == height-1 && j == 0 {
				result.WriteByte('\\')
			} else if i == height-1 && j == width-1 {
				result.WriteByte('/')
			} else if i == 0 || i == height-1 {
				result.WriteByte('-')
			} else if j == 0 || j == width-1 {
				result.WriteByte('|')
			} else {
				result.WriteByte(' ')
			}
		}
		result.WriteByte('\n')
	}
	fmt.Print(result.String())
}
`,
	"quadC": `package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: quadC width height")
		return
	}
	width, _ := strconv.Atoi(os.Args[1])
	height, _ := strconv.Atoi(os.Args[2])
	if width <= 0 || height <= 0 {
		return
	}
	var result strings.Builder
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == 0 {
				result.WriteByte('A')
			} else {
				result.WriteByte('C')
			}
		}
		result.WriteByte('\n')
	}
	fmt.Print(result.String())
}
`,
	"quadD": `package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: quadD width height")
		return
	}
	width, _ := strconv.Atoi(os.Args[1])
	height, _ := strconv.Atoi(os.Args[2])
	if width <= 0 || height <= 0 {
		return
	}
	var result strings.Builder
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == 0 {
				result.WriteByte('A')
			} else {
				result.WriteByte('C')
			}
		}
		result.WriteByte('\n')
	}
	fmt.Print(result.String())
}
`,
	"quadE": `package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: quadE width height")
		return
	}
	width, _ := strconv.Atoi(os.Args[1])
	height, _ := strconv.Atoi(os.Args[2])
	if width <= 0 || height <= 0 {
		return
	}
	var result strings.Builder
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == 0 {
				result.WriteByte('A')
			} else {
				result.WriteByte('C')
			}
		}
		result.WriteByte('\n')
	}
	fmt.Print(result.String())
}
`,
}

func main() {
	// Δημιουργία των εκτελέσιμων αρχείων
	for name, source := range quadSources {
		fileName := name + ".go"
		executableName := name

		// Δημιουργία του αρχείου .go
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Printf("Error creating %s: %v\n", fileName, err)
			return
		}

		// Γράψιμο του κώδικα στο αρχείο
		_, err = file.WriteString(source)
		if err != nil {
			fmt.Printf("Error writing to %s: %v\n", fileName, err)
			file.Close()
			return
		}

		file.Close()

		// Μεταγλώττιση του αρχείου για να δημιουργηθεί το εκτελέσιμο
		cmd := exec.Command("go", "build", "-o", executableName, fileName)
		err = cmd.Run()
		if err != nil {
			fmt.Printf("Error building %s: %v\n", executableName, err)
			return
		}

		// Διαγραφή του προσωρινού αρχείου .go
		err = os.Remove(fileName)
		if err != nil {
			fmt.Printf("Error deleting %s: %v\n", fileName, err)
			return
		}
	}

	// Δημιουργία του quadchecker εκτελέσιμου αρχείου
	cmd := exec.Command("go", "build", "-o", "quadchecker", "main.go")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error building quadchecker: %v\n", err)
		return
	}

	// Διάβασμα της εισόδου
	var input strings.Builder
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input.WriteString(scanner.Text() + "\n")
	}
	inputStr := strings.TrimSuffix(input.String(), "\n")

	// Έλεγχος των quad functions
	checkQuads(inputStr)
}

func checkQuads(inputStr string) {
	// Ορισμός των quad patterns ως πίνακες γραμμών
	quads := map[string][][]string{
		"quadA": {{"o--o", "|  |", "o--o"}},
		"quadB": {{"/*", "*\\", "/*"}, {"\\*", "*\\", "/*"}},
		"quadC": {{"A", "C"}},
		"quadD": {{"A", "C"}},
		"quadE": {{"A", "C"}},
	}

	// Εύρεση των ταιριαστών quad
	var matches []string
	for name, patterns := range quads {
		for _, pattern := range patterns {
			if inputStr == strings.Join(pattern, "\n") {
				matches = append(matches, name)
			}
		}
	}

	// Ταξινόμηση των ταιριασμάτων αλφαβητικά
	sort.Strings(matches)

	// Εμφάνιση αποτελεσμάτων
	if len(matches) > 0 {
		var results []string
		for _, match := range matches {
			results = append(results, fmt.Sprintf("[%s] [%d] [%d]", match, len(strings.Split(inputStr, "\n")), len(strings.Split(inputStr, "\n")[0])))
		}
		fmt.Println(strings.Join(results, " || "))
	} else {
		fmt.Println("Not a quad function")
	}
}
