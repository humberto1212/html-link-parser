package utils

import (
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/net/html"
)

// =======================
// 		 parse func
// =======================
func parse(bs string) []string {
	// =======================
	//		Conditionals
	//========================
	//var isLink bool
	//var isEndLink bool
	isLinkN := 0
	isEndLinkN := 0

	//========================
	//  	Slice
	//========================
	var v []string

	tkn := html.NewTokenizer(strings.NewReader(bs))

	// START LOOP
	for {
		tt := tkn.Next()

		switch {

		case tt == html.ErrorToken:

			return v

		case tt == html.StartTagToken:
			t := tkn.Token()

			// ------------ Search <a> ---------
			if t.Data == "a" {
				// FOUND <a> IS ON
				isLinkN = 1
				// CLEAN VALUE AND START OVER
				isEndLinkN = 0
			}

			//......Print example.......
			//isLink = t.Data == "a"
			// if isLink {
			// 	fmt.Println(t.Data)
			// }

			// ------------ Print the value of href ---------
			for _, a := range t.Attr {
				if a.Key == "href" {
					//fmt.Println("Found href:", a.Val)
					v = append(v, a.Val)
					break
				}
			}

		case tt == html.EndTagToken:
			t := tkn.Token()

			//......Print example.......
			// isEndLink = t.Data == "a"
			// if isEndLink {
			// 	fmt.Println(t.Data)
			// }

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

		}

	}

}

// =======================
// Read html & call parse func
// Create a Map & list results
// =======================
func LinkParser() map[string]interface{} {
	// Maps
	m := make(map[string]interface{})
	var mList []map[string]interface{}
	mResult := make(map[string]interface{})

	// Variables
	var textValues []string

	// 	Read html file
	bs, err := ioutil.ReadFile("C:\\Users\\hangelucci\\Desktop\\playground\\go-exer\\html-link-parser\\template\\index.html")
	if err != nil {
		log.Fatalln(err)
	}

	// 	Call parse func
	parsed := parse(string(bs))

	for i, v := range parsed {

		s := strings.TrimSpace(v)

		if len(s) > 0 {
			if strings.Contains(s, "/") {

				_, ok := m["href"]
				if ok {

					mList = append(mList, m)

					m = make(map[string]interface{})
					textValues = nil
				}

				m["href"] = s

			} else if !strings.Contains(s, "/") {
				textValues = append(textValues, s)
				m["Text"] = textValues

			}

		}

		if len(parsed)-1 == i {

			mList = append(mList, m)

		}

	} // END OF RANGE

	mResult["Link"] = mList

	return mResult
}
