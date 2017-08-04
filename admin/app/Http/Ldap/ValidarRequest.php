<?php

namespace App\Http\Ldap;

class ValidarRequest
{
  /**
   * @var string
   */
  protected $Email;

  /**
   * @var string
   */
  protected $Clave;

  /**
   * ValidarRequest constructor.
   *
   * @param string $Email
   * @param string $Clave
   */
  public function __construct($Email, $Clave)
  {
    $this->Email = $Email;
    $this->Clave = $Clave;
  }

  /**
   * @return string
   */
  public function getEmail()
  {
    return $this->Email;
  }

  /**
   * @return string
   */
  public function getClave()
  {
    return $this->Clave;
  }
}