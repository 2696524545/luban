package k8s

import "strings"

type K8sListObj struct {
	Kind       string        `json:"kind"`
	ApiVersion string        `json:"apiVersion"`
	Metadata   interface{}   `json:"metadata"`
	Items      []interface{} `json:"items"`
}

type Request struct {
	Search        string `form:"search"`
	FieldSelector string `form:"fieldSelector"`
	LabelSelector string `form:"labelSelector"`
	Page          int    `form:"page"`
	PageSize      int    `form:"pageSize"`
}

type Result struct {
	Items    interface{} `json:"items"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

func pageFilter(page, pageSize int, data []interface{}) (t int64, r []interface{}, err error) {
	total := len(data)
	result := make([]interface{}, 0)
	if page*pageSize < len(data) {
		result = data[(page-1)*pageSize : (page * pageSize)]
	} else {
		result = data[(page-1)*pageSize:]
	}
	return int64(total), result, nil
}

// pagerAndSearch 分页/搜索
func pagerAndSearch(page, pageSize int, items []interface{}, keywords string) (*Result, error) {
	var result Result
	if keywords != "" {
		items = fieldFilter(items, withNamespaceAndNameMatcher(keywords))
	}
	if page != 0 && pageSize != 0 {
		total, items, err := pageFilter(page, pageSize, items)
		if err != nil {
			return nil, err
		}
		result.Total = total
		result.Items = items
		result.Page, result.PageSize = page, pageSize

		return &result, nil
	}

	result.Total = int64(len(items))
	result.Items = items
	result.Page = page
	result.PageSize = pageSize
	return &result, nil
}

type fieldMatcher interface {
	Match(item interface{}) bool
}

type keywordsMatcher struct {
	keywords string
}

func (n keywordsMatcher) Match(item interface{}) bool {
	pageItem := item.(map[string]interface{})
	if pageItem["metadata"].(map[string]interface{})["namespace"] != nil && pageItem["metadata"].(map[string]interface{})["namespace"].(string) == n.keywords {
		return true
	}
	if strings.Contains(pageItem["metadata"].(map[string]interface{})["name"].(string), n.keywords) {
		return true
	}
	if pageItem["message"] != nil && strings.Contains(strings.ToLower(pageItem["message"].(string)), strings.ToLower(n.keywords)) {
		return true
	}
	return false
}

func withNamespaceAndNameMatcher(keywords string) fieldMatcher {
	return &keywordsMatcher{
		keywords: keywords,
	}
}

// fieldFilter 字段过滤
func fieldFilter(data []interface{}, fms ...fieldMatcher) []interface{} {
	var result []interface{}
	for i := range data {
		for j := range fms {
			if fms[j].Match(data[i]) {
				result = append(result, data[i])
				break
			}
		}
	}
	return result
}
