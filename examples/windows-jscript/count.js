// Simple example script that counts to 10 at ~2Hz, then stops.
for (var i = 1; i <= 10; i++) {
  WScript.echo(i);
  WScript.sleep(500);
}
