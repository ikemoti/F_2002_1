package database

import "github.com/jphacks/F_2002_1/go/domain/entity"

var Plants = entity.Plants{
	entity.Plant{
		ID:          1,
		Name:        "じゃがいも",
		NickName:    "ジャガーくん",
		Price:       1200,
		Period:      90,
		Difficulty:  1,
		Description: "僕はおいしさ満点のジャガイモだいも！！！育てばホクホクフィーバーできるいもよ！！！！",
		KitName:     "初めてでも安心！収穫キット",
		SeasonFrom:  1,
		SeasonTo:    12,
	},
	entity.Plant{
		ID:          2,
		Name:        "いちご",
		NickName:    "ベリーちゃん",
		Price:       1200,
		Period:      210,
		Difficulty:  3,
		Description: "ハァイ、私はベリーちゃん。甘いけど、私を育てるのは甘くないわよ。試してみる？",
		KitName:     "初めてでも安心！収穫キット",
		SeasonFrom:  1,
		SeasonTo:    12,
	},
	entity.Plant{
		ID:          3,
		Name:        "なす",
		NickName:    "なっくん",
		Price:       1200,
		Period:      120,
		Difficulty:  4,
		Description: "やぁ！！！！！！！ナスだよ！！！！！田楽、焼き浸し、なんでも美味しいよ！！！！！僕を選んでおくれ！！！！！",
		KitName:     "初めてでも安心！収穫キット",
		SeasonFrom:  1,
		SeasonTo:    12,
	},
	entity.Plant{
		ID:          4,
		Name:        "たまねぎ",
		NickName:    "たまちゃん",
		Price:       1200,
		Period:      80,
		Difficulty:  3,
		Description: "こんにちは。体によくて、サラダでもカレーでも何にでも使える玉ねぎ、育てたことある？",
		KitName:     "初めてでも安心！収穫キット",
		SeasonFrom:  1,
		SeasonTo:    12,
	},
	entity.Plant{
		ID:          5,
		Name:        "きゅうり",
		NickName:    "キューちゃん",
		Price:       1200,
		Period:      90,
		Difficulty:  3,
		Description: "きゅきゅ！！僕は漬物にしてよし、サラダにしてよしのジューシーきゅうりだキュ！試してみるきゅ？",
		KitName:     "初めてでも安心！収穫キット",
		SeasonFrom:  1,
		SeasonTo:    12,
	},
	entity.Plant{
		ID:          6,
		Name:        "にんじん",
		NickName:    "ジンさん",
		Price:       1200,
		Period:      120,
		Difficulty:  3,
		Description: "人参、俺の仮の名前さ……甘くて芯があるのさ……君だけの名前を、僕にくれるかい……？",
		KitName:     "初めてでも安心！収穫キット",
		SeasonFrom:  1,
		SeasonTo:    12,
	},
}
