package model

import (
	"pkg/necessary"
)

type GetTagsReq struct {
	Necessary       necessary.NecessaryUserInformation
	AccountGroupIDs []uint32 `json:"accountGroupIDs" schema:"accountGroupIDs" minimum:"1"` // Идентификаторы групп счетов
}
