/*
Copyright 2021 BaiLian.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cnbailian/Advanced-Go-Programming-examples/rpc/4.4.2/types"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal("dial connects error:", err)
	}
	defer conn.Close()

	client := types.NewHelloServiceClient(conn)
	res, err := client.Hello(context.TODO(), &types.HelloMessage{Value: "World!"})
	if err != nil {
		log.Fatal()
	}
	fmt.Println(res.GetValue())
}
