package boluobaoLib

import (
	"github.com/AlexiaVeronica/boluobaoLib/boluobaomodel"
)

const (
	username      = `androiduser`
	password      = "1a#$51-yt69;*Acv@qxq"
	timeout       = 30
	baseURL       = "https://api.sfacg.com"
	androidApiKey = "FN_Q29XHVmfV3mYX"
	userAgent     = "boluobao/5.0.62(android;25)/OPPO/%s/OPPO"
)

type continueFunction func(bookInfo *boluobaomodel.BookInfoData, chapter boluobaomodel.ChapterList) bool
type contentFunction func(bookInfo *boluobaomodel.BookInfoData, chapter *boluobaomodel.ContentData)
