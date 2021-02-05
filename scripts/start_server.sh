#!/bin/bash
echo "SERVER_PORT=80" >> /home/ec2-user/GoSample/.env 
/home/ec2-user/GoSample/hello >logs.txt 2>errors.txt &
