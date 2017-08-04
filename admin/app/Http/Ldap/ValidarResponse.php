<?php

namespace App\Http\Ldap;

class ValidarResponse
{
  /**
   * @var string
   */
  protected $return;

  /**
   * ValidarResponse constructor.
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
  public function getReturn()
  {
    return $this->return;
  }
}