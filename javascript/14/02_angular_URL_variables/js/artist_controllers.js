"use strict";

var artistControllers = angular.module('artistControllers', ['firebase']);

artistControllers.controller('ListController', ['$scope', 'GetData', function ($scope, GetData) {
    var ref = new Firebase("https://test-swbc-14-01-rout.firebaseio.com/");
    console.log($firebaseArray(ref));
    $scope.data = $firebaseArray(ref);
    //$scope.data = GetData;
    //console.log(GetData);
}]);

//artistControllers.controller('DetailsController', ['$scope', 'GetData', '$routeParams',
//    function ($scope, GetData, $routeParams) {
//        $scope.data = GetData;
//        $scope.whichItem = $routeParams.itemId;
//    }]);
