#!/bin/bash

kind create cluster  --image kindest/node:v1.23.3 --wait 60s
kubectl get pods -A
