export class CategoryCard extends HTMLElement {
  static get observedAttributes() {
    return ["name", "slug", "spot", "blogs"];
  }

  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.render();
  }

  attributeChangedCallback(name, oldValue, newValue) {
    this.render();
  }

  render() {
    const name = this.getAttribute("name") ?? "";
    const slug = this.getAttribute("slug") ?? "";
    const spot = this.getAttribute("spot") ?? "";
    const blogs = this.getAttribute("blogs") ?? "";
    const preview = this.getAttribute("preview") === "true";

    function formatDate(date) {
      if (!date) return "";
      const options = { year: "numeric", month: "long", day: "numeric" };
      return new Date(date).toLocaleDateString("en-US", options);
    }

    this.shadowRoot.innerHTML = `
      <style>
        a {
          display: flex;
          padding: 0.75rem;
          flex-direction: column;
          word-break: break-word;
          overflow: hidden;
          border-radius: 0.5rem;
          background-color: var(--smoke-50);
          text-decoration: none;
          color: inherit;
          transition: all 0.2s ease-in-out;
        }

        a:hover {
          box-shadow: 0 0 0 0.125rem var(--smoke-500);
        }

        div {
          display: flex;
          align-items: center;
          justify-content: space-between;
          gap: 0.25rem;
        }

        p {
          margin: 0.5rem 0;
          color: var(--smoke-500);
        }

        span {
          display: flex;
          flex-direction: column;
          word-break: normal;
          align-items: center;
          color: var(--smoke-500);
          background-color: var(--smoke-100);
          padding: 0.25rem 0.5rem;
          border-radius: 0.25rem;
        }

        @media (prefers-color-scheme: dark) {
          a {
            background-color: var(--smoke-950);
            color: var(--smoke-100);
          }
        }
      </style>

      <a href="${preview ? "#" : `/categories/edit/${slug}`}">
        <div>
          <strong>${name}</strong>

          <span>${blogs}</span>
        </div>

        <p>${spot}</p>
      </a>
    `;
  }
}

customElements.define("category-card", CategoryCard);
