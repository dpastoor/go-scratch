package main

import (
	"os"

	"github.com/alecthomas/chroma/quick"
)

func main() {
	f, _ := os.Create("test1.html")
	_ = quick.Highlight(f, `
		test <- function() {
			return(a + b)
		}
	
		df %>% group_by(col) %>%
		summarize(cmax = max(conc))
	
	`, "go", "html", "friendly")
}
