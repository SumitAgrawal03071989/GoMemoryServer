
# Server execution log
➜  server go run main.go
2021/12/05 15:04:49 server listening at [::]:50052
2021/12/05 15:05:14 Received: 1
2021/12/05 15:05:19 server_memory 1
2021/12/05 15:05:19 Received: 2
2021/12/05 15:05:21 server_memory 1
2021/12/05 15:05:21 server_memory 2
2021/12/05 15:05:21 Received: 3
2021/12/05 15:05:25 server_memory 1
2021/12/05 15:05:25 server_memory 2
2021/12/05 15:05:25 server_memory 3
2021/12/05 15:05:25 Received: 1
2021/12/05 15:05:27 server_memory 1
2021/12/05 15:05:27 server_memory 2
2021/12/05 15:05:27 server_memory 3
2021/12/05 15:05:27 Received: 2
2021/12/05 15:05:29 server_memory 1
2021/12/05 15:05:29 server_memory 2
2021/12/05 15:05:29 server_memory 3
2021/12/05 15:05:29 Received: 3
2021/12/05 15:05:31 server_memory 1
2021/12/05 15:05:31 server_memory 2
2021/12/05 15:05:31 server_memory 3
2021/12/05 15:05:31 Received: 4
2021/12/05 15:05:33 server_memory 1
2021/12/05 15:05:33 server_memory 2
2021/12/05 15:05:33 server_memory 3
2021/12/05 15:05:33 server_memory 4
2021/12/05 15:05:33 Received: 1
2021/12/05 15:05:39 server_memory 1
2021/12/05 15:05:39 server_memory 2
2021/12/05 15:05:39 server_memory 3
2021/12/05 15:05:39 server_memory 4
2021/12/05 15:05:39 server_memory 1
2021/12/05 15:05:39 Received: 4
2021/12/05 15:05:42 server_memory 1
2021/12/05 15:05:42 server_memory 2
2021/12/05 15:05:42 server_memory 3
2021/12/05 15:05:42 server_memory 4
2021/12/05 15:05:42 server_memory 1
2021/12/05 15:05:42 Received: 3
2021/12/05 15:05:45 server_memory 1
2021/12/05 15:05:45 server_memory 2
2021/12/05 15:05:45 server_memory 3
2021/12/05 15:05:45 server_memory 4
2021/12/05 15:05:45 server_memory 1
2021/12/05 15:05:45 Received: 1
2021/12/05 15:05:47 server_memory 1
2021/12/05 15:05:47 server_memory 2
2021/12/05 15:05:47 server_memory 3
2021/12/05 15:05:47 server_memory 4
2021/12/05 15:05:47 server_memory 1
2021/12/05 15:05:47 Received: 2



# Client Execution log

➜  client go run main.go 1
2021/12/05 14:56:23 Greeting: It might be that you are telling me this number first time or probably I am getting old its hard for me to remember more than 3 number 1
➜  client go run main.go 1
2021/12/05 14:56:25 Greeting: Yes I remember 1
➜  client go run main.go 1
2021/12/05 14:56:31 Greeting: Yes I remember 1
➜  client go run main.go 2
2021/12/05 14:56:34 Greeting: It might be that you are telling me this number first time or probably I am getting old its hard for me to remember more than 3 number 2
➜  client go run main.go 1
2021/12/05 14:58:51 Greeting: Yes I remember 1
➜  client go run main.go 1
2021/12/05 14:59:08 Greeting: It might be that you are telling me this number first time or probably I am getting old its hard for me to remember more than 3 number 1
➜  client go run main.go 1
2021/12/05 14:59:49 Greeting: It might be that you are telling me this number first time or probably I am getting old its hard for me to remember more than 3 number 1
➜  client go run main.go 1
2021/12/05 14:59:57 Greeting: Yes I remember 1
➜  client go run main.go 4
2021/12/05 15:00:05 Greeting: It might be that you are telling me this number first time or probably I am getting old its hard for me to remember more than 3 number 4
➜  client go run main.go 8
2021/12/05 15:00:13 Greeting: It might be that you are telling me this number first time or probably I am getting old its hard for me to remember more than 3 number 8
➜  client go run main.go 8
2021/12/05 15:00:19 Greeting: Yes I remember 8
➜  client go run main.go 9
2021/12/05 15:00:25 Greeting: It might be that you are telling me this number first time or probably I am getting old its hard for me to remember more than 3 number 9
➜  client go run main.go 1
2021/12/05 15:00:32 Greeting: It might be that you are telling me this number first time or probably I am getting old its hard for me to remember more than 3 number 1
➜  client go run main.go 9
2021/12/05 15:00:38 Greeting: Yes I remember 9
➜  client go run main.go 99
2021/12/05 15:01:20 Greeting: It might be that you are telling me this number first time or probably I am getting old its hard for me to remember more than 3 number 99
➜  client go run main.go 9 
2021/12/05 15:01:23 Greeting: Yes I remember 9
➜  client go run main.go 1
2021/12/05 15:01:26 Greeting: Yes I remember 1
➜  client go run main.go 10
2021/12/05 15:01:30 Greeting: It might be that you are telling me this number first time or probably I am getting old its hard for me to remember more than 3 number 10
➜  client go run main.go 1 
^Csignal: interrupt
➜  client go run main.go 1
2021/12/05 15:05:14 Greeting: It might be that you are telling me this number first time or probably I am getting old its hard for me to remember more than 3 number, I will try to remember 1
➜  client go run main.go 2
2021/12/05 15:05:19 Greeting: It might be that you are telling me this number first time or probably I am getting old its hard for me to remember more than 3 number, I will try to remember 2
➜  client go run main.go 3
2021/12/05 15:05:21 Greeting: It might be that you are telling me this number first time or probably I am getting old its hard for me to remember more than 3 number, I will try to remember 3
➜  client go run main.go 1
2021/12/05 15:05:25 Greeting: Yes I remember 1
➜  client go run main.go 2
2021/12/05 15:05:27 Greeting: Yes I remember 2
➜  client go run main.go 3
2021/12/05 15:05:29 Greeting: Yes I remember 3
➜  client go run main.go 4
2021/12/05 15:05:31 Greeting: It might be that you are telling me this number first time or probably I am getting old its hard for me to remember more than 3 number, I will try to remember 4
➜  client go run main.go 1
2021/12/05 15:05:33 Greeting: It might be that you are telling me this number first time or probably I am getting old its hard for me to remember more than 3 number, I will try to remember 1
➜  client go run main.go 4
2021/12/05 15:05:39 Greeting: Yes I remember 4
➜  client go run main.go 3
2021/12/05 15:05:42 Greeting: Yes I remember 3
➜  client go run main.go 1
2021/12/05 15:05:45 Greeting: Yes I remember 1
➜  client go run main.go 2
2021/12/05 15:05:47 Greeting: It might be that you are telling me this number first time or probably I am getting old its hard for me to remember more than 3 number, I will try to remember 2
➜  client 