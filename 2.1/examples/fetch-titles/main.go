package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/zeevallin/go-ninjs/2.1/ninjs"
)

var files = []string{
	"https://raw.githubusercontent.com/iptc/newsinjson/refs/heads/main/examples/2.1/SIPA_image_2.json",
	"https://raw.githubusercontent.com/iptc/newsinjson/refs/heads/main/examples/2.1/ap_audio.json",
	"https://raw.githubusercontent.com/iptc/newsinjson/refs/heads/main/examples/2.1/ap_image.json",
	"https://raw.githubusercontent.com/iptc/newsinjson/refs/heads/main/examples/2.1/ap_video.json",
	"https://raw.githubusercontent.com/iptc/newsinjson/refs/heads/main/examples/2.1/businesswire-newsml-20130605006126.json",
	"https://raw.githubusercontent.com/iptc/newsinjson/refs/heads/main/examples/2.1/businesswire-newsml-20130731006140.json",
	"https://raw.githubusercontent.com/iptc/newsinjson/refs/heads/main/examples/2.1/dpa_text.json",
	"https://raw.githubusercontent.com/iptc/newsinjson/refs/heads/main/examples/2.1/imageEncodedRights.json",
	"https://raw.githubusercontent.com/iptc/newsinjson/refs/heads/main/examples/2.1/imageLinkedRights.json",
	"https://raw.githubusercontent.com/iptc/newsinjson/refs/heads/main/examples/2.1/ninjsExSimpleText_2.json",
	"https://raw.githubusercontent.com/iptc/newsinjson/refs/heads/main/examples/2.1/ntb_text.json",
	"https://raw.githubusercontent.com/iptc/newsinjson/refs/heads/main/examples/2.1/tt_text_image_2.json",
}

func main() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)

	for _, url := range files {
		article, err := getArticle(url)
		if err != nil {
			fmt.Fprintf(w, "error fetching: %v\n", err)
			continue
		}

		fmt.Fprintf(w, "url:\t%s\n", url)

		for i, headline := range article.Headlines {
			fmt.Fprintf(w, "headline [%d]:\t%s\n", i+1, headline.Value)
		}

		w.Write([]byte("\n"))
	}

	w.Flush()
}

func getArticle(url string) (ninjs.Item, error) {
	resp, err := http.Get(url)
	if err != nil {
		return ninjs.Item{}, err
	}
	defer resp.Body.Close()

	var article ninjs.Item
	if err := json.NewDecoder(resp.Body).Decode(&article); err != nil {
		return ninjs.Item{}, err
	}

	return article, nil
}
