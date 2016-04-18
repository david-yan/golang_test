package main

import (
		"fmt"
		// "net/http"
		// "io/ioutil"
		"os"
		"github.com/PuerkitoBio/goquery"
		"strings"
        "strconv"
		)

/* Prints out help message */
func help() {
    fmt.Println("Usage: gosample [-help|-h]")
    fmt.Println("       gosample -urls=<comma-seperated-one-or-more-urls>")
}
/* Check for error */
func check(err error) {
	if err != nil {
		fmt.Printf("%s", err);
		os.Exit(1);
	}
}
/* Create url#.txt file of word count */
func ScrapeWords(url string, index int) {
	doc, err := goquery.NewDocument(url) 
	check(err)

	var wordCount map[string]int
	wordCount = make(map[string]int)

	doc.Find("html").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		pwords := strings.Split(text, " ")
		for _, pword := range pwords {
			words := strings.Split(pword, "\n")
			for _, word := range words {
				word = strings.ToLower(word)
				if word != "" && !strings.ContainsAny(word, "',./?!<>{}[]|/\\\";:@#$`~©^&*()_·+-=%") {
					if count, ok := wordCount[word]; ok {
							wordCount[word] = count + 1;
					} else {
							wordCount[word] = 1;
					}
				}
			}
		}
	})
	f, err := os.Create("url"+strconv.Itoa(index + 1)+".txt")
    check(err)
    defer f.Close()

    data := []byte("url: " + url + "\n")
    _, err = f.Write(data)
    check(err)
	for word, count := range wordCount {
        data = []byte("  "+word+": "+strconv.Itoa(count)+"\n")
        _, err := f.Write(data)
        check(err)
	}
    f.Sync()
}

func main() {
    if len(os.Args) > 2 {
        fmt.Println("Too many arguments")
        help()
        os.Exit(1)
    }
    if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "-help" {
        help()
        os.Exit(0)
    }
    if strings.HasPrefix(os.Args[1], "-urls=") {
        urls := strings.Split(strings.TrimPrefix(os.Args[1],"-urls="), ",")
        for i,url := range urls {
            ScrapeWords(strings.Trim(url," "), i)
        }
    } else {
        fmt.Println("Invalid arguments")
        help()
        os.Exit(1)
    }
}
