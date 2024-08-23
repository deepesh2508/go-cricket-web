let cart = [];
let subtotal = 0;

function toggleCart() {
    const cartElement = document.querySelector(".cart");
    cartElement.classList.toggle("open");
}

function addToCart(productName, productPrice, productImage) {
    // Add item to cart array with the correct price
    let existingItem = cart.find(item => item.name === productName);

    if (existingItem) {
        existingItem.quantity += 1;
    } else {
        const newItem = {
            name: productName,
            price: productPrice,
            image: productImage,
            quantity: 1
        };
        cart.push(newItem);
    }

    subtotal += productPrice;
    updateCartDisplay();
    toggleCart();
}

function updateCartDisplay() {
    const cartItemsDiv = document.querySelector(".cart-items");
    cartItemsDiv.innerHTML = ""; // Clear existing items

    cart.forEach(item => {
        const cartItem = document.createElement("div");
        cartItem.classList.add("cart-item");

        cartItem.innerHTML = `
            <img src="${item.image}" alt="Product Image">
            <div class="item-details">
                <p>${item.name}</p>
                <div class="quantity-controls">
                    <button class="quantity-btn" onclick="updateQuantity('${item.name}', ${item.quantity - 1})">-</button>
                    <span class="quantity">${item.quantity}</span>
                    <button class="quantity-btn" onclick="updateQuantity('${item.name}', ${item.quantity + 1})">+</button>
                </div>
            </div>
            <div class="item-price">₹${item.price}</div>
        `;
        cartItemsDiv.appendChild(cartItem);
    });

    document.querySelector(".subtotal").innerText = `₹${subtotal}`;
}


function updateQuantity(productName, newQuantity) {
    let item = cart.find(item => item.name === productName);
    if (item) {
        if (newQuantity <= 0) {
            removeFromCart(productName);
            return;
        }
        let difference = item.price * (newQuantity - item.quantity);
        item.quantity = newQuantity;
        subtotal += difference;
        updateCartDisplay();
    }
}

function removeFromCart(productName) {
    let itemIndex = cart.findIndex(item => item.name === productName);
    if (itemIndex > -1) {
        subtotal -= cart[itemIndex].price * cart[itemIndex].quantity;
        cart.splice(itemIndex, 1);
        updateCartDisplay();
    }
}

document.querySelector(".close-cart").addEventListener("click", toggleCart);

// Handle checkout button click
document.querySelector('.checkout-btn').addEventListener('click', function() {
    window.location.href = 'checkout.html';
});
