package main

import (
    "context"
    "fmt"
    "github.com/aws/aws-lambda-go/lambda"
)

// Define a estrutura do request e do response
type Request struct {
    Name string `json:"name"`
}

type Response struct {
    Message string `json:"message"`
}

// Handler da Lambda
func handler(ctx context.Context, req Request) (Response, error) {
    // Formata a mensagem de resposta
    message := fmt.Sprintf("Hello, %s!", req.Name)
    return Response{Message: message}, nil
}

func main() {
    // Inicia o handler da Lambda
    lambda.Start(handler)
}
