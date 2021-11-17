#!/usr/bin/env bash
docker run -it -d -p 8080:8080 --env-file=env.txt -v /home/honghuiqiang/log/:/var/log/yang yang-backend:1.0.0