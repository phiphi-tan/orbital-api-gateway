<h1> [NUS ORBITAL] API Gateway based on CloudWeGo Projects</h1>

by [Phinehas Tan](https://github.com/phiphi-tan/ "also known as Phi Phi") and [Johan Soo](https://github.com/delishad21 "pronounced Yo-han")


**Project Objective**: The primary goal of this project is to implement an API Gateway that accepts HTTP requests encoded in JSON format and uses the Generic-Call feature of Kitex to translate these requests into Thrift binary format requests. The API Gateway will then forward the request to one of the backend RPC servers discovered from the registry center. 

\
**Deliverables**
1. Implement an API Gateway that accepts HTTP requests with JSON encoded bodies.
2. Use the Generic Call feature in Kitex to translate JSON requests into Thrift binary
format.
3. Integrate a load balancing mechanism to distribute requests among backend RPC
servers.
4. Integrate a service registry and discovery mechanism for the API Gateway and RPC
servers.
5. Develop backend RPC servers using Kitex for testing the API Gateway.
6. Document the project, including design decisions, implementation details, and usage
instructions.
___

**<h2>System Design:</h2>**

![image](https://github.com/phiphi-tan/orbital-api-gateway/assets/103935416/7c8e00ef-eaea-4f5d-acc3-b52bd9eed8f9)

**<h3>Running the System:</h3>**
This system was built to run on windows

> root directory is ```orbital-api-gateway```.

**Installing Nacos**

Install nacos [here](https://github.com/alibaba/nacos/releases) and extract it into the root directory

**Starting System**

start system with run.bat

> Verify that all services (echo, mathsvc, echosvc) are registered on Nacos by visiting [this page](http://127.0.0.1:8848/nacos) and looking under *Service Management*
___

**<h2>Testing System:</h2>**

The current Hertz gateway responds to HTTP POST requests with a JSON body

3 RPC services are included: echo, mathsvc, timesvc

Services can be accessed through ```/:service/:method```

**Methods and Parameters**

echo:
| method | params           |
| ------ | ---------------- |
| echo   | message (string) |

mathsvc:
| method   | params                          |
| -------- | ------------------------------- |
| add      | first (number), second (number) |
| subtract | first (number), second (number) |
| multiply | first (number), second (number) |

timesvc:
| method | biz_params                    |
| ------ | ----------------------------- |
| time   | twentyfourhour (boolean)      |
| date   | american_formatting (boolean) |
___

**<h2>Timeline:</h2>**

*13/05/23*: Made the Github Repo and this cool looking `.md` file
- attended workshops too

*17/05/23*: Implemented a basic RPC server with math services and client testing using Kitex

*22/05/23*: Implemented a basic Hertz based sever that takes in HTTP request with forms

*26/05/23*: Finished a skeleton for a full Hertz-Kitex API Gateway -> RPC server system

*01/06/23*: Implemented service discovery for the gateway using [Nacos](https://nacos.io/en-us/)

*06/06/23*: Implemented 3 RPC servers with differing services for testing

*09/06/23*: Implemented IDL parsing for Hertz Server

*12/06/23*: Finished implementation of a working system with 
- Service Discovery
- Generic Call

*15/06/23*: Finished implementation of
- JSON request handling
- Parameter Validation
- Weighted Load Balancer
___

**<h2>Curl Commands for Testing</h2>**

Echo:
```shell
curl --location --request POST 'http://127.0.0.1/echo/echo' --header 'Content-Type: application/json' --data '{"message": "Hello World!"}'
```

MathSvc:
```shell
curl --location --request POST 'http://172.19.80.1/mathsvc/add' --header 'Content-Type: application/json' --data '{"first": 1, "second": 2}'
```

TimeSvc:
```shell
curl --location --request POST 'http://172.19.80.1/timesvc/time' --header 'Content-Type: application/json' --data '{"twentyfourhour": true}'
```
___

**<h2>Adding/Updating IDLs:</h2>**

Add your IDL files under '/idl' directory. View [IDL annotation standards](https://www.cloudwego.io/docs/kitex/tutorials/advanced-feature/generic-call/thrift_idl_annotation_standards/).
Note: the name of your IDL file should be the same as the name of your service

**Hertz Scaffolding**
Under `/hertz_gateway` run these commands to generate/update hertz scaffolding for your IDL
```shell
hz update -module "github.com/phiphi-tan/orbital-api-gateway/hertz_gateway" -idl ../idl/[IDL file]
```
Update handler logic in `hertz_gateway/handler/[Service Name]

**Kitex Scaffolding**
Under the root directory, run these commands to generate kitex code for your IDL
```shell
kitex -module "github.com/phiphi-tan/orbital-api-gateway" ../idl/[IDL file]
```

**Kitex RPC Servers**
If you would like to generate a server for your RPC server under this directory, run this command under `/backend_RPC_servers/[Your Server Directory]`
```shell
kitex -module "github.com/phiphi-tan/orbital-api-gateway" -service [Service Name] -use ./../../kitex_gen ./../../idl/[IDL file]
```
Under run.bat, include a run command for your newly generated RPC server
___
\
Quote of the day:
>There's 104 days of summer vacation. And school comes along just to end it
