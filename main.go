// gd-demo-5-previews is step 5 of the git-deploy demo: development branches
// deployed as applications of their own.
//
// A preview does not know which branch it came from — the platform injects
// neither the branch nor the commit. So what makes two preview URLs look
// different is the code on them: the block of constants below is the only thing
// that changes from branch to branch, which is exactly what a feature branch is.
package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	// branchLabel names the branch this code lives on.
	branchLabel = "feat/search"
	// headline is what this version of the product is about.
	headline = "Search"
	// accent is the colour that makes this deployment recognisable at a glance
	// when two of them are open side by side.
	accent = "#e8710a"
	// pitch is one sentence about what changed here.
	pitch = "Find a product without scrolling for it."
)

const version = "v1"

var page = template.Must(template.New("page").Parse(`<!doctype html>
<title>{{.Headline}} — gd-demo-5-previews</title>
<style>
  :root { color-scheme: light dark }
  body { font: 16px/1.6 system-ui, sans-serif; margin: 0 }
  header { background: {{.Accent}}; color: #fff; padding: 3rem 2rem }
  h1 { margin: 0 0 .3rem; font-size: 2.4rem }
  .branch { font-family: ui-monospace, monospace; opacity: .85 }
  main { padding: 2rem; max-width: 42rem }
  dt { font-weight: 600; margin-top: .8rem }
  dd { margin: 0 }
</style>
<header>
  <h1>{{.Headline}}</h1>
  <div class="branch">branch {{.Branch}}</div>
</header>
<main>
  <p>{{.Pitch}}</p>
  <dl>
    <dt>pod</dt><dd>{{.Pod}}</dd>
    <dt>version</dt><dd>{{.Version}}</dd>
    <dt>served</dt><dd>{{.Served}}</dd>
  </dl>
  <p>Push a branch, get a URL of its own. Delete the branch, it goes away.
  Every preview runs the same pipeline and the same gates as the application it
  came from — a branch that breaks the tests never gets a URL at all.</p>
</main>
`))

func main() {
	log.Printf("gd-demo-5-previews %s (branch %s) starting on :8080", version, branchLabel)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		hostname, _ := os.Hostname()
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		err := page.Execute(w, map[string]string{
			"Headline": headline,
			"Branch":   branchLabel,
			"Accent":   accent,
			"Pitch":    pitch,
			"Pod":      hostname,
			"Version":  version,
			"Served":   time.Now().Format(time.RFC3339),
		})
		if err != nil {
			log.Printf("rendering the page failed: %v", err)
		}

		log.Printf("%s %s from %s (%s)", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start).Round(time.Millisecond))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
