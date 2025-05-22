const sections = document.querySelectorAll("#content section");
const buttons = document.querySelectorAll(".guide-button");
const topbar = document.querySelector("header");
const codeBlocks = document.querySelectorAll("code");
const getTopOffset = () =>
  window.matchMedia("(max-width: 1170px)").matches ? 80 : 140;

let scrollLock = false;

buttons.forEach((button) => {
  button.addEventListener("click", () => {
    const targetId = button.getAttribute("data-target");
    const target = document.getElementById(targetId);
    const topOffset = getTopOffset();
    const targetPosition = target.offsetTop - topOffset;

    scrollLock = true;

    buttons.forEach((btn) => btn.classList.remove("active"));
    button.classList.add("active");

    window.scrollTo({
      top: targetPosition,
      behavior: "smooth",
    });

    setTimeout(() => {
      scrollLock = false;
      updateActiveButton();
    }, 1000);
  });
});

function updateActiveButton() {
  if (scrollLock) return;

  const viewportHeight = window.innerHeight;
  const topOffset = getTopOffset();
  let maxVisibleArea = 0;
  let currentSectionId = "";

  sections.forEach((section) => {
    const rect = section.getBoundingClientRect();
    const visibleTop = Math.max(rect.top, topOffset);
    const visibleBottom = Math.min(rect.bottom, viewportHeight);
    const visibleHeight = Math.max(0, visibleBottom - visibleTop);

    if (visibleHeight > maxVisibleArea) {
      maxVisibleArea = visibleHeight;
      currentSectionId = section.id;
    }
  });

  buttons.forEach((button) => {
    button.classList.toggle(
      "active",
      button.getAttribute("data-target") === currentSectionId
    );
  });
}

codeBlocks.forEach((codeBlock) => {
  codeBlock.addEventListener("click", () => {
    const text = codeBlock.innerText;
    navigator.clipboard
      .writeText(text)
      .then(() => {
        const toast = document.getElementById("toast");
        toast.classList.add("toast-visible");
        setTimeout(() => {
          toast.classList.remove("toast-visible");
        }, 2000);
      })
      .catch((err) => console.error("Failed to copy: ", err));
  });
});

window.addEventListener("scroll", updateActiveButton);
window.addEventListener("resize", updateActiveButton);

function toggleMobileGuideMenu() {
  document.getElementById("guide-menu").classList.toggle("active");
  document.getElementById("overlay").classList.toggle("active");
  console.log(document.getElementById("overlay"));
}

document
  .getElementById("guide-menu-toggle")
  .addEventListener("click", toggleMobileGuideMenu);
