package idea

import (
	"bytes"
	"encoding/xml"
	"github.com/go-xmlfmt/xmlfmt"
	"io"
	"os"
	"reflect"
	"strings"
)

const workspaceXmlFilePath string = ".idea/workspace.xml"

func isRunManagerStartToken(token xml.Token) bool {
	if reflect.TypeOf(token).String() != "xml.StartElement" {
		return false
	}

	if token.(xml.StartElement).Name.Local != "component" {
		return false
	}

	if token.(xml.StartElement).Attr[0].Value != "RunManager" {
		return false
	}

	return true
}

func isRunConfiguration(token xml.Token) bool {
	if reflect.TypeOf(token).String() != "xml.StartElement" {
		return false
	}

	if token.(xml.StartElement).Name.Local != "configuration" {
		return false
	}

	return true
}

func AddRunConfiguration(config *RunConfiguration) {
	file, _ := os.OpenFile(workspaceXmlFilePath, os.O_RDONLY, os.ModePerm)
	output := bytes.NewBuffer([]byte{})

	decoder := xml.NewDecoder(file)
	encoder := xml.NewEncoder(output)

	for {
		token, err := decoder.Token()

		if err == io.EOF {
			break
		}

		_ = encoder.EncodeToken(token)

		if isRunManagerStartToken(token) {
			_ = encoder.Flush()
			output.WriteString(marshalRunConfiguration(config))
		}
	}

	_ = encoder.Flush()
	formatted := strings.TrimSpace(xmlfmt.FormatXML(output.String(), "", "  "))
	_ = os.WriteFile(workspaceXmlFilePath, []byte(formatted), os.ModePerm)
}

func HasRunConfiguration(name string) bool {
	file, _ := os.OpenFile(workspaceXmlFilePath, os.O_RDONLY, os.ModePerm)
	decoder := xml.NewDecoder(file)

	for {
		token, err := decoder.Token()

		if err == io.EOF {
			return false
		}

		if !isRunConfiguration(token) {
			continue
		}

		for _, attr := range token.(xml.StartElement).Attr {
			if attr.Name.Local == "name" && attr.Value == name {
				return true
			}
		}
	}
}
