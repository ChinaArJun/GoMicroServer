package handler

import (
	"HouseProject/IHomeWeb/models"
	"context"
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/julienschmidt/httprouter"

	//"google.golang.org/grpc"
	"net/http"

	"github.com/micro/go-grpc"

	GETAREA "HouseProject/GetArea/proto/GetArea"
	//Ihomeweb "github.com/micro/examples/template/srv/proto/example"
)

// 获取地区
func GetAreas(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("获取地区请求客户端 url:api/v1.0/areas")
	//创建新的grpc返回句柄
	server := grpc.NewService()
	// 服务初始化
	server.Init()

	// 创建获取地区的服务并返回句柄
	exampleClient := GETAREA.NewGetAreaService("go.micro.srv.GetArea", server.Client())

	// 调用函数并获得返回的数据
	rsp, err := exampleClient.GetArea(context.TODO(), &GETAREA.Request{})
	if err != nil {
		http.Error(w, err.Error(), 502)
		return
	}
	// 创建返回类型的切片
	area_list := []models.Area{}
	// 循环读取接收的数据
	for _, value := range rsp.Data {
		tmp := models.Area{Id: int(value.Aid), Name: value.Aname}
		area_list = append(area_list, tmp)
	}

	// 返回数据给前端的map
	response := map[string]interface{}{
		"err":  rsp.Error,
		"msg":  rsp.Errmsg,
		"data": area_list,
	}

	// 会传数据的时候三直接发送过去的并没有设置数据格式
	//w.Header().Set("Content-Type","application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}

//func IhomewebCall(w http.ResponseWriter, r *http.Request) {
//	// decode the incoming request as json
//	var request map[string]interface{}
//	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
//		http.Error(w, err.Error(), 500)
//		return
//	}
//
//	// call the backend service
//	IhomewebClient := Ihomeweb.NewIhomewebService("go.micro.srv.Ihomeweb", client.DefaultClient)
//	rsp, err := IhomewebClient.Call(context.TODO(), &Ihomeweb.Request{
//		Name: request["name"].(string),
//	})
//	if err != nil {
//		http.Error(w, err.Error(), 500)
//		return
//	}
//
//	// we want to augment the response
//	response := map[string]interface{}{
//		"msg": rsp.Msg,
//		"ref": time.Now().UnixNano(),
//	}
//
//	// encode and write the response as json
//	if err := json.NewEncoder(w).Encode(response); err != nil {
//		http.Error(w, err.Error(), 500)
//		return
//	}
//}
