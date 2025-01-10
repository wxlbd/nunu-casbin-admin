package response

// MenuTree 菜单树节点
type MenuTreeNode struct {
	ID        uint64          `json:"id"`
	ParentID  uint64          `json:"parent_id"`
	Name      string          `json:"name"`
	Meta      Meta            `json:"meta"`
	Path      string          `json:"path"`
	Component string          `json:"component"`
	Redirect  string          `json:"redirect"`
	Status    int8            `json:"status"`
	Sort      int16           `json:"sort"`
	CreatedBy int             `json:"created_by"`
	UpdatedBy int             `json:"updated_by"`
	CreatedAt string          `json:"created_at"`
	UpdatedAt string          `json:"updated_at"`
	Remark    string          `json:"remark"`
	Children  []*MenuTreeNode `json:"children"`
}

type Meta struct {
	I18N             string `json:"i18n"`
	Icon             string `json:"icon"`
	Type             string `json:"type"`
	Affix            bool   `json:"affix"`
	Cache            bool   `json:"cache"`
	Title            string `json:"title"`
	Hidden           bool   `json:"hidden"`
	Copyright        bool   `json:"copyright"`
	ComponentPath    string `json:"componentPath"`
	ComponentSuffix  string `json:"componentSuffix"`
	BreadcrumbEnable bool   `json:"breadcrumbEnable"`
}

// MenuTreeResponse 菜单列表响应
type MenuTreeResponse []*MenuTreeNode

// Menu 菜单信息响应
type Menu struct {
	ID         uint64 `json:"id"`
	ParentID   uint64 `json:"parent_id"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	Component  string `json:"component"`
	Sort       int16  `json:"sort"`
	Status     int8   `json:"status"`
	Mate       any    `json:"mate"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
