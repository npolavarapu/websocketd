#!/usr/bin/php
<?php

// Simple example script that counts to 10 at ~2Hz, then stops.

for ($count = 1; $count <= 10; $count++) {
	echo $count . "\n";
	usleep(500000);
}

?>
