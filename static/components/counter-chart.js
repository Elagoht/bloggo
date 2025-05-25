class CounterChart extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });

    this.data = [];
    try {
      this.data = JSON.parse(this.getAttribute("data"));
    } catch (error) {
      console.error("Invalid data received:", error);
    }

    this.shadowRoot.innerHTML = `
      <style>
        .counter-chart {
          display: grid;
          grid-template-columns: repeat(${this.data.length}, 1fr);
          align-items: center;
          justify-content: center;
          background-color: var(--smoke-0);
          color: var(--smoke-900);
          text-align: center;
          height: 100%;
          border-radius: 0.5rem;
        }
        .counter {
          padding: 0.5rem;
          gap: 0.5rem;
          display: flex;
          flex-direction: column;
          align-items: center;
          justify-content: center;
        }
        .counter span {
          font-size: 0.875rem;
        }
        .counter strong {
          font-size: 2rem;
          font-weight: 600;
        }

        @media (prefers-color-scheme: dark) {
          .counter-chart {
            background-color: var(--smoke-950);
            color: var(--smoke-100);
          }
        }
      </style>
      <div class="counter-chart">
        ${this.data
          .map(
            (item) => `
          <div class="counter">
            <strong>${item.value}</strong>
            <span>${item.label}</span>
          </div>
        `
          )
          .join("")}
      </div>
    `;
  }
}

customElements.define("counter-chart", CounterChart);
