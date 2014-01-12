#!/usr/bin/php
<?php

// For each line FOO received on STDIN, respond with "Hello FOO!".
$stdin = fopen('php://stdin', 'r');
while ($line = fgets($stdin)) {
	echo 'Hello ' . trim($line) . "!\n";
}

?>
