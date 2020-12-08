package main

import (
	"html/template"
	"net/http"
	"strings"
)

func StoryHandler(story Story) http.HandlerFunc {
	fallbackHandler := DefaultMux()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pathToChapter := strings.ToLower(r.URL.Path[1:])

		if chapter, ok := story[pathToChapter]; ok {
			TemplatedStory(w, chapter)
		} else {
			fallbackHandler.ServeHTTP(w, r)
		}
	})
}

func DefaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}

func TemplatedStory(w http.ResponseWriter, chapter Chapter) {
	const htmlTemplate string = `
<!doctype html>
<html lang="">
<head>
  <meta charset="utf-8">
  <title></title>
  <meta name="description" content="">
  <meta name="viewport" content="width=device-width, initial-scale=1">

  <meta property="og:title" content="">
  <meta property="og:type" content="">
  <meta property="og:url" content="">
  <meta property="og:image" content="">

  <meta name="theme-color" content="#fafafa">
</head>
<body>

  <h1>{{ .Title }}</h1>
{{ range $i, $v := .Story }}
  <p>{{ $v }}</p>
{{ end }}

  <br/>

  <ul>
{{ range $i, $v := .Options }}
   <li><a href="http://localhost:8080/{{ $v.Arc }}">{{ $v.Text }}</a></li>
{{ end }}
  </ul>

</body>
</html>
`

	t, _ := template.New("chapter").Parse(htmlTemplate)
	err := t.Execute(w, chapter)
	if err != nil {
		panic(err)
	}
}
