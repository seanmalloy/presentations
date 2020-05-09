#!/bin/bash

node_exporter &
prometheus --config.file="./code/prometheus.yml" &
alertmanager --config.file="./code/alertmanager.yml" &

sleep 15
echo "Prometheus http://localhost:9090"
echo "Alertmanager http://localhost:9093"
