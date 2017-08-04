<?php

namespace App\Http\Ldap;

class ValidarRequest
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
   * ValidarRequest constructor.
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
  public function getEmail()
  {
    return $this->email;
  }

  /**
   * @return string
   */
  public function getClave()
  {
    return $this->clave;
  }
}