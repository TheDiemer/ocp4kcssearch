package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/TheDiemer/ocp4kcssearch/structs"
)

func main() {
	fmt.Println("This tool is to collect all KCS Solutions and Articles for OCP 4.X")
	encodedURL := urlBuilder()
	resp, err := http.Get(encodedURL)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("experienced error:", err, "\nwhile reaching out to ", encodedURL)
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			//fmt.Println(body)
			var result structs.Result
			json.Unmarshal([]byte(body), &result)

			size := len(result.Response.Docs)
			fmt.Println("We found", size, "and hydra told us we'd find", result.Response.NumFound)
			for _, s := range result.Response.Docs {
				row := []string{s.Title, s.ViewURI, s.DocumentKind, s.AuthorSSOName, s.LastModifiedBySSOName, s.CreatedDate, s.LastModifiedDate, s.HasPublishedRevision}
				fmt.Println(row)
				//fmt.Println("KCS:", s.ViewURI, "lists the following products: ", s.Product, " and the following tags: ", s.Tags, "\n---")
			}

			/* Now we've got a structure of a ton of KCS Solutions/Articles where the
			Product includes RHOCP and the Internal Tags include ocp_4. */
		}
	}
}

func urlBuilder() string {
	base, err := url.Parse("https://access.redhat.com/hydra/rest/search/kcs")
	if err != nil {
		return "Failure"
	}

	params := url.Values{}
	params.Add("q", "(documentKind:Solution OR documentKind:Article) AND product:((\"Red Hat OpenShift Container Platform\")) AND internalTags:((\"ocp_4\"))")
	params.Add("rows", strconv.Itoa(10000))
	base.RawQuery = params.Encode()

	// fmt.Println(base.String())
	return base.String()
}
