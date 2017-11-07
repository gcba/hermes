<?php

namespace App\Extensions;

use App\User;
use App\Http\Ldap\validar;
use App\Http\Ldap\validarResponse;
// use App\Http\Ldap\validar_porcuit;
// use App\Http\Ldap\validar_porcuitResponse;
// use App\Http\Ldap\buscarporemail;
// use App\Http\Ldap\buscarporemailResponse;
// use App\Http\Ldap\buscarporcuit;
// use App\Http\Ldap\buscarporcuitResponse;

use Illuminate\Contracts\Auth\Authenticatable;
use Illuminate\Auth\EloquentUserProvider;
use Illuminate\Hashing\BcryptHasher;
use Illuminate\Support\Facades\Log;
use Artisaninweb\SoapWrapper\SoapWrapper;

class LdapUserProvider extends EloquentUserProvider
{

  /**
    * @var SoapWrapper
    */
  protected $soapWrapper;

   /**
    * Create a new database user provider.
    *
    * @param  \Illuminate\Contracts\Hashing\Hasher  $hasher
    * @param  string  $model
    * @return void
    */
    public function __construct(BcryptHasher $hasher, $model)
    {
        parent::__construct($hasher, $model);

        $this->soapWrapper = new SoapWrapper();

        $this->setupLDAP();
    }

    private function setupLDAP() {
        $url = env('LDAP_URL', null);

        try {
            $this->soapWrapper->add('LDAP', function ($service) use($url) {
                $service
                    ->wsdl($url)
                    ->trace(false)
                    ->classmap([
                        validar::class,
                        validarResponse::class,
                        // validar_porcuit::class,
                        // validar_porcuitResponse::class,
                        // buscarporemail::class,
                        // buscarporemailResponse::class,
                        // buscarporcuit::class,
                        // buscarporcuitResponse::class
                ]);
            });
        } catch (\SoapFault $fault) {
            Log::error($fault);

            return;
        }
    }

    /**
     * Retrieve a user by the given credentials.
     *
     * @param  array  $credentials
     * @return \Illuminate\Contracts\Auth\Authenticatable|null
     */
    public function retrieveByCredentials(array $credentials)
    {
        if (empty($credentials) || !$this->soapWrapper->has('LDAP')) {
            return null;
        }

        $user = parent::retrieveByCredentials($credentials);

        if (!$user) {
            return null;
        }

        try {
            $validationResponse = $this->soapWrapper->call('LDAP.validar', [
                new validar($credentials['email'], $credentials['password'])
            ]);
        } catch (\SoapFault $fault) {
            Log::error($fault);

            return null;
        }

        if ($validationResponse->return == 1) {
            return $user;
        }

        return null;
    }

   /**
     * Validate a user against the given credentials.
     *
     * @param  \Illuminate\Contracts\Auth\Authenticatable  $user
     * @param  array  $credentials
     * @return bool
     */
    public function validateCredentials(Authenticatable $user, array $credentials)
    {
        return true;
    }
}