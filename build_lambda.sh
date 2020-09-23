#!/bin/bash

GOOS=linux go build main.go secrets_manager.go
zip jwt.zip ./main
