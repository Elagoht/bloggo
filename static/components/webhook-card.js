class WebhookCard extends HTMLElement {
  static get observedAttributes() {
    return ["action", "description", "secret", "url", "headers", "body", "method"];
  }

  constructor() {
    super();
  }

  connectedCallback() {
    this.render();
  }

  attributeChangedCallback(name, oldValue, newValue) {
    this.render();
  }

  fireManually() {
    const url = this.getAttribute("url");
    const headers = this.getAttribute("headers");
    const body = this.getAttribute("body");
    const method = this.getAttribute("method");

    fetch(url, {
      method,
      headers: JSON.parse(headers),
      body: JSON.parse(body),
    });
  }

  render() {
    const action = this.getAttribute("action");
    const description = this.getAttribute("description");
    const url = this.getAttribute("url");
    const headers = this.getAttribute("headers");
    const body = this.getAttribute("body");
    const method = this.getAttribute("method");

    this.innerHTML = `
    <style>
      div.webhook-card {
        background-color: var(--smoke-50);
        padding: 0.75rem;
        border-radius: 1rem;
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
      }

      hgroup {
        display: flex;
        flex-wrap: wrap;
        justify-content: space-between;
        align-items: flex-start;
        gap: 0.5rem;
      }

      .content {
        display: flex;
        flex-direction: column;
        gap: 0.25rem;
      }

      h3 {
        font-size: 1.25rem;
        font-weight: 600;
        color: var(--smoke-700);
        margin: 0;
      }

      p {
        font-size: 0.875rem;
        color: var(--smoke-400);
        margin: 0;
      }

      button {
        background-color: var(--gopher-500);
        color: white;
        border: none;
        padding: 0.5rem 1rem;
        border-radius: 0.5rem;
        font-weight: 500;
        cursor: pointer;
        white-space: nowrap;
        transition: background-color 0.2s;
      }

      button:hover {
        background-color: var(--gopher-600);
      }

      .details {
        display: flex;
        flex-direction: column;
        background-color: var(--smoke-100);
        gap: 0.5rem;
        padding: 0.5rem;
        border-radius: 0.5rem;
      }

      .endpoint {
        display: flex;
        gap: 0.5rem;
      }

      .endpoint > * {
        background-color: var(--smoke-100);
        color: var(--smoke-700);
        padding: 0.5rem;
        border-radius: 0.5rem;
      }

      pre {
        font-family: monospace;
        font-size: 0.875rem;
        color: var(--smoke-500);
        white-space: pre-wrap;
        word-break: break-all;
      }

      a {
        width: 100%;
        color: var(--smoke-0);
        text-decoration: none;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
      }

      strong {
        color: var(--smoke-500);
        font-size: 0.75rem;
        font-weight: 600;
        text-transform: uppercase;
        display: block;
      }

      @media(prefers-color-scheme: dark) {
        .webhook-card {
          background-color: var(--smoke-950) !important;
        }

        h3 {
          color: var(--smoke-100);
        }

        .details {
          background-color: var(--smoke-900);
        }

        .endpoint > * {
          background-color: var(--smoke-900);
          color: var(--smoke-100);
        }

        pre {
          color: var(--smoke-100);
        }

        strong {
          color: var(--smoke-300);
        }
      }

      .patch {
        background-color: purple;
        color: white;
      }

      .post {
        background-color: green;
        color: white;
      }

      .get {
        background-color: blue;
        color: white;
      }
    </style>

    <div class="webhook-card">
      <hgroup>
        <div class="content">
          <h3>${action}</h3>
          <p>${description}</p>
        </div>

        <button type="button" onclick="this.closest('webhook-card').fireManually()">Fire Manually</button>
      </hgroup>

      <div class="endpoint">
        <span class="method ${method.toLowerCase()}">${method}</span>
        <a>${url}</a>
      </div>

      <div class="details">
        ${headers ? `
          <strong>HEADERS</strong>
          <pre>${JSON.stringify(JSON.parse(headers), null, 2)}</pre>
        ` : ""}

        ${body ? `
          <strong>BODY</strong>
          <pre>${JSON.stringify(JSON.parse(body), null, 2)}</pre>
        ` : ""}
      </div>
    </div>
    `;
  }
}

customElements.define("webhook-card", WebhookCard);