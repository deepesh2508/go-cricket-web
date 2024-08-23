const leftButton = document.querySelectorAll('.scroll-button.left');
const rightButton = document.querySelectorAll('.scroll-button.right');
const categoryGrid = document.querySelector('.category-grid');
const productGrid = document.querySelector('.product-grid');

// Scroll categories left
leftButton[0].addEventListener('click', () => {
    categoryGrid.scrollBy({
        top: 0,
        left: -200,
        behavior: 'smooth'
    });
});

// Scroll categories right
rightButton[0].addEventListener('click', () => {
    categoryGrid.scrollBy({
        top: 0,
        left: 200,
        behavior: 'smooth'
    });
});

// Scroll products left
leftButton[1].addEventListener('click', () => {
    productGrid.scrollBy({
        top: 0,
        left: -200,
        behavior: 'smooth'
    });
});

// Scroll products right
rightButton[1].addEventListener('click', () => {
    productGrid.scrollBy({
        top: 0,
        left: 200,
        behavior: 'smooth'
    });
});

// Get the modal elements
const loginModal = document.getElementById('login-modal');
const signupModal = document.getElementById('signup-modal');

// Get the button that opens the login modal
const loginBtn = document.querySelector('.right-side .login');

// Get the <span> elements that close the modals
const closeButtons = document.querySelectorAll('.modal .close');

// Get the signup link inside the login modal
const signupLink = document.querySelector('.signup-link');

// When the user clicks on the login button, open the login modal
loginBtn.addEventListener('click', () => {
    loginModal.style.display = 'block';
});

// When the user clicks on <span> (x), close the modal
closeButtons.forEach(button => {
    button.addEventListener('click', () => {
        loginModal.style.display = 'none';
        signupModal.style.display = 'none';
    });
});

// When the user clicks anywhere outside of the modal, close it
window.addEventListener('click', event => {
    if (event.target === loginModal) {
        loginModal.style.display = 'none';
    }
    if (event.target === signupModal) {
        signupModal.style.display = 'none';
    }
});

// When the user clicks on the signup link inside the login modal, switch to the signup modal
signupLink.addEventListener('click', event => {
    event.preventDefault();
    loginModal.style.display = 'none';
    signupModal.style.display = 'block';
});

const testimonials = document.querySelectorAll('.testimonial');
const dots = document.querySelectorAll('.dot');

let activeIndex = 0;

// Function to update testimonials and dots
function updateTestimonials(index) {
    testimonials.forEach((testimonial, i) => {
        testimonial.classList.remove('active');
        dots[i].classList.remove('active');
        if (i === index) {
            testimonial.classList.add('active');
            dots[i].classList.add('active');
        }
    });
}

// Initialize first testimonial as active
updateTestimonials(activeIndex);

// Add event listeners to dots
dots.forEach((dot, index) => {
    dot.addEventListener('click', () => {
        activeIndex = index;
        updateTestimonials(activeIndex);
    });
});

// Handle signup
document.getElementById("signup-btn").addEventListener("click", function(event) {
    console.log("Signup button clicked");
    event.preventDefault();

    const name = document.getElementById("signup-name").value;
    const email = document.getElementById("signup-email").value;
    const mobile = document.getElementById("signup-mobile").value;
    const password = document.getElementById("signup-password").value;
    const confirmPassword = document.getElementById("signup-confirm-password").value;

    if (password !== confirmPassword) {
        alert("Passwords do not match!");
        return;
    }

    fetch('http://127.0.0.1:8080/users/signup', {  // Ensure this matches your route
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            name: name,
            email: email,
            mobile: mobile,
            password: password
        })
    })
    .then(response => {
        if (!response.ok) {
            return response.json().then(err => { throw err; });
        }
        return response.json();
    })
    .then(data => {
        if (data.id) {
            alert('Signup successful');
            // Optionally redirect or close modal
        } else {
            alert('Signup failed: ' + data.error);
        }
    })
    .catch(error => {
        console.error('Error:', error);
        alert('Signup failed: ' + error.message);
    });
});


// Handle login
document.getElementById("login-btn").addEventListener("click", function(event) {
    event.preventDefault();
    const email = document.getElementById("login-email").value;
    const password = document.getElementById("login-password").value;

    fetch('localhost:8000/login', {  // Ensure this matches your route
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            email: email,
            password: password
        })
    })
    .then(response => response.json())
    .then(data => {
        if (data.user_id) {
            alert('Login successful');
            // Optionally redirect or close modal
        } else {
            alert('Login failed: ' + data.error);
        }
    })
    .catch(error => console.error('Error:', error));
});
