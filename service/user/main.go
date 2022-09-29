package main

import (
    "go-framework/service/user/api"
    "go-framework/service/user/rpc"
)

func main() {
    
    go api.Main()
    
    rpc.Main()
    
}
