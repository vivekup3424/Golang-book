/*SPDX-License-Identifier : GPL 2.0*/
#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>
#include <linux/tcp.h>
#include <linux/ip.h>
#include <linux/in.h>
#include <linux/if_ether.h>
#include <unistd.h>
SEC("cgroup_skb/ingress")

struct data_t
{
    __u32 pid;
    char progran_name[16]
};
struct bpf_map_def SEC("maps") packets = {
    .type = BPF_MAP_TYPE_ARRAY,
    .key_size = sizeof(__u32),
    .value_size = sizeof(__u64),
    .max_entries = 1,
};

SEC("cgroup_skb/egress")
int block_process_ports(struct __sk_buff *skb)
{
    // check for safety of pointer, for verication
    __u32 ipsize = 0;
    void *data = (void *)(long)(skb->data);
    void *data_end = (void *)(long)(skb->data_end);
    struct iphdr *ip;
    struct tcphdr *th;
    struct ethhdr *eth = data;
    ipsize = sizeof(*eth);
    ip = data + ipsize;
    ipsize += sizeof(struct iphdr);
    if (data + ipsize > data_end)
    {
        return BPF_DROP;
    }

    // getting the pid
    struct data_t program;
    program.pid = bpf_get_current_pid_tgid() >> 32;
    // bpf_get_current_comm loads the current executable name
    bpf_get_current_comm(&program.progran_name, sizeof(program.progran_name));

    // check if the process name matches "myprocess"
    if (__builtin_memcmp(program.progran_name,
                         "myprocess",
                         sizeof("myprocess") - 1) == 0)
    {
        // check the protocol on the packets
        if (skb->protocol != IPPROTO_TCP)
        {
            return BPF_DROP;
        }
        else
        {
            th = (struct tcphdr *)(ip + 1);
            if ((void *)(th + 1) > data_end)
            {
                return BPF_DROP;
            }
            else if (th->dest != htons(4040))
            {
                return BPF_DROP;
            }
        }
    }
    return BPF_OK;
}
char _license[] SEC("license") = "GPL";