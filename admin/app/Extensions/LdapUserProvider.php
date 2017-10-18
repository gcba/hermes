<?php

namespace App\Extensions;

use App\User;
use App\Http\Ldap\validar;
use App\Http\Ldap\validarResponse;
// use App\Http\Ldap\validar_porcuit;
// use App\Http\Ldap\validar_porcuitResponse;
use App\Http\Ldap\buscarporemail;
use App\Http\Ldap\buscarporemailResponse;
// use App\Http\Ldap\buscarporcuit;
// use App\Http\Ldap\buscarporcuitResponse;
use Artisaninweb\SoapWrapper\SoapWrapper;

use Illuminate\Contracts\Auth\Authenticatable;
use Illuminate\Auth\EloquentUserProvider;
use Illuminate\Hashing\BcryptHasher;
use Illuminate\Support\Facades\Log;

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
        $url = env('LDAP_URL', 'https://esb-qa.gcba.gob.ar/ad/consulta?wsdl');

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
                        buscarporemail::class,
                        buscarporemailResponse::class,
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

        try {
            $validationResponse = $this->soapWrapper->call('LDAP.validar', [
                new validar($credentials['email'], $credentials['password'])
            ]);
        } catch (\SoapFault $fault) {
            Log::error($fault);

            return null;
        }

        if ($validationResponse->return == 1) {
            $user = parent::retrieveByCredentials($credentials);

            if (!$user) {
                try {
                    $userDataResponse = $this->soapWrapper->call('LDAP.buscarporemail', [
                        new buscarporemail($credentials['email'])
                    ]);
                } catch (\SoapFault $fault) {
                    Log::error($fault);

                    return null;
                }

                $newUser = new User;

                $newUser->name = $userDataResponse->return->nombre . ' ' . $userDataResponse->return->apellido;
                $newUser->email = $credentials['email'];

                $newUser->save();

                return $newUser;
            }

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