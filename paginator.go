package webutils

import (
	"math"
	"net/http"
	"net/url"
	"strconv"
)

// Pagintor 包含了 http request.
type Paginator struct {
	Request     *http.Request
	PerPageNums int
	MaxPages    int

	nums      int64
	pageRange []int
	pageNums  int
	page      int
}

// 返回总共的页数
func (p *Paginator) PageNums() int {
	if p.pageNums != 0 {
		return p.pageNums
	}
	pageNums := math.Ceil(float64(p.nums) / float64(p.PerPageNums))
	if p.MaxPages > 0 {
		pageNums = math.Min(pageNums, float64(p.MaxPages))
	}
	p.pageNums = int(pageNums)
	return p.pageNums
}

//返回子项的个数
func (p *Paginator) Nums() int64 {
	return p.nums
}

// 设置子项的个数
func (p *Paginator) SetNums(nums interface{}) {
	p.nums, _ = ToInt64(nums)
}

// 返回当前的页
func (p *Paginator) Page() int {
	if p.page != 0 {
		return p.page
	}
	if p.Request.Form == nil {
		p.Request.ParseForm()
	}
	p.page, _ = strconv.Atoi(p.Request.Form.Get("p"))
	if p.page > p.PageNums() {
		p.page = p.PageNums()
	}
	if p.page <= 0 {
		p.page = 1
	}
	return p.page
}

// 返回全部页数的计划
// 使用方法
//  {{range $index, $page := .paginator.Pages}}
//    <li{{if $.paginator.IsActive .}} class="active"{{end}}>
//      <a href="{{$.paginator.PageLink $page}}">{{$page}}</a>
//    </li>
//  {{end}}

func (p *Paginator) Pages() []int {
	if p.pageRange == nil && p.nums > 0 {
		var pages []int
		pageNums := p.PageNums()
		page := p.Page()
		switch {
		case page >= pageNums-4 && pageNums > 9:
			start := pageNums - 9 + 1
			pages = make([]int, 9)
			for i := range pages {
				pages[i] = start + i
			}
		case page >= 5 && pageNums > 9:
			start := page - 5 + 1
			pages = make([]int, int(math.Min(9, float64(page+4+1))))
			for i := range pages {
				pages[i] = start + i
			}
		default:
			pages = make([]int, int(math.Min(9, float64(pageNums))))
			for i := range pages {
				pages[i] = i + 1
			}
		}
		p.pageRange = pages
	}
	return p.pageRanges
}

// 根据请求返回 URL
func (p *Paginator) PageLink(page int) string {
	link, _ := url.ParseRequestURI(p.Request.URL.String())
	values := link.Query()
	if page == 1 {
		values.Del("p")
	} else {
		values.Set("p", strconv.Itoa(page))
	}
	link.RawQuery = values.Encode()
	return link.String()
}

// 返回上一页
func (p *Paginator) PageLinkPrev() (link string) {
	if p.HasPrev() {
		link = p.PageLink(p.Page() - 1)
	}
	return
}

// 返回下一页
func (p *Paginator) PageLinkNext() (link string) {
	if p.HasNext() {
		link = p.PageLink(p.Page() + 1)
	}
	return
}

// 返回第一页
func (p *Paginator) PageLinkFirst() (link string) {
	return p.PageLink(1)
}

// 返回最后一页
func (p *Paginator) PageLinkLast() (link string) {
	return p.PageLink(p.PageNums())
}

// 检查当前页是否有上一页
func (p *Paginator) HasPrev() bool {
	return p.Page() > 1
}

// 检查当前页是否有下一页
func (p *Paginator) HasNext() bool {
	return p.Page() < p.PageNums()
}

// 检查当前页数是不是指向该页
func (p *Paginator) IsActive(page int) bool {
	return p.Page() == page
}

// 返回当前的偏移量
func (p *Paginator) Offset() int {
	return (p.Page() - 1) * p.PerPageNums
}

// 返回页数是否超过一页
func (p *Paginator) HasPages() bool {
	return p.PageNums() > 1
}

// 实例化
func NewPaginator(req *http.Request, per int, num interface{}) *Paginator {
	p := Paginator{}
	p.Request = req
	if per <= 0 {
		per = 10
	}
	p.PerPageNums = per
	p.SetNums(nums)
	return &p
}
