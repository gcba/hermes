<?php

namespace App\Console\Commands;

use App\User;
use App\Role;
use Illuminate\Console\Command;

class AdminCreate extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'admin:create
    {name : The user\'s name }
    {email : The user\'s email }';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Create a new admin';

    /**
     * Execute the console command.
     *
     * @return mixed
     */
    public function handle()
    {
        $name = $this->argument('name');
        $email = $this->argument('email');

        if (!$name && !$email) {
            $this->error('No arguments passed');

            return;
        }

        if (!$name) {
            $this->error('No name specified');

            return;
        }

        if (!$email) {
            $this->error('No email address');

            return;
        }

        if (!filter_var($email, FILTER_VALIDATE_EMAIL)) {
            $this->error('Invalid email address');

            return;
        }

        $this->setRole($name, $email);
    }

    private function setRole(String $userName, String $email) {
        $role = Role::where('name', 'admin')->first();

        if ($role === null) {
            $this->error('Admin role does not exist');

            return;
        }

        $user  = new User;

        $user->name = $userName;
        $user->email = $email;
        $user->role_id = $role->id;

        if (!$user->save()) {
            $this->error('Error creating admin');

            return;
        }

        $this->info('Admin created successfully');
    }
}
