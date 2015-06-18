"use strict";

var artistControllers = angular.module('artistControllers', ['firebase']);

artistControllers.controller('ListController', ['$scope', 'GetData', 'FIREBASE_URL', function ($scope, GetData, FIREBASE_URL) {

    $scope.data = GetData;

    $scope.onDelete = function(deleteID) {
        var deleteRef = new Firebase(FIREBASE_URL + deleteID);
        deleteRef.remove();
    };

    $scope.onAdd = function() {
        var ref = new Firebase(FIREBASE_URL);
        ref.push($scope.user);
    };
}]);

artistControllers.controller('DetailsController', function ($scope, GetData, $routeParams) {
    $scope.data = GetData;
    $scope.whichItem = $routeParams.itemId;
});
