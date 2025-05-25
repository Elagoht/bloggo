class VerticalBarChart extends HTMLElement {
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

    const colors = [
      "#1765af",
      "#1e6bba",
      "#2571c4",
      "#2c77ce",
      "#337dd9",
      "#3a83e4",
      "#4e8fe8",
      "#629bed",
      "#76a7f1",
      "#8ab3f5",
      "#9dbff9",
    ];

    const maxValue = Math.max(...this.data);

    this.shadowRoot.innerHTML = `
      <style>
        .vertical-bar-chart {
          display: flex;
          flex-direction: column;
          gap: 1rem;
          background-color: var(--smoke-50);
          border-radius: 1rem;
          padding: 1rem;
        }

        h3 {
          margin: 0;
          font-weight: 500;
        }

        .chart-container {
          display: grid;
          gap: 0.25rem;
          grid-template-columns: repeat(24, 1fr);
          min-height: 300px;
        }

        .bar-item {
          display: flex;
          flex-direction: column;
          gap: 1rem;
          flex: 1;
        }

        .bar-value {
          font-size: 0.8rem;
          line-height: 0;
          transform: translateY(-0.5rem);
          text-align: center;
          transition: all 0.1s ease-in-out;
        }

        .bar-label {
          transform: rotate(-90deg) translateX(-100%);
          width: 0.5rem;
          margin: 0 auto;
          font-size: 0.6rem;
          font-family: monospace;
        }

        #progress {
          position: relative;
          height: 100%;
          width: 100%;
          max-width: 4rem;
          margin: 0 auto;
          border-radius: 0.5rem;
          background-color: var(--smoke-100);
        }

        #progress-value {
          position: absolute;
          bottom: 0;
          left: 0;
          right: 0;
          background-color: var(--gopher-500);
          border-radius: 0.5rem;
          transition: all 0.1s ease-in-out;
        }

        #progress:hover #progress-value {
          background-color: var(--gopher-600);
        }

        #progress:hover .bar-value {
          color: var(--gopher-500);
          scale: 1.25;
        }

        @media (prefers-color-scheme: dark) {
          .vertical-bar-chart {
            background-color: var(--smoke-950);
            color: var(--smoke-100);
          }

          #progress {
            background-color: var(--smoke-1000);
          }
        }
      </style>

      <div class="vertical-bar-chart">
        <h3>${this.title}</h3>
        <div class="chart-container">
          ${this.data
            .map(
              (item, index) => `
              <div class="bar-item">
                <div id="progress">
                  <div id="progress-value" style="height: ${
                    (item / maxValue) * 100
                  }%;background-color: ${
                colors[Math.round((item / maxValue) * 10)]
              }">
                    <div class="bar-value">
                      ${item}
                    </div>
                  </div>
                </div>
                <div class="bar-label">${String(index).padStart(
                  2,
                  "0"
                )}:00</div>
              </div>
            `
            )
            .join("")}
        </div>
      </div>
    `;
  }
}

customElements.define("vertical-bar-chart", VerticalBarChart);
