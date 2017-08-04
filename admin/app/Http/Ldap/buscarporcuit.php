<?php

namespace App\Http\Ldap;

class buscarporcuit
{
  /**
   * @var string
   */
  protected $numero;

  /**
   * buscarporcuit constructor.
   *
   * @param string $numero
   */
  public function __construct($numero)
  {
    $this->numero = $numero;
  }

  /**
   * @return string
   */
  public function getnumero()
  {
    return $this->numero;
  }
}