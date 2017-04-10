package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

//  https://api.tronalddump.io/random/quote
var (
	version = "v2"
	r       = rand.New(rand.NewSource(time.Now().UnixNano()))
	quotes  = []string{
		"Chuck Norris does not hunt because the word hunting infers the probability of failure. Chuck Norris goes killing",
		"Paper beats rock, rock beats scissors, and scissors beats paper, but Chuck Norris beats all 3 at the same time.",
		"The Chuck Norris Eclipse plugin made alien contact.",
		"A guy was furious at Chuck Norris. He fought with Chuck Norris with all of his rage. Chuck Norris happily fed on his rage until the guy became a hippie. Then Chuck Norris obliterated him, because Chuck Norris doesn't like hippies.",
		"Chuck Norris can sleep standing up...while laying down......awake.",
		"If someone says, nobodys perfect, Chuck Norris takes that as a personal insult.",
		"Chuck Norris taught lions and tigers how to chase down, kill and eat their prey. However, despite his best efforts on proper carnivorous table fare etiquette, they continue to start lavishly eating the assholes and genitals first.",
	}
	htmlTemplate = `<html>
		<head><title>waqquoter</title>
		<link href="https://fonts.googleapis.com/css?family=Palanquin+Dark" rel="stylesheet">
		</head>
		<body style="font-family: 'Palanquin Dark', sans-serif;font-size:2em;display:flex;align-items: center;flex-direction:column;justify-content: center;background-color:#efefef"><h1>The WAQquoter</h1><div style="display:flex;align-items: center;justify-content: center;background-color:white;min-height:50%;min-width:50%;max-width:300px;padding:2em 2em">{{.Quote}}</div></body>
	</html>`
)

func getRandomQuote() string {
	return quotes[r.Intn(len(quotes))]
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("version", version)
		w.Header().Set("content-type", "text/html; charset=utf-8")
		tmpl, _ := template.New("t").Parse(htmlTemplate)
		tmpl.Execute(w, map[string]string{"Quote": getRandomQuote()})
	})

	port := "0.0.0.0:8080"
	log.Println("Listening on: ", port)
	log.Panic(http.ListenAndServe(port, nil))
}

var ()
