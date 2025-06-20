package test

import (
	"encoding/json"
	"log"
	"testing"
)

var s = `[
  {
    "ID": 1,
    "Name": "电子产品",
    "Level": 1,
    "IsTab": true
  },
  {
    "ID": 2,
    "Name": "服装鞋帽",
    "Level": 1,
    "ParentCategoryID": 1,
    "IsTab": true
  },
  {
    "ID": 3,
    "Name": "家居生活",
    "Level": 1,
    "ParentCategoryID": 2,
    "IsTab": true
  },
  {
    "ID": 4,
    "Name": "美妆护肤",
    "Level": 1,
    "ParentCategoryID": 0,
    "IsTab": true
  },
  {
    "ID": 5,
    "Name": "食品饮料",
    "Level": 1,
    "ParentCategoryID": 4,
    "IsTab": true
  },
  {
    "ID": 6,
    "Name": "运动户外",
    "Level": 1,
    "ParentCategoryID": 5,
    "IsTab": true
  },
  {
    "ID": 6,
    "Name": "运动户外",
    "Level": 1,
    "ParentCategoryID": 4,
    "IsTab": true
  }
]`

type Category struct {
	ID       int32       `json:"ID"`
	Name     string      `json:"Name"`
	Level    int32       `json:"Level"`
	Pid      int32       `json:"ParentCategoryID"`
	IsTab    bool        `json:"IsTab"`
	Children []*Category `json:"children"`
}

func TestGenTree(t *testing.T) {
	// 1. 解析 JSON
	var flatList []Category
	err := json.Unmarshal([]byte(s), &flatList)
	if err != nil {
		panic(err)
	}

	// 2. 用 map[id] 节点保存，方便查找父节点
	m := make(map[int32]*Category)
	for i, v := range flatList {
		m[v.ID] = &flatList[i]
	}

	var roots []*Category
	for i, category := range flatList {
		node, ok := m[category.Pid]
		// 如果他的父级存在，则
		if ok {
			if node.Children == nil {
				node.Children = make([]*Category, 0)
			}
			// 将他挂在父级下面
			node.Children = append(node.Children, &flatList[i])
		} else {
			// 如果父级不存在说明它既是父级
			roots = append(roots, m[category.ID])
		}
	}

	// 4. 打印结果
	bs, _ := json.MarshalIndent(roots, "", "  ")
	log.Println(string(bs))
}
