package main

import "fmt"

const numDisks = 3

func main() {
	// Make three posts.
	posts := [][]int{}

	// Push the disks onto post 0 biggest first.
	posts = append(posts, []int{})
	for disk := numDisks; disk > 0; disk-- {
		posts[0] = push(posts[0], disk)
	}

	// Make the other posts empty.
	for p := 1; p < 3; p++ {
		posts = append(posts, []int{})
	}

	// Draw the initial setup.
	drawPosts(posts)

	// Move the disks.
	moveDisks(posts, numDisks, 0, 1, 2)
}

// Add a disk to beginning of the post.
func push(post []int, disk int) []int {
	return append([]int{disk}, post...)
}

// Remove the first disk from the post.
// Return that disk and the revised post.
func pop(post []int) (int, []int) {
	return post[0], post[1:]
}

// Move one disk from fromPost to toPost.
func moveDisk(posts [][]int, fromPost, toPost int) {
	var disk int
	disk, posts[fromPost] = pop(posts[fromPost])
	posts[toPost] = push(posts[toPost], disk)
}

func drawPosts(posts [][]int) {
	// Add 0s to the end of each post, so they all have numDisks entries.
	for p := 0; p < 3; p++ {
		for len(posts[p]) < numDisks {
			posts[p] = push(posts[p], 0)
		}
	}

	// Draw the posts.
	for row := 0; row < numDisks; row++ {
		// Draw this row.
		for p := 0; p < 3; p++ {
			// Draw the disk on post p's row.
			fmt.Printf("%d ", posts[p][row])
		}
		fmt.Println()
	}

	// Draw a line between moves.
	fmt.Println("-----")

	// Remove the 0s.
	for p := 0; p < 3; p++ {
		for len(posts[p]) > 0 && posts[p][0] == 0 {
			_, posts[p] = pop(posts[p])
		}
	}
}

func moveDisks(posts [][]int, numToMove, fromPost, toPost, tempPost int) {
	// If there is more than one disk to move,
	// recursively move numToMove - 1 disks from the source post to the temporary post.
	if numToMove > 1 {
		moveDisks(posts, numToMove-1, fromPost, tempPost, toPost)
	}

	// Move the remaining disk from the source post to the destination post.
	moveDisk(posts, fromPost, toPost)

	// Draw the posts to show the current state.
	drawPosts(posts)

	// If disks were moved to the temporary post,
	// move them from the temporary post to the destination post.
	if numToMove > 1 {
		moveDisks(posts, numToMove-1, tempPost, toPost, fromPost)
	}
}
