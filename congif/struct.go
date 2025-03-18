package main

import (
    "time"
)

type User struct {
    ID       int
    Username string
    Email    string
    Password string 
    Avatar   string
    Role     string 
    Created  time.Time
}


type Post struct {
    ID          int
    Title       string
    Content     string
    UserID      int
    Username    string
    Created     time.Time
    Modified    time.Time
    ImagePath   string
    LikesCount  int
    DislikeCount int
    Categories  []Category
}


type Category struct {
    ID   int
    Name string
}


type Comment struct {
    ID        int
    Content   string
    PostID    int
    UserID    int
    Username  string
    Created   time.Time
    Modified  time.Time
}

type Like struct {
    ID     int
    PostID int
    UserID int
    Type   int 
}

type Session struct {
    UUID      string
    UserID    int
    Username  string
    Role      string
    CreatedAt time.Time
    ExpiresAt time.Time
}


type FilterOptions struct {
    Category    string
    User        string
    LikedByUser string
    SortBy      string
}

type Response struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}