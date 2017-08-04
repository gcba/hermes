<?php

namespace App\Extensions;

use App\User;
use App\Http\Ldap\validar;
use App\Http\Ldap\validarResponse;
use App\Http\Ldap\validar_porcuit;
use App\Http\Ldap\validar_porcuitResponse;
use App\Http\Ldap\buscarporemail;
use App\Http\Ldap\buscarporemailResponse;
use App\Http\Ldap\buscarporcuit;
use App\Http\Ldap\buscarporcuitResponse;
use Artisaninweb\SoapWrapper\SoapWrapper;
use Illuminate\Auth\EloquentUserProvider;

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
    public function __construct(HasherContract $hasher, $model)
    {
        $this->model = $model;
        $this->hasher = $hasher;
        $this->soapWrapper = new SoapWrapper();

        $this->soapWrapper->add('Ldap', function ($service) {
            $service
                ->wsdl('https://esb-hml.gcba.gob.ar/ad/consulta?wsdl') // TODO: Mind the environments
                ->trace(true)
                ->classmap([
                    validar::class,
                    validarResponse::class,
                    validar_porcuit::class,
                    validar_porcuitResponse::class,
                    buscarporemail::class,
                    buscarporemailResponse::class,
                    buscarporcuit::class,
                    buscarporcuitResponse::class
                ]);
            });
    }

    /**
     * Retrieve a user by the given credentials.
     *
     * @param  array  $credentials
     * @return \Illuminate\Contracts\Auth\Authenticatable|null
     */
    public function retrieveByCredentials(array $credentials)
    {
        if (empty($credentials)) {
            return;
        }

        // Rest of the implementation
    }

  /**
   * Use the SoapWrapper
   */
    public function validateCredentials(UserInterface $user, array $credentials)
    {
        $response = $this->soapWrapper->call('Ldap.validar', [
            new validar($credentials['email'], $credentials['password'])
        ]);
    }
}