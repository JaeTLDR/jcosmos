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
		jcosmos.LogLevelTrace,
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
	pks := []string{
		pk + "",
		pk + "1",
		pk + "2",
	}
	id := "MYUNIQUEID"
	doc := egDoc{
		ID:      id,
		Pk:      pks[0],
		Pk1:     pks[1],
		Pk2:     pks[2],
		Message: "Hello World",
	}
	var rDoc egDoc
	err := cosmosClient.CreateDocument(pks, false, &doc)
	if err != nil {
		log.Fatal(err)
	}
	err = cosmosClient.ReadDocument(id, pks, &rDoc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("message is ", rDoc.Message)
	doc.Message = "updated message"
	err = cosmosClient.UpdateDocument(id, pks, &doc)
	if err != nil {
		log.Fatal(err)
	}
	err = cosmosClient.ReadDocument(id, pks, &rDoc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("message is ", rDoc.Message)
	// err = cosmosClient.DeleteDocument(id, pks)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("done")
}
