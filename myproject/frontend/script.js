document.addEventListener('DOMContentLoaded', () => {
    const registrationForm = document.getElementById('registrationForm');
    const loginForm = document.getElementById('loginForm');

    registrationForm.addEventListener('submit', async (event) => {
        event.preventDefault();

        const nickname = document.getElementById('nickname').value;
        const email = document.getElementById('email').value;
        const password = document.getElementById('password').value;
        const confirmPassword = document.getElementById('confirmPassword').value;

        if (password !== confirmPassword) {
            document.getElementById('registrationMessage').textContent = 'Passwords do not match';
            return;
        }

        try {
            const response = await fetch('/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ nickname, email, password }),
            });

            const data = await response.json();
            if (response.ok) {
                document.getElementById('registrationMessage').textContent = data.message;
            } else {
                document.getElementById('registrationMessage').textContent = `Error: ${data.error}`;
            }
        } catch (error) {
            console.error('Registration error:', error);
            document.getElementById('registrationMessage').textContent = 'An error occurred during registration';
        }
    });

    loginForm.addEventListener('submit', async (event) => {
        event.preventDefault();

        const nickname = document.getElementById('nickname').value;
        const password = document.getElementById('password').value;

        try {
            const response = await fetch('/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ nickname, password }),
            });

            const data = await response.json();
            if (response.ok) {
                document.getElementById('loginMessage').textContent = data.message;
                window.location.href = '/users.html'; // Redirect to users page on successful login
            } else {
                document.getElementById('loginMessage').textContent = `Error: ${data.error}`;
            }
        } catch (error) {
            console.error('Login error:', error);
            document.getElementById('loginMessage').textContent = 'An error occurred during login';
        }
    });
});
