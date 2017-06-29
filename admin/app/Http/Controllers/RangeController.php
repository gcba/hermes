<?php

namespace App\Http\Controllers;

use App\Range;
use Illuminate\Http\Request;
use TCG\Voyager\Http\Controllers;

class RangeController extends \TCG\Voyager\Http\Controllers\VoyagerBreadController
{
    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index(Request $request)
    {
        return parent::index($request);
    }

    /**
     * Show the form for creating a new resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function create(Request $request)
    {
        return parent::create($request);
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function store(Request $request)
    {
        return parent::storage($request);
    }

    /**
     * Display the specified resource.
     *
     * @param  \App\Range  $range
     * @return \Illuminate\Http\Response
     */
    public function show(Request $request, $id)
    {
        return parent::show($request, $id);
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  \App\Range  $range
     * @return \Illuminate\Http\Response
     */
    public function edit(Request $request, $id)
    {
        return parent::edit($request, $id);
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \App\Range  $range
     * @return \Illuminate\Http\Response
     */
    public function update(Request $request, $id)
    {
        return parent::update($request, $id);
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  \App\Range  $range
     * @return \Illuminate\Http\Response
     */
    public function destroy(Request $request, $id)
    {
        return parent::destroy($request, $id);
    }
}
