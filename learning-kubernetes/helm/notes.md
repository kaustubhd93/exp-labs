
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
    provides access to all non special files in a chart. cannot be used to access templates. 
    ```
- Capabilities 
    ```
    provides info about what capabilities the k8s cluster supports. 
    ```
- Template
    ```
    contains info about the current template.
    ```

## Values Files

- values.yaml is the default, which can be overridden by a parent chart's values.yaml, which can in turn be overridden by a user-supplied values file, which can in turn be overridden by --set parameters.

## Template functions and pipeline

### Template function syntax : 

```funcName arg1 arg2...```

### Template pipeline syntax : 

```arg | funcName1 | funcname2```
- this syntax is usually used.
