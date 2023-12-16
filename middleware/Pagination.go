package middleware

import (
	"github.com/lfardell1/Go-Web-App-Blog/Database"
	"github.com/lfardell1/Go-Web-App-Blog/models"
)

func PaginateBlogResults(page uint) (models.Page, error) {
	if page == 0 {
		page = 1
	}

	var postsPerPage uint = 5

	// Retrieve posts from the database
	posts, err := Database.PaginateBlogResults(page, postsPerPage)
	if err != nil {
		return models.Page{}, err
	}

	// Ensure `TotalPages` is accurately calculated in posts

	prevPage := calculatePrevPage(page)
	nextPage := calculateNextPage(page, posts.TotalPages)

	// Handle potential errors in calculatePrevPage and calculateNextPage functions

	return models.Page{
		Posts:      posts.Posts,
		PrevPage:   prevPage,
		NextPage:   nextPage,
		CurPage:    page,
		TotalPages: posts.TotalPages,
		TotalPosts: posts.TotalPosts,
	}, nil
}

// Function to calculate the previous page number
func calculatePrevPage(page uint) uint {
	if page <= 1 {
		return 0 // Indicates no previous page
	}
	return page - 1
}

// Function to calculate the next page number
func calculateNextPage(page uint, totalPages uint) uint {
	if page >= totalPages {
		return 0 // Indicates no next page
	}
	return page + 1
}
