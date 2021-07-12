package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler.HandleMain)
	http.HandleFunc("/LogToLeboncoin", handler.HandleGoogleLogin)

	h := `
<html>
 <body>
  <a name="0"> text </a>
  <a name="1"> abc </a>
  <a name="2"> def ghi jkl </a>
  <a name="3"> vert </a>
  <a name="4"> Some text </a>
 </body>
</html>`

        pattern := "vert"

        doc, err := goquery.NewDocumentFromReader(strings.NewReader(h))
        if err != nil {
                panic(err)
        }

        doc.Find("a").Each(func(i int, s *goquery.Selection) {
                if strings.TrimSpace(s.Text()) == pattern {
                        name, ok := s.Attr("bike")
                        if ok {
                                fmt.Println(bike)
                        }
                }
        })

}

	log.Fatal(http.ListenAndServe(":8081", nil))
}

