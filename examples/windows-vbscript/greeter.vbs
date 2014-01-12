' For each line FOO received on STDIN, respond with "Hello FOO!".
while true
  line = WScript.stdIn.readLine
  WScript.echo "Hello " & line & "!"
wend
