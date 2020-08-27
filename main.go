package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/alifpay/plg/feedpb"
	"github.com/tidwall/gjson"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	c := feedpb.NewFeedsClient(conn)
	//  get client stream
	stream, err := c.Broadcast(context.Background())
	if err != nil {
		log.Fatalf("failed to call Broadcast: %v", err)
	}

	go func() {
		sig := make(chan os.Signal)
		signal.Notify(sig, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
		<-sig
		cancel()
		signal.Stop(sig)
		close(sig)
	}()

	for {
		select {
		case <-ctx.Done():
			break
		default:
			msg, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalf("failed to recieve: %v", err)
				return
			}

			httpResp, err := http.Get("https://api.ratesapi.io/api/latest")
			if err != nil {
				stream.Send(&feedpb.FeedRequest{Feed: err.Error(), ReplyTo: msg.ReplyTo})
				stream.CloseSend()
				continue
			}

			defer httpResp.Body.Close()
			data, err := ioutil.ReadAll(httpResp.Body)
			usd := gjson.GetBytes(data, "rates.USD").Num

			stream.Send(&feedpb.FeedRequest{Feed: "1 EUR = " + fmt.Sprintf("%f", usd) + " USD", ReplyTo: msg.ReplyTo})

			fmt.Println("message from server:", msg.Feed, msg.ReplyTo)
		}
	}
}
