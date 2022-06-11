# MDBX for Go

libmdbx wrapper for Go which uses an Assembly trampoliine to bypass CGO for hot path functions. Assembly trampoline is 15x faster than CGO trampoline.

Assembly trampoline works on AMD64 and ARM64 CPUs. Other CPUs silently degrade to standard CGO calls.