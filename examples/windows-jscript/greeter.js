// For each line FOO received on STDIN, respond with "Hello FOO!".
while (true) {
  var input= WScript.stdIn.readLine();
  WScript.echo('Hello ' + input+ '!');
}
