package models

import (
	"Gin/db"
	"fmt"
	"log"
)

type User_room struct {
	Useridently string `json:"useridently"`
	Roomidently string `json:"roomidently"`
	Create_time int64  `json:"create_time"`
	Update_time int64  `json:"update_time"`
	Room_type   string `json:"room_type"`
}
type Userroom struct {
	Id          int    `json:"id"`
	Useridently string `json:"useridently"`
	Roomidently string `json:"roomidently"`
	Create_time int64  `json:"create_time"`
	Update_time int64  `json:"update_time"`
	Room_type   string `json:"room_type"`
}

// InsertUseridently 加入群聊
func InsertUseridently(user_room *User_room) error {
	_, err := db.DB.Exec("insert into user_room (useridently,roomidently,create_time,update_time,room_type) value (?,?,?,?,?)", user_room.Useridently, user_room.Roomidently, user_room.Create_time, user_room.Update_time, user_room.Room_type)
	if err != nil {
		return err
	}
	return nil
}

// GetUserbyIdentlyRoomId 查
func GetUserbyIdentlyRoomId(roomidently string) []User_room {
	fmt.Println(roomidently)
	user_room := []User_room{}
	err := db.DB.Select(&user_room, "select * from user_room where  roomidently=?", roomidently)
	if err != nil {
		log.Println(err)
		return nil
	}
	fmt.Println(user_room)
	return user_room
}

// PrivateInsertUseridently 加好友
func PrivateInsertUseridently(use1 User_room) error {
	_, err := db.DB.Exec("insert into user_room (useridently,roomidently,create_time,update_time,room_type) value (?,?,?,?,?)", use1.Useridently, use1.Roomidently, use1.Create_time, use1.Room_type)
	if err != nil {
		return err
	}
	return nil
}

// 是否已经互为好友
func GetOther(indently1, indently2 string) (bool, error) {
	user_room1 := Userroom{}
	user_room2 := Userroom{}
	err := db.DB.Get(&user_room1, "select * from user_room where  Useridently=?", indently1)
	if err != nil {
		return true, err
	}

	err = db.DB.Get(&user_room2, "select * from user_room where  Useridently=?", indently2)
	if err != nil {
		return true, err
	}
	fmt.Println(user_room1, user_room2)
	if user_room1.Roomidently == user_room2.Roomidently {
		return true, nil
	}
	return false, nil
}

// 删除
func Del(indently1, indently2 string) error {
	exec, err := db.DB.Exec("delete from user_room where useridently=?", indently1)
	if err != nil {
		return err
	}
	fmt.Println(exec)
	exec, err = db.DB.Exec("delete from user_room where useridently=?", indently2)
	if err != nil {
		return err
	}
	return nil
}
