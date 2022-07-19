<?php
$t1 = microtime(true);
$arr = array();
for ($i = 0; $i < 1000000; $i++) {
	$value = time();
	$key = $i . '_' . $value;
	$arr[$key] = $value;
}
$t2 = microtime(true);
echo ($t2-$t1)*1000 . " ms\n";

