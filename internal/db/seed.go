package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"social/internal/store"
)

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()
	tx, _ := db.BeginTx(ctx, nil)

	users := generateUsers(100)
	for _, user := range users {
		if err := store.Users.Create(ctx, tx, user); err != nil {
			_ = tx.Rollback()
			log.Printf("Error creating user: %v", err)
			return
		}
	}

	tx.Commit()

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Printf("Error creating post: %v", err)
			log.Printf("Post: %+v", post)
			return
		}
	}

	comments := generateComments(500, posts, users)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Printf("Error creating comment: %v", err)
			return
		}
	}

	log.Println("Database seeded successfully")
}

func generateUsers(n int) []*store.User {
	users := make([]*store.User, n)
	for i := 0; i < n; i++ {
		users[i] = &store.User{
			Username: "user" + fmt.Sprint(i),
			Email:    "user" + fmt.Sprint(i) + "@example.com",
			Role: store.Role{
				Name: "user",
			}, // Assuming role ID 1 is for regular users
		}
	}
	return users
}

func generatePosts(n int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, n)
	for i := 0; i < n; i++ {
		user := users[rand.Intn(len(users))] // Randomly select a user
		posts[i] = &store.Post{
			Title:   "Post Title " + fmt.Sprint(i),
			Content: "Post Content " + fmt.Sprint(i),
			UserId:  user.ID, // assumes users have valid IDs
			Tags:    []string{"tag1", "tag2"},
		}
	}
	return posts
}

func generateComments(n int, posts []*store.Post, users []*store.User) []*store.Comment {
	comments := make([]*store.Comment, n)
	for i := 0; i < n; i++ {
		comments[i] = &store.Comment{
			Content: "Comment Content " + fmt.Sprint(i),
			UserID:  users[i%len(users)].ID,
			PostID:  posts[i%len(posts)].ID,
		}
	}
	return comments
}
