<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;
use Illuminate\Filesystem\FileNotFoundException;
use Illuminate\Support\Facades\File;

abstract class Token extends Command
{
    /**
     * The private key.
     *
     * @var string
     */
    abstract protected function key();

    /**
     * Execute the console command.
     *
     * @return mixed
     */
    public function handle()
    {
        try
        {
            $key = File::get($this->key());
        }
        catch (FileNotFoundException $exception)
        {
            $this->error('Cannot read key file');
        }

        // Adapted from https://stackoverflow.com/questions/33773477/jwt-json-web-token-in-php-without-using-3rd-party-library-how-to-sign

        $headers = ['alg' => 'RS256', 'typ' => 'JWT'];
        $payload = json_encode([ 'at' => time()]);
        $signature = '';
        $headersEncoded = $this->base64UrlEncode(json_encode($headers));
        $payloadEncoded = $this->base64UrlEncode($payload);

        openssl_sign("$headersEncoded.$payloadEncoded", $signature, $key, OPENSSL_ALGO_SHA256);

        $signatureEncoded = $this->base64UrlEncode($signature);
        $token = "$headersEncoded.$payloadEncoded.$signatureEncoded";
        $this->info($token);
    }

    protected function base64UrlEncode($data)
    {
        return rtrim(strtr(base64_encode($data), '+/', '-_'), '=');
    }
}
