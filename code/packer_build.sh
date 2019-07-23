#!/bin/bash

cat code/gce.json
export GOOGLE_APPLICATION_CREDENTIALS=~/cred.json
packer build -force -color=false code/gce.json
