package confagen

import "strings"

// YAMLStructField is representation of a field in a structure.
type YAMLStructField struct {
	Name      string
	ValType   string
	ViperType string
	Path      string
}

// YAMLStruct is representation of a structure.
type YAMLStruct struct {
	Name   string
	Fields []YAMLStructField
}

func (strct *YAMLStruct) parseStruct(node *YAMLNode) *YAMLStruct {
	if len(node.childs) == 0 {
		processname := func(name string) string {
			splitedname := strings.Split(name, "_")
			for i := 0; i < len(splitedname); i++ {
				splitedname[i] = strings.Title(splitedname[i])
			}
			newname := strings.Join(splitedname, "")

			splitedPath := strings.Split(node.path, ".")
			if len(splitedPath) > 2 {
				newname = strings.Join([]string{strings.Title(splitedPath[len(splitedPath)-2]), newname}, "")
			}
			return newname
		}
		processTypeViper := func(tp string) string {
			switch tp {
			case "[]string":
				return "StringSlice"
			case "[]int":
				return "IntSlice"
			default:
				return strings.Title(tp)
			}
		}
		strct.Fields = append(strct.Fields, YAMLStructField{
			Name:      processname(node.name),
			ValType:   node.valType,
			ViperType: processTypeViper(node.valType),
			Path:      node.path,
		})
		return strct
	}
	for _, child := range node.childs {
		strct.parseStruct(child)
	}
	return strct
}
