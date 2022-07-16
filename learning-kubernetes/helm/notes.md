
# Chart structure

sample-chart/
  Chart.yaml            -- Description of the chart. 
  values.yaml           -- default values for a chart
  charts/               -- subcharts for dependencies
  templates/            -- directory for template files

## Template format

- A template directive is enclosed in {{ and }} blocks
- The template directive {{ .Release.Name }} injects the release name into the template. The values that are passed into a template can be thought of as namespaced objects, where a dot (.) separates each namespaced element
- The leading dot before Release indicates that we start with the top-most namespace for this scope (we'll talk about scope in a bit). So we could read .Release.Name as "start at the top namespace, find the Release object, then look inside of it for an object called Name"

## Built in objects :

- Release
- Values
- Chart
- Files 
    ```

    ```
- Capabilities 
    ```

    ```
- Template
    ```
    
    ```