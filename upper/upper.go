package upper

import (
	"log"

	"strings"

	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedUpperServiceServer
}

func (s *Server) ToUpper(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	upperCasedString := strings.ToUpper(message.Body)
	log.Printf("Returning message body to client: %s", upperCasedString)
	return &Message{Body: upperCasedString}, nil
}
