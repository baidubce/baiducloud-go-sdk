# GO SDK Documentation

# Overview

This document mainly introduces the Baidu Cloud Go language SDK (Software Development Kit). Users can use this SDK to access various Baidu Cloud products in Go (see the supported product list for details). The SDK encapsulates convenient calling interfaces, maintains similar usage methods and calling interfaces across multiple programming language versions, and provides unified error codes and return formats for easy debugging.

# Installing the SDK

## Runtime Environment

The GO SDK can run in go1.11 and above environments.

## SDK Installation

**Download directly from GitHub**

Use the `go get` tool to download from GitHub:

```shell
go get github.com/baidubce/baiducloud-go-sdk
```

**SDK Directory Structure**

```text
baiducloud-go-sdk
|--bce                    //BCE common base components
|--core               //BCE HTTP communication module
|--services               //BCE related service directories
|--sample                 //BCE related service examples
```

## Uninstalling the SDK

To uninstall the SDK, simply delete the downloaded source code.

# Usage Steps

## Confirm Endpoint

Before using the SDK, you need to confirm the Endpoint (service domain) of the Baidu Cloud product you will be accessing. Taking Baidu Object Storage as an example, you can read the [VPC Access Domain](https://cloud.baidu.com/doc/VPC/s/xjwvyuhpw) section to understand Endpoint-related concepts. Other services are similar, and you need to understand and confirm the corresponding service's Endpoint.

## Create Client Object

Each specific service has a `Client` object that encapsulates a series of easy-to-use methods for developers to interact with the corresponding service. Developers can refer to the documentation in the directory corresponding to the specific service in the SDK to use the respective service.

## Call Function Interfaces

Based on the created `Client` object for the corresponding service, developers can call the corresponding function interfaces to use Baidu Cloud product features.

## Example

Below is a basic usage example using Baidu Cloud Virtual Private Cloud (VPC) service. For detailed usage instructions, please refer to the detailed documentation of each service.

```go
import (
    "fmt"
    "bytes"
    "encoding/json"
    "github.com/baidubce/baiducloud-go-sdk/services/vpc"
    "github.com/baidubce/baiducloud-go-sdk/core/util"

)

func CreateVpc() {
    // Set the Client's Access Key ID and Secret Access Key. For details on obtaining AKSK, see: https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
    ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
    client, err := vpc.NewClient(ak, sk, endpoint)
    if err != nil {
        fmt.Println("create client err:", err)
        return
    }
    createVpcRequest := &vpc.CreateVpcRequest{
        Name : util.PtrString("vpcName"),
        Cidr : util.PtrString("10.1.1.0/24"),
    }
    result := &vpc.CreateVpcResponse{}
    result, err = client.CreateVpc(createVpcRequest)
    if err != nil {
        // This is only for display purposes. Please handle exceptions carefully and never ignore exceptions directly in engineering projects.
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.Marshal(result)
    if err != nil {
        fmt.Println("json marshal failed:", err)
        return
    }
    var out bytes.Buffer
    err = json.Indent(&out, data, "", "  ")
    if err != nil {
        fmt.Println("json indent failed:", err)
        return
    }
    fmt.Println(out.String())
}
```

# Configuration

## Using HTTPS Protocol

This SDK supports using the HTTPS protocol to access Baidu Cloud service products. To use the HTTPS protocol, you only need to specify a domain name using the https protocol in the Endpoint when creating the `Client` object for the corresponding service. The SDK will automatically recognize and use the HTTPS protocol for access.

## Detailed Configuration

When developers use the GO SDK, the exported `Config` field of the created `Client` object for the corresponding service provides the following parameters to support detailed configuration:

Configuration Item |  Type   | Description
-----------|---------|--------
Endpoint   |  string | Service request domain name
ProxyUrl   |  string | Client request proxy address
Region     |  string | Request resource region
UserAgent  |  string | User name, User-Agent header of HTTP request
Credentials| \*auth.BceCredentials | Request authentication object, divided into normal AK/SK and STS
SignOption | \*auth.SignOptions    | Authentication string signature options
Retry      | RetryPolicy | Connection retry policy
ConnectionTimeoutInMillis| int     | Connection timeout in milliseconds, default is 20 minutes

Notes:

1. The `Credentials` field is created using the `auth.NewBceCredentials` and `auth.NewSessionBceCredentials` functions. The former is used by default, and the latter is used for STS authentication.
2. The `SignOption` field is an option for generating signature strings. See the table below for details:

Name          | Type  | Description
--------------|-------|-----------
HeadersToSign |map[string]struct{} | HTTP headers used when generating signature strings
Timestamp     | int64 | Timestamp used in the generated signature string, defaults to the value at request time
ExpireSeconds | int   | Validity period of the signature string

     Among them, HeadersToSign defaults to `Host`, `Content-Type`, `Content-Length`, `Content-MD5`; TimeStamp is generally zero, indicating that the timestamp when the authentication string is generated is used. Users generally should not explicitly specify the value of this field; ExpireSeconds defaults to 1800 seconds (30 minutes).
3. The `Retry` field specifies the retry policy. Currently two types are supported: `NoRetryPolicy` and `BackOffRetryPolicy`. The latter is used by default. This retry policy specifies the maximum number of retries, maximum retry time, and retry base, and retries exponentially by multiplying the retry base by 2 until the maximum number of retries or maximum retry time is reached.


Developers can configure detailed parameters accordingly. Below are some configuration examples:

```go
// client is a `Client` object for a specific service

// Configure request proxy address
client.Config.ProxyUrl = "127.0.0.1:8080"

// Configure no retry, default is Back Off retry
client.Config.Retry = bce.NewNoRetryPolicy()

// Configure connection timeout to 30 seconds
client.Config.ConnectionTimeoutInMillis = 30 * 1000

// Configure signature to use HTTP request header as `Host`
client.Config.SignOption.HeadersToSign = map[string]struct{}{"Host": struct{}{}}

// Configure signature validity period to 30 seconds
client.Config.SignOption.ExpireSeconds = 30
```

# Error Handling

The Go language identifies errors with the error type and defines the following two error types:

Error Type        |  Description
----------------|-------------------
BceClientError  | Errors generated by user operations
BceServiceError | Errors returned by the service

When users use the SDK to call relevant interfaces of various services, in addition to returning the required results, errors will also be returned. Users can obtain detailed error information for processing. Example as follows:

```go
// vpcClient is a created Client object for VPC service
result, err = vpcClient.CreateVpc(createVpcRequest)
if err != nil {
	switch realErr := err.(type) {
	case *bce.BceClientError:
		fmt.Println("client occurs error:", realErr.Error())
	case *bce.BceServiceError:
		fmt.Println("service occurs error:", realErr.Error())
	default:
		fmt.Println("unknown error:", err)
	}
}
fmt.Println("create vpc success")
```

## Client Exceptions

Client exceptions represent exceptions encountered by the client when attempting to send requests and transfer data to Baidu Cloud services. For example, when the network connection is unavailable when sending a request, BceClientError will be returned; when an IO exception occurs when uploading a file, BceClientError will also be thrown.

## Server Exceptions

When a server exception occurs, the Baidu Cloud server will return corresponding error information to the user to help locate the problem. Each type of server exception should refer to the official documentation of each service.

## SDK Logging

The GO SDK implements its own logging module that supports six levels, three outputs (standard output, standard error, file), and basic format settings. The import path is `github.com/baidubce/baiducloud-go-sdk/core/util/log`. When outputting to a file, it supports setting five log rotation methods (no rotation, by day, by hour, by minute, by size), and in this case, you also need to set the output log file directory.

This logging module has no external dependencies. Developers using the GO SDK to develop projects can directly reference this logging module for use in their projects. You can use the package-level log object used by the GO SDK, or create a new log object. See the following example for details:

```go
// Directly use the package-level global log object (will be output together with the GO SDK's own logs)
log.SetLogHandler(log.STDERR)
log.Debugf("%s", "logging message using the log package in the sdk")

// Create a new log object (output logs according to custom settings, separate from GO SDK log output)
myLogger := log.NewLogger()
myLogger.SetLogHandler(log.FILE)
myLogger.SetLogDir("/home/log")
myLogger.SetRotateType(log.ROTATE_SIZE)
myLogger.Info("this is my own logger from the sdk")
```