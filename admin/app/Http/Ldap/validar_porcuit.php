<?php

namespace App\Http\Ldap;

class validar_porcuitRequest
{
  /**
   * @var string
   */
  protected $numero;

  /**
   * @var string
   */
  protected $clave;

  /**
   * validar_porcuitRequest constructor.
   *
   * @param string $numero
   * @param string $clave
   */
  public function __construct($numero, $clave)
  {
    $this->numero = $numero;
    $this->clave = $clave;
  }

  /**
   * @return string
   */
  public function getnumero()
  {
    return $this->numero;
  }

  /**
   * @return string
   */
  public function getclave()
  {
    return $this->clave;
  }
}