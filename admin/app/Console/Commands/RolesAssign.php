<?php

namespace App\Console\Commands;

use App\User;
use App\Role;
use Illuminate\Console\Command;

class RolesAssign extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'roles:assign
    {role : The role to apply. Can be \'admin\', \'supervisor\', \'support\' or \'user\'  }
    {email : The user\'s email }';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Assign a role to a user';

    /**
     * Execute the console command.
     *
     * @return mixed
     */
    public function handle()
    {
        $role = $this->argument('role');
        $email = $this->argument('email');

        if (!$role && !$email) {
            $this->error('No arguments passed');

            return;
        }

        if (!$role) {
            $this->error('No role specified');

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

        $this->setRole($role, $email);
    }

    private function setRole(String $roleName, String $email) {
        $role = Role::where('name', $roleName)->first();

        if ($role === null) {
            $this->error("'$roleName' is not a valid role");

            return;
        }

        $user = User::where('email', $email)->first();

        if ($user === null) {
            $this->error('User does not exist');

            return;
        }

        $user->role_id = $role->id;

        if (!$user->save()) {
            $this->error('Error assigning role to user');

            return;
        }

        $this->info('Role assigned successfully');
    }
}
