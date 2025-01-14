package nodes

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/theabdullahalam/ava-go/internal/utils"
)

const NODES_FILE = "nodes.json"

type Node struct {
	Name string
	Topic string
}

func NewNode(name string, topic string) Node {
	return Node{
		Name: name,
		Topic: topic,
	}
}

func GetNodes() []Node {
	ava_folder := utils.GetAvaFolder()
	nodes_file_path := filepath.Join(ava_folder, NODES_FILE)

	data, err := os.ReadFile(nodes_file_path)
	if err != nil {
		fmt.Println(err)
	}

	var nodes []Node
	err = json.Unmarshal(data, &nodes)
	if err != nil {
		fmt.Println(err)
	}

	return nodes
}

func SaveNodes(nodes []Node) {
	ava_folder := utils.GetAvaFolder()
	nodes_file_path := filepath.Join(ava_folder, NODES_FILE)

	data, err := json.MarshalIndent(nodes, "", "    ")
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile(nodes_file_path, data, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func AddNode(node Node) {
	nodes := GetNodes()
	nodes = append(nodes, node)

	SaveNodes(nodes)
}

func RemoveNode(node Node) {
	nodes := GetNodes()
	for i, n := range nodes {
		if n.Name == node.Name {
			nodes = append(nodes[:i], nodes[i+1:]...)
			break
		}
	}

	SaveNodes(nodes)
}