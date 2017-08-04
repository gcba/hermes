<?php

namespace App\Http\Ldap;

class buscarporemailResponse
{
  /**
   * @var string
   */
  protected $nombre;

  /**
   * @var string
   */
  protected $apellido;

  /**
   * @var string
   */
  protected $numero_cui;

  /**
   * @var string
   */
  protected $tipo_cui;

  /**
   * @var string
   */
  protected $rlaboral;

  /**
   * @var string
   */
  protected $tipo_cuenta;

  /**
   * buscarporemailResponse constructor.
   *
   * @param string $nombre
   * @param string $apellido
   * @param string $numero_cui
   * @param string $tipo_cui
   * @param string $rlaboral
   * @param string $tipo_cuenta
   */
  public function __construct($nombre, $apellido, $numero_cui, $tipo_cui, $rlaboral, $tipo_cuenta)
  {
    $this->nombre = $nombre;
    $this->apellido = $apellido;
    $this->numero_cui = $numero_cui;
    $this->tipo_cui = $tipo_cui;
    $this->rlaboral = $rlaboral;
    $this->tipo_cuenta = $tipo_cuenta;
  }

  /**
   * @return string
   */
  public function getnombre()
  {
    return $this->nombre;
  }

  /**
   * @return string
   */
  public function getapellido()
  {
    return $this->apellido;
  }

  /**
   * @return string
   */
  public function getnumero_cui()
  {
    return $this->numero_cui;
  }

  /**
   * @return string
   */
  public function gettipo_cui()
  {
    return $this->tipo_cui;
  }

  /**
   * @return string
   */
  public function getrlaboral()
  {
    return $this->rlaboral;
  }

  /**
   * @return string
   */
  public function gettipo_cuenta()
  {
    return $this->tipo_cuenta;
  }
}