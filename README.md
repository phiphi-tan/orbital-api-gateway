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

**<h3>System Design:</h3>**

![image](https://github.com/phiphi-tan/orbital-api-gateway/assets/103935416/7c8e00ef-eaea-4f5d-acc3-b52bd9eed8f9)

**<h3>Running the System:</h3>**
This system was built to run on windows

> root directory is ```orbital-api-gateway```.

**Installing Nacos**

Install nacos [here](https://github.com/alibaba/nacos/releases) and extract it into the root directory

**Starting System**

start system with run.bat

> Verify that all services (echo, mathsvc, echosvc) are registered on Nacos by visiting [this page](http://127.0.0.1:8848/nacos) and looking under *Service Management*

**<h3>Testing System:</h3>**

The current Hertz gateway responds to HTTP POST requests

3 RPC services are included: echo, mathsvc, timesvc

Services can be accessed through /echo, /mathsvc, /timesvc and parameters are sent through a http form (to be changed to JSON)

**Methods and Parameters**

echo:
| method | biz_params |
| ------ | ---------- |
| echo   | message    |

mathsvc:
| method   | biz_params    |
| -------- | ------------- |
| add      | first, second |
| subtract | first, second |
| multiply | first, second |

timesvc:
| method | biz_params          |
| ------ | ------------------- |
| time   | twentyfourhour      |
| date   | american_formatting |

**<h3>Timeline:</h3>**

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


\
Quote of the day:
>There's 104 days of summer vacation. And school comes along just to end it
