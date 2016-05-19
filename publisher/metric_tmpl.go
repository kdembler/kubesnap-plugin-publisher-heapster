/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2016 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package publisher

const builtinMetricTemplate = `{
	"id":"!!",
	"name":"!!",
	"aliases":[
	],
	"labels":{
		"io.kubernetes.pod.name":"__tmpl|/labels/io_kubernetes_pod_name/value||str",
		"io.kubernetes.container.name":"__tmpl|/labels/io_kubernetes_container_name/value||str",
		"io.kubernetes.pod.uid":"__tmpl|/labels/io_kubernetes_pod_uid/value||str",
		"io.kubernetes.pod.namespace":"__tmpl|/labels/io_kubernetes_pod_namespace/value||str"
	},
	"subcontainers":[
	],
	"spec":{
		"creation_time":"__tmpl|/creation_time|2016-05-16T03:25:47Z|str",
		"labels":{
			"io.kubernetes.pod.name":"__tmpl|/labels/io_kubernetes_pod_name/value||str",
			"io.kubernetes.container.name":"__tmpl|/labels/io_kubernetes_container_name/value||str",
			"io.kubernetes.pod.uid":"__tmpl|/labels/io_kubernetes_pod_uid/value||str",
			"io.kubernetes.pod.namespace":"__tmpl|/labels/io_kubernetes_pod_namespace/value||str"
		},
		"has_cpu":true,
		"cpu":{
			"limit":2,
			"max_limit":2,
			"mask":"0-1"
		},
		"has_memory":true,
		"memory":{
			"limit":0,
			"swap_limit":0
		},
		"has_network":true,
		"has_filesystem":true,
		"has_diskio":false,
		"has_custom_metrics":false,
		"image":"__tmpl|/image_name|docker|str"
	},
	"stats":[
		{
			"timestamp":"!!",
			"cpu":{
				"usage":{
					"total":"__tmpl|/cgroups/cpu_stats/cpu_usage/total_usage|0|int",
					"user":"__tmpl|/cgroups/cpu_stats/cpu_usage/usage_in_usermode|0|int",
					"system":"__tmpl|/cgroups/cpu_stats/cpu_usage/usage_in_kernelmode|0|int"
				},
				"load_average":"__tmpl|/sched_load|0|int"
			},
			"diskio":{
			},
			"memory":{
				"usage":"__tmpl|/cgroups/memory_stats/usage/usage|0|int",
				"cache":"__tmpl|/cgroups/memory_stats/cache|0|int",
				"rss":"__tmpl|/cgroups/memory_stats/stats/rss|0|int",
				"working_set":"__tmpl|/cgroups/memory_stats/stats/working_set|0|int",
				"failcnt":0,
				"container_data":{
					"pgfault":"__tmpl|/cgroups/memory_stats/stats/pgfault|0|int",
					"pgmajfault":"__tmpl|/cgroups/memory_stats/stats/pgmajfault|0|int"
				},
				"hierarchical_data":{
					"pgfault":0,
					"pgmajfault":0
				}
			},
			"network":{
				"name":"",
				"rx_bytes":"__tmpl|/network/eth0/rx_bytes|0|int",
				"rx_packets":"__tmpl|/network/eth0/rx_packets|0|int",
				"rx_errors":"__tmpl|/network/eth0/rx_errors|0|int",
				"rx_dropped":"__tmpl|/network/eth0/rx_dropped|0|int",
				"tx_bytes":"__tmpl|/network/eth0/tx_bytes|0|int",
				"tx_packets":"__tmpl|/network/eth0/tx_packets|0|int",
				"tx_errors":"__tmpl|/network/eth0/tx_errors|0|int",
				"tx_dropped":"__tmpl|/network/eth0/tx_dropped|0|int",
				"interfaces":[
				],
				"tcp":{
					"Established":0,
					"SynSent":0,
					"SynRecv":0,
					"FinWait1":0,
					"FinWait2":0,
					"TimeWait":0,
					"Close":0,
					"CloseWait":0,
					"LastAck":0,
					"Listen":0,
					"Closing":0
				},
				"tcp6":{
					"Established":0,
					"SynSent":0,
					"SynRecv":0,
					"FinWait1":0,
					"FinWait2":0,
					"TimeWait":0,
					"Close":0,
					"CloseWait":0,
					"LastAck":0,
					"Listen":0,
					"Closing":0
				}
			},
			"filesystem":[
				{
					"device":"/dev/sdx1",
					"type":"vfs",
					"capacity":0,
					"usage":0,
					"base_usage":0,
					"available":0,
					"inodes_free":0,
					"reads_completed":0,
					"reads_merged":0,
					"sectors_read":0,
					"read_time":0,
					"writes_completed":0,
					"writes_merged":0,
					"sectors_written":0,
					"write_time":0,
					"io_in_progress":0,
					"io_time":0,
					"weighted_io_time":0
				}
			]
		}
	]
}
`
