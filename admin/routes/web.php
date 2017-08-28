<?php

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/

Route::get('/', function () {
    return view('welcome');
});

Route::group(['prefix' => 'admin'], function () {
    Voyager::routes();

    Route::get('ratings.api', 'DataTablesController@ratingsAPI')->name('ratings.api');
    Route::get('messages.api', 'DataTablesController@messagesAPI')->name('messages.api');
    Route::get('devices.api', 'DataTablesController@devicesAPI')->name('devices.api');
    Route::get('appusers.api', 'DataTablesController@appusersAPI')->name('appusers.api');

    Route::post('messages/receive', 'MailgunController@receive')
        ->name('messages.receive')
        ->middleware(\App\Http\Middleware\ValidateMailgun::class, \App\Http\Middleware\Sessionless::class);
});