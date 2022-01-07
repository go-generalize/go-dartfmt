import 'package:js/js.dart';
import "package:dart_style/dart_style.dart";

@JS('formatDart')
external set _formatDart(String Function(String code) f);

String formatDart(String code) {
  final formatter = DartFormatter();

  return formatter.format(code);
}

void main() {
  _formatDart = allowInterop(formatDart);
}
