package chart

const ChartTmp = `
<html>

<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width">

<script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/3.9.1/chart.min.js" integrity="sha512-ElRFoEQdI5Ht6kZvyzXhYG9NqjtkmlkfYk0wr6wHxU9JEHakS7UJZNeml5ALk+8IKlU6jDgMabC3vkumRokgJA==" crossorigin="anonymous" referrerpolicy="no-referrer"></script>

<script>

window.addEventListener('load', (event) => {
  loadCharts()
});

const loadCharts = function() {
  loadLineChart(0)
}

const loadLineChart = function(i) {
  const labels = [
    {{range .X}}
    {{.}},
    {{end}}
  ];

  const data = {
    labels: labels,
    datasets: [
      {
        label: '{{.Title}}',
        backgroundColor: 'rgb(204, 255, 153)',
        borderColor: 'rgb(204, 255, 153)',
        data: [
          {{range .Y}}
	  {{.}},
	  {{end}}
        ],
      },
      {
        label: '{{.Title}}-projected',
        backgroundColor: 'rgb(40, 88, 99)',
        borderColor: 'rgb(150, 255, 153)',
        data: [
          {{range .YProject}}
	  {{.}},
	  {{end}}
        ],
      },
      {
        label: '{{.Title}}-projected-upper',
        backgroundColor: 'rgb(230, 140, 99)',
        borderColor: 'rgb(95, 168, 9)',
        data: [
          {{range .YProjectUpper}}
	  {{.}},
	  {{end}}
        ],
      },
      /*
      {
        label: '{{.Title}}-projected-lower',
        backgroundColor: 'rgb(9, 9, 9)',
        borderColor: 'rgb(9, 9, 9)',
        data: [
          {{range .YProjectLower}}
	  {{.}},
	  {{end}}
        ],
      },
      */
    ]
  }

  const config = {
    type: 'line',
    data: data,
    options: {}
  };

  let elementId = "lineChart" + i.toString()
  console.log("elementId is", elementId) 

  const myChart = new Chart(document.getElementById(elementId), config)
}

const loadBarChart = function() {
  const ctx = document.getElementById('barChart').getContext('2d');

  const myChart = new Chart(ctx, {
    type: 'bar',
    data: {
        labels: ['Red', 'Blue', 'Yellow', 'Green', 'Purple', 'Orange'],
        datasets: [{
            label: '# of Votes',
            data: [12, 19, 3, 5, 2, 3],
            backgroundColor: 'rgb(255, 99, 132)',
            borderColor: 'rgb(255, 99, 132)',
            borderWidth: 1
        }]
    },
    options: {}
  });
}
</script>
      
<style>
      html {
        font-family: sans-serif;
      }

      body {
        margin: 0;
      }

      header {
        background: purple;
        height: 100px;
      }

      h1 {
        text-align: center;
        color: white;
        line-height: 100px;
        margin: 0;
      }

      article {
        padding: 10px;
        margin: 10px;
        width: 80%;
        height: 80%;
      }

      /* Add your flexbox CSS below here */

      section {
        display: flex;
	flex-direction: row;
        flex-wrap: wrap
      }

      
</style>

</head>

<body>
    <section>

      <article>
        <h2>Test line chart1</h2>

	<canvas id="lineChart0"></canvas>
      </article>

      <article>
        <h2>End of sample charts</h2>
    </section>
	<h1> Is there a chart below this </h1>
</body>
</html>
`
