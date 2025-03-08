# Commands to analyze bpf object files
llvm-objdump --section-headers loadbalancer.bpf.o
llvm-objdump -D loadbalancer.bpf.o
llvm-objdump -d loadbalancer.bpf.o