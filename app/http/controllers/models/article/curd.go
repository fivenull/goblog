package article

import (
	"fmt"
	"goblog/pkg/model"
	"goblog/pkg/types"
)

// Get 通过 ID 获取文章
func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToUint64(idstr)
	fmt.Println(123)
	fmt.Println("id", id)
	fmt.Println("articleerr", model.DB.First(&article))
	if err := model.DB.First(&article).Error; err != nil {
		fmt.Println("article", err)
		return article, err
	}

	return article, nil
}
