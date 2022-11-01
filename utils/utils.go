package utils

import "github.com/Puyodead1/fosscord-server-go/initializers"

// generates a snowflake id
func GenerateID() string {
	return initializers.Node.Generate().String()
}
