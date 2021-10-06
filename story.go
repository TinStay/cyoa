package cyoa

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"text/template"
)

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTemplate))
}

var tpl *template.Template

var defaultHandlerTemplate = `
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
          <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
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

type HandlerOption func(h *handler) 

func WithTemplate(t *template.Template) HandlerOption {
  return func(h *handler){
    h.t = t
  }
}

func WithPathFn(fn func(r *http.Request) string) HandlerOption{
  return func(h *handler){
    h.pathFn = fn
  }
}

func NewHandler(s Story, opts ...HandlerOption) http.Handler {
  // Set default handler
  h := handler{s, tpl, defaultPathFn}

  // Apply options
  for _, opt := range opts{
    opt(&h)
  }

	return h
}

type handler struct {
	s Story
  t *template.Template
  pathFn func(r *http.Request) string
}

func defaultPathFn(r *http.Request) string {
  path := strings.TrimSpace(r.URL.Path)

	if path == "" || path == "/" {
		path = "/intro"
	}

	//  Remove "/" in the beginning
	return path[1:]
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  path := h.pathFn(r)

	if chapter, ok := h.s[path]; ok {
		err := h.t.Execute(w, chapter)

		if err != nil {
			fmt.Printf("Error: %v", err)
			http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		}

		return
	}

	// Chaper was not found
	http.Error(w, "Chapter was not found.", http.StatusNotFound)

}

func JsonStory(r io.Reader) (Story, error) {
	dec := json.NewDecoder(r)
	var story Story
	if err := dec.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title,omitempty"`
	Paragraphs []string `json:"story,omitempty"`
	Options    []Option `json:"options,omitempty"`
}

type Option struct {
	Text    string `json:"text,omitempty"`
	Chapter string `json:"arc,omitempty"`
}
