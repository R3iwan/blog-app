document.addEventListener('DOMContentLoaded', function () {
    const loginForm = document.getElementById('login-form');
    if (loginForm) {
        loginForm.addEventListener('submit', async function (event) {
            event.preventDefault(); 

            const username = document.getElementById('username').value;
            const password = document.getElementById('password').value;
            const role = document.getElementById('role').value;

            try {
                const response = await fetch('http://localhost:8080/api/v1/login', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ username, password, role}), 
                });

                if (!response.ok) {
                    throw new Error('An error occurred during login.');
                }

                const data = await response.json();
                if (data.access_token) {
                    localStorage.setItem('token', data.access_token);
                    window.location.href = 'blog.html'; 
                } else {
                    throw new Error('Login failed. Please check your credentials.');
                }
            } catch (error) {
                document.getElementById('login-error').textContent = error.message;
            }
        });
    }
});
