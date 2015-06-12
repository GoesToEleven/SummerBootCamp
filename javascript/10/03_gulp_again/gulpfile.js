var gulp = require('gulp');
var webserver = require('gulp-webserver');

gulp.task('default', ['webserver']);

gulp.task('webserver', function() {
    gulp.src('builds/development')
        .pipe(webserver({
            livereload: true,
            directoryListing: false,
            open: true
        }));
});