package service

import (
	"MyWebProject/config"
	"MyWebProject/dao"
	"MyWebProject/models"
	"html/template"
	"log"
)

//	type PostMore struct {
//		Pid          int           `json:"pid"`          // 文章ID
//		Title        string        `json:"title"`        // 文章ID
//		Slug         string        `json:"slug"`         // 自定也页面 path
//		Content      template.HTML `json:"content"`      // 文章的html
//		CategoryId   int           `json:"categoryId"`   // 文章的Markdown
//		CategoryName string        `json:"categoryName"` // 分类名
//		UserId       int           `json:"userId"`       // 用户id
//		UserName     string        `json:"userName"`     // 用户名
//		ViewCount    int           `json:"viewCount"`    // 查看次数
//		Type         int           `json:"type"`         // 文章类型 0 普通，1 自定义文章
//		CreateAt     string        `json:"createAt"`
//		UpdateAt     string        `json:"updateAt"`
//	}
//
// GetAllIndexInfo
//
//	@Description:
//	@param page
//	@param pageSize
//	@return *models.HomeData
//	@return error
func GetPostDetail(pId int) (*models.PostRes, error) {
	post, err := dao.GetPostById(pId)
	if err != nil {
		return nil, err
	}

	categoryName := dao.GetCategoryNameById(post.CategoryId)
	userName := dao.GetUserNameById(post.UserId)
	postMore := models.PostMore{
		Pid:          post.Pid,
		Title:        post.Title,
		Slug:         post.Slug,
		Content:      template.HTML(post.Content),
		CategoryId:   post.CategoryId,
		CategoryName: categoryName,
		UserId:       post.UserId,
		UserName:     userName,
		ViewCount:    post.ViewCount,
		Type:         post.Type,
		CreateAt:     models.DateDay(post.CreateAt),
		UpdateAt:     models.DateDay(post.UpdateAt),
	}
	//postMore.Pid = post.Pid
	var postRes = &models.PostRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		postMore,
	}
	return postRes, err

}
func SearchPost(condition string) []models.SearchResp {
	posts, _ := dao.GetPostSearch(condition)
	var searchResps []models.SearchResp
	for _, post := range posts {
		searchResps = append(searchResps, models.SearchResp{
			Pid:   post.Pid,
			Title: post.Title,
		})

	}
	return searchResps

}

func SavePost(post *models.Post) {
	dao.SavePost(post)
}
func UpdatePost(post *models.Post) {
	dao.UpdatePost(post)
}

func Writing() (wr models.WritingRes) {
	wr.Title = config.Cfg.Viewer.Title
	wr.CdnURL = config.Cfg.System.CdnURL
	category, err := dao.GetAllCategory()
	if err != nil {
		log.Println(err)
		return
	}
	wr.Categorys = category
	return
}
