#!/usr/bin/python
# Copied from https://gist.github.com/lizrice/47ad44a15cce912502f8667a403f5649
# Must run as root
from bcc import BPF

prog = """
int hello(void *ctx) {
    bpf_trace_printk("Hello world\\n");
    return 0;
}
"""

b = BPF(text=prog)
clone = b.get_syscall_fnname("read")
b.attach_kprobe(event=clone, fn_name="hello")
b.trace_print()

# This prints out a trace line every time the read system call is called

# If you rename hello() to kprobe__sys_clone() you can delete the b.attach_kprobe() line, because bcc can work
# out what event to attach this to from the function name.

