let previewContainers = document.querySelectorAll(".preview-img-container")
let activeIndex = 0
let nextIndex = 1
let prevIndex = 4

function updateCard() {
    getNextCardIndex()
    getPrevCardIndex()
    previewContainers.forEach(preview => {
        preview.classList.remove("active", "next", "prev")
    })
    previewContainers[activeIndex].classList.add("active")
    previewContainers[nextIndex].classList.add("next")
    previewContainers[prevIndex].classList.add("prev")
}

function getNextCardIndex() {
    nextIndex = activeIndex + 1
    if (nextIndex > (previewContainers.length - 1)) {
        nextIndex = 0
    }
}

function getPrevCardIndex() {
    prevIndex = activeIndex - 1
    if (prevIndex < 0) {
        prevIndex = (previewContainers.length - 1)
    }
}

function showPrevCard() {
    activeIndex -= 1
    if (activeIndex < 0) {
        activeIndex = (previewContainers.length - 1)
    }
    updateCard()
}

function showNextCard() {
    activeIndex += 1
    if (activeIndex > (previewContainers.length - 1)) {
        activeIndex = 0
    }
    updateCard()
}
document.getElementById("prev-button").addEventListener("click", showPrevCard)
document.getElementById("next-button").addEventListener("click", showNextCard)

updateCard()

