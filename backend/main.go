package main

import "jinshiho/server"

func main() {
	svr := server.New(9000)

	if err := svr.ListenAndServe(); err != nil {
		panic(err)
	}
}
