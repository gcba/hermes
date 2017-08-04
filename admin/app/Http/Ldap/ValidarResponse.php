<?php

namespace App\Http\Ldap;

class validarResponse
{
  /**
   * @var string
   */
  protected $return;

  /**
   * validarResponse constructor.
   *
   * @param string $return
   */
  public function __construct($return)
  {
    $this->return = $return;
  }

  /**
   * @return string
   */
  public function getreturn()
  {
    return $this->return;
  }
}