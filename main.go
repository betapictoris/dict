package main

import (
    "log"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "strings"

    "github.com/charmbracelet/glamour"
)

var lang = "en"
var api_url = "https://api.dictionaryapi.dev/api/v2/entries/"

func parseWord(word string) string {
    /* 
    Function to return a URL, for dictionary API.

    word        string      ==>     Word to return URL for
    Returns     string      <==     Full URL (with lang code and word)
    */
    return api_url + lang + "/" + word 
}

func main() {
    var words []Word
    word  := ""
    mdDef := ""

    if len(os.Args) >= 2 {
        word = os.Args[1]

        if len(os.Args) >= 3 {
            lang = os.Args[2]
        }
    } else {
        log.Fatal("Usage: dict <Word> [<Language>]")
    }

    resp, err := http.Get(parseWord(word))
    cont, err := ioutil.ReadAll(resp.Body)
    err = json.Unmarshal([]byte(string(cont)), &words)

    if err != nil {
        fmt.Printf("Couldn't get that word for you...\nFull debug: ")
        log.Fatal(err)
        log.Fatal(1)
    }
    
    mdDef += "# " + strings.Title(words[0].Word) + "\n"
    
    for i:=0; i < len(words[0].Meanings[0].Definitions); i++ {
        if strings.HasPrefix(string(words[0].Meanings[0].Definitions[i].Definition[0]), "(") != true {
            mdDef += " - " + string(words[0].Meanings[0].Definitions[i].Definition) + "\n"
        }
    }
    
    out, err := glamour.Render(mdDef, "dark")

    fmt.Printf(out)
}


// Type for JSON
type Word struct {
	Word      string `json:"word"`
	Phonetics []struct {
		Audio     string `json:"audio"`
		SourceURL string `json:"sourceUrl,omitempty"`
		License   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"license,omitempty"`
		Text string `json:"text,omitempty"`
	} `json:"phonetics"`
	Meanings []struct {
		PartOfSpeech string `json:"partOfSpeech"`
		Definitions  []struct {
			Definition string        `json:"definition"`
			Synonyms   []interface{} `json:"synonyms"`
			Antonyms   []interface{} `json:"antonyms"`
		} `json:"definitions"`
		Synonyms []string      `json:"synonyms"`
		Antonyms []interface{} `json:"antonyms"`
	} `json:"meanings"`
	License struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"license"`
	SourceUrls []string `json:"sourceUrls"`
}
