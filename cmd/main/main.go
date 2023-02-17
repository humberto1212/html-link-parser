package main

import (
	"fmt"
	//"htmlLinkParser/models"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func main() {

	// =======================
	// 		Read html file
	// =======================
	bs, err := ioutil.ReadFile("C:\\Users\\hangelucci\\Desktop\\playground\\go-exer\\html-link-parser\\template\\index.html")
	if err != nil {
		log.Fatalln(err)
	}

	// =======================
	//		Conditionals
	//========================
	var isLink bool
	var isEndLink bool
	isLinkN := 0
	isEndLinkN := 0

	//========================
	//  	Map
	//========================
	var v []string
	//vals := make(map[string]models.LinksList)

	//========================
	// 		Parse File
	//========================
	tkn := html.NewTokenizer(strings.NewReader(string(bs)))

	// START LOOP
	for {
		tt := tkn.Next()

		switch {

		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			t := tkn.Token()

			// ------------ Search <a> ---------
			isLink = t.Data == "a"

			if t.Data == "a" {
				// FOUND <a> IS ON
				isLinkN = 1
				// CLEAN VALUE AND START OVER
				isEndLinkN = 0
			}

			if isLink {
				fmt.Println(t.Data)
			}

			// ------------ Print the value of href ---------
			for _, a := range t.Attr {
				if a.Key == "href" {
					fmt.Println("Found href:", a.Val)
					v = append(v, a.Val)
					break
				}
			}

		case tt == html.EndTagToken:
			t := tkn.Token()

			isEndLink = t.Data == "a"

			if isEndLink {
				fmt.Println(t.Data)
			}

			if t.Data == "a" {
				// FOUND </a> IS ON
				isEndLinkN = 1
				// CLEAN VALUE AND START OVER
				isLinkN = 0
			}

		case tt == html.TextToken:

			t := tkn.Token()

			//vals = append(vals, t.Data)

			if isLinkN == 1 && isEndLinkN == 0 {
				//fmt.Println(t.Data)
				v = append(v, t.Data)
			}

			// ***========== FIX ============***
			fmt.Println(v)
			// ***========== FIX ============***
		}
	}

}

/*
=====================================================================================
	Example:

	Link{
	Href: "/dog",
	Text: "Something in a span Text not in a span Bold text!",
	}

	=>

	*** {
			href: "/cat",
			Text: " other span text other Text not in span other Bold text HI!"
	},
	*** {
			href: "/dog",
			Text: " other span text other Text not in span other Bold text HI!"
	}

	=

	JSON
=====================================================================================
*/
