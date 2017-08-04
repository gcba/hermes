<?php

namespace App\Ldap;

class ValidarResponse
{
  /**
   * @var string
   */
  protected $Result;

  /**
   * ValidarResponse constructor.
   *
   * @param string
   */
  public function __construct($Result)
  {
    $this->Result = $Result;
  }

  /**
   * @return string
   */
  public function getResult()
  {
    return $this->Result;
  }
}