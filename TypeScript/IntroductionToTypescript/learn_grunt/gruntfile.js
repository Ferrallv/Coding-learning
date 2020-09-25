module.exports = function(grunt) {
	grunt.loadNpmTasks('grunt-ts');

	grunt.initConfig({
		ts: {
			main: {
				src: ['typescript/main.ts'],
				dest: 'javascript/',
			}
		}
	});

	grunt.registerTask('default', ['ts'])
}
