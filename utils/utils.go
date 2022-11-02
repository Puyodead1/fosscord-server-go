package utils

import (
	"github.com/bwmarrin/snowflake"
)

var Node *snowflake.Node

func InitSnowflake() {
	var err error
	Node, err = snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
}

// generates a snowflake id
func GenerateID() string {
	return Node.Generate().String()
}
