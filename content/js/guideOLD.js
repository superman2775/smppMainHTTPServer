const codeBlocks = document.querySelectorAll("code.copyable");
const tabBtns = document.querySelectorAll(".tab-btn");
const docSections = document.querySelectorAll(".doc-section");
const sidebar = document.getElementById("sidebar");
const mobileMenuTrigger = document.getElementById("mobileMenuTrigger");
const menuCloseBtn = document.getElementById("menuCloseBtn");
const sidebarOverlay = document.getElementById("sidebarOverlay");
const content = document.getElementById("content");
let isMobile = window.innerWidth <= 768;

function updateLayout() {
  isMobile = window.innerWidth <= 768;
  if (isMobile) {
    mobileMenuTrigger.classList.add("mobile-visible");
    menuCloseBtn.classList.add("mobile-visible");
    sidebar.classList.add("mobile-sidebar");
    content.classList.add("mobile-content");
  } else {
    mobileMenuTrigger.classList.remove("mobile-visible");
    menuCloseBtn.classList.remove("mobile-visible");
    sidebar.classList.remove("mobile-sidebar");
    content.classList.remove("mobile-content");
    sidebarOverlay.style.display = "none";
  }
}

function openMobileMenu() {
  sidebar.classList.add("sidebar-open");
  sidebarOverlay.style.display = "block";
  document.body.style.overflow = "hidden";
  mobileMenuTrigger.classList.remove("mobile-visible");
}

function closeMobileMenu() {
  sidebar.classList.add("sidebar-closing");
  setTimeout(() => {
    if (isMobile) {
      sidebar.classList.remove("sidebar-open");
    }
    sidebar.classList.remove("sidebar-closing");
    sidebarOverlay.style.display = "none";
    document.body.style.overflow = "";
    mobileMenuTrigger.classList.add("mobile-visible");
  }, 280);
}

function activateTab(tabBtn, section) {
  tabBtns.forEach((btn) => btn.classList.remove("active"));
  docSections.forEach((sec) => sec.classList.remove("active"));
  tabBtn.classList.add("active");
  section.classList.add("active");
  if (isMobile) closeMobileMenu();
}

tabBtns.forEach((btn) => {
  btn.addEventListener("click", () => {
    const target = btn.dataset.target;
    const section = document.getElementById(target);
    activateTab(btn, section);
  });
});

if (tabBtns.length > 0 && docSections.length > 0) {
  activateTab(tabBtns[0], docSections[0]);
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

mobileMenuTrigger.addEventListener("click", openMobileMenu);
menuCloseBtn.addEventListener("click", closeMobileMenu);
sidebarOverlay.addEventListener("click", closeMobileMenu);
window.addEventListener("resize", updateLayout);
updateLayout();
