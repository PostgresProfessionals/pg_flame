package html

import (
	"html/template"
)

var templateHTML *template.Template = template.Must(template.New("html").Parse(`
    <div class="flame-chart-container">
      <div id="chart" class="flame-chart-style">
      </div>
      <hr>
      <div id="details" class="flame-detail-style">
      </div>
    </div>

    <script type="text/javascript">
    var flameGraph = d3.flamegraph()
      .width(960)
      .cellHeight(18)
      .transitionDuration(750)
      .minFrameSize(5)
      .transitionEase(d3.easeCubic)
      .sort(false)
      .title("")
      .differential(false)
      .selfValue(false)
      .setColorMapper(function(d, originalColor) {
        return d.data.color || originalColor;
      });

    var tip = d3.tip()
      .direction("s")
      .offset([8, 0])
      .attr('class', 'd3-flame-graph-tip')
      .html(function(d) {
        return d.data.name + " | " + d.data.time + " ms";
      });
    flameGraph.tooltip(tip);

    var details = document.getElementById("details");
    flameGraph.setDetailsElement(details);

    var label = function(d) {
      return d.data.detail;
    }
    flameGraph.label(label);

    var data = {{.}};

    d3.select("#chart")
      .datum(data)
      .call(flameGraph);

    </script>
`))

var templateTable *template.Template = template.Must(template.New("table").Parse(`
<table class="table table-striped table-bordered">
  <tbody>
    {{if .Method}}
      <tr>
        <th>Method</th>
        <td>{{.Method}}</td>
      </tr>
    {{end}}
    {{if .Table}}
      <tr>
        <th>Table</th>
        <td>{{.Table}}</td>
      </tr>
    {{end}}
    {{if .Index}}
      <tr>
        <th>Index</th>
        <td>{{.Index}}</td>
      </tr>
    {{end}}
    {{if .Alias}}
      <tr>
        <th>Alias</th>
        <td>{{.Alias}}</td>
      </tr>
    {{end}}
    {{if .ParentRelationship}}
      <tr>
        <th>Parent Relationship</th>
        <td>{{.ParentRelationship}}</td>
      </tr>
    {{end}}
    {{if .PlanCost}}
      <tr>
        <th>Plan Cost</th>
        <td>{{.PlanCost}}</td>
      </tr>
    {{end}}
    {{if .PlanRows}}
      <tr>
        <th>Plan Rows</th>
        <td>{{.PlanRows}}</td>
      </tr>
    {{end}}
    {{if .PlanWidth}}
      <tr>
        <th>Plan Width</th>
        <td>{{.PlanWidth}}</td>
      </tr>
    {{end}}
    {{if .ActualTotalTime}}
      <tr>
        <th>Actual Total Time</th>
        <td>{{.ActualTotalTime}} ms</td>
      </tr>
    {{end}}
    {{if .ActualRows}}
      <tr>
        <th>Actual Rows</th>
        <td>{{.ActualRows}}</td>
      </tr>
    {{end}}
    {{if .ActualLoops}}
      <tr>
        <th>Actual Loops</th>
        <td>{{.ActualLoops}}</td>
      </tr>
    {{end}}
    {{if .Filter}}
      <tr>
        <th>Filter</th>
        <td>{{.Filter}}</td>
      </tr>
    {{end}}
    {{if .JoinFilter}}
      <tr>
        <th>Join Filter</th>
        <td>{{.JoinFilter}}</td>
      </tr>
    {{end}}
    {{if .HashCond}}
      <tr>
        <th>Hash Cond</th>
        <td>{{.HashCond}}</td>
      </tr>
    {{end}}
    {{if .IndexCond}}
      <tr>
        <th>Index Cond</th>
        <td>{{.IndexCond}}</td>
      </tr>
    {{end}}
    {{if .RecheckCond}}
      <tr>
        <th>Recheck Cond</th>
        <td>{{.RecheckCond}}</td>
      </tr>
    {{end}}
    {{if .BuffersHit}}
      <tr>
        <th>Buffers Shared Hit</th>
        <td>{{.BuffersHit}}</td>
      </tr>
    {{end}}
    {{if .BuffersRead}}
      <tr>
        <th>Buffers Shared Read</th>
        <td>{{.BuffersRead}}</td>
      </tr>
    {{end}}
    {{if .HashBuckets}}
      <tr>
        <th>Hash Buckets</th>
        <td>{{.HashBuckets}}</td>
      </tr>
    {{end}}
    {{if .HashBatches}}
      <tr>
        <th>Hash Batches</th>
        <td>{{.HashBatches}}</td>
      </tr>
    {{end}}
    {{if .MemoryUsage}}
      <tr>
        <th>Memory Usage</th>
        <td>{{.MemoryUsage}} kB</td>
      </tr>
    {{end}}
    {{if .SortKey}}
      <tr>
        <th>Sort Key</th>
        <td>{{.SortKey}}</td>
      </tr>
    {{end}}
    {{if .SortMethod}}
      <tr>
        <th>Sort Method</th>
        <td>{{.SortMethod}}</td>
      </tr>
    {{end}}
    {{if .SortSpaceUsed}}
      <tr>
        <th>Sort Space Used</th>
        <td>{{.SortSpaceUsed}} kB</td>
      </tr>
    {{end}}
    {{if .SortSpaceType}}
      <tr>
        <th>Sort Space Type</th>
        <td>{{.SortSpaceType}}</td>
      </tr>
    {{end}}
  </tbody>
</table>
`))
