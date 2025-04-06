document.addEventListener('scroll', () => {
    const isScrolled = window.scrollY > 10;
    document.body.classList.toggle('is-scrolled', isScrolled);
});

const menuButtons = document.querySelectorAll('.top-nav-menu-button');
menuButtons.forEach(button => {
    button.addEventListener("click", () => {
        menuButtons.forEach(otherButton => {
            if (otherButton != button) {
                otherButton.classList.remove("open-menu")
            }
        })
        if (button.classList.contains("open-menu")) {
            button.classList.remove("open-menu")
        } else {
            button.classList.add("open-menu")
        }
    })
})

document.addEventListener("click", function (e) {
    if (!e.target.closest('.top-nav-menu-button')) {
        const menuButtons = document.querySelectorAll('.top-nav-menu-button');
        menuButtons.forEach(button => {
            button.classList.remove("open-menu")
        })
    }
})
