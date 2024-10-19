// blog.js

document.addEventListener('DOMContentLoaded', function () {
    fetchPosts();

    // Function to fetch and display posts
    async function fetchPosts() {
        try {
            const response = await fetch('http://localhost:8080/api/v1/posts', {
                headers: {
                    'Authorization': `Bearer ${localStorage.getItem('token')}`
                }
            });

            if (!response.ok) {
                throw new Error('Failed to fetch posts.');
            }

            const posts = await response.json();
            displayPosts(posts);
        } catch (error) {
            document.getElementById('error-message').textContent = error.message;
        }
    }

    // Function to display posts
    function displayPosts(posts) {
        const postsContainer = document.getElementById('posts-container');
        postsContainer.innerHTML = '';

        posts.forEach(post => {
            const postElement = document.createElement('div');
            postElement.classList.add('post');
            postElement.innerHTML = `
                <h3>${post.title}</h3>
                <p><strong>By:</strong> ${post.username || 'Unknown'}</p>
                <p><strong>Posted on:</strong> ${new Date(post.created_at).toLocaleDateString()}</p>
                <p>${post.content}</p>
                <div class="post-actions">
                    <button class="edit-btn" data-id="${post.id}">Edit</button>
                    <button class="delete-btn" data-id="${post.id}">Delete</button>
                </div>
            `;
            postsContainer.appendChild(postElement);
        });

        // Add event listeners to edit and delete buttons
        document.querySelectorAll('.edit-btn').forEach(button => {
            button.addEventListener('click', () => editPost(button.getAttribute('data-id')));
        });

        document.querySelectorAll('.delete-btn').forEach(button => {
            button.addEventListener('click', () => deletePost(button.getAttribute('data-id')));
        });
    }

    // Function to handle post editing
    async function editPost(postId) {
        const title = prompt('Enter new title:');
        const content = prompt('Enter new content:');
        if (!title || !content) return;

        try {
            const response = await fetch(`http://localhost:8080/api/v1/posts/${postId}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${localStorage.getItem('token')}`
                },
                body: JSON.stringify({ title, content })
            });

            if (!response.ok) {
                throw new Error('Failed to edit the post.');
            }

            fetchPosts(); // Refresh posts after editing
        } catch (error) {
            alert(error.message);
        }
    }

    // Function to handle post deletion
    async function deletePost(postId) {
        if (!confirm('Are you sure you want to delete this post?')) return;

        try {
            const response = await fetch(`http://localhost:8080/api/v1/posts/${postId}`, {
                method: 'DELETE',
                headers: {
                    'Authorization': `Bearer ${localStorage.getItem('token')}`
                }
            });

            if (!response.ok) {
                throw new Error('Failed to delete the post.');
            }

            fetchPosts(); // Refresh posts after deletion
        } catch (error) {
            alert(error.message);
        }
    }

    // Function to handle new post creation
    document.getElementById('create-post-form').addEventListener('submit', async function (event) {
        event.preventDefault();
        const title = document.getElementById('post-title').value;
        const content = document.getElementById('post-content').value;

        if (!title || !content) {
            alert('Please fill out both fields.');
            return;
        }

        try {
            const response = await fetch('http://localhost:8080/api/v1/posts', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${localStorage.getItem('token')}`
                },
                body: JSON.stringify({ title, content })
            });

            if (!response.ok) {
                throw new Error('Failed to create a new post.');
            }

            document.getElementById('create-post-form').reset();
            fetchPosts(); // Refresh posts after creation
        } catch (error) {
            alert(error.message);
        }
    });
});
