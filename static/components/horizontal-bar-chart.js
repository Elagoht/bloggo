class HorizontalBarChart extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });

    this.title = this.getAttribute("title");
    this.data = [];
    try {
      this.data = JSON.parse(this.getAttribute("data"));
    } catch (error) {
      console.error("Invalid data received:", error);
    }

    const maxValue = Math.max(...this.data.map((item) => item.value));

    const colors = [
      "#7c3aed",
      "#fbbf24",
      "#10b981",
      "#f43f5e",
      "#f59e42",
      "#3b82f6",
      "#14b8a6",
      "#eab308",
      "#ef4444",
      "#6366f1",
    ];

    this.shadowRoot.innerHTML = `
      <style>
        .horizontal-bar-chart {
          display: flex;
          flex-direction: column;
          background-color: var(--smoke-0);
          padding: 1rem;
          gap: 0.5rem;
          border-radius: 0.5rem;
        }

        h3 {
          font-size: 1.25rem;
          margin: 0;
          margin-bottom: 0.5rem;
          font-weight: 500;
        }

        #progress {
          width: 100%;
          min-width: 50px;
          flex-shrink: 0;
          background-color: var(--smoke-50);
          border-radius: 0.5rem;
        }

        #progress-value {
          height: 0.75rem;
          border-radius: 0.5rem;
          background-color: var(--gopher-500);
        }

        .bar-chart-item {
          display: grid;
          grid-template-columns: auto 1fr auto;
          gap: 0.5rem;
          align-items: center;
          justify-content: space-between;
        }

        .bar-chart-label {
          line-height: 1.15;
          text-align: left;
          max-width: 10rem;
          overflow: hidden;
          font-size: 0.875rem;
          text-overflow: ellipsis;
          -webkit-line-clamp: 2;
          -webkit-box-orient: vertical;
          display: -webkit-box;
        }

        .bar-chart-value {
          text-align: right;
          flex-shrink: 0;
          min-width: 2rem;
        }

        .bar-chart-item:hover #progress-value {
          background-color: var(--gopher-600);
        }

        @media (prefers-color-scheme: dark) {
          .horizontal-bar-chart {
            background-color: var(--smoke-950);
            color: var(--smoke-100);
          }

          #progress {
            background-color: var(--smoke-1000);
          }
        }
      </style>

      <div class="horizontal-bar-chart">
        <h3>${this.title}</h3>
          ${this.data
            .map(
              (item, index) => `
              <div class="bar-chart-item">
                <div class="bar-chart-label"><b>${index + 1}#</b> ${
                item.label
              }</div>
                <div id="progress">
                  <div id="progress-value" style="width: ${
                    (item.value / maxValue) * 100
                  }%; background-color: ${colors[index % colors.length]}"></div>
                </div>
                <div class="bar-chart-value">${item.value}${
                item.unit ? ` ${item.unit}` : ""
              }</div>
              </div>
            `
            )
            .join("")}
      </div>
    `;
  }
}

customElements.define("horizontal-bar-chart", HorizontalBarChart);
