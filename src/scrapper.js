const leboncoin = require('leboncoin-api');
var search = new leboncoin.Search()
    .setPage(1)
    .setQuery("renove")
    .setFilter(leboncoin.FILTERS.PARTICULIER)
    .setCategory("locations")
    .setRegion("ile_de_france")
    
    .addSearchExtra("price", {min: 1500, max: 2000}) // range of price
    
search.run().then(function (data) {
    console.log(data.page); //  current page
    console.log(data.pages); //  number of pages
    console.log(data.nbResult); //  number of results for this search
    console.log(data.results); //  array of results
    data.results[0].getDetails().then(function (details) {
        console.log(details); //  item 0 with more data
    }, function (err) {
        console.error(err);
    });
}, function (err) {
    console.error(err);
});