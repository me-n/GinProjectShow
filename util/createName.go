package util

import (
	"math/rand"
	"time"
)

//随机生成名字
func CreateName() string {
	familyName := []string{"孟", "赵", "钱", "孙", "李", "周",
		"吴", "郑", "王", "冯", "陈", "楚", "卫",
		"蒋", "沈", "韩", "杨", "张",
		"欧阳", "东门", "西门", "上官", "诸葛",
		"司徒", "司空", "夏侯"}
	lastName := []string{"春", "夏", "秋", "冬",
		"风", "霜", "雨", "雪",
		"木", "禾", "米", "竹",
		"山", "石", "田", "土", "福",
		"禄", "寿", "喜", "文", "武", "才", "华",
		"离月", "风华", "长剑", "琼华", "尚德"}
	rand.Seed(time.Now().Unix())
	index01 := rand.Intn(len(familyName))
	index02 := rand.Intn(len(lastName))
	name := familyName[index01] + lastName[index02]
	return name
}
