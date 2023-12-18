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
        <span class="Paginate" hx-get="/api/Blogs/%d" hx-target="#leftcolumn" hx-swap="innerHTML">PREVIOUS</span>
                        <span class="Paginate">Page %d</span>	
		<span class="Paginate" hx-get="/api/Blogs/%d" hx-target="#leftcolumn" hx-swap="innerHTML">NEXT</span>
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

		html += "<div class='card' id='blog'>"
		html += "<section class='Details'>"
		html += "<h3 class='authorbox'>"
		html += "<div class='beforename'>Author:</div>"
		html += "<div class='author'>" + post.Author + "</div>"
		html += "</h3>"
		html += "<section class='title'>"
		html += "<h4 class='titlelead'>Title:</h4>"
		html += "<h2 class='title'>" + post.Title + "</h2>"
		html += "</section>"
		html += "</section>"
		html += "<hr>"
		html += "<p class='Content'><b>" + post.Content + "</b></p>"
		html += "<h5>Published: " + post.DateCreated.Format("Jan 02, 2006") + "</h5>"
		html += "<div class='bottom-row'>"
		html += "<img class='img' width='500' height='700' src='" + post.ImageURL + "'></img>"
		html += "<span class='comment-section'><span class='comment-heading'>Comments:</span></span>"
		html += "</div>"
		html += "</div>"
	}

	html += "</div>"

	finalHTML := paginationHTML + html

	w.Header().Set("Content-Type", "text/html")

	_, err = w.Write([]byte(finalHTML))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
