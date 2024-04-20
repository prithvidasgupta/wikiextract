/*
Copyright © 2024 Prithvijit Dasgupta <prithvid@umich.edu>
*/
package xsd

type Revision struct {
	Text string `xml:"text"`
}

type RedirectType struct {
	Title string `xml:"title,attr"`
}

type Page struct {
	Id        string       `xml:"id"`
	Title     string       `xml:"title"`
	Namespace int64        `xml:"ns"`
	Revisions []Revision   `xml:"revision"`
	Redirect  RedirectType `xml:"redirect"`
}
