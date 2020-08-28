package http

import "encoding/xml"

const XMLString = `
<user name="las", age=19>
    <loves>
        <love name="programming" score=1></love>
        <love name="python" score=2></love>
    </loves>
</user>
`

type User struct {
	XMLName xml.Name `xml:"user"`
	Name    string   `xml:"name,attr"`
	Age     string   `xml:"age,attr"`
	Loves   struct {
		Love []struct {
			Name  string `xml:"name,attr"`
			Score string `xml:"score,attr"`
		} `xml:"love"`
	} `xml:"loves"`
}


func simple(){
	user := User{}
	xml.Unmarshal([]byte(XMLString), &user)
	println(user.Loves.Love[0].Name)
}