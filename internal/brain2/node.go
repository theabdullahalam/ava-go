package brain2

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/theabdullahalam/ava-go/internal/utils"
)

type Payload struct {
	Name string
	Description string
}

type Action struct {
	Name string
	Description string
	Payload []Payload
	Type string
}

type Node struct {
	Name string
	Topic string
	Description string
	Actions []Action
}

func (node Node) FilePath() string {
	_node, _ := GetThisNode()
	if _node.Name == node.Name {
		return filepath.Join(utils.GetAvaFolder(), "node.json")
	}
	return filepath.Join(utils.GetAvaFolder(), "network", node.Name + ".json")
}

func GetNetworkNode(name string) (Node, bool) {

	var node Node = Node{};
	avarDir := utils.GetAvaFolder()
	nodefile := filepath.Join(avarDir, "network", name + ".json")
	data, err := os.ReadFile(nodefile)
	if err != nil {
		return Node{}, false
	}
	
	err = json.Unmarshal(data, &node)
	if err != nil {
		return Node{}, false
	}
	
	if node.Name == name {
		return node, true
	}

	return Node{}, false

}

func GetNode(name string) (Node, bool) {
	node, ok := GetThisNode()
	if ok {
		if node.Name == name {
			return node, true
		}
		return Node{}, false
	}

	node, ok = GetNetworkNode(name)
	if ok {
		return node, true
	}

	return Node{}, false
}

func GetThisNode() (Node, bool) {
	var node Node = Node{};
	avarDir := utils.GetAvaFolder()
	nodefile := filepath.Join(avarDir, "node.json")
	data, err := os.ReadFile(nodefile)
	if err != nil {
		return Node{}, false
	}
	
	err = json.Unmarshal(data, &node)
	if err != nil {
		return Node{}, false
	}
	
	return node, true
}