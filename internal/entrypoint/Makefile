.PHONY: init build
init:
	dart pub get

build:
	dart2js -O2 -o dist/entrypoint.js bin/entrypoint.dart
