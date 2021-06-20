package confagen

import (
	"bytes"
	"embed"
	"io/ioutil"
	"log"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

//go:embed templates/*.tpl
var tpls embed.FS

type Confagen struct {
	tmpls *template.Template
}

func New() *Confagen {
	templates, err := template.ParseFS(tpls, "templates/*")
	if err != nil {
		log.Fatal(err)
	}

	return &Confagen{
		tmpls: templates,
	}
}

func (c *Confagen) Generate(srcFilePath string, dstFilePath string) {
	fileB, err := ioutil.ReadFile(srcFilePath)
	if err != nil {
		log.Fatal(err)
	}

	file := make(map[interface{}]interface{})
	err = yaml.Unmarshal(fileB, &file)
	if err != nil {
		log.Fatal(err)
	}

	structData := c.parseStructs(c.parseNodes(file))

	var buf bytes.Buffer
	err = c.tmpls.ExecuteTemplate(&buf, "config.tpl", map[string]interface{}{
		"Struct": structData,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(dstFilePath, buf.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Confagen) parseNodes(data interface{}) []*YAMLNode {
	YAMLNodes := make([]*YAMLNode, 0)
	for k, v := range data.(map[interface{}]interface{}) {
		node := &YAMLNode{
			name: k.(string),
			path: k.(string),
		}
		YAMLNodes = append(YAMLNodes, node.parseNode(v))
	}
	return YAMLNodes
}

func (c *Confagen) parseStructs(nodes []*YAMLNode) []*YAMLStruct {
	res := make([]*YAMLStruct, 0)
	for _, node := range nodes {
		strct := &YAMLStruct{
			Name: strings.Title(node.name),
		}
		res = append(res, strct.parseStruct(node))
	}
	return res
}

func (c *Confagen) printNodes(node []*YAMLNode) {
	for _, item := range node {
		item.printNode(0)
	}
}
