class PieChart extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });

    this.title = this.getAttribute("title");
    let data = [];
    try {
      data = JSON.parse(this.getAttribute("data"));
    } catch (e) {
      console.error("Invalid data for pie-chart", e);
    }
    if (!Array.isArray(data)) data = [];

    const total = data.reduce((sum, item) => sum + item.value, 0);
    const size = 160;
    const radius = size / 2 - 16;
    const center = size / 2;
    let angle = 0;
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

    const paths = data.map((item, i) => {
      const value = item.value;
      const percent = value / total;
      const startAngle = angle;
      const endAngle = angle + percent * 2 * Math.PI;
      const x1 = center + radius * Math.cos(startAngle);
      const y1 = center + radius * Math.sin(startAngle);
      const x2 = center + radius * Math.cos(endAngle);
      const y2 = center + radius * Math.sin(endAngle);
      const largeArc = percent > 0.5 ? 1 : 0;
      const d = [
        `M ${center} ${center}`,
        `L ${x1} ${y1}`,
        `A ${radius} ${radius} 0 ${largeArc} 1 ${x2} ${y2}`,
        "Z",
      ].join(" ");
      angle = endAngle;
      return `<path d="${d}" fill="${colors[i % colors.length]}" />`;
    });

    const legend = data
      .map(
        (item, i) => `
      <div class="legend-item">
        <span class="legend-color" style="background:${
          colors[i % colors.length]
        }"></span>
        <span class="legend-label">${item.label}</span>
        <span class="legend-value">${item.value}</span>
      </div>
    `
      )
      .join("");

    this.shadowRoot.innerHTML = `
      <style>
        .pie-chart-container {
          display: flex;
          flex-wrap: wrap;
          align-items: center;
          background: var(--smoke-50);
          border-radius: 1rem;
          padding:1rem;
          flex-grow: 1;
          flex: 1;
        }
        svg {
          display: block;
          margin: 0 auto;
          background: none;
        }
        .legend {
          display: flex;
          flex-direction: column;
          align-items: flex-start;
          gap: 0.15rem;
          margin-top: 0.5rem;
        }
        path {
          transform-origin: center;
          transition: transform 0.2s ease;
        }
        path:hover {
          transform: scale(1.1);
        }
        .legend-item {
          display: flex;
          align-items: center;
          gap: 0.3rem;
          font-size: 0.75rem;
          line-height: 1.2;
          padding: 0.1rem 0;
        }
        .legend-color {
          width: 0.75rem;
          height: 0.75rem;
          border-radius: 0.2rem;
          display: inline-block;
        }
        .legend-label {
          font-weight: 500;
          margin-right: 0.2rem;
        }
        .legend-value {
          color: var(--smoke-500);
          font-family: monospace;
          margin-left: 0.2rem;
        }
        h3 {
          font-size: 1.2rem;
          font-weight: 500;
          margin: 0;
        }
        @media (prefers-color-scheme: dark) {
          .pie-chart-container {
            background: var(--smoke-950);
            color: var(--smoke-100);
          }
        }
      </style>
      <div class="pie-chart-container">
        <h3>${this.title}</h3>
        <svg width="${size}" height="${size}" viewBox="0 0 ${size} ${size}">
          ${paths.join("")}
        </svg>
        <div class="legend">
          ${legend}
        </div>
      </div>
    `;
  }
}

customElements.define("pie-chart", PieChart);
