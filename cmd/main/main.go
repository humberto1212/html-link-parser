package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func main() {

	// Read file
	bs, err := ioutil.ReadFile("ADD FILE PATH HERE")
	if err != nil {
		log.Fatalln(err)
	}

	//var vals []string
	var isLink bool

	// Parse file
	//tkn := html.NewTokenizer(strings.NewReader(string(bs)))
	tkn := html.NewTokenizer(strings.NewReader(string(bs)))

	for {
		tt := tkn.Next()

		switch {

		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			t := tkn.Token()
			// ------------ Search <a> ---------
			isLink = t.Data == "a"
			if isLink {
				fmt.Println("there is a link", t.Data)
			}

			// ------------ Print the value of href ---------
			for _, a := range t.Attr {
				if a.Key == "href" {
					fmt.Println("Found href:", a.Val)
					break
				}
			}

			// case tt == html.TextToken:

			// 	t := tkn.Token()

			// 	if isLink {
			// 		// 		//vals = append(vals, t.Data)
			// 		// 		//fmt.Println(vals)

			// 		//fmt.Println(t.Data)

			// 		fmt.Println(t)

			// 	}

		}

	}

}
