<h2>Provider</h2>
<ul>
<li>Makefile -> Creates the provider binary and stores it under ~/.terraform.d/plugins/</li>
<li>main.go -> Entry point of the program (faciltates connection between Terraform Core and the provider)</li>
<li>provider/provider.go -> Defines the initial provider and the names for each data source/resource</li>
<br/>
<li>vendor/ -> Contains provider dependencies</li>
<li>examples/ -> Contains example .tf file</li>
<li>provider/datasource/ -> Contains datasource files</li>
<li>provider/resource/ -> Contains resource files</li>
</ul>

<h2>Adding new data sources/resources</h2>
Data Sources: Create a <DATA-SOURCE-NAME>.go file under provider/datasource. Use provider/datasource/service.go as an example.
Resources: Create a <RESOURCE-NAME>.go file under provider/resource. Use provider/resource/resource.go as an example.
