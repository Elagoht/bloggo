export class BlogCard extends HTMLElement {
  static get observedAttributes() {
    return ["title", "slug", "image", "published-at", "created-at", "category", "read-time", "preview", "views"];
  }

  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.render();
  }

  attributeChangedCallback(name, oldValue, newValue) {
    this.render();
  }

  getPreviewValue() {
    const value = this.getAttribute("preview");
    return value === "" || value === "true";
  }

  render() {
    const preview = this.getPreviewValue();
    const title = this.getAttribute("title") ?? "";
    const slug = this.getAttribute("slug") ?? "";
    const image = this.getAttribute("image") ?? "";
    const publishedAt = this.getAttribute("published-at");
    const createdAt = this.getAttribute("created-at");
    const category = this.getAttribute("category") ?? "";
    const views = this.getAttribute("views") ?? "";
    const readTime = this.getAttribute("read-time") ?? "";

    function formatDate(date) {
      if (!date) return "";
      const options = { year: "numeric", month: "long", day: "numeric" };
      return new Date(date).toLocaleDateString("en-US", options);
    }

    this.shadowRoot.innerHTML = `
      <style>
        a {
          display: flex;
          word-break: break-word;
          overflow-wrap: break-word;
          flex-direction: column;
          overflow: hidden;
          border-radius: 0.5rem;
          background-color: var(--smoke-50);
          text-decoration: none;
          color: inherit;
          transition: all 0.2s ease-in-out;
          height: 100%;
        }

        figure {
          position: relative;
          aspect-ratio: 16 / 9;
          margin: 0;
          overflow: hidden;
        }

        figcaption {
          position: absolute;
          bottom: 0;
          right: 0;
          border-top-left-radius: 0.5rem;
          opacity: 0.8;
          padding: 0.25rem 0.5rem;
          background-color: var(--gopher-900);
          color: var(--gopher-50);
          font-size: 0.875rem;
        }

        span {
          position: absolute;
          bottom: 0;
          left: 0;
          border-top-right-radius: 0.5rem;
          background-color: var(--gopher-900);
          color: var(--gopher-50);
          font-size: 0.875rem;
          padding: 0.25rem 0.5rem;
          display: flex;
          align-items: center;
          gap: 0.25rem;
          opacity: 0.8;
        }

        img {
          width: 100%;
          transition: all 0.2s ease-in-out;
        }

        div {
          padding: 0.5rem;
        }

        p {
          margin: 0;
          color: var(--smoke-500);
        }

        time {
          color: var(--smoke-500);
          font-size: 0.875rem;
        }

        div:has(time) {
          margin-top: auto;
          display: flex;
          justify-content: space-between;
          align-items: center;
        }

        strong {
          word-break: break-word;
        }

        a:hover img {
          transform: scale(1.05);
        }

        @media (prefers-color-scheme: dark) {
          a {
            background-color: var(--smoke-950);
            color: var(--smoke-100);
          }
        }
      </style>

      <a href="${preview ? "#" : `/blogs/edit/${slug}`}">
        <figure>
          <img src="${image}" alt="${title}" />

          <figcaption>${category}</figcaption>

          <span><icon-eye></icon-eye>${views}</span>
        </figure>

        <div>
          <strong>${title}</strong>
        </div>

        <div>
          <time>${formatDate(publishedAt ?? createdAt)}</time>

          <time>${readTime}</time>
        </div>
      </a>
    `;
  }
}

customElements.define("blog-card", BlogCard);
