package main

import (
	"log"
	"net/http"

	pb "github.com/pwlae/playground/rpc/protobuf/image/v1"
	"google.golang.org/protobuf/proto"
)

func main() {
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		image := &pb.Image{}

		image.Original = "https://s3.com/original.png"
		image.Small = "https://s3.com/small.png"
		image.Medium = "https://s3.com/medium.png"
		image.Large = "https://s3.com/large.png"

		out, err := proto.Marshal(image)
		if err != nil {
			log.Fatalln("Failed to encode image:", err)
		}

		w.Write(out)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
