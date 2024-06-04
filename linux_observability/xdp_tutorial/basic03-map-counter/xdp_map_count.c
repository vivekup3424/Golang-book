#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>
#define XDP_ACTION_MAX (XDP_REDIRECT + 1)

struct datarec
{
    __u64 rx_packets;
    __u32 byte_counter;
};
struct
{
    __uint(type, BPF_MAP_TYPE_ARRAY);
    __type(key, __u32);
    __type(value, struct datarec);
    __uint(max_enteries, XDP_ACTION_MAX);
} xdp_stats_map SEC(".maps");

/**
 * If lock_xadd is not already defined, it defines it using
 * the __sync_fetch_and_add built-in function, which performs
 * an atomic addition. This ensures safe increment operations
 * in a multi-CPU environment.
 */
#ifndef lock_xadd
#define lock_xadd(ptr, val) ((void)__sync_fetch_and_add(ptr, val))
#endif

SEC("xdp")
int xdp_stats1(struct xdp_md *ctx)
{
    void *data = (void *)(long)ctx->data;
    void *data_end = (void *)(long)ctx->data_end;
}