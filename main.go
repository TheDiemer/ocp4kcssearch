package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
			var result structs.Result
			json.Unmarshal([]byte(body), &result)

			size := len(result.Response.Docs)
			fmt.Println("We found", size, "and hydra told us we'd find", result.Response.NumFound)
			var data [][]string
			tmpRow := []string{"Title", "ViewURI", "DocumentKind", "AuthorSSOName", "LastModifiedBySSOName", "CreatedDate", "LastModifiedDate", "HasPublishedRevision"}
			data = append(data, tmpRow)
			for _, s := range result.Response.Docs {
				var row []string
				if s.Title == "" {
					row = []string{s.AllTitle, s.ViewURI, s.DocumentKind, s.AuthorSSOName, s.LastModifiedBySSOName, s.CreatedDate, s.LastModifiedDate, s.HasPublishedRevision}
				} else {
					row = []string{s.Title, s.ViewURI, s.DocumentKind, s.AuthorSSOName, s.LastModifiedBySSOName, s.CreatedDate, s.LastModifiedDate, s.HasPublishedRevision}
				}
				data = append(data, row)
			}
			success := csvWriter(data)
			if success {
				fmt.Println("The csv file containing all", size, "results in the new file: /tmp/OCP4KCS.csv")
			} else {
				fmt.Println("We encountered an error. Hopefully it should be recorded above.")
			}
		}
	}
}

func csvWriter(data [][]string) bool {
	csvFileName := "/tmp/OCP4KCS.csv"
	csvFile, err := os.Create(csvFileName)
	if err != nil {
		fmt.Println("Apologies, but while trying to write", csvFileName, "and encountered the following error:\n", err)
		return false
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	writer.WriteAll(data)

	return true
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

	return base.String()
}
