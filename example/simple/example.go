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
	Pk1     string `json:"pk1"`
	Pk2     string `json:"pk2"`
	Message string `json:"message"`
}

func main() {
	pk := "MyPK"
	id := "MYUNIQUEID"
	doc := egDoc{
		ID:      id,
		Pk:      pk,
		Pk1:     pk + "1",
		Pk2:     pk + "2",
		Message: "Hello World",
	}
	var rDoc egDoc
	err := cosmosClient.CreateDocument([]string{pk}, false, &doc)
	if err != nil {
		log.Fatal(err)
	}
	err = cosmosClient.ReadDocument(id, []string{pk}, &rDoc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("message is ", rDoc.Message)
	doc.Message = "updated message"
	err = cosmosClient.UpdateDocument(id, []string{pk}, &doc)
	if err != nil {
		log.Fatal(err)
	}
	err = cosmosClient.ReadDocument(id, []string{pk}, &rDoc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("message is ", rDoc.Message)
	// err = cosmosClient.DeleteDocument(id, pk)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("done")
}
