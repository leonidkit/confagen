package confagen

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// YAMLNode is a node of a YAML file with children.
type YAMLNode struct {
	childs  []*YAMLNode
	name    string
	value   interface{}
	valType string
	path    string
}

// parseNode parses YAML node.
func (node *YAMLNode) parseNode(data interface{}) *YAMLNode {
	switch data.(type) {
	case map[interface{}]interface{}:
		if len(data.(map[interface{}]interface{})) == 0 {
			return node
		}
		for k, v := range data.(map[interface{}]interface{}) {
			newNode := &YAMLNode{
				name: k.(string),
				path: strings.Join([]string{node.path, k.(string)}, "."),
			}
			node.childs = append(node.childs, newNode)
			newNode.parseNode(v)
		}
	case interface{}:
		switch t := reflect.TypeOf(data).Kind(); t {
		case reflect.String:
			if stringDuration(data.(string)) {
				node.valType = "Duration"
			} else {
				node.valType = "string"
			}
		case reflect.Slice:
			for _, v := range data.([]interface{}) {
				node.valType = fmt.Sprintf("[]%s", reflect.TypeOf(v).String())
			}
		default:
			node.valType = t.String()
		}
		node.value = data
	case nil:
		node.valType = "string"
		node.value = ""
	}
	return node
}

// printNode print node for testing and debugging.
func (node *YAMLNode) printNode(depth int) {
	fmt.Println()
	fmt.Print(strings.Repeat("\t", depth))
	fmt.Printf("%v: ", node.name)

	if len(node.childs) == 0 {
		switch t := reflect.TypeOf(node.value).Kind(); t {
		case reflect.Slice:
			for _, v := range node.value.([]interface{}) {
				fmt.Println()
				fmt.Print(strings.Repeat("\t", depth+1))
				fmt.Print("- ", v)
			}
			return
		case reflect.Int, reflect.String, reflect.Bool:
			fmt.Printf("%v", node.value)
			return
		}
	}
	for _, child := range node.childs {
		child.printNode(depth + 1)
	}
}

// stringDuration checks whether the string is a duration.
func stringDuration(str string) bool {
	for _, ch := range []string{"h", "m", "s"} {
		if strings.Count(str, ch) == 1 {
			if splited := strings.Split(str, ch); len(splited) == 2 && splited[1] == "" {
				_, err := strconv.ParseFloat(splited[0], 64)
				if err == nil {
					return true
				}
				_, err = strconv.ParseInt(splited[0], 10, 64)
				if err == nil {
					return true
				}
			}
		}
	}
	return false
}
