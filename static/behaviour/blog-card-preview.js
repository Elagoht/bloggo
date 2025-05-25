const blogCard = document.querySelector("#blog-card");
if (!blogCard) {
  console.error("Blog card not found!");
}

const titleInput = document.querySelector("input[name='title']");
const imageInput = document.querySelector("input[name='image']");
const categoryInput = document.querySelector("select[name='category']");
const contentInput = document.querySelector("textarea[name='content']");

if (!contentInput) {
  console.error("Content textarea not found!");
}

function calculateReadTime(content) {
  const wordsPerMinute = 200;
  const wordCount = content.split(" ").length;
  const readTime = Math.ceil(wordCount / wordsPerMinute);
  return `${readTime} min`;
}

titleInput?.addEventListener("input", (event) => {
  console.log("Title changed:", event.target.value);
  blogCard.setAttribute("title", event.target.value);
});

imageInput?.addEventListener("change", (event) => {
  blogCard.setAttribute("image", URL.createObjectURL(event.target.files[0]));
});

categoryInput?.addEventListener("change", (event) => {
  blogCard.setAttribute("category", event.target.options[event.target.selectedIndex].text);
});

contentInput?.addEventListener("input", (event) => {
  blogCard.setAttribute("read-time", calculateReadTime(event.target.value));
});

document.addEventListener("DOMContentLoaded", () => {
  blogCard.setAttribute("published-at", new Date().toISOString());
  blogCard.setAttribute("created-at", new Date().toISOString());
  blogCard.setAttribute("read-time", "0 min read");
  blogCard.setAttribute("preview", "true");
  blogCard.setAttribute("category", "?");
  blogCard.setAttribute("views", "0");
});
