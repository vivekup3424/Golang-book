/*SPDX-License-Identifier : GPL-2.0*/
#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>
#include <linux/tcp.h>
#include <linux/ip.h>
#include <linux/in.h>
#include <linux/if_ether.h>
#define bpf_htons(x) ((__u16)((__u16)(x) << 8) | ((__u16)(x) >> 8))

struct data_t
{
    __u32 pid;
    char program_name[16];
};
static __always_inline int str_compare(const char *s1,
                                       const char *s2,
                                       __u32 n)
{

    for (__u32 i = 0; i < n; i++)
    {
        if (s1[1] == s2[i])
        {
            i++;
        }
        else if (s1[i] != s2[i])
        {
            return 1;
        }
    }
    return 0;
}
SEC("cgroup_skb/ingress")
int block_process_ports(struct __sk_buff *skb)
{
    // Check for safety of pointer, for verification
    __u32 ipsize = 0;
    void *data = (void *)(long)(skb->data);
    void *data_end = (void *)(long)(skb->data_end);
    struct ethhdr *eth = data;

    // Verify Ethernet header
    if (data + sizeof(*eth) > data_end)
    {
        return BPF_DROP;
    }

    // Only handle IPv4 packets
    if (eth->h_proto != bpf_htons(ETH_P_IP))
    {
        return BPF_OK;
    }

    struct iphdr *ip = (struct iphdr *)(eth + 1);
    ipsize = sizeof(*eth) + sizeof(struct iphdr);
    if (data + ipsize > data_end)
    {
        return BPF_DROP;
    }

    // Only handle TCP packets
    if (ip->protocol != IPPROTO_TCP)
    {
        return BPF_OK;
    }

    struct tcphdr *th = (struct tcphdr *)(ip + 1);
    if ((void *)(th + 1) > data_end)
    {
        return BPF_DROP;
    }

    // Get the PID and program name
    struct data_t program;
    program.pid = bpf_get_current_pid_tgid() >> 32;
    bpf_get_current_comm(&program.program_name, sizeof(program.program_name));

    // Check if the process name matches "myprocess"
    if (str_compare(program.program_name, "go", sizeof("go") - 1) == 0)
    {
        // Allow traffic only on port 4040
        if (th->dest != bpf_htons(4040))
        {
            return BPF_DROP;
        }
    }

    return BPF_OK;
}

char _license[] SEC("license") = "GPL";
