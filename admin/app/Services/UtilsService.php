<?php

namespace App\Services;

use DateTime;

class UtilsService
{
    public static function formatDate($dateString) {
        $date = new DateTime($dateString);

        return $date->format(env('APP_DATETIME_FORMAT', 'd/m/Y H:i:s'));
    }

    public static function beginsWith($string, $prefix) {
        return strncmp($string, $prefix, strlen($prefix)) === 0;
    }
}