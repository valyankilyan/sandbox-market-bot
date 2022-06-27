#!/bin/bash

cp ../../go.mod .
sudo docker build . -t marketbot-golang