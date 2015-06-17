var myApp = angular.module('myApp', []);

myApp.controller('MyController', ['$scope', '$http', function($scope, $http){
    $http.get('https://test-swbc-13-02-04.firebaseio.com/').success(function(data) {
        $scope.artists = data;
    });
}]);