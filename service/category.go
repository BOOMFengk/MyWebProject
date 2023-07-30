package service

import (
	"MyWebProject/config"
	"MyWebProject/dao"
	"MyWebProject/models"
	"fmt"
	"html/template"
)

func GetPostByCategoryId(cId, page, pageSize int) (*models.CategoryResponse, error) {

	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	posts, err := dao.GetPostPageByCategoryId(cId, page, pageSize)
	var postMores []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[0:100]
		}

		postMore := models.PostMore{
			post.Pid,
			post.Title,
			post.Slug,
			template.HTML(content),
			post.CategoryId,
			categoryName,
			post.UserId,
			userName,
			post.ViewCount,
			post.Type,
			models.DateDay(post.CreateAt),
			models.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postMore)
	}
	total := dao.CountGetAllPostByCategoryId(cId)
	pageCount := (total-1)/10 + 1
	var pages []int
	for i := 0; i < pageCount; i++ {
		pages = append(pages, i+1)

	}
	var hr = &models.HomeData{
		config.Cfg.Viewer,
		categorys,
		postMores,
		total,
		page,
		pages,
		page != pageCount,
	}
	categoryName := dao.GetCategoryNameById(cId)
	categoryResponse := &models.CategoryResponse{
		hr,
		categoryName,
	}
	fmt.Printf("categoryName 是 %s ", categoryName)
	return categoryResponse, nil
}
