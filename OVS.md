# OVS Cheet Sheet
## Links
[Faucet SDN Controller](https://github.com/faucetsdn/faucet/blob/master/docker/runtests.sh)<BR/>

## OVS Controllers
```
ovs-vsctl add-br hwbr &&
ovs-vsctl set-controller hwbr tcp:127.0.0.1:6653 tcp:127.0.0.1:6654 

```
## Magma
[Magma](https://github.com/magma/magma/blob/master/lte/gateway/python/scripts/setup_metering_ovs)
```
#!/bin/bash
# simple script to set up a test environment for metering.
# two namespaces connected by a virtual switch, pointing at a local OF
# controller.
sudo ip netns add left
sudo ip netns add right
sudo ip link add veth01 type veth peer name veth10
sudo ip link set veth01 netns left                                                                                                                                                             
sudo ip link set veth10 up
sudo ip link add veth02 type veth peer name veth20                                                                                                                                             
sudo ip link set veth02 netns right                                                                                                                                                            
sudo ip link set veth20 up                                                                                                                                                                     
sudo ovs-vsctl add-br vswitch
sudo ovs-vsctl set-fail-mode vswitch secure                                                                                                                                                    
sudo ovs-vsctl add-port vswitch veth10
sudo ovs-vsctl add-port vswitch veth20
sudo ip netns exec left ifconfig veth01 192.168.201.1 up                                                                                                                                       
sudo ip netns exec left route add -net 192.168.202.0/24 gw 192.168.201.1
sudo ip netns exec right ifconfig veth02 192.168.202.1 up
sudo ip netns exec right route add -net 192.168.201.0/24 gw 192.168.202.1
# set the controller to a default localhost controller
sudo ovs-vsctl set-controller vswitch tcp:127.0.0.1:6633
# set the vswitch to use openflow 1.0 and 1.4
# we need 1.0 apparently for ovs-ofctl to work, and flow_stats
# doesn't work in 1.5
sudo ovs-vsctl set bridge vswitch protocols=OpenFlow10,OpenFlow14
```

```
#!/bin/bash
# simple script to set up a test environment for metering.
# two namespaces connected by a virtual switch, pointing at a local OF
# controller.
sudo ip netns add left
sudo ip netns add right
sudo ip link add veth01 type veth peer name veth10
sudo ip link set veth01 netns left                                                                                                                                                             
sudo ip link set veth10 up
sudo ip link add veth02 type veth peer name veth20                                                                                                                                             
sudo ip link set veth02 netns right                                                                                                                                                            
sudo ip link set veth20 up                                                                                                                                                                     
sudo ovs-vsctl add-br vswitch
sudo ovs-vsctl set-fail-mode vswitch secure                                                                                                                                                    
sudo ovs-vsctl add-port vswitch veth10
sudo ovs-vsctl add-port vswitch veth20
sudo ip netns exec left ifconfig veth01 192.168.201.1 up                                                                                                                                       
sudo ip netns exec left route add -net 192.168.202.0/24 gw 192.168.201.1
sudo ip netns exec right ifconfig veth02 192.168.202.1 up
sudo ip netns exec right route add -net 192.168.201.0/24 gw 192.168.202.1
# set the controller to a default localhost controller
sudo ovs-vsctl set-controller vswitch tcp:127.0.0.1:6633
# set the vswitch to use openflow 1.0 and 1.4
# we need 1.0 apparently for ovs-ofctl to work, and flow_stats
# doesn't work in 1.5
sudo ovs-vsctl set bridge vswitch protocols=OpenFlow10,OpenFlow14

```

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



## Create a bridge by OpenvSwitch and bind N tap device
```
#!/bin/bash

BR=br0                          # Bridge name
DPID=0000000000000001           # Bridge DPID
IPADDR=10.168.1.1/16            # IP address of the internel port of the bridge (ie. Bridge IP Address)
LINKNUM=3                       # Igress Link number of the bridge
HOSTNUM=3                       # Egress Link (Host) number of the bridge
CTRLIP=127.0.0.1                # Controller IP address
CTRLPORT=6653                   # OpenFlow listening port
OPENFLOW=OpenFlow10             # OpenFlow version


if [ `id -u` != "0" ]; then
    echo "*** This script must run as root ***"
    exit
fi

# Create bridge
ovs-vsctl --if-exist del-br $BR
ovs-vsctl --may-exist add-br $BR \
    -- set Bridge $BR fail-mode=secure \
    -- set Bridge $BR other-config:datapath-id=$DPID \
    -- set Bridge $BR protocols=$OPENFLOW \
    -- set Bridge $BR stp_enable=true \
    -- set-controller $BR tcp:$CTRLIP:$CTRLPORT \
    -- set controller $BR connection-mode=out-of-band

#ovs-vsctl set bridge $BR protocols=$OPENFLOW
#ovs-vsctl set bridge $BR stp_enable=true
#ovs-vsctl set-controller $BR tcp:$CTRLIP:$CTRLPORT
#ip address add $IPADDR dev $BR

# Create tap interfaces and bind each to the corresponding port
for ((i=0; i<$LINKNUM; i++)); do
    TAP=tap$i
    ip tuntap del mode tap $TAP
    ip tuntap add mode tap $TAP
    ip link set $TAP up
    ovs-vsctl add-port $BR $TAP -- set interface $TAP ofport_request=$(( i + HOSTNUM + 1 ))
done

# Up bridge
ip link set $BR up

echo "------------------------------ Bridge ------------------------------"
ovs-vsctl -- --columns=name,datapath_id,ports,fail_mode,protocols,stp_enable list Bridge $BR
echo "--------------------------------------------------------------------"

echo "----------------------------- Interface ----------------------------"
ovs-vsctl -- --columns=name,ofport,link_state,link_speed list Interface
echo "--------------------------------------------------------------------"

echo "----------------------------- Controller ----------------------------"
ovs-vsctl -- --columns=connection_mode,target list Controller
echo "--------------------------------------------------------------------"

echo "----------------------------- Flows ----------------------------"
ovs-ofctl dump-flows $BR
```


## docker-quagga.sh
```
#!/bin/bash
# clean
docker rm -f h1 h2 s1 s2 s3

# hosts
docker run -ti -d --name h1 --net none --privileged snlab/dovs
docker run -ti -d --name h2 --net none --privileged snlab/dovs

# switches
docker run -ti -d --name s1 --privileged snlab/dovs-quagga
docker run -ti -d --name s2 --privileged snlab/dovs-quagga
docker run -ti -d --name s3 --privileged snlab/dovs-quagga

# get network namespace
h1ns=$(docker inspect --format '{{.State.Pid}}' h1)
h2ns=$(docker inspect --format '{{.State.Pid}}' h2)
s1ns=$(docker inspect --format '{{.State.Pid}}' s1)
s2ns=$(docker inspect --format '{{.State.Pid}}' s2)
s3ns=$(docker inspect --format '{{.State.Pid}}' s3)

# add links
nsenter -t $h1ns -n ip link add e0 type veth peer name e0 netns $s1ns
nsenter -t $s1ns -n ip link add e1 type veth peer name e1 netns $s3ns
nsenter -t $s1ns -n ip link add e2 type veth peer name e0 netns $s2ns
nsenter -t $s2ns -n ip link add e1 type veth peer name e0 netns $s3ns
nsenter -t $s3ns -n ip link add e2 type veth peer name e0 netns $h2ns

# configure ovs
docker exec s1 service openvswitch-switch start
docker exec s1 ovs-vsctl add-br s
docker exec s1 ovs-vsctl add-port s e0
docker exec s1 ovs-vsctl add-port s e1
docker exec s1 ovs-vsctl add-port s e2
docker exec s1 ovs-vsctl add-port s i0 -- set interface i0 type=internal
docker exec s1 ovs-vsctl add-port s i1 -- set interface i1 type=internal
docker exec s1 ovs-vsctl add-port s i2 -- set interface i2 type=internal
docker exec s1 ifconfig r
docker exec s1 ifconfig i1 10.0.2.1/24
docker exec s1 ifconfig i2 10.0.3.1/24

docker exec s2 service openvswitch-switch start
docker exec s2 ovs-vsctl add-br s
docker exec s2 ovs-vsctl add-port s e0
docker exec s2 ovs-vsctl add-port s e1
docker exec s2 ovs-vsctl add-port s i0 -- set interface i0 type=internal
docker exec s2 ovs-vsctl add-port s i1 -- set interface i1 type=internal
docker exec s2 ifconfig i0 10.0.3.2/24
docker exec s2 ifconfig i1 10.0.4.1/24
docker exec s2 ifconfig e0 0.0.0.0
docker exec s2 ifconfig e1 0.0.0.0

docker exec s3 service openvswitch-switch start
docker exec s3 ovs-vsctl add-br s
docker exec s3 ovs-vsctl add-port s e0
docker exec s3 ovs-vsctl add-port s e1
docker exec s3 ovs-vsctl add-port s e2
docker exec s3 ovs-vsctl add-port s i0 -- set interface i0 type=internal
docker exec s3 ovs-vsctl add-port s i1 -- set interface i1 type=internal
docker exec s3 ovs-vsctl add-port s i2 -- set interface i2 type=internal
docker exec s3 ifconfig i0 10.0.4.2/24
docker exec s3 ifconfig i1 10.0.2.2/24
docker exec s3 ifconfig i2 10.0.1.254/24
docker exec s3 ifconfig e0 0.0.0.0
docker exec s3 ifconfig e1 0.0.0.0
docker exec s3 ifconfig e2 0.0.0.0

# bring up ethx ports
docker exec s1 ifconfig e0 0.0.0.0
docker exec s1 ifconfig e1 0.0.0.0
docker exec s1 ifconfig e2 0.0.0.0
docker exec s2 ifconfig e0 0.0.0.0
docker exec s2 ifconfig e1 0.0.0.0
docker exec s3 ifconfig e0 0.0.0.0
docker exec s3 ifconfig e1 0.0.0.0
docker exec s3 ifconfig e2 0.0.0.0

# configure host network
nsenter -t $h1ns -n ifconfig e0 10.0.0.1/24
nsenter -t $h2ns -n ifconfig e0 10.0.1.1/24
nsenter -t $h1ns -n route add default gw 10.0.0.254
nsenter -t $h2ns -n route add default gw 10.0.1.254

# configure host iface mac
docker exec h1 ifconfig e0 hw ether 00:00:00:00:00:01
docker exec h2 ifconfig e0 hw ether 00:00:00:00:00:02

# configure quagga
nsenter -t $s1ns -m bash -c "echo $'interface i0\ninterface i1\ninterface i2\nrouter ospf\n network 10.0.0.0/24 area 0\n network 10.0.2.0/24 area 0\n network 10.0.3.0/24 area 0' >> /etc/quagga/ospfd.conf"
nsenter -t $s1ns -m bash -c "echo $'interface i0\n ip address 10.0.0.254/24' >> /etc/quagga/zebra.conf"
nsenter -t $s1ns -m bash -c "echo $'interface i1\n ip address 10.0.2.1/24' >> /etc/quagga/zebra.conf"
nsenter -t $s1ns -m bash -c "echo $'interface i2\n ip address 10.0.3.1/24' >> /etc/quagga/zebra.conf"
nsenter -t $s2ns -m bash -c "echo $'interface i0\ninterface i1\nrouter ospf\n network 10.0.3.0/24 area 0\n network 10.0.4.0/24 area 0' >> /etc/quagga/ospfd.conf"
nsenter -t $s2ns -m bash -c "echo $'interface i0\n ip address 10.0.3.2/24' >> /etc/quagga/zebra.conf"
nsenter -t $s2ns -m bash -c "echo $'interface i1\n ip address 10.0.4.1/24' >> /etc/quagga/zebra.conf"
nsenter -t $s3ns -m bash -c "echo $'interface i0\ninterface i1\ninterface i2\nrouter ospf\n network 10.0.1.0/24 area 0\n network 10.0.4.0/24 area 0\n network 10.0.2.0/24 area 0' >> /etc/quagga/ospfd.conf"
nsenter -t $s3ns -m bash -c "echo $'interface i0\n ip address 10.0.4.2/24' >> /etc/quagga/zebra.conf"
nsenter -t $s3ns -m bash -c "echo $'interface i1\n ip address 10.0.2.2/24' >> /etc/quagga/zebra.conf"
nsenter -t $s3ns -m bash -c "echo $'interface i2\n ip address 10.0.1.254/24' >> /etc/quagga/zebra.conf"

# start quagga
nsenter -t $s1ns -m -p -n -i zebra -d -f /etc/quagga/zebra.conf --fpm_format protobuf
nsenter -t $s1ns -m -p -n -i ospfd -d -f /etc/quagga/ospfd.conf
nsenter -t $s2ns -m -p -n -i zebra -d -f /etc/quagga/zebra.conf --fpm_format protobuf
nsenter -t $s2ns -m -p -n -i ospfd -d -f /etc/quagga/ospfd.conf
nsenter -t $s3ns -m -p -n -i zebra -d -f /etc/quagga/zebra.conf --fpm_format protobuf
nsenter -t $s3ns -m -p -n -i ospfd -d -f /etc/quagga/ospfd.conf

# set flow rules
docker exec s1 ovs-ofctl del-flows s
docker exec s1 ovs-ofctl add-flow s ip,in_port=1,ip_proto=89,actions=output:4
docker exec s1 ovs-ofctl add-flow s arp,in_port=1,arp_tpa=10.0.0.254,actions=output:4
docker exec s1 ovs-ofctl add-flow s in_port=4,actions=output:1
docker exec s1 ovs-ofctl add-flow s ip,in_port=2,ip_proto=89,actions=output:5
docker exec s1 ovs-ofctl add-flow s arp,in_port=2,arp_tpa=10.0.2.1,actions=output:5
docker exec s1 ovs-ofctl add-flow s in_port=5,actions=output:2
docker exec s1 ovs-ofctl add-flow s ip,in_port=3,ip_proto=89,actions=output:6
docker exec s1 ovs-ofctl add-flow s arp,in_port=3,arp_tpa=10.0.3.1,actions=output:6
docker exec s1 ovs-ofctl add-flow s in_port=6,actions=output:3

docker exec s2 ovs-ofctl del-flows s
docker exec s2 ovs-ofctl add-flow s ip,in_port=1,ip_proto=89,actions=output:3
docker exec s2 ovs-ofctl add-flow s arp,in_port=1,arp_tpa=10.0.3.2,actions=output:3
docker exec s2 ovs-ofctl add-flow s in_port=3,actions=output:1
docker exec s2 ovs-ofctl add-flow s ip,in_port=2,ip_proto=89,actions=output:4
docker exec s2 ovs-ofctl add-flow s arp,in_port=2,arp_tpa=10.0.4.1,actions=output:4
docker exec s2 ovs-ofctl add-flow s in_port=4,actions=output:2

docker exec s3 ovs-ofctl del-flows s
docker exec s3 ovs-ofctl add-flow s ip,in_port=1,ip_proto=89,actions=output:4
docker exec s3 ovs-ofctl add-flow s arp,in_port=1,arp_tpa=10.0.4.2,actions=output:4
docker exec s3 ovs-ofctl add-flow s in_port=4,actions=output:1
docker exec s3 ovs-ofctl add-flow s ip,in_port=2,ip_proto=89,actions=output:5
docker exec s3 ovs-ofctl add-flow s arp,in_port=2,arp_tpa=10.0.2.2,actions=output:5
docker exec s3 ovs-ofctl add-flow s in_port=5,actions=output:2
docker exec s3 ovs-ofctl add-flow s ip,in_port=3,ip_proto=89,actions=output:6
docker exec s3 ovs-ofctl add-flow s arp,in_port=3,arp_tpa=10.0.1.254,actions=output:6
docker exec s3 ovs-ofctl add-flow s in_port=6,actions=output:3

# copy fpmserver
docker cp fpmserver s1:/
docker cp fpmserver s2:/
docker cp fpmserver s3:/

# start fpmserver
docker exec s1 /fpmserver/main.py &
docker exec s2 /fpmserver/main.py &
docker exec s3 /fpmserver/main.py &

# configure openflow app
docker exec s1 ovs-ofctl add-group s group_id:1,type=ff,bucket=watch_port=2,actions=output:2,bucket=watch_port=3,actions=resubmit\(,100\) -O OpenFlow11
docker exec s1 ovs-ofctl add-flow s table=1,ip,nw_dst=10.0.1.0/24,actions=mod_dl_dst:00:00:00:00:00:02,group:1 -O OpenFlow11
docker exec s1 ovs-ofctl add-flow s table=1,ip,nw_dst=10.0.0.0/24,actions=mod_dl_dst:00:00:00:00:00:01,output:1
docker exec s2 ovs-ofctl add-flow s table=1,actions=resubmit\(,100\)
docker exec s3 ovs-ofctl add-flow s table=1,actions=resubmit\(,100\)

# configure openflow app
docker exec s1 ovs-ofctl add-group s group_id:1,type=ff,bucket=watch_port=3,actions=output:3,bucket=watch_port=2,actions=resubmit\(,100\) -O OpenFlow11
docker exec s1 ovs-ofctl add-flow s table=1,ip,nw_dst=10.0.1.0/24,actions=mod_dl_dst:00:00:00:00:00:02,group:1 -O OpenFlow11
docker exec s1 ovs-ofctl add-flow s table=1,ip,nw_dst=10.0.0.0/24,actions=mod_dl_dst:00:00:00:00:00:01,output:1
docker exec s2 ovs-ofctl add-flow s table=1,actions=resubmit\(,100\)
docker exec s3 ovs-ofctl add-flow s table=1,actions=resubmit\(,100\)

docker exec s1 ovs-ofctl add-flow s table=1,actions=resubmit\(,100\)
docker exec s2 ovs-ofctl add-flow s table=1,actions=resubmit\(,100\)
docker exec s3 ovs-ofctl add-flow s table=1,actions=resubmit\(,100\)

```

## Set up vxlan tunnel vxlan.sh
[vxlan Tunnel] (https://gist.github.com/yulis/5c20aa7695fc859d357d2285d4c51e7d)
```
#!/bin/bash
# On Host 192.168.100.25
ovs-vsctl add-br testbr0
ovs-vsctl add-port testbr0 tun1 -- set interface tun1 type=internal
ifconfig tun1 172.17.17.1 netmask 255.255.255.0
ovs-vsctl add-port testbr0 testvxlan0 -- set interface testvxlan0 type=vxlan options:remote_ip=192.168.100.24

# On Host 192.168.100.24
ovs-vsctl add-br testbr0
ovs-vsctl add-port testbr0 tun1 -- set interface tun1 type=internal
ifconfig tun1 172.17.17.2 netmask 255.255.255.0
ovs-vsctl add-port testbr0 testvxlan0 -- set interface testvxlan0 type=vxlan options:remote_ip=192.168.100.25
```


## Set up vxlan tunnel vxlan.sh
[vxlan Tunnel](https://gist.githubusercontent.com/tfherb71/a40206f8450a5a12f7a07eb36208a143/raw/4185803de98c231cf35d3c81a6bb9a1a62db9e05/vxlan%2520tunnels) 
```
#!/bin/bash
#
echo
echo create .ssh directory on enpoint machine using private key 
echo created by ssh-keygen on your two card.
echo copy .ssh directory so you will have same key pair on both servers server1 and server2
echo
#
# Endpoint1
#
ENDPOINT_1="192.168.122.90"
ENDPOINT_2="192.168.122.91"

echo ENDPOINT_1="192.168.122.90"
echo ENDPOINT_2="192.168.122.91"

TUNNEL_IP_1="40.1.1.1"
TUNNEL_IP_2="40.1.1.2"

echo TUNNEL_IP_1="40.1.1.1"
echo TUNNEL_IP_2="40.1.1.2"

ENDPOINTS_IP=("192.168.122.90" "192.168.122.91")


PKEY=/home/therbert/.ssh/vmhost/id_vm_rsa

SSH_OPTIONS="-i ${PKEY} -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -o BatchMode=yes -o LogLevel=error"

function ssh_do() {
    echo
    echo "### "  ssh $@
    ssh ${SSH_OPTIONS} $@
}


for server in "${!ENDPOINTS_IP[@]}"; do
    ssh_do root@${ENDPOINTS_IP[${server}]}\
        /opt/bin/ovs-ofctl del-flows br0
    ssh_do root@${ENDPOINTS_IP[${server}]}\
        /opt/bin/ovs-ofctl del-flows br1
    ssh_do root@${ENDPOINTS_IP[${server}]}\
        /opt/bin/ovs-vsctl del-br br1
done

for server in "${!ENDPOINTS_IP[@]}"; do
    ssh_do root@${ENDPOINTS_IP[${server}]}\
        /opt/bin/ovs-vsctl del-port br0 sw0p0
    ssh_do root@${ENDPOINTS_IP[${server}]}\
        /opt/bin/ovs-vsctl add-br br1
    ssh_do root@${ENDPOINTS_IP[${server}]}\
        /opt/bin/ovs-vsctl add-port br1 sw0p0 -- set interface sw0p0 ofport_request=101
    ssh_do root@${ENDPOINTS_IP[${server}]}\
        /opt/bin/ovs-vsctl add-port br1 tep0 -- set interface tep0 type=internal
done
#
# Server 1 create vtep
#
ssh ${SSH_OPTIONS} root@${ENDPOINT_1}\
    ip addr add ${TUNNEL_IP_1} dev tep0
#
# Server 2 create vtep
#
ssh ${SSH_OPTIONS} root@${ENDPOINT_2}\
    ip addr add ${TUNNEL_IP_2} dev tep0
#
# Servers 1 and 2 set ports and tunnels
#
for server in "${!ENDPOINTS_IP[@]}"; do
    ssh_do root@${ENDPOINTS_IP[${server}]}\
        ip link set tep0 up
    ssh_do root@${ENDPOINTS_IP[${server}]}\
        /opt/bin/ovs-vsctl add-port br0 vxlan0 -- set interface vxlan0 type=vxlan\
        options:remote_ip=flow options:key=flow options:dest_port=4789
done

#
# Set flows in both endpoints for vxlan tunnel
#
echo
echo Set flows on endpoints.
echo
for p in {0..1} ; do
    tid=$(printf "%03d" ${p} 
    ssh_do root@${ENDPOINT_1} /opt/bin/ovs-ofctl add-flow br0 \
"in_port=sw0pf0vf${p},actions=set_tunnel=5${tid},set_field:${TUNNEL_IP_2}\-\>tun_dst,output:vxlan0"

    ssh_do root@${ENDPOINT_2} /opt/bin/ovs-ofctl add-flow br0 \
"in_port=sw0pf0vf${p},actions=set_tunnel=5${tid},set_field:${TUNNEL_IP_1}\-\>tun_dst,output:vxlan0"

    ssh_do root@${ENDPOINT_1}\
        /opt/bin/ovs-ofctl add-flow br0 in_port=vxlan0,tun_id=5${tid},actions=sw0pf0vf${p}
    ssh_do root@${ENDPOINT_2}\
        /opt/bin/ovs-ofctl add-flow br0 in_port=vxlan0,tun_id=5${tid},actions=sw0pf0vf${p}
done

```

## set-multins-vlan-bond.sh
```
#!/bin/bash -eu
#
# VERSION: 0.7
# AUTHOR: Mauro S. Oddi
#
# DESCRIPTION:
#
# Create a network environment similar to what OpenStack Neutron with the default
# Neutorn/OvS ML2 plugin would create on a compute node for testing purposes.
# The script will create bonding device, vlans, OvS and linux bridges, veth pairs,
# 10 net namespaces with internal IPs that will be tunneled
# in a VXLAN overlay. Here OpenStack would normally use OpenFlow flows to determine the path but this
# case is simpler by using ports as alternative.
#
# NETWORK LAYOUT:
#
# PF - BOND - VLAN - OVS BRIDGE(br-ex) - OVS BRIDGE(br-tun) - VXLAN - OVS BRIDGE (br-int) - VETH - LINUXBRIDGE - TAP - VM
#
# TODO:
#
#  - Alternatively VLAN Tagging can be managed by OvS (br-ex) directly but this is not implemented here yet
#
#
# VARIABLES:
#
# IP:     <VTEP IP address>
# INTIP:  <Internal NS IP>
# EXNIC:  <bond_name>
# SLAVES: <bond slaves>
#
#

IP=172.16.0.3
INTIP=192.168.200.3
# To set wihtout bond use p2p1 as EXNIC
EXNIC=bond0
SLAVES="p2p1 p2p2"
VLANID=1001

MTU=1500
#NETNS=vmns0

NSLIST=( vmns10 vmns9 vmns8 vmns7 vmns6 vmns5 vmns4 vmns3 vmns2 vmns1 vmns0 )


function reset() {

	echo "Cleanup all"
        # Delete namespaces
	for I_NS in ${NSLIST[@]}; do
		ip netns del $I_NS
		ip link del veth0-$I_NS
	done
        # Delete OvS bridges br-int and br-tun
	ovs-vsctl del-br br-int
	ip link del br-int
	ovs-vsctl del-br br-tun
	ip link del br-tun
	ovs-vsctl del-port br-ex vtep0
	if [[ -z $VLANID ]]; then
	    # Unplug from br-ex and delete vlan device
	    ovs-vsctl del-port br-ex ${EXNIC}.${VLANID}
	    ip link del ${EXNIC}.{$VLANID}
	else
	    # Unplug from br-ex and remove external interface
            ovs-vsctl del-port br-ex $EXNIC
	fi
	if [[ $EXNIC = "bond0" ]]; then
		 ip link del $EXNIC
		 for I_SLAVE in $SLAVES; do
			ip a f dev $I_SLAVE
		 done
	else
		ip a f dev $EXNIC
	fi
        # Delete vxlan dev if it is was created outside OvS
	#ip link del vxlan10
        # Delete OvS bridge br-ex
	ovs-vsctl del-br br-ex
	ip link del br-tun

}

# Set IP on external interface
function setphysnet() {

	echo "Setting phyisical NIC $NIC ($MTU)"
	ip address add ${IP}/24 dev $EXNIC
	ip link set $ENNIC mtu $MTU
	ip link set $EXNIC up

}


# Setup bond interface
function setbond() {

	local BONDNIC=$1
	shift 1
	local SLAVENICS=$@

	echo "Setting active-passive bond $BONDNIC with $SLAVENICS ($MTU)"

	# Clean if exists
	ip link del $BONDNIC
	ip link add $BONDNIC type bond
	ip link set $BONDNIC type bond miimon 100 mode active-backup
	for I_NIC in $SLAVENICS; do
		ip link set $I_NIC down
		ip link set $I_NIC master $BONDNIC
	done
	ip link set $BONDNIC mtu $MTU
	ip link set $BONDNIC up
#	ip address add ${IP}/24 dev $BONDNIC

}


# Setup vxlan interface (if we use OvS this is not required)
function setvxlan() {
	echo "Setting VXLAN tunnels in linux br"
	# set tunnel to C
	ip link add vxlan10 type vxlan id 10 dev $NIC remote 172.16.0.4 dstport 4789
	ip link set vxlan10 mtu $(( $MTU - 50 ))
	ip link set vxlan10 up
	# set tunnel to A
	#ip link add vxlan11 type vxlan id 11 dev $NIC remote 172.16.0.2 dstport 4789
	#ip link set vxlan11 mtu $(( $MTU - 50 ))
	#ip link set vxlan11 up
}


# integration bridge with linux br
function setbrint() {
	echo "Creating linux bridge with VTEP"
	ip link add br-int type bridge
	ip link set br-int up
	# add the vxlan port
	ip link set vxlan10 master br-int
	#ip link set vxlan11 master br-int
}

# Create vlan device
function setvlan() {
	echo "Creating kernel vlan device"
        ip link add name $EXNIC link ${EXNIC}.${VLANID} type vlan id $VLANID
}

# Create external bridge
function setbrex() {
	echo "Creating OvS bridge for external traffic br-ex and VTEP"
	ovs-vsctl add-br br-ex
	if [[ -z $VLANID ]]; then
	    ovs-vsctl add-port br-ex $EXNIC.$VLANID
	else
	    ovs-vsctl add-port br-ex $EXNIC
	fi
	ovs-vsctl add-port br-ex vtep0  \
	    	-- set interface vtep0 type=internal
	ip link set vtep0 up
	ip a a ${IP}/24 dev vtep0
}


# Create tunnel bridge
function setbrtun() {
	echo "Creating OvS bridge for tunnel traffic br-tun and VXLAN tunnel"
	ovs-vsctl add-br br-tun
	ovs-vsctl add-port br-tun vxlan10 \
		-- set interface vxlan10 type=vxlan \
		   options:remote_ip=172.16.0.4 \
		   options:key=5000 \
		   options:dst_port=4789
	#ovs-vsctl add-port br-tun vtep0 \
	#	-- set interface vtep0 type=internal
	#ip link set vtep0 mtu $(( $MTU - 50 ))
	# ip address add 192.168.1.2/24 dev vtep0

}

# Create interconnection bridge and patch it to the br-tun
function setbrint() {
	echo "Creating OcS bridge for internal traffic br-int and patch to br-tun"
	ovs-vsctl add-br br-int
	ovs-vsctl add-port br-int patch-br-tun \
		-- set interface patch-br-tun type=patch \
		   options:peer=patch-br-int
	ovs-vsctl add-port br-tun patch-br-int \
		-- set interface patch-br-int type=patch \
		   options:peer=patch-br-tun
}

# Create the netns and veths for test interface based on ID passed
function setns() {

	local NETNS=vmns$1
	local IPNS=192.168.$(( 100 + $1 )).3
	local REMOTEIPNS=192.168.$(( 100 + $1 )).4

	echo "Craating NS $NETNS - $IPNS"
	ip netns add $NETNS
	#NSLIST+=( $NETNS )
	ip link add veth0-$NETNS type veth peer name veth1-$NETNS
	ip link set veth0-$NETNS mtu $(( $MTU - 50 ))
	ip link set veth1-$NETNS mtu $(( $MTU - 50 ))
	ovs-vsctl add-port br-int veth0-$NETNS
	ip link set veth0-$NETNS up
	ip link set veth1-$NETNS netns  $NETNS
	ip netns exec $NETNS ip link set veth1-$NETNS name eth0
	ip netns exec $NETNS ip address add ${IPNS}/24 dev eth0
	ip netns exec $NETNS ip link set eth0 up
	return 0

}

# Run iperf on existing NS
function runns() {
	local NETNS=vmns$1
	local IPNS=192.168.$(( 100 + $1 )).3
	local REMOTEIPNS=192.168.$(( 100 + $1 )).4

	echo "Running on NS $NETNS - Command: iperf -s $IPNS \&"
	ip netns exec $NETNS iperf3 -s $IPNS  &
	RETVAL=$?
	[ $RETVAL -eq 0 ] && echo "iperf3 listening on IP  $IPNS" || echo "failed to run iperf3 -s on $IPNS" >&2
	return $RETVAL
}

# Create multiple test namespaces
function set_multiple_ns() {
	for I_NS in {1..10} ;  do
		setns  $I_NS
	done
}

# Run command in multiple test namespaces
function run_multiple_ns() {
	for I_NS in {1..10} ;  do
		runns  $I_NS
		sleep 1
	done

}

# MAIN
function main() {
	reset
	setbond $EXNIC $SLAVES
	#setvlan
	setbrex
	setbrtun
	setbrint
	#setns
	set_multiple_ns
	run_multiple_ns

}
main
#EOS

```
