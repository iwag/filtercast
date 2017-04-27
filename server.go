package main

import (
	"encoding/xml"
	"fmt"
)

func main() {
	type Enclosure struct {
		Url string `xml:"url,attr"`
	}
	type Item struct {
		XMLName xml.Name `xml:"item"`
		Link    string `xml:"link"`
		Title string `xml:"title"`
		Description string `xml:"description"`
		Copyright string `xml:"copyright"`
		PubDate string `xml:"pubDate"`
		Enclosure Enclosure `xml:"enclosure"`
	}
	v := Item{}
	data := `
<item>
<link>
http://example.com/20170404/1
</link>
<title>
<![CDATA[title title title]]>
</title>
<description>
</description>
<pubDate>Sun, 04 Apr 2017 00:00:00 +0000</pubDate>
<copyright>iwag</copyright>
<enclosure url="http://example.com/201704041.mp3" length="0" type="audio/mpeg"/>
<itunes:summary>
</itunes:summary>
<itunes:author>iwag</itunes:author>
</item>
`
	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("err %v", err)
		return
	}
	fmt.Printf("%v", v)

}
