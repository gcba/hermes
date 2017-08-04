<?php

namespace App\Http\Ldap;

class validar
{
  /**
   * @var string
   */
  protected $email;

  /**
   * @var string
   */
  protected $clave;

  /**
   * validar constructor.
   *
   * @param string $email
   * @param string $clave
   */
  public function __construct($email, $clave)
  {
    $this->email = $email;
    $this->clave = $clave;
  }

  /**
   * @return string
   */
  public function getemail()
  {
    return $this->email;
  }

  /**
   * @return string
   */
  public function getclave()
  {
    return $this->clave;
  }
}