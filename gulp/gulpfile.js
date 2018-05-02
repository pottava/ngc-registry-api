"use strict";

var gulp = require('gulp');
var shell = require('gulp-shell');

gulp.task("api", function() {
  gulp.src("/monitor/generated/cmd/ngc-registry-server/main.go")
      .pipe(shell(['docker restart api'], {ignoreErrors: true}));
});

gulp.task("default", function() {
    gulp.watch("/monitor/controllers/**/*.go", ["api"]);
    gulp.watch("/monitor/lib/**/*.go",         ["api"]);
    gulp.watch("/monitor/ngc/**/*.go",         ["api"]);
});
