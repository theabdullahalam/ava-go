package brain2

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/theabdullahalam/ava-go/internal/context"
	"github.com/theabdullahalam/ava-go/internal/utils"
)

type Payload struct {
	Key string
	Value string
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

func GetTaggedString(message string, tag string) string {
	return fmt.Sprintf("{{%s}}%s{{/%s}}", tag, message, tag)
}

func (node Node) Publish(message string) {

	topic := node.Topic
	topic_url := fmt.Sprintf("https://ntfy.sh/%s", topic)
	resp, err := http.Post(topic_url, "text/plain", strings.NewReader(message))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

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

func GetAva() (Node, bool) {
	ava_topic, ok := context.GetFromContext("node.json", "ava")
	if !ok {
		return Node{}, false
	}
	return Node{Name: "Ava", Topic: ava_topic}, true
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