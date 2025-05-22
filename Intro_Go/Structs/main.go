package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (authI authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", authI.username, authI.password)
}

func test(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
}

func main() {
	test(authenticationInfo{
		username: "Wall",
		password: "123123123123",
	})
}
