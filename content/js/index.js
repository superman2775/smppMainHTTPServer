const previewContainers = document.querySelectorAll(".preview-img-container");
const prevButton = document.getElementById("prev-button");
const nextButton = document.getElementById("next-button");
let activeIndex = 1;

if (!previewContainers.length || !prevButton || !nextButton) console.error("Required elements not found");

const getNextIndex = (currentIndex) => (currentIndex + 1) % previewContainers.length;
const getPrevIndex = (currentIndex) => (currentIndex - 1 + previewContainers.length) % previewContainers.length;

function updateCard() {
    previewContainers.forEach(preview => {
        preview.classList.remove("active", "next", "prev");
    });

    const nextIndex = getNextIndex(activeIndex);
    const prevIndex = getPrevIndex(activeIndex);

    previewContainers[activeIndex].classList.add("active");
    previewContainers[nextIndex].classList.add("next");
    previewContainers[prevIndex].classList.add("prev");
    previewContainers[activeIndex].addEventListener("transitionend", function transitionEndHandler() {
        previewContainers.forEach(preview => { preview.classList.remove("ready"); });
        previewContainers[activeIndex].classList.add("ready");
        previewContainers[activeIndex].removeEventListener("transitionend", transitionEndHandler);
    });
}

function showPrevCard() {
    activeIndex = getPrevIndex(activeIndex);
    updateCard();
}

function showNextCard() {
    activeIndex = getNextIndex(activeIndex);
    updateCard();
}

prevButton.addEventListener("click", showPrevCard);
nextButton.addEventListener("click", showNextCard);

updateCard();