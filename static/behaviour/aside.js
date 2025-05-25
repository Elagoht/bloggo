const body = document.body;
const toggler = document.querySelector("header > button");
const aside = document.querySelector("aside");
let isOpen = window.innerWidth > 768;

toggler.addEventListener("click", () => {
  isOpen = !isOpen;
  handleCollapse(isOpen);
});

function handleCollapse(isOpen) {
  aside.style.transform = isOpen ? "translateX(0)" : "translateX(-100%)";
}

window.addEventListener("load", () => {
  if (window.innerWidth < 768) {
    handleCollapse(false);
  }
  setTimeout(() => {
    aside.style.transition = "all 0.2s ease-in-out";
  }, 200);
});

window.addEventListener("resize", () => {
  if (window.innerWidth > 768) {
    handleCollapse(true);
  } else {
    handleCollapse(false);
  }
});
