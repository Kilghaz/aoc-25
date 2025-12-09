package idea

import (
	"encoding/xml"
	"strings"
)

type XMLRunManagerConfiguration struct {
	XMLName          xml.Name                      `xml:"configuration"`
	Name             string                        `xml:"name,attr"`
	Type             string                        `xml:"type,attr"`
	FactoryName      string                        `xml:"factoryName,attr"`
	FolderName       string                        `xml:"folderName,attr"`
	Module           XMLModule                     `xml:"module"`
	WorkingDirectory XMLKeyValue                   `xml:"working_directory"`
	Parameters       XMLKeyValue                   `xml:"parameters"`
	Kind             XMLKeyValue                   `xml:"kind"`
	Package          XMLKeyValue                   `xml:"package"`
	Directory        XMLKeyValue                   `xml:"directory"`
	FilePath         XMLKeyValue                   `xml:"filePath"`
	Method           XMLRuntimeConfigurationMethod `xml:"method"`
}

type XMLRuntimeConfigurationMethod struct {
	Version int `xml:"v,attr"`
}

type XMLModule struct {
	XMLName xml.Name `xml:"module"`
	Name    string   `xml:"name,attr"`
}

type XMLKeyValue struct {
	Value string `xml:"value,attr"`
}

type RunConfiguration struct {
	Name             string
	Type             string
	FactoryName      string
	Module           string
	WorkingDirectory string
	Parameters       []string
	Kind             string
	Package          string
	Directory        string
	FilePath         string
	FolderName       string
	Method           int
}

func marshalRunConfiguration(configuration *RunConfiguration) string {
	xmlConfig := XMLRunManagerConfiguration{
		XMLName:     xml.Name{Local: "configuration"},
		Name:        configuration.Name,
		Type:        configuration.Type,
		FactoryName: configuration.FactoryName,
		FolderName:  configuration.FolderName,
		Module: XMLModule{
			Name: configuration.Module,
		},
		WorkingDirectory: XMLKeyValue{
			Value: configuration.WorkingDirectory,
		},
		Parameters: XMLKeyValue{
			Value: strings.Join(configuration.Parameters, " "),
		},
		Kind: XMLKeyValue{
			Value: configuration.Kind,
		},
		Package: XMLKeyValue{
			Value: configuration.Package,
		},
		Directory: XMLKeyValue{
			Value: configuration.Directory,
		},
		FilePath: XMLKeyValue{
			Value: configuration.FilePath,
		},
		Method: XMLRuntimeConfigurationMethod{
			Version: configuration.Method,
		},
	}

	marshalled, err := xml.Marshal(xmlConfig)
	if err != nil {
		panic(err)
	}
	return string(marshalled)
}
