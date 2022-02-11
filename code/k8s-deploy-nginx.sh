#!/bin/bash

kubectl apply -f ./code/k8s-deployment-nginx.yml
#kubectl port-forward -n nginx service/nginx 3000:80
