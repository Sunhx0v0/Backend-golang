package models

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type LoginInfo struct {
	Password string `json:"password"` // 密码
	UserID   int64  `json:"userId"`   // 用户编号
	UserName string `json:"userName"` // 用户名
}

type UserClaim struct { // 登录验证
	UserName string
	Claims   []LoginInfo
}

// 收藏的笔记
type Collects struct {
	Cover       string `json:"cover"`
	LikedNum    int64  `json:"likedNum"` // 点赞数
	NoteID      int64  `json:"noteId"`   // 笔记编号
	Title       string `json:"title"`
	CreatorID   int64  `json:"creatorID"`
	CreatorName string `json:"creatorName"` // 作者姓名
	Portrait    string `json:"portrait"`
}

// 点赞的笔记
type Likes struct {
	Cover       string `json:"cover"`
	LikedNum    int64  `json:"likedNum"` // 点赞数
	NoteID      int64  `json:"noteId"`   // 笔记编号
	Title       string `json:"title"`
	CreatorID   int64  `json:"creatorID"`
	CreatorName string `json:"creatorName"` // 作者姓名
	Portrait    string `json:"portrait"`
}

// 发布的笔记
type Notes struct {
	Cover       string `json:"cover"`
	LikedNum    int64  `json:"likedNum"` // 点赞数
	NoteID      int64  `json:"noteId"`   // 笔记编号
	Title       string `json:"title"`
	CreatorID   int64  `json:"creatorID"`
	CreatorName string `json:"creatorName"` // 作者姓名
	Portrait    string `json:"portrait"`
}

// 用户基本信息
type UserInfo struct {
	Birthday     string `json:"birthday"`
	CollectedNum int64  `json:"collectedNum"` // 被收藏的笔记数
	CollectNum   int64  `json:"collectNum"`   // 收藏数
	FansNum      int64  `json:"fansNum"`
	FollowNum    int64  `json:"followNum"`    // 关注数
	Gender       string `json:"gender "`      // 性别
	Introduction string `json:"introduction"` // 简介
	LikedNum     int64  `json:"likedNum"`     // 被点赞数量
	Mail         string `json:"mail"`
	NoteNum      int64  `json:"noteNum "`
	Password     string `json:"password"`
	PhoneNumber  string `json:"phoneNumber"`
	Portrait     string `json:"portrait"` // 头像
	RegistTime   string `json:"registTime"`
	UserID       int64  `json:"userAccount"`
	UserName     string `json:"userName "`
}

func UserInfoDB(id int) UserInfo { // 从数据库获得用户信息
	// fmt.Println("传入ID", id)
	// var ui UserInfo
	// sqlStr := `select userAccount, userName, gender, portrait, introduction, fansNum, noteNum, collectNum, followNum, collectedNum, likedNum, phoneNumber, mail, password from userinfo where userAccount = ?`
	// err := db.QueryRow(sqlStr, id).Scan(&ui.UserID, &ui.UserName, &ui.Gender, &ui.Portrait, &ui.Introduction, &ui.FansNum, &ui.NoteNum, &ui.CollectNum, &ui.FollowNum, &ui.CollectedNum, &ui.LikedNum, &ui.PhoneNumber, &ui.Mail, &ui.Password)
	// ui.Birthday = time.Now().Format("2006-01-02 15:04:05")
	// ui.RegistTime = time.Now().Format("2006-01-02 15:04:05")
	// if err != nil {
	// 	fmt.Printf("query failed, err:%v\n", err)
	// 	return ui
	// }
	// return ui

	// var ui UserInfo
	// sqlStr := `select userAccount, userName, gender, portrait, introduction, fansNum, noteNum, collectNum, followNum, collectedNum, likedNum, phoneNumber, mail, password from userinfo where userAccount = ?`
	// rows, err := db.Query(sqlStr, id)
	// if err != nil {
	// 	fmt.Println("没有数据1")
	// 	return ui
	// }
	// //defer rows.Close()
	// // for rows.Next() {
	// rows.Next()
	// ui.Birthday = time.Now().Format("2006-01-02 15:04:05")
	// ui.RegistTime = time.Now().Format("2006-01-02 15:04:05")
	// rows.Scan(&ui.UserID, &ui.UserName, &ui.Gender, &ui.Portrait, &ui.Introduction, &ui.FansNum, &ui.NoteNum, &ui.CollectNum, &ui.FollowNum, &ui.CollectedNum, &ui.LikedNum, &ui.PhoneNumber, &ui.Mail, &ui.Password)
	// //}
	// defer rows.Close()
	// fmt.Println("到最后了吗")
	// return ui

	var ui UserInfo
	ui.UserID = 10001
	ui.UserName = "zahgsad"
	ui.Gender = "sha"
	ui.Portrait = "/"
	ui.Introduction = "jntm"
	ui.Birthday = time.Now().Format("2006-01-02 15:04:05")
	ui.RegistTime = time.Now().Format("2006-01-02 15:04:05")
	ui.FansNum = 666
	ui.NoteNum = 777
	ui.CollectNum = 888
	ui.FollowNum = 999
	ui.CollectedNum = 1010
	ui.LikedNum = 115
	ui.PhoneNumber = "1233663"
	ui.Mail = "1235"
	ui.Password = "123456"
	return ui
}

func NoteInfoDB(id int) []Notes { // 从数据库获得用户信息
	// var notes []Notes
	// sqlStr := `select noteId, title, cover,creatorAccount,likeNum, creatorName
	// from noteInfo
	// where creatorAccount = '"+id+"'`
	// rows, err := db.Query(sqlStr)
	// if err != nil {
	// 	fmt.Printf("query failed, err:%v\n", err)
	// 	return notes
	// }
	// // 关闭rows释放持有的数据库链接
	// defer rows.Close()

	// // 循环读取结果集中的数据
	// for rows.Next() {
	// 	var nt Notes
	// 	err := rows.Scan(&nt.NoteID, &nt.Title, &nt.Cover, &nt.CreatorID, &nt.LikedNum, &nt.CreatorName)
	// 	nt.Portrait = "/"
	// 	if err != nil {
	// 		fmt.Printf("scan failed, err:%v\n", err)
	// 		return notes
	// 	}
	// 	notes = append(notes, nt)
	// }
	// return notes

	var notes []Notes
	var nt Notes
	nt.NoteID = 10001
	nt.Title = "l就别"
	nt.Cover = "站发给"
	nt.CreatorID = 10001
	nt.LikedNum = 500
	nt.Portrait = "/"
	nt.CreatorName = "张菲"
	notes = append(notes, nt)
	return notes
}

func CollectInfoDB(id int) (notes []Collects) { // 从数据库获得用户信息
	// sqlStr := `select n.noteId, n.title, n.cover,n.creatorAccount,n.likeNum, u.portrait
	// from noteInfo n,userInfo u
	// where n.creatorAccount = '"+id+"' and u.userAccount = '"+id+"'`
	// rows, err := db.Query(sqlStr)
	// if err != nil {
	// 	fmt.Printf("query failed, err:%v\n", err)
	// 	return
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var nt Collects
	// 	err := rows.Scan(&nt.NoteID, &nt.Title, &nt.Cover, &nt.CreatorID, &nt.LikedNum, &nt.Portrait)
	// 	if err != nil {
	// 		fmt.Printf("scan failed, err:%v\n", err)
	// 		return
	// 	}
	// 	notes = append(notes, nt)
	// }
	// return
	var nt Collects
	nt.NoteID = 10001
	nt.Title = "l就别"
	nt.Cover = "站发给"
	nt.CreatorID = 10001
	nt.LikedNum = 500
	nt.Portrait = "/"
	nt.CreatorName = "张菲"
	notes = append(notes, nt)
	return
}

func LikeInfoDB(id int) (notes []Likes) { // 从数据库获得用户信息
	// sqlStr := `select n.noteId, n.title, n.cover,n.creatorAccount,n.likeNum, u.portrait
	// from noteInfo n,userInfo u
	// where n.creatorAccount = '"+id+"' and u.userAccount = '"+id+"'`
	// rows, err := db.Query(sqlStr)
	// if err != nil {
	// 	fmt.Printf("query failed, err:%v\n", err)
	// 	return
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var nt Likes
	// 	err := rows.Scan(&nt.NoteID, &nt.Title, &nt.Cover, &nt.CreatorID, &nt.LikedNum, &nt.Portrait)
	// 	if err != nil {
	// 		fmt.Printf("scan failed, err:%v\n", err)
	// 		return
	// 	}
	// 	notes = append(notes, nt)
	// }
	// return
	var nt Likes
	nt.NoteID = 10001
	nt.Title = "l就别"
	nt.Cover = "站发给"
	nt.CreatorID = 10001
	nt.LikedNum = 500
	nt.Portrait = "/"
	nt.CreatorName = "张菲"
	notes = append(notes, nt)
	return
}

func CheckUser(userName, password string) bool {
	//用户的登录信息
	var buser LoginInfo
	sqlstr := "select userAccount from userInfo where userName=? and password=?"
	err := db.QueryRow(sqlstr, userName, password).Scan(&buser.UserID)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return false
	}
	if buser.UserID > 0 {
		// fmt.Print(buser.UserID)
		return true
	}

	return false
}
