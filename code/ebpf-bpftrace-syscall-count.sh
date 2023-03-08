#!/bin/bash

# Must run as root
bpftrace -e 'tracepoint:raw_syscalls:sys_enter { @[comm] = count(); }'
