START %cd%/nacos/bin/startup.cmd -m standalone
START go run %cd%/backend_RPC_servers/echo
START go run %cd%/backend_RPC_servers/mathsvc
START go run %cd%/backend_RPC_servers/timesvc
START go run %cd%/hertz_gateway
EXIT
