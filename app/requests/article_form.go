package requests

import (
	"goblog/app/models/article"

	"github.com/thedevsaddam/govalidator"
)

// ValidateArticleForm 验证表单，返回 errs 长度等于零即通过
func ValidateArticleForm(data article.Article) map[string][]string {

	// 1. 定制认证规则
	rules := govalidator.MapData{
		"title":       []string{"required", "min:3", "max:40"},
		"body":        []string{"required", "min:10"},
		"category_id": []string{"reuqired", "numeric", "exists:categories,id"},
	}

	// 2. 定制错误消息
	messages := govalidator.MapData{
		"title": []string{
			"required:标题为必填项",
			"min_cn:标题长度需大于 3",
			"max_cn:标题长度需小于 40",
		},
		"body": []string{
			"required:文章内容为必填项",
			"min_cn:长度需大于 10",
		},
		"category_id": []string{
			"required:所属分类为必填项",
			"numeric:填写正确类型的所属分类",
			"exists:所属分类必须存在",
		},
	}

	// 3. 配置初始化
	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}

	// 4. 开始验证
	return govalidator.New(opts).ValidateStruct()
}
