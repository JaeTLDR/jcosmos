package main

import (
	"fmt"
	"log"
	"os"

	"github.com/JaeTLDR/jcosmos"
	"github.com/joho/godotenv"
)

var (
	cosmosClient *jcosmos.Jcosmos
)

func init() {
	godotenv.Load()
	cosmosClient = jcosmos.Init(
		os.Getenv("JCOSMOS_HOST"),
		"master",
		os.Getenv("JCOSMOS_KEY"),
		os.Getenv("JCOSMOS_DB"),
		os.Getenv("JCOSMOS_COLL"),
		jcosmos.LogLevelInfo,
		false,
		false,
		log.Default(),
	)
}

type egDoc struct {
	ID      string `json:"id"`
	Pk      string `json:"pk"`
	Message string `json:"message"`
}

func main() {
	pk := "MyPK"
	id := "MYUNIQUEID"
	doc := egDoc{
		ID:      id,
		Pk:      pk,
		Message: "Hello World",
	}
	var rDoc egDoc
	err := cosmosClient.CreateDocument(pk, false, &doc)
	if err != nil {
		log.Fatal(err)
	}
	err = cosmosClient.ReadDocument(id, pk, &rDoc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("message is ", rDoc.Message)
	doc.Message = "updated message"
	err = cosmosClient.UpdateDocument(id, pk, &doc)
	if err != nil {
		log.Fatal(err)
	}
	err = cosmosClient.ReadDocument(id, pk, &rDoc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("message is ", rDoc.Message)
	err = cosmosClient.DeleteDocument(id, pk)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("done")
}
