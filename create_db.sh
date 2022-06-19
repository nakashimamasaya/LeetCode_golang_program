#!/bin/bash
docker compose up -d mysql
sleep 50
docker compose up -d
docker compose logs
