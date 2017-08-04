<?php

namespace App\Http\Controllers;

use Artisaninweb\SoapWrapper\SoapWrapper;
use App\Http\Ldap\LdapRequest;
use App\Http\Ldap\LdapResponse;

class LdapController
{
  /**
   * @var SoapWrapper
   */
  protected $soapWrapper;

  /**
   * LdapController constructor.
   *
   * @param SoapWrapper $soapWrapper
   */
  public function __construct(SoapWrapper $soapWrapper)
  {
    $this->soapWrapper = $soapWrapper;
  }

  /**
   * Use the SoapWrapper
   */
  public function show()
  {
    $this->soapWrapper->add('Currency', function ($service) {
      $service
        ->wsdl('https://esb-hml.gcba.gob.ar/ad/consulta?wsdl') // TODO: Mind the environments
        ->trace(true)
        ->classmap([
          LdapRequest::class,
          LdapResponse::class,
        ]);
    });

    /*
    // Without classmap
    $response = $this->soapWrapper->call('Currency.GetConversionAmount', [
      'CurrencyFrom' => 'USD',
      'CurrencyTo'   => 'EUR',
      'RateDate'     => '2014-06-05',
      'Amount'       => '1000',
    ]);

    var_dump($response);
    */

    // With classmap
    $response = $this->soapWrapper->call('validacionldap.validar', [
      new LdapRequest('USD', 'EUR')
    ]);

    var_dump($response);
    exit;
  }
}