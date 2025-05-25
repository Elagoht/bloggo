const categoryCard = document.querySelector("#category-card");
if (!categoryCard) {
  console.error("Category card not found!");
}

const nameInput = document.querySelector("input[name='name']");
const spotInput = document.querySelector("input[name='spot']");

nameInput?.addEventListener("input", (event) => {
  console.log("Title changed:", event.target.value);
  categoryCard.setAttribute("name", event.target.value);
});

spotInput?.addEventListener("input", (event) => {
  console.log("Spot changed:", event.target.value);
  categoryCard.setAttribute("spot", event.target.value);
});

document.addEventListener("DOMContentLoaded", () => {
  categoryCard.setAttribute("blogs", "0");
});
