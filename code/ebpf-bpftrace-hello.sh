#!/bin/bash

# Must run as root
bpftrace -e 'BEGIN { printf("hello world\n"); }'
