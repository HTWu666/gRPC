<h1 align="center">gRPC</h1>

The primary objective of this experiment is to compare the data transmission speeds of gRPC and HTTP.

## Experimental Condition

The diagram below illustrates the architecture of our experiment. Both the client and server are cloud servers in Golang. The process for a single request is as follows: A request is sent from the local k6 to the client, which then transfers a text file of approximately 1MB size to the server via gRPC or HTTP. The server responds back to the client using gRPC or HTTP, and finally, the client responds back to k6, completing the request.

![image](https://github.com/HTWu666/Restaurant-Reservation-System-Outline/assets/126232123/8caec752-7d58-46dc-a497-cae960283aa9)

The requests sent from k6 follow the stress test scheme as figure below.

![Stress Test](https://github.com/HTWu666/Restaurant-Reservation-System-Outline/assets/126232123/a355ec58-0ecc-4503-998b-af1d64ac2b44)

## Result

### http

![image](https://github.com/HTWu666/Restaurant-Reservation-System-Outline/assets/126232123/1b9e59cc-961e-4503-866a-06098efb6da5)

### gRPC

![image](https://github.com/HTWu666/OUTLiNE/assets/126232123/3a8faace-0d09-4042-913c-49f812844a0a)

## Conclusion

1. gRPC is 100 % better than http in http_req_duration

![image](https://github.com/HTWu666/Restaurant-Reservation-System-Outline/assets/126232123/3c88f9ba-1934-41ad-81ce-30153f39fbea)

![image](https://github.com/HTWu666/Restaurant-Reservation-System-Outline/assets/126232123/486fc2d2-28b8-4366-b3e4-dde8bbb0a42c)
