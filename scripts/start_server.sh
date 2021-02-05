#!/bin/bash
env_file="/home/ec2-user/GoSample/.env"

if [ ! -e "$env_file" ]
then
    echo "creating .env file..."
    echo "SERVER_PORT=80" >> "$env_file"
else
    echo "no need to create env file"
fi

/home/ec2-user/GoSample/hello >logs.txt 2>errors.txt &
