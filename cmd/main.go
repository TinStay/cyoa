package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	cyoa "github.com/Basics/src/github.com/TinStay/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the web server on")
	fileName := flag.String("file", "gopher.json", "JSON file that stores your CYOA story")
	flag.Parse()

	fmt.Printf("Story in %s\n", *fileName)

	f, err := os.Open(*fileName)

	// Throw error
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	tpl := template.Must(template.New("").Parse(storyTmpl))

	h := cyoa.NewHandler(story,
		cyoa.WithTemplate(tpl),
		cyoa.WithPathFn(pathFn))

	fmt.Printf("Starting the server at: %d\n", *port)

	mux := http.NewServeMux()

	mux.Handle("/story/", h)

	// Display error if something happens
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))

}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)

	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}

	//  Remove "/" in the beginning
	return path[len("/story/"):]
}

var storyTmpl = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" contet="IE=edge" />
    <metaname="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Choose your own adventure</title>
  </head>
 <body>
  <section class="page">
    <h1>{{.Title}}</h1>
    {{range.Paragraphs}}
      <p>{{.}}</p>
      {{end}}
      <ul>
      {{range.Options}}
          <li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
      {{end}}
      </ul>
    </section>
	<style>
      bdy {
       font-family: helvetica, arial;
     }
      h1 {
        text-aligncenter;
       position:relative;
     }
      .page {
        idth: 80%;
       max-width: 500px;
       margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
       padding: 80px;
        background: #FFFCF6;
        border: 1x solid #eee;
       box-shadow: 0 10px 6px -6px #777;
     }
      ul {
        border-tp: 1px dotted #ccc;
       padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
     li {
        padding-tp: 10px;
      }
      a,
     a:visited {
       text-decoration: none;
        olor: #6295b5;
     }
     a:active,
      a:hover {
        color: #7792a2;
     }
     p {
       text-indent: 1em;
      }
    </style>
 </body>
</html>
`
