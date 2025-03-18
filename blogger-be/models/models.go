package models

import (
	"github.com/Akash-Manikandan/blogger-be/internal/utils"

	"github.com/lucsky/cuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string        `gorm:"primaryKey;size:25" json:"id"`
	CreatedAt utils.ISOTime `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt utils.ISOTime `gorm:"autoUpdateTime" json:"updatedAt"`

	Username    string `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Email       string `gorm:"uniqueIndex;size:100;not null" json:"email"`
	Password    string `gorm:"not null" json:"-"`
	Age         int    `json:"age,omitempty"`
	Description string `gorm:"type:text;" json:"description,omitempty"`

	// Relationships with CASCADE delete
	Blogs    []Blog    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"blogs,omitempty"`
	Shares   []Share   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"shares,omitempty"`
	Comments []Comment `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"comments,omitempty"`
	Likes    []Like    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"likes,omitempty"`
}

// BeforeCreate - Hash password before inserting into DB
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	if u.ID == "" {
		u.ID = cuid.New()
	}
	u.Password = hashedPassword
	return nil
}

// BeforeUpdate - Hash password only if changed
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	var existingUser User
	if err := tx.First(&existingUser, "id = ?", u.ID).Error; err != nil {
		return err
	}

	// Hash only if password has changed
	if u.Password != existingUser.Password {
		hashedPassword, err := utils.HashPassword(u.Password)
		if err != nil {
			return err
		}
		u.Password = hashedPassword
	}
	return nil
}

type Blog struct {
	ID            string        `gorm:"primaryKey;size:25" json:"id"`
	CreatedAt     utils.ISOTime `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt     utils.ISOTime `gorm:"autoUpdateTime" json:"updatedAt"`
	UserID        string        `gorm:"index;not null" json:"userId"`
	Title         string        `gorm:"size:200;not null" json:"title"`
	Content       string        `gorm:"type:text;not null" json:"content"`
	IsPrivate     bool          `gorm:"default:true;not null" json:"isPrivate"`
	TrendingCount int           `gorm:"default:0" json:"trendingCount"`
	ViewCount     int           `gorm:"default:0" json:"viewCount"`

	// Add CASCADE delete constraints to all child relationships
	User     User      `gorm:"foreignKey:UserID" json:"user"`
	Shares   []Share   `gorm:"foreignKey:BlogID;constraint:OnDelete:CASCADE" json:"shares,omitempty"`
	Comments []Comment `gorm:"foreignKey:BlogID;constraint:OnDelete:CASCADE" json:"comments,omitempty"`
	Likes    []Like    `gorm:"foreignKey:BlogID;constraint:OnDelete:CASCADE" json:"likes,omitempty"`
	Tags     []Tag     `gorm:"many2many:blog_tags;"`
}

func (b *Blog) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = cuid.New()
	}
	return nil
}

type Share struct {
	ID        string        `gorm:"primaryKey;size:25" json:"id"`
	CreatedAt utils.ISOTime `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt utils.ISOTime `gorm:"autoUpdateTime" json:"updatedAt"`
	BlogID    string        `gorm:"index;not null" json:"blogId"`
	UserID    string        `gorm:"index;not null" json:"userId"`

	// Define constraints here
	Blog Blog `gorm:"foreignKey:BlogID;constraint:OnDelete:CASCADE" json:"blog"`
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"`
}

func (s *Share) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == "" {
		s.ID = cuid.New()
	}
	return nil
}

type Comment struct {
	ID              string        `gorm:"primaryKey;size:25" json:"id"`
	CreatedAt       utils.ISOTime `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt       utils.ISOTime `gorm:"autoUpdateTime" json:"updatedAt"`
	BlogID          string        `gorm:"index;not null" json:"blogId"`
	UserID          string        `gorm:"index;not null" json:"userId"`
	ParentCommentID *string       `gorm:"index" json:"parentCommentId,omitempty"`
	Content         string        `gorm:"type:text;not null" json:"content"`

	// Define constraints here
	Blog    Blog      `gorm:"foreignKey:BlogID;constraint:OnDelete:CASCADE" json:"blog"`
	User    User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"`
	Parent  *Comment  `gorm:"foreignKey:ParentCommentID;constraint:OnDelete:SET NULL" json:"parent,omitempty"`
	Replies []Comment `gorm:"foreignKey:ParentCommentID" json:"replies,omitempty"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		c.ID = cuid.New()
	}
	return nil
}

type Like struct {
	ID        string        `gorm:"primaryKey;size:25" json:"id"`
	CreatedAt utils.ISOTime `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt utils.ISOTime `gorm:"autoUpdateTime" json:"updatedAt"`
	BlogID    string        `gorm:"uniqueIndex:idx_user_blog;not null" json:"blogId"`
	UserID    string        `gorm:"uniqueIndex:idx_user_blog;not null" json:"userId"`

	// Define constraints here
	Blog Blog `gorm:"foreignKey:BlogID;constraint:OnDelete:CASCADE" json:"blog"`
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"`
}

func (l *Like) BeforeCreate(tx *gorm.DB) (err error) {
	if l.ID == "" {
		l.ID = cuid.New()
	}
	return nil
}

type Tag struct {
	ID    string `gorm:"primarykey;type:varchar(25)"`
	Name  string `gorm:"unique;not null"`
	Blogs []Blog `gorm:"many2many:blog_tags;"`
}

func (t *Tag) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == "" {
		t.ID = cuid.New()
	}
	return nil
}
