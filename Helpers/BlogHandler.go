package Helpers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lfardell1/Go-Web-App-Blog/middleware"
)

func renderPaginationSection(page int) string {

	paginationHTML := fmt.Sprintf(`
    <section class="Pagination">
        <button class="Paginate" hx-get="/api/Blogs/%d" hx-target="#blogs" hx-swap="innerHTML">PREVIOUS</button>
                        <span class="Paginate">Page %d</span>	
		<button class="Paginate" hx-get="/api/Blogs/%d" hx-target="#blogs" hx-swap="innerHTML">NEXT</button>
    </section>`, (page - 1), page, (page + 1))

	return paginationHTML
}

func RenderBlogPostsPage(w http.ResponseWriter, page int) {
	// Fetch blog posts for the requested page using your middleware function
	posts, err := middleware.PaginateBlogResults(uint(page))
	if err != nil {
		log.Printf("Error retrieving posts: %v", err)
		http.Error(w, "Error retrieving posts", http.StatusInternalServerError)
		return
	}

	paginationHTML := renderPaginationSection(page)

	html := "<div class='blog-posts'>"

	for _, post := range posts.Posts {

		html += fmt.Sprintf(`<div class="w3-card">
		<img src="%s alt="Nature">
		<div class="w3-container">
		  <h3><b>%s</b></h3>
		  <h5>%s <span class="w3-opacity">%s</span></h5>
		</div>
	
		<div class="w3-container">
		  <p>%s</p>
		  <div class="w3-row">
			<div class="w3-col m8 s12">
			  <p><button class="flat-button"><b>READ MORE »</b></button></p>
			</div>
			<div class="w3-col m4 w3-hide-small">
			  <p><span class="w3-padding-large w3-right"><b>Comments  </b> <span class="w3-tag">0</span></span></p>
			</div>
		  </div>
		</div>
	  </div>
	  <hr>`, post.ImageURL, post.Title, post.Author, post.DateCreated.Format("January 2, 2006"), post.Content)
	}

	finalHTML := paginationHTML + html

	w.Header().Set("Content-Type", "text/html")

	_, err = w.Write([]byte(finalHTML))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
