package models

import (
	"Gin/db"
	"fmt"
	"log"
)

type Room_id struct {
	Roomidently string `json:"roomidently"`
	Useridently string `json:"useridently"`
	Room_type   string `json:"room_type"`
	Creaet_time int64  `json:"creaet_time"`
	Create_uesr string `json:"create_uesr"`
	Info        string `json:"info"`
}

func (r Room_id) GetRoomidently() string {
	return "Room_id"
}

func GetRoomidentlyUser(roomidently *Room_id) {
	room := &Room_id{}
	err := db.DB.Select(room, "select * from  roomidently where roomidently=?,", roomidently)
	if err != nil {
		return
	}

}

// 创建群聊
func CreateRoom(id *Room_id) bool {
	_, err := db.DB.Exec("insert into room_id(roomidently,useridently,room_type,creaet_time,create_uesr,info) value (?,?,?,?,?,?)", id.Roomidently, id.Useridently, id.Room_type, id.Creaet_time, id.Create_uesr, id.Info)
	if err != nil {
		log.Println("INSERT ERR:", err)
		return false
	}
	return true
}

// 创建群聊个数
func SelectGroupNumber(create_uesr, types string) int {
	room := []Room_id{}
	err := db.DB.Select(&room, "select * from room_id where create_uesr=? and room_type=?", create_uesr, types)
	if err != nil {
		log.Println("SELECT ERR", err)
		return 0
	}
	fmt.Println(len(room))
	return len(room)

}

// 获取房间是否存在
func GetRoomId(roomidently string) bool {
	room := Room_id{}
	err := db.DB.Get(&room, "select * from room_id where roomidently =?", roomidently)
	if err != nil {
		log.Println(err)
		return false
	}
	fmt.Println(room)
	return true

}
