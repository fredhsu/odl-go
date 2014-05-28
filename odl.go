package main

import (
    "fmt"
    "net/http"
    "strings"
    "io/ioutil"
    "encoding/json"
)

type Node struct {
    Type string
    Id string
}

type NodeConnector struct {
    Type string
    Node Node
    Id string
}

type Edge struct {
    // for decoding it seems to be case insensitive
    TailNodeConnector NodeConnector //`json:"tailNodeConnector"`
    HeadNodeConnector NodeConnector //`json:"headNodeConnector"`
}

type Properties struct {
    TimeStamp TimeStamp
    Name ValueString
    State ValueInt
    Config ValueInt
    Bandwidth ValueInt
}

type TimeStamp struct {
    Value int
    Name string
}

type UserLinks struct {
    UserLinks []UserLink
}

type UserLink struct {
    Status string
    Name string
    SrcNodeConnector string
    DstNodeConnector string
}

type ValueString struct {
    Value string
}

type ValueInt struct {
    Value int
}

type EdgeProperty struct {
    Edge Edge //`json:"edge"`
    Properties Properties //`json:"properties"`
}

type EdgeProperties struct {
    EdgeProperties []EdgeProperty //`json:"edgeProperties"`
}

func main() {
    baseurl := "http://admin:admin@odl.aristanetworks.com:8080/controller/nb/v2"

    // The URL to get the topology of the default slice
    url := strings.Join([]string{baseurl, "topology/default"}, "/")
    fmt.Println(url)
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error")
    }
    contents, err := ioutil.ReadAll(resp.Body)
    var e EdgeProperties
    err = json.Unmarshal(contents, &e)
    fmt.Println(e.EdgeProperties[0].Edge)
    fmt.Println(e.EdgeProperties[0].Edge.TailNodeConnector.Id)
    fmt.Println(e.EdgeProperties[0].Properties.Name)
    fmt.Println(e.EdgeProperties[0].Properties.TimeStamp)
}
