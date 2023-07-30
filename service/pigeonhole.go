package service

import (
	"MyWebProject/config"
	"MyWebProject/dao"
	"MyWebProject/models"
	"log"
)

func FindPostPigeonhole() models.PigeonholeRes {
	posts, err := dao.GetAllPost()
	if err != nil {
		log.Println(err)
	}
	pigeonholeMap := make(map[string][]models.Post)
	for _, post := range posts {
		at := post.CreateAt
		month := at.Format("2006-01")
		pigeonholeMap[month] = append(pigeonholeMap[month], post)

	}
	categorys, _ := dao.GetAllCategory()
	return models.PigeonholeRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		categorys,
		pigeonholeMap,
	}

}
