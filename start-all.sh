#!/bin/bash -e

nowWorkingDir=`pwd`;

# Start Frontier
echo "[Frontier-Service] Starting up...\n"
cd "$nowWorkingDir/frontier";
docker-compose up -d
echo "[Frontier-Service] Started\n"

# Start Kafka
echo "[Kafka-Service] Starting up...\n"
cd "$nowWorkingDir/kafka";
docker-compose up -d
echo "[Kafka-Service] Started\n"

# Start Mach
echo "[Mach-Service] Starting up...\n"
cd "$nowWorkingDir/mach";
docker-compose up -d
echo "[Mach-Service] Started\n"
