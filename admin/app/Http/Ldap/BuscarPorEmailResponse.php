<?php

namespace App\Http\Ldap;

class BuscarPorEmailResponse
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
   * BuscarPorEmailResponse constructor.
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
  public function getNombre()
  {
    return $this->nombre;
  }

  /**
   * @return string
   */
  public function getApellido()
  {
    return $this->apellido;
  }

  /**
   * @return string
   */
  public function getNumeroCui()
  {
    return $this->numero_cui;
  }

  /**
   * @return string
   */
  public function getTipoCui()
  {
    return $this->tipo_cui;
  }

  /**
   * @return string
   */
  public function getRlaboral()
  {
    return $this->rlaboral;
  }

  /**
   * @return string
   */
  public function getTipoCuenta()
  {
    return $this->tipo_cuenta;
  }
}