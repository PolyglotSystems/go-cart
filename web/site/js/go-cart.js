
function getNamespaces(vueContext) {
  // Make a request for a user with a given ID
  axios.get(goCartServiceEndpoint + goCartServicePath + "/namespaces")
  .then(function (response) {
    // handle success
    returnedData = response.data.data;

    var arrayLength = returnedData.length;
    for (var i = 0; i < arrayLength; i++) {
      vueContext.clusterNamespaces.push({name: returnedData[i]});
    }
  })
  .catch(function (error) {
    // handle error
    console.log(error);
  })
  .then(function () {
    // always executed
  });  
}

function loadDashboard(vueContext) {
  // Make a request for a user with a given ID
  axios.get(goCartServiceEndpoint + goCartServicePath + "/dashboard")
  .then(function (response) {
    // handle success
    returnedData = response.data.data;
    
    vueContext.dashboardPage.namespaceCount = returnedData.numberOfNamespaces;
    vueContext.dashboardPage.serviceCount = returnedData.numberOfServices;
    vueContext.dashboardPage.podCount = returnedData.numberOfPods;
  })
  .catch(function (error) {
    // handle error
    console.log(error);
  })
  .then(function () {
    // always executed
  });
}

var app = new Vue({
  el: '#app',
  data: {
    curPage: 'dashboard',
    dashboardPage: {
      namespaceCount: '--',
      serviceCount: '--',
      podCount: '--',
    },
    clusterNamespaces: [],
    namespaceMap: {}
  },
  updated: function() {
    if (this.curPage == 'dashboard') {
      loadDashboard(this);
    }
  },
  mounted: function() {
    document.onreadystatechange = () => {
      if (document.readyState == "complete") {
        console.log('Page completed with image and files!')
        // fetch
        getNamespaces(this);
        switch (this.curPage) {
          case 'dashboard':
            // Load the dashboard
            loadDashboard(this)
          break;
          case 'table':
            // Load the table
          break;
          case 'graph':
            // Load the graph
          break;
        }
      }
    }
  },
  //watch: {
  //  curPage: function (val) {
  //    switch (val) {
  //      case 'dashboard':
  //        // Load the dashboard
  //        //loadDashboard()
  //      break;
  //      case 'table':
  //        // Load the table
  //      break;
  //      case 'graph':
  //        // Load the graph
  //      break;
  //    }
  //  }
  //}
});