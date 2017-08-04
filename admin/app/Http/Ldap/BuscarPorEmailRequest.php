<?php

namespace App\Http\Ldap;

class BuscarPorEmailRequest
{
  /**
   * @var string
   */
  protected $Email;

  /**
   * BuscarPorEmailRequest constructor.
   *
   * @param string $Email
   */
  public function __construct($Email, $Clave)
  {
    $this->Email = $Email;
  }

  /**
   * @return string
   */
  public function getEmail()
  {
    return $this->Email;
  }
}