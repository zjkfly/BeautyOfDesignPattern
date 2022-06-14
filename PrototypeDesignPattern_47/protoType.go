package PrototypeDesignPattern_47

import (
	"encoding/json"
	"time"
)

/*
其中，单例模式⽤来创建全局唯⼀的对象。
⼯⼚模式⽤来创建不同但是相关类型的对象（继承同⼀⽗类或者接⼝的⼀组⼦类），由给定的参数来决定创建哪种类型的对象。
建造者模式是⽤来创建复杂对象，可以通过设置不同的可选参数，“定制化”地创建不同的对象。
原型模式针对创建成本⽐较⼤的对象，利⽤对已有对象进⾏复制的⽅式进⾏创建，以达到节省创建时间的⽬的。
*/
type KeyWordsDemo struct {
	LastUt      int64
	AllKeyWords map[string]keyWordDetail
}

func NewKeyWordsDemo(keyWords map[string]keyWordDetail) *KeyWordsDemo {
	return &KeyWordsDemo{AllKeyWords: keyWords}
}

type keyWordDetail struct {
	ut        int64
	searchCnt int64
}

func (k *KeyWordsDemo) refresh() {
	KeyWords := getSnapShootModifyKeyWordsByUt(k.LastUt)
	for keyWord, detail := range KeyWords {
		if originalDetail, ok := k.AllKeyWords[keyWord]; ok {
			if detail.searchCnt != originalDetail.searchCnt {
				originalDetail.searchCnt = detail.searchCnt
			}
			originalDetail.ut = detail.ut
		} else {
			k.AllKeyWords[keyWord] = detail
		}
	}
	k.LastUt = time.Now().Unix()
}

// 获取某个时间点之后多出的数据
func getSnapShootModifyKeyWordsByUt(ut int64) map[string]keyWordDetail {
	return map[string]keyWordDetail{}
}

// 原型模式-序列化进行
func AllKeyWordsCloneBySerialization(origin map[string]keyWordDetail) map[string]keyWordDetail {
	var newKeywords map[string]keyWordDetail
	b, _ := json.Marshal(origin)
	json.Unmarshal(b, &newKeywords)
	return newKeywords
}

// todo 原型模式-深拷贝递归对象copy进行
func AllKeyWordsCloneByObject(origin map[string]keyWordDetail) map[string]keyWordDetail {
	return nil
}