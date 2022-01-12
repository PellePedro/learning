# OVS Cheet Sheet

## VLAN trunk
```
#  |A|\                       /|C|
#      |ovs-br0|-----|ovs-br1|
#  |B|/                       \|D|

# netns
ip netns add A
ip netns add B
ip netns add C
ip netns add D

# netns - openvswitch
ip link add a0 type veth peer name ovs0-0
ip link add b0 type veth peer name ovs0-1
ip link add c0 type veth peer name ovs1-0
ip link add d0 type veth peer name ovs1-1

# netns
ip link set a0 netns A up
ip link set b0 netns B up
ip link set c0 netns C up
ip link set d0 netns D up

# IP
ip netns exec A ip addr add 1.0.0.1/24 dev a0
ip netns exec B ip addr add 1.0.0.2/24 dev b0
ip netns exec C ip addr add 1.0.0.3/24 dev c0
ip netns exec D ip addr add 1.0.0.4/24 dev d0

# openvswitch bridges
ovs-vsctl add-br ovs-br0
ovs-vsctl add-br ovs-br1

#
ovs-vsctl add-port ovs-br0 ovs0-0 tag=10
ovs-vsctl add-port ovs-br0 ovs0-1 tag=20
ovs-vsctl add-port ovs-br1 ovs0-0 tag=10
ovs-vsctl add-port ovs-br1 ovs1-0 tag=10
ovs-vsctl add-port ovs-br1 ovs1-1 tag=20

#
ip link add t0 type veth peer name t1

#
ovs-vsctl add-port ovs-br0 t0 trunk=10,20
ovs-vsctl add-port ovs-br1 t1 trunk=10,20

ip netns exec A ping 1.0.0.3
ip netns exec B ping 1.0.0.4

ip netns exec A ping 1.0.0.4
ip netns exec B ping 1.0.0.3

```

# frequent used command
```

ovs-dpctl show -s
ovs-ofctl show, dump-ports, dump-flows, add-flow, mod-flows, del-flows
ovsdb-tools show-log -m
ovs-vsctl
  show - show the ovsdb
  bridge -  add-br, list-br, del-br, br-exists.
  port   -  list-ports, add-port, del-port, add-bond, port-to-br.
  interface -  list-ifaces, iface-to-br
  ovs-vsctl list/set/get/add/remove/clear/destroy table record column [value], tables like "bridge", "controller","interface","mirror","netflow","open_vswitch","port","qos","queue","ssl","sflow".
  ovs-appctl list-commands, fdb/show, qos/show
  
  *Find all bridges in ovs*
  ovs-vsctl list-br
  br-eth1
br-ex
br-int
br-tun

*Port belong to which bridge*

root@openstack:~# ovs-vsctl port-to-br "qg-496054f8-ed"
br-ex

*Dump open flows on 1 bridge*
ovs-ofctl dump-flows br-ex

*Find the portid port_name mapping*
 ovs-ofctl show br-ex
 OFPT_FEATURES_REPLY (xid=0x2): dpid:00000015179267ea
n_tables:254, n_buffers:256
capabilities: FLOW_STATS TABLE_STATS PORT_STATS QUEUE_STATS ARP_MATCH_IP
actions: OUTPUT SET_VLAN_VID SET_VLAN_PCP STRIP_VLAN SET_DL_SRC SET_DL_DST SET_NW_SRC SET_NW_DST SET_NW_TOS SET_TP_SRC SET_TP_DST ENQUEUE
 1(eth0): addr:00:15:17:92:67:ea
     config:     0
     state:      0
     current:    1GB-FD COPPER AUTO_NEG
     advertised: 10MB-HD 10MB-FD 100MB-HD 100MB-FD 1GB-FD COPPER AUTO_NEG
     supported:  10MB-HD 10MB-FD 100MB-HD 100MB-FD 1GB-FD COPPER AUTO_NEG
     speed: 1000 Mbps now, 1000 Mbps max
 2(qg-72bd35e9-ef): addr:00:00:00:00:00:00
     config:     PORT_DOWN
     state:      LINK_DOWN
     speed: 0 Mbps now, 0 Mbps max
 30(qg-a11b2dbd-70): addr:00:00:00:00:00:00
     config:     PORT_DOWN
     state:      LINK_DOWN
     speed: 0 Mbps now, 0 Mbps max
 41(qg-496054f8-ed): addr:00:00:00:00:00:00
     config:     PORT_DOWN
     state:      LINK_DOWN
     speed: 0 Mbps now, 0 Mbps max
 43(qg-385c6c99-ba): addr:00:00:00:00:00:00
     config:     PORT_DOWN
     state:      LINK_DOWN
     speed: 0 Mbps now, 0 Mbps max
 LOCAL(br-ex): addr:00:15:17:92:67:ea
     config:     0
     state:      0
     speed: 0 Mbps now, 0 Mbps max
OFPT_GET_CONFIG_REPLY (xid=0x4): frags=normal miss_send_len=0

*View port stat*
root@openstack:~# ovs-dpctl show
system@ovs-system:
        lookups: hit:124186956 missed:12639251 lost:0
        flows: 39
        port 0: ovs-system (internal)
        port 1: br-eth1 (internal)
        port 2: br-ex (internal)
        port 3: eth0
        port 4: br-tun (internal)
        port 5: qvo08a69421-b7
        port 6: tap30e070dd-44 (internal)
        port 7: tapc40f1d22-99 (internal)
        port 8: tap7f052b44-f3 (internal)
        port 9: br-int (internal)
        port 10: tap788443ef-f3 (internal)
        port 11: tap4c5caa08-a0 (internal)
        port 12: tap651a559c-f8 (internal)
        port 13: tapf2058b84-ff (internal)
        port 14: tapfaab6341-0c (internal)
        port 15: qr-9d205827-91 (internal)
        port 16: qg-72bd35e9-ef (internal)
        port 17: tapf422c067-6f (internal)
        port 18: qvoca912ee1-7b
        port 19: qg-a11b2dbd-70 (internal)
        port 20: tap7b8a1621-e4 (internal)
        port 21: qvof8ec5847-f4
        port 22: qvo2c1dce02-b8
        port 23: qvo4b9194da-f1
        port 24: qvo38f7cdd8-6e
        port 25: qvo95df97b6-a2
        port 26: qr-90f302be-a2 (internal)
        port 27: qvoba75ecb6-d7
        port 28: tap60cfc6df-82 (internal)
        port 29: tapdc18468c-52 (internal)
        port 30: qvoc9e43ae4-7c
        port 31: qg-496054f8-ed (internal)
        port 32: qvo7a630380-99
        port 33: qvoe2326bbe-6c
        port 34: qvo404b9071-6d
        port 35: qvo0925b440-53
        port 36: qvobc9b001a-5c
        port 37: qvob67002bb-da
        port 38: qvofcdb6a9c-e5
        port 39: qr-d1434069-b0 (internal)
        port 40: qvod60d6b5d-b2
        port 41: qvob37b77bb-bc
        port 42: dummy0
        port 46: tap65329559-ed (internal)
        port 47: qg-385c6c99-ba (internal)
        port 48: tapfc1372b9-51 (internal)
        port 49: qvo2df99bfa-d7
        port 50: qvo2e6dff89-29
        port 51: qvo63d6265f-23
        port 52: qvod6236a87-81
        port 53: qvob216027e-9a
        port 54: qvo4f88d853-99
        port 55: qr-3b5a4828-c1 (internal)
*view the openflow for 1 port*
ovs-ofctl dump-ports br [port]
root@openstack:~# ovs-ofctl dump-ports br-ex br-ex
OFPST_PORT reply (xid=0x4): 1 ports
  port LOCAL: rx pkts=21306010, bytes=17596300096, drop=0, errs=0, frame=0, over=0, crc=0
           tx pkts=35898428, bytes=31761622928, drop=0, errs=0, coll=0
           
*View fdb for certain bridge*
  ovs-appctl fdb/show br-ex
  root@openstack:~# ovs-appctl fdb/show br-ex
 port  VLAN  MAC                Age
   30     0  fa:16:3e:42:49:25  226
    1     0  00:0c:29:f9:3d:f0  153
    1     0  00:21:b7:f1:c3:21   92
    1     0  80:ee:73:b3:98:38   83
   41     0  fa:16:3e:d3:3e:2c   83
    1     0  f8:ca:b8:22:ab:55   17
    1     0  00:1e:c9:4c:4d:fb   14
    1     0  34:e6:d7:07:19:0e   14
    1     0  00:25:64:b7:e1:99    6
    1     0  00:21:70:9d:b4:02    3
    1     0  f8:b1:56:5b:18:db    0
LOCAL     0  00:15:17:92:67:ea    0

*Watch openflow match*
watch -d -n 1 "ovs-ofctl dump-flows <bridge>"

Every 1.0s: ovs-ofctl dump-flows br-ex                                                        Mon May 16 18:34:16 2016

NXST_FLOW reply (xid=0x4):
 cookie=0x0, duration=3038302.945s, table=0, n_packets=60158083, n_bytes=50818748765, idle_age=0, hard_age=65534, prio
rity=0 actions=NORMAL

*set mirror port*
ovs-vsctl -- --id=@m create mirror name=you_mirror_name -- add bridge br-int mirrors @m

root@openstack:~# ovs-vsctl list port *qvo79ee7320-da*
_uuid               : 9cbdda7a-44c4-4e7e-b21b-b19d5f013fac
bond_downdelay      : 0
bond_fake_iface     : false
bond_mode           : []
bond_updelay        : 0
external_ids        : {}
fake_bridge         : false
interfaces          : [f057d0b8-b2c7-407f-8ecf-c5632ce106e6]
lacp                : []
mac                 : []
name                : "*qvo79ee7320-da*"
other_config        : {}
qos                 : []
statistics          : {}
status              : {}
tag                 : 84
trunks              : []
vlan_mode           : []

ovs-vsctl set mirror you_mirror_name  output_port=9cbdda7a-44c4-4e7e-b21b-b19d5f013fac

set mirror you_mirror_name select_all=1

*Check ovs logs*
ovsdb-tool show-log -m
record 0: "Open_vSwitch" schema, version="7.3.0", cksum="2483452374 20182"

record 1: 2016-03-09 03:36:15.503 "ovs-vsctl: ovs-vsctl --no-wait -- init -- set Open_vSwitch . db-version=7.3.0"
        table Open_vSwitch insert row 98524532:

record 2: 2016-03-09 03:36:15.506 "ovs-vsctl: ovs-vsctl --no-wait set Open_vSwitch . ovs-version=2.0.2 "external-ids:s
ystem-id=\"23733e54-fdf5-4960-ad33-e7c8efa64166\"" "system-type=\"Ubuntu\"" "system-version=\"14.04-trusty\"""
        table Open_vSwitch row 98524532 (98524532):

record 3: 2016-03-09 03:36:18.086 "ovs-vsctl: ovs-vsctl add-br br-int"
        table Port insert row "br-int" (19763c11):
        table Interface insert row "br-int" (c24e9a47):
        table Bridge insert row "br-int" (183b4bdd):
        table Open_vSwitch row 98524532 (98524532):

record 4: 2016-03-09 03:36:18.093
        table Interface row "br-int" (c24e9a47):
        table Open_vSwitch row 98524532 (98524532):

record 5: 2016-03-09 03:36:18.095 "ovs-vsctl: ovs-vsctl add-br br-eth1"
        table Port insert row "br-eth1" (7958a484):
        table Bridge insert row "br-eth1" (d8421435):
        table Interface insert row "br-eth1" (b85884d6):
        table Open_vSwitch row 98524532 (98524532):

record 6: 2016-03-09 03:36:18.097
        table Interface row "br-eth1" (b85884d6):
        table Open_vSwitch row 98524532 (98524532):

record 7: 2016-03-09 03:36:18.099 "ovs-vsctl: ovs-vsctl add-br br-ex"
        table Port insert row "br-ex" (8ae5b57b):

```


# OVS cheat sheet
## DB
```
ovs-vsctl list open_vswitch
ovs-vsctl list interface
ovs-vsctl list interface vxlan-ac000344
ovs-vsctl --columns=options list interface vxlan-ac000344
ovs-vsctl --columns=ofport,name list Interface
ovs-vsctl --columns=ofport,name --format=table list Interface
ovs-vsctl -f csv --no-heading --columns=_uuid list controller
ovs-vsctl -f csv --no-heading -d bare --columns=other_config list port
ovs-vsctl --format=table --columns=name,mac_in_use find Interface name=br-dpdk1
ovs-vsctl get interface vhub656c3cb-23 name

ovs-vsctl set port vlan1729 tag=1729
ovs-vsctl get port vlan1729 tag
ovs-vsctl remove port vlan1729 tag 1729

# not sure this is best
ovs-vsctl set interface vlan1729 mac='5c\:b9\:01\:8d\:3e\:9d'

ovs-vsctl clear Bridge br0 stp_enable

ovs-vsctl --may-exist add-br br0 -- set bridge br0 datapath_type=netdev
ovs-vsctl --if-exists del-br br0

```

## Flows
```
ovs-ofctl dump-flows br-int

# include hidden flows
ovs-appctl bridge/dump-flows br0

# remove stats on older versions that don't have --no-stats
ovs-ofctl dump-flows br-int | cut -d',' -f3,6,7-
ovs-ofctl -O OpenFlow13 dump-flows br-int | cut -d',' -f3,6,7-

ovs-appctl dpif/show
ovs-ofctl show br-int | egrep "^ [0-9]"

ovs-ofctl add-flow brbm priority=1,in_port=11,dl_src=00:05:95:41:ec:8c/ff:ff:ff:ff:ff:ff,actions=drop
ovs-ofctl --strict del-flows brbm priority=0,in_port=11,dl_src=00:05:95:41:ec:8c

# Kernel Datapath

ovs-dpctl dump-flows
ovs-appctl dpctl/dump-flows
ovs-appctl dpctl/dump-flows system@ovs-system
ovs-appctl dpctl/dump-flows netdev@ovs-netdev
```

# DPDK
```
ovs-appctl dpif/show
ovs-ofctl dump-ports br-int
ovs-appctl dpctl/dump-flows
ovs-appctl dpctl/show --statistics
ovs-appctl dpif-netdev/pmd-stats-show
ovs-appctl dpif-netdev/pmd-stats-clear
ovs-appctl dpif-netdev/pmd-rxq-show
```


# Debug log
```
ovs-appctl vlog/list | grep dpdk
ovs-appctl vlog/set dpdk:file:dbg

# log openflow
ovs-appctl vlog/set vconn:file:dbg
```

# Misc
```
ovs-appctl list-commands
ovs-appctl fdb/show brbm

ovs-appctl ofproto/trace br-int in_port=6

ovs-appctl ofproto/trace br-int tcp,in_port=3,vlan_tci=0x0000,dl_src=fa:16:3e:8d:26:61,dl_dst=fa:16:3e:0d:f5:e6,nw_src=10.0.0.26,nw_dst=10.0.0.9,nw_tos=0,nw_ecn=0,nw_ttl=0,tp_src=0,tp_dst=22,tcp_flags=0

# history
ovsdb-tool -mm show-log /etc/openvswitch/conf.db

top -p `pidof ovs-vswitchd` -H -d1

# port and dp cache stats
ovs-appctl dpctl/show -s
ovs-appctl memory/show
ovs-appctl upcall/show
ovs-appctl coverage/show
```



# neutron ml2/ovs tracing
```
PORT=tapfdd73231-29
tag=$(ovs-vsctl get port $PORT tag)
ofport=$(ovs-vsctl get interface $PORT ofport)
mac=$(ovs-vsctl get interface $PORT external_ids:attached-mac | sed -e 's/"//g')

# will flood to all tunnels
ovs-appctl ofproto/trace br-int in_port=${ofport},dl_src=${mac}

# unicast
dhcp_mac=fa:16:3e:46:07:82
ovs-appctl ofproto/trace br-int in_port=${ofport},dl_src=${mac},dl_dst=${dhcp_mac}

# Inbound from tunnel
ovs-ofctl show br-tun | grep -E "^ [0-9]"
tun_ofport=2
tun_id=$(ovs-vsctl get port $PORT other_config:segmentation_id | sed -e 's/"//g')
ovs-appctl ofproto/trace br-tun in_port=${tun_ofport},dl_src=${dhcp_mac},dl_dst=${mac},tun_id=$tun_id

```
