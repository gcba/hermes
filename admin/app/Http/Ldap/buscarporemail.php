<?php

namespace App\Http\Ldap;

class buscarporemail
{
  /**
   * @var string
   */
  protected $email;

  /**
   * BuscarPorEmailRequest constructor.
   *
   * @param string $email
   */
  public function __construct($email)
  {
    $this->email = $email;
  }

  /**
   * @return string
   */
  public function getemail()
  {
    return $this->email;
  }
}