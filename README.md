# How to Use

Open a browser and run

```
$ go run cmd/server/server.go
```

Now all the files in the current directory can be viewed over the browser.

The csv file unique_visitors.csv can be plotted into a graph as follow. 

```
$ go run cmd/chartMain/chartCsv.go unique_visitors.csv > out.html 

```

Open out.html by looking at http://localhost:8080/out.html

http://localhost:8080/list gives you a view of the files in that directory

## TODO

* Use arima to project csv's 
>> Have a rough integration. Try to plot the data and see if it makes any sense.
* Show projections against original somehow.
