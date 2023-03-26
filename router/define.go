package router

import (
	"gorm.io/plugin/soft_delete"
)

type Content struct {
	Rules           []Rule      `json:"rules"`
	HTTPDirectives  []Directive `json:"http_directives"`
	HTTPSDirectives []Directive `json:"https_directives"`
}

type Directive struct {
	Name      string   `json:"name" binding:"required"`
	Arguments []string `json:"arguments"`
}

type Rule struct {
	Name           string      `json:"name" binding:"required"`
	MatchType      string      `json:"match_type" binding:"required"`
	MatchPath      string      `json:"match_path" binding:"required"`
	UsedInHTTP     bool        `json:"used_in_http"`
	UsedInHTTPS    bool        `json:"used_in_https"`
	TargetName     string      `json:"target_name"`
	TargetProtocol string      `json:"target_protocol"`
	Directives     []Directive `json:"directives"`
}

type Router struct {
	ID           uint                  `gorm:"column:id" json:"-"`
	InstanceName string                `gorm:"column:instance_name" json:"instance_name" binding:"required"`
	Domain       string                `gorm:"column:domain" json:"domain" binding:"required"`
	EnableHTTP   bool                  `gorm:"column:enable_http" json:"enable_http"`
	EnableHTTPS  bool                  `gorm:"column:enable_https" json:"enable_https"`
	CertName     string                `gorm:"column:cert_name" json:"cert_name"`
	Content      Content               `gorm:"column:content" json:"content"`
	CreatedAt    int64                 `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    int64                 `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy    string                `gorm:"column:updated_by" json:"updated_by"`
	DeletedAt    soft_delete.DeletedAt `json:"-"`
}
