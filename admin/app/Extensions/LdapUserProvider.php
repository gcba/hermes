<?php

namespace App\Extensions;

use App\User;
use App\Http\Ldap\ValidarRequest;
use App\Http\Ldap\ValidarResponse;
use App\Http\Ldap\BuscarPorEmailRequest;
use App\Http\Ldap\BuscarPorEmailResponse;
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
                    LdapRequest::class,
                    LdapResponse::class,
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
            new ValidarRequest($credentials['email'], $credentials['password'])
        ]);
    }
}