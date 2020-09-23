#!/bin/bash

# export GOPATH=`echo $PWD | sed -e '/\(.*\)\/src\/.*/ s//\1/'`
# go get ./
# export AWS_PROFILE=dev
# export AWS_SDK_LOAD_CONFIG=1
export AWS_REGION=us-east-1
# export PNP_SERVICE_NAME=settlement_data
#payload='{"SecretKey" : "AWS_JWT_SECRET_KEY"}'
go run main.go secrets_manager.go test
