MODULE = github.com/phiphi-tan/orbital-api-gateway

# generate client code by IDL.
.PHONY: gen
gen:
	kitex -module $(MODULE) idl/echo.thrift
	kitex -module $(MODULE) idl/mathsvc.thrift
	kitex -module $(MODULE) idl/timesvc.thrift

