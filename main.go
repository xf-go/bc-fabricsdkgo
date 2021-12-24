/**
  author: hanxiaodong
  QQ群（专业Fabric交流群）：862733552
*/
package main

import (
	"bc-fabricsdkgo/sdkInit"
	"bc-fabricsdkgo/service"
	"bc-fabricsdkgo/web"
	"bc-fabricsdkgo/web/controller"
	"fmt"
	"os"
)

const (
	configFile  = "config.yaml"
	initialized = false
	SimpleCC    = "simplecc"
)

func main() {
	initInfo := &sdkInit.InitInfo{
		ChannelID:     "kevinkongyixueyuan",
		ChannelConfig: os.Getenv("GOWORKSPACE") + "/src/bc-fabricsdkgo/fixtures/artifacts/channel.tx",

		OrgAdmin:       "Admin",
		OrgName:        "Org1",
		OrdererOrgName: "orderer.kevin.kongyixueyuan.com",

		ChaincodeID:     SimpleCC,
		ChaincodeGoPath: os.Getenv("GOWORKSPACE"),
		ChaincodePath:   "bc-fabricsdkgo/chaincode", //因为会自动在GOPATH后加src，所以需要../
		UserName:        "User1",
	}

	sdk, err := sdkInit.SetupSDK(configFile, initialized)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer sdk.Close()

	err = sdkInit.CreateChannel(sdk, initInfo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	channelClient, err := sdkInit.InstallAndInstantiateCC(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(channelClient)

	serviceSetup := service.ServiceSetup{
		ChaincodeID: SimpleCC,
		Client:      channelClient,
	}

	// 调用链码添加状态
	msg, err := serviceSetup.SetInfo("hanxiaodong", "kongyixueyuan")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}

	// 调用链码查询状态
	msg, err = serviceSetup.GetInfo("hanxiaodong")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}

	app := controller.Application{
		Fabric: &serviceSetup,
	}
	web.WebStart(&app)
}
