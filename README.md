This repository contains examples of using the Sumsup KYC service in Go. The examples are based on the official Sumsup Golang example at https://github.com/SumSubstance/AppTokenUsageExamples.

In addition to the examples provided by Sumsup, this repository includes a Webhook client example and test. The Webhook client example shows how to listen for events from the Sumsup KYC service using a RESTful API and a webhook. The test demonstrates how to simulate webhook events for testing purposes.

To run the examples and test, you'll need to obtain an App Token and App Secret from the Sumsup Developer Dashboard. Once you have these credentials, you can create a .env file using the .env.example template and run the examples and test using the go command:

```go
go run main.go
go test -v
```

Note that the Webhook client example requires a publicly accessible URL for the webhook endpoint. You can use a service like ngrok (https://ngrok.com/) to expose a local endpoint to the Internet for testing purposes.

For more information on using the Sumsup KYC service in Go, refer to the official documentation at https://developers.sumsub.com/api-reference/.
