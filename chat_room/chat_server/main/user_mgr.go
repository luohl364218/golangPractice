package main

import "golangPractice/chat_room/chat_server/model"

var(
	mgr *model.UserMgr
)

func initUserMgr() {
	mgr = model.NewUserMgr(pool)
}

